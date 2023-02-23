package api

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	db "goskill/db/sqlc"
	docs "goskill/docs"
	"log"
)

type Server struct {
	db     *db.Store
	router *gin.Engine
}

type ServerError struct {
	error error
}

func NewServerError(err error) ServerError {
	return ServerError{err}
}

func NewServer(store *db.Store) *Server {
	server := Server{db: store}
	router := gin.Default()

	// REGISTER THE ROUTES

	docs.SwaggerInfo.BasePath = "/"
	v1 := router.Group("/api/v1")
	{

		v1.POST("/createSkill", server.CreateSkillHandler)

	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	server.router = router
	return &server
}

func (s *Server) Start(address string) {
	log.Fatal(s.router.Run(address))
}

func errorMessage(e error) gin.H {
	return gin.H{"error": e.Error()}
}
