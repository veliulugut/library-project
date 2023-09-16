package server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"library/cmd/api/handler/v1/book"
	loginhnd "library/cmd/api/handler/v1/login"
	userhnd "library/cmd/api/handler/v1/user"
	_ "library/docs"
	"library/ent"
	golangjwt "library/pkg/jwt/golang_jwt"
	"library/pkg/passwd/bcrypt"
	"library/pkg/repository/entadp"
	booksrv "library/service/book"
	loginsrv "library/service/login"
	usersrv "library/service/user"
)

func (s *Server) initDB() error {
	var err error
	s.dbClient, err = ent.Open("mysql", "root:123turkTR562@tcp(localhost:3306)/librarydb?parseTime=true")
	if err != nil {
		return fmt.Errorf("init db :%w", err)
	}

	if err = s.dbClient.Schema.Create(context.Background()); err != nil {
		return fmt.Errorf("init db :%w", err)
	}

	return nil
}

/*func (s *Server) initCache() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
		Protocol: 3,  // specify 2 for RESP 2 or 3 for RESP 3
	})

}*/

func (s *Server) initMiddlewares() {
	s.router.Use(gin.Logger())
	s.router.Use(gin.Recovery())
	s.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

}

func (s *Server) initHandlers() {
	secret := "verysecretkey"
	repo := entadp.NewRepository(s.dbClient)

	//service
	bc := bcrypt.NewBcrypt(secret, 10)
	j := golangjwt.New(secret, 24)
	userService := usersrv.New(bc, repo)
	loginService := loginsrv.New(repo, j, bc)
	bookService := booksrv.New(repo)

	//handlers
	s.hnd.user = userhnd.NewUser(userService)
	s.hnd.login = loginhnd.New(loginService)
	s.hnd.book = book.New(bookService)

}
