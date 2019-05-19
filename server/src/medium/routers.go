package routers

import (
	"CRUD_DEMO/server/src/controllers"
	"fmt"
	"github.com/gorilla/mux"
	// "io/ioutil"
	"net/http"
	"time"

	// "encoding/json"
	// "github.com/auth0/go-jwt-middleware"
	"github.com/codegangsta/negroni"
	"github.com/dgrijalva/jwt-go"
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

	// router.HandleFunc("/home", isAuthorized(homePage))

	router.HandleFunc("/api/students/list", controllers.ListStudents).Methods("GET")
	router.HandleFunc("/api/students/create/", controllers.InsertStudents).Methods("POST")
	router.HandleFunc("/api/students/delete/", controllers.DeleteStudents).Methods("DELETE")
	router.HandleFunc("/api/students/update/", controllers.UpdateStudents).Methods("PUT")
	router.HandleFunc("/api/students/find/", controllers.FindStudents).Methods("GET")
	router.HandleFunc("/api/students/", controllers.Search).Methods("GET")

	// router.HandleFunc("/login", controllers.Login).Methods("POST")
	// router.HandleFunc("/home", homePage).Methods("GET")

	// router.HandleFunc("/protected", controllers.ProtectedEndpoint).Methods("GET")
	// router.HandleFunc("/test", controllers.ValidateMiddleware(controllers.TestEndpoint)).Methods("GET")

	router.HandleFunc("/ping", PingHandler)
	router.Handle("/secured/", negroni.New(
		negroni.HandlerFunc(ValidateTokenMiddleware),
		negroni.Wrap(http.HandlerFunc(SecuredPingHandler)),
	))

	return router
}

var mySigningKey = []byte("captainjacksparrowsayshi")

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["client"] = "Elliot Forbes"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
	validToken, err := GenerateJWT()
	if err != nil {
		fmt.Println("Failed to generate token")
	}
	w.Write([]byte(validToken))
}

func SecuredPingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("All good. You only get this message if you're authenticated"))
}

func ValidateTokenMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	if r.Header["Token"] != nil {

		token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error")
			}
			return mySigningKey, nil
		})

		if err != nil {
			fmt.Fprintf(w, err.Error())
		}

		if token.Valid {
			next(w, r)
		}
	} else {

		fmt.Fprintf(w, "Not Authorized")
	}
	// })
}
