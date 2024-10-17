package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/huannguyen2114/go-toy-project/db/sqlc"
	"github.com/huannguyen2114/go-toy-project/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Server serve HTTP requests for our services
type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer creates new HTTP server and setup routing
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	docs.SwaggerInfo.BasePath = "/api/v1"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)
	router.POST("/transfers", server.createTransfer)

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
