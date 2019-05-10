package main

import (
	"net/http"
	// "fmt"
	"CRUD_DEMO/server/src/medium"
	"CRUD_DEMO/server/src/migrations"
	// "github.com/urfave/negroni"
	"log"
	"os"
)

func main() {
	migrations.Migrate()
	port := os.Args[1]
	log.Println("Server running on port : " + port)
	http.ListenAndServe(":8080", routers.Route())
}
