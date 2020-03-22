package main

import (
	"log"
	"servicemanager/handler"
)

func main() {
	localServer := ":8080"

	engine := handler.InitEngine()

	if err := engine.Run(localServer); err != nil {
		log.Fatal(err.Error())
	}

}
