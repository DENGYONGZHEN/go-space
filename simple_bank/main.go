package main

import (
	"context"
	"os"

	"github.com/hibiken/asynq"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"net"
	"net/http"
	"simple-bank/api"
	db "simple-bank/db/sqlc"
	_ "simple-bank/doc/statik"
	"simple-bank/gapi"
	"simple-bank/mail"
	"simple-bank/pb"
	"simple-bank/util"
	"simple-bank/worker"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rakyll/statik/fs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Msg("cannot load config")
	}
	if config.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}
	pool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal().Msg("Unable to connect to database")
	}

	//run db migration
	runDBMigration(config.MigrationURL, config.DBSource)

	testStore := db.NewStore(pool)

	redisOpt := asynq.RedisClientOpt{
		Addr: config.RedisAddress,
	}
	taskDistributor := worker.NewRedisTaskDistibutor(redisOpt)

	go runTaskProcessor(config, redisOpt, testStore)
	go runGatewayServer(config, testStore, taskDistributor)
	runGrpcServer(config, testStore, taskDistributor)
}

func runDBMigration(migrationURL string, dbSource string) {

	m, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatal().Msg("can not create new migrate instance")
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal().Msg("failed to run migrate up")
	}

	log.Info().Msg("db migrated successfully")

}

func runTaskProcessor(config util.Config, redisOpt asynq.RedisClientOpt, store db.Store) {

	mailer := mail.NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EMAILEmailSenderPassword)
	taskProcessor := worker.NewRedisTaskProcessor(redisOpt, store, mailer)
	log.Info().Msg("start task processor")
	if err := taskProcessor.Start(); err != nil {
		log.Fatal().Err(err).Msg("failed to start task processor")
	}

}

func runGrpcServer(config util.Config, testStore db.Store, taskDistributor worker.TaskDistributor) {
	server, err := gapi.NewServer(config, testStore, taskDistributor)
	if err != nil {
		log.Fatal().Msg("cannot create server")
	}

	//grpc Logger
	grpcLogger := grpc.UnaryInterceptor(gapi.GrpcLogger)
	grpcServer := grpc.NewServer(grpcLogger)
	pb.RegisterSimpleBankServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GrpcServerAddress)
	if err != nil {
		log.Fatal().Msg("cannot create listener")
	}
	log.Info().Msgf("start gRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal().Msg("cannot start gRPC server")
	}
}

func runGatewayServer(config util.Config, testStore db.Store, taskDistributor worker.TaskDistributor) {
	server, err := gapi.NewServer(config, testStore, taskDistributor)
	if err != nil {
		log.Fatal().Msg("cannot create server")
	}
	grpcMux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	}))

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = pb.RegisterSimpleBankHandlerServer(ctx, grpcMux, server)
	if err != nil {
		log.Fatal().Msg("cannot register handler server")
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	// fs := http.FileServer(http.Dir("./doc/swagger"))
	// mux.Handle("/swagger/", http.StripPrefix("/swagger/", fs))
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal().Msg("cannot create statik fs")
	}
	swaggerHandler := http.StripPrefix("/swagger/", http.FileServer(statikFS))
	mux.Handle("/swagger/", swaggerHandler)

	listener, err := net.Listen("tcp", config.HTTPServerAddress)
	if err != nil {
		log.Fatal().Msg("cannot create listener")
	}
	log.Info().Msgf("start HTTP gateway server at %s", listener.Addr().String())

	err = http.Serve(listener, gapi.HttpLogger(mux))
	if err != nil {
		log.Fatal().Msg("cannot start HTTP gateway server")
	}
}

func runGinServer(config util.Config, testStore db.Store) {
	server, err := api.NewServer(config, testStore)
	if err != nil {
		log.Fatal().Msg("cannot create server")
	}
	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal().Msg("cannot start server")
	}
}
