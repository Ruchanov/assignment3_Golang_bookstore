package main

import (
	"fmt"
	"github.com/Ruchanov/assignment3_Golang_bookstore/routes"
	"github.com/Ruchanov/assignment3_Golang_bookstore/utils"
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
