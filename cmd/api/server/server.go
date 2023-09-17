package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"library/cmd/api/handler/v1/book"
	"library/cmd/api/handler/v1/login"
	"library/cmd/api/handler/v1/user"
	"library/ent"
)

func New(port int) *Server {
	return &Server{
		port: port,
	}
}

type Server struct {
	router   *gin.Engine
	dbClient *ent.Client
	port     int
	hnd      Handlers
}

type Handlers struct {
	user  user.Handler
	login login.Handler
	book  book.Handler
}

func (s *Server) Init() error {
	s.router = gin.New()
	err := s.initDB()
	if err != nil {
		return err
	}
	s.initHandlers()
	s.initMiddlewares()
	s.linkRoutes()

	return nil
}

func (s *Server) Run() error {
	err := s.router.Run(fmt.Sprintf(":%d", s.port))
	if err != nil {
		return fmt.Errorf("server run :%w", err)
	}

	return nil

}