package api

import (
	"database/sql"
	"go-api/service/product"
	"go-api/service/user"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {

	return &APIServer{
		addr: addr,
		db:   db,
	}

}

func (s *APIServer) Run() error {

	router := mux.NewRouter()

	subrouter := router.PathPrefix("/api/v1").Subrouter()
	//user routes
	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	//product routes
	productStore := product.NewStore(s.db)
	productHandler := product.NewHandler(productStore)
	productHandler.RegisterRoutes(subrouter)

	return http.ListenAndServe(s.addr, router)
}
