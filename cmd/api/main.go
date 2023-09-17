package main

import (
	"fmt"
	"library/cmd/api/server"
	"library/ent"
	"library/pkg/utils"
	"log"
)

// @title          Library API
// @version         1.0
// @description     library
// @contact.name   Veli Ulugut
// @contact.email  veliulugut1@gmail.com
// @host      localhost:8080
// @BasePath  /v1
// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	s := server.New(8080)
	err := s.Init()
	if err != nil {
		log.Fatalln(err)
	}

	err = s.Run()
	if err != nil {
		log.Fatalln(err)

	}

	var db *ent.Client
	data := utils.NewBookData(db)

	fmt.Println(data)
}
