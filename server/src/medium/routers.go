package routers

import (
	"CRUD_DEMO/server/src/controllers"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// registers all routes for the application.
var mySigningKey = []byte("captainjacksparrowsayshi")

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] != nil {

			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return mySigningKey, nil
			})

			if err != nil {
				fmt.Fprintf(w, "TOken not valid")
				fmt.Fprintf(w, err.Error())
			}

			if token.Valid {
				endpoint(w, r)
			}
		} else {

			fmt.Fprintf(w, "Not Authorized")
		}
	})
}

func Route() *mux.Router {
	router := mux.NewRouter()

	router.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		w.WriteHeader(http.StatusNoContent)
		return
	})

	router.Methods("GET").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		w.WriteHeader(http.StatusNoContent)
		return
	})

	router.StrictSlash(true)

	// http.Handle("/home", isAuthorized(homePage))

	// router.HandleFunc("/api/students/list", controllers.ListStudents)).Methods("GET")
	// router.HandleFunc("/api/students/create/", controllers.InsertStudents).Methods("POST")
	// router.HandleFunc("/api/students/delete/", controllers.DeleteStudents).Methods("DELETE")
	// router.HandleFunc("/api/students/update/", controllers.UpdateStudents).Methods("PUT")
	// router.HandleFunc("/api/students/find/", controllers.FindStudents).Methods("GET")
	// router.HandleFunc("/api/students/", controllers.Search).Methods("GET")

	router.HandleFunc("/login", controllers.Login).Methods("POST")
	// router.HandleFunc("/home", isAuthorized(homePage)).Methods("GET")

	// router.HandleFunc("/protected", controllers.ProtectedEndpoint).Methods("GET")
	// router.HandleFunc("/test", controllers.ValidateMiddleware(controllers.TestEndpoint)).Methods("GET")

	return router
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
	fmt.Println("Endpoint Hit: homePage")
	log.Println("LOgged In")
}
