package main

import (
	"log"
	"newProject/internal/delivery/http"
	"newProject/internal/service"
)

func main() {
	app := http.NewApp(service.New("https://7103.api.greenapi.com"))
	err := app.Route().Run(":4000")
	if err != nil {
		log.Fatal(err)
		return
	}
}
