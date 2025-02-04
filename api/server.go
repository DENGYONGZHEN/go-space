package api

import (
	"fmt"
	db "simple-bank/db/sqlc"
	"simple-bank/token"
	"simple-bank/util"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Server serves HTTP requests for banking service
type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{config: config, store: store, tokenMaker: tokenMaker}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}
	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()
	//add routes to router
	router.POST("/user", server.createUser)
	router.POST("/user/login", server.loginUser)
	router.POST("/tokens/renew_acess", server.renewAccessToken)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	authRoutes.POST("/account", server.createAccount)
	authRoutes.GET("/account/:id", server.getAccountById)
	authRoutes.GET("/accounts", server.listAccounts)
	authRoutes.POST("/transfer", server.createTransfer)

	server.router = router
}

// Start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
