package api

import (
	db "github.com/Go-Docker-postgreSQL-project-test/simplebank/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)

	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)//没有实现



	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}