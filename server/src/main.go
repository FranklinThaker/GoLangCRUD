package main

import (
	"CRUD_DEMO/server/src/medium"
	"CRUD_DEMO/server/src/migrations"
	// "fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	migrations.Migrate()
	port := os.Args[1]

	log.Println("Server running on port : " + port)
	http.ListenAndServe(":8080", routers.Route())
}
