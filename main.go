package main

import (
	"fmt"

	"github.com/DANCANKARANI/tyson/database"
	"github.com/DANCANKARANI/tyson/endpoints"
	"github.com/DANCANKARANI/tyson/model"
)

func main() {
	fmt.Println("hello world!")
	database.ConnectDB()
	model.MigrateDB()
	endpoints.CreateEndpoint()

}