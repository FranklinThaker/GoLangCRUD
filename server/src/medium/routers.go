package routers

import (
	"CRUD_DEMO/server/src/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

// registers all routes for the application.

func Route() *mux.Router {
	router := mux.NewRouter()

	router.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		w.WriteHeader(http.StatusNoContent)
		return
	})

	router.StrictSlash(true)

	router.HandleFunc("/api/students/list", controllers.ListStudents).Methods("GET")
	router.HandleFunc("/api/students/create/", controllers.InsertStudents).Methods("POST")
	router.HandleFunc("/api/students/delete/", controllers.DeleteStudents).Methods("DELETE")
	router.HandleFunc("/api/students/update/", controllers.UpdateStudents).Methods("PUT")
	router.HandleFunc("/api/students/find/", controllers.FindStudents).Methods("GET")
	router.HandleFunc("/api/students/", controllers.Search).Methods("GET")

	return router
}
