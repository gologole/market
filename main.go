package main

import (
	"log"
	"main.go/pkg/myhandler"
	"main.go/pkg/repository"
	"main.go/pkg/service"
	"main.go/server"
)

func main() {
	//initconfigs()
	port := "8080"

	db := new(repository.SQLiteDB)
	db.ConnectDB()
	service := service.NewAuthservice(db)
	handler := new(myhandler.MyHandler)
	handler.Service = service

	server := new(server.Server)
	if err := server.RunServer(port, handler.InitRouts()); err != nil {
		log.Fatal("Server start error: ", err)
	}

}
