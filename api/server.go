package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/huannguyen2114/go-toy-project/db/sqlc"
)

// Server serve HTTP requests for our services
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// NewServer creates new HTTP server and setup routing
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)

	server.router = router
	return server
}

// StartServer will run the server given the address string
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

// errorResponse take in an error and return a map of "error": error
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
