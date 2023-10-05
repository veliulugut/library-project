package server

func (s *Server) linkRoutes() {

	v1Grp := s.router.Group("v1")
	v1Grp.POST("/login", s.hnd.login.Login)
	v1Grp.POST("/auth/reset-password", s.hnd.respass.SendResetPasswordValidation)
	v1Grp.POST("/auth/validate-reset-password", s.hnd.respass.Validate)
	v1Grp.POST("/user", s.hnd.user.CreateUser)

	authGrp := v1Grp.Group("")
	authGrp.Use(s.mw.Check())
	authGrp.GET("/user/:id", s.hnd.user.GetUserByID)
	authGrp.DELETE("/user/:id", s.hnd.user.DeleteUser)
	authGrp.POST("/user/:id", s.hnd.user.UpdateUser)
	authGrp.POST("/book", s.hnd.book.CreateBook)
	authGrp.DELETE("/book/:id", s.hnd.book.DeleteBook)
	authGrp.PUT("/book/:id", s.hnd.book.UpdatedBook)
	authGrp.GET("/book/:name", s.hnd.book.GetBookByName)
	authGrp.GET("/book/get/:id", s.hnd.book.GetBookByID)
	authGrp.GET("/books", s.hnd.book.ListBook)

}
