package main

import (
	"log"

	"back/app"
)

func main() {
	router := app.SetRouter()
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
