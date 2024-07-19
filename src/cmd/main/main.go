package main

import (
	"fmt"
	"log"

	"go-fiber-mongo-hrms/src/cmd/server"
	"go-fiber-mongo-hrms/src/db"
)

func main() {
	// DataBase - MongoDB Client
	client, db, err := db.GetClient()
	if err != nil {
		panic("Failed to Connect Database")
	} else {
		log.Println("DataBase Succecfully Connected")
	}

	// Server
	server := server.NewServer(fmt.Sprintf(":%s", "9010"), client, db)

	// RUN Server
	err = server.Run()
	if err != nil {
		log.Fatal(err)
	}
}
