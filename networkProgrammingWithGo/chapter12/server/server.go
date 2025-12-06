package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/deng/go-space/networkProgrammingWithGo/chapter12/housework/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var addr, certFn, keyFn, dataFile string

func init() {
	flag.StringVar(&addr, "address", "localhost:34443", "listen address")
	flag.StringVar(&certFn, "cert", "cert.pem", "certificate file")
	flag.StringVar(&keyFn, "key", "key.pem", "private key file")
	flag.StringVar(&dataFile, "file", "housework.db", "data file")
}

func main() {
	flag.Parse()

	creds, err := credentials.NewServerTLSFromFile(certFn, keyFn)
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer(grpc.Creds(creds))
	housework.RegisterRobotMaidServer(grpcServer, NewRosie(dataFile))
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Listening for TLS connection on %s ...", addr)
	grpcServer.Serve(listener)

	// server := grpc.NewServer()
	// housework.RegisterRobotMaidServer(server, &Rosie{})

	// cert, err := tls.LoadX509KeyPair(certFn, keyFn)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// listener, err := net.Listen("tcp", addr)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("Listening for TLS connection on %s ...", addr)

	// log.Fatal(
	// 	server.Serve(
	// 		tls.NewListener(
	// 			listener,
	// 			&tls.Config{
	// 				Certificates:     []tls.Certificate{cert},
	// 				CurvePreferences: []tls.CurveID{tls.CurveP256},
	// 				MinVersion:       tls.VersionTLS12,
	// 				NextProtos:       []string{"h2"},
	// 			},
	// 		),
	// 	),
	// )

}
