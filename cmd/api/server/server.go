package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"library/cmd/api/handler/v1/book"
	"library/cmd/api/handler/v1/login"
	"library/cmd/api/handler/v1/reset_password"
	"library/cmd/api/handler/v1/user"
	"library/cmd/api/middlewares/auth"
	"library/ent"
	"library/pkg/utils"
	"os"
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
	mw       auth.Auth
}

type Handlers struct {
	user    user.Handler
	login   login.Handler
	book    book.Handler
	respass reset_password.Handler
}

func (s *Server) Init() error {
	s.router = gin.New()
	err := s.initDB()
	if err != nil {
		return err
	}
	err = s.initHandlers()
	if err != nil {
		return fmt.Errorf("init :%w", err)
	}
	s.initMiddlewares()
	s.linkRoutes()

	return nil
}

func (s *Server) Run() error {
	env, _ := os.LookupEnv("TEST_DATA")
	if env == "TRUE" || env == "true" {
		err := utils.CreateTestData(s.dbClient)
		if err != nil {
			return fmt.Errorf("server run / create test data :%w", err)
		}
	}

	err := s.router.Run(fmt.Sprintf(":%d", s.port))
	if err != nil {
		return fmt.Errorf("server run :%w", err)
	}

	return nil

}
