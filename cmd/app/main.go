package main

import (
	"log"
	"newProject/internal/delivery/http"
	"newProject/internal/service"
	"os"
)

func main() {
	app := http.NewApp(service.New("https://7103.api.greenapi.com"))
	port := os.Getenv("PORT")
	if port == "" {
		port = ":4000"
	}
	err := app.Route().Run(port)
	if err != nil {
		log.Fatal(err)
		return
	}
}
