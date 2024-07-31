package main

import (
	"log"
	"net/http"
	"todolist/database"
	"todolist/routers"
)

func main() {
	database.InitDatabase()

	r := routers.InitializeRouter()
	
	log.Fatal(http.ListenAndServe(":8080", r))
}
