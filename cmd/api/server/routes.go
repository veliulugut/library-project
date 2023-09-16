package server

func (s *Server) linkRoutes() {

	v1Grp := s.router.Group("v1")
	{
		v1Grp.POST("/login", s.hnd.login.Login)
		v1Grp.POST("/user", s.hnd.user.CreateUser)
		v1Grp.GET("/user/:id", s.hnd.user.GetUserByID)
		v1Grp.DELETE("/user/:id", s.hnd.user.DeleteUser)
		v1Grp.POST("/user/:id", s.hnd.user.UpdateUser)
		v1Grp.POST("/book", s.hnd.book.CreateBook)
		v1Grp.DELETE("/book/:id", s.hnd.book.DeleteBook)
		v1Grp.PUT("/book/:id", s.hnd.book.UpdatedBook)
		v1Grp.GET("/book/:name", s.hnd.book.GetBookByName)
		v1Grp.GET("/book/get/:id", s.hnd.book.GetBookByID)
	}
}
