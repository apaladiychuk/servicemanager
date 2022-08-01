package main

import (
	"log"

	"servicemanager/handler"
	"servicemanager/storage"
)

func main() {
	localServer := ":8081"

	engine := handler.InitEngine()
	db := storage.NewLocal()
	serviceHandler := handler.NewService(db)
	serviceHandler.Mount(engine)

	if err := engine.Run(localServer); err != nil {
		log.Fatal(err.Error())
	}

}
