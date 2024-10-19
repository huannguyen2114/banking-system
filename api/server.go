package api

import (
	"fmt"

	"github.com/huannguyen2114/go-toy-project/token"
	"github.com/huannguyen2114/go-toy-project/util"

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
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

// NewServer creates new HTTP server and setup routing
func NewServer(store db.Store, config util.Config) (*Server, error) {
	token, err := token.NewPasetoMaker(config.TokenSymetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: token,
	}
	router := gin.Default()

	docs.SwaggerInfo.BasePath = "/api/v1"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	authRoutes.POST("/accounts", server.createAccount)
	authRoutes.GET("/accounts/:id", server.getAccount)
	authRoutes.GET("/accounts", server.listAccount)

	authRoutes.POST("/transfers", server.createTransfer)

	server.router = router
}

// StartServer will run the server given the address string
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

// errorResponse take in an error and return a map of "error": error
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
