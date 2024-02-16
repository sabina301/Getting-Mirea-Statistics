package main

import (
	"getting-statistics-mirea/handler"
	"getting-statistics-mirea/router"
	"getting-statistics-mirea/service"
)

func main() {

	s := service.NewService()
	h := handler.NewHandler(s)
	r := router.InitRouter(h)

	router.Start(r, "0.0.0.0:8080")

}
