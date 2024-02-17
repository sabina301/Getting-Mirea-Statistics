package main

import (
	"getting-statistics-mirea/server/handler"
	"getting-statistics-mirea/server/router"
	service "getting-statistics-mirea/server/service"
	"log"
)

func main() {

	s := service.NewService()
	h := handler.NewHandler(s)
	r := router.InitRouter(h)

	if err := router.Start(r, "0.0.0.0:8080"); err != nil {
		log.Fatal("Cant start")
	}

}
