package main

import (
	"fmt"
	"github.com/Ruchanov/Golang_2023/assignment3/routes"
	"github.com/Ruchanov/Golang_2023/assignment3/utils"
	"log"
	"net/http"
)

func main() {
	db, err := utils.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	contr := routes.NewController(db)
	routes.Start(contr)
	fmt.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
