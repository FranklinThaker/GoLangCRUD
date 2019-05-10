package controllers

import (
	"CRUD_DEMO/server/src/db"
	"CRUD_DEMO/server/src/models"
	"encoding/json"
	"fmt"
	// "golang.org/x/text/search"
	"log"
	"net/http"
	// "strconv"
	jwt "github.com/dgrijalva/jwt-go"
	// "io/ioutil"
	"time"
)

var mySigningKey = []byte("captainjacksparrowsayshi")

func SetupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

//Listing students...
func ListStudents(res http.ResponseWriter, req *http.Request) {
	SetupResponse(&res, req)
	db := db.DbConn()
	Stud := []models.Student{}

	db.Find(&Stud)
	fmt.Println(Stud)

	json.NewEncoder(res).Encode(Stud)
	log.Println("Data fetched")
	defer db.Close()
}

func InsertStudents(res http.ResponseWriter, req *http.Request) {
	SetupResponse(&res, req)
	db := db.DbConn()
	var stud models.Student
	req.ParseForm()

	if err := json.NewDecoder(req.Body).Decode(&stud); err != nil {
		fmt.Println("err", err)
	}

	Stud := models.Student{Fname: stud.Fname, Lname: stud.Lname, Email: stud.Email, Age: stud.Age, Mobile: stud.Mobile}
	// db.NewRecord(&Stud)
	db.Create(&Stud)

	log.Println("Data Inserted", stud)
	json.NewEncoder(res).Encode(Stud)
	defer db.Close()

}

func DeleteStudents(res http.ResponseWriter, req *http.Request) {
	SetupResponse(&res, req)
	db := db.DbConn()
	id := req.URL.Query().Get("id")

	Stud := models.Student{}

	db.Unscoped().Delete(&Stud, id)

	json.NewEncoder(res).Encode(Stud)
	log.Println("Data Deleted")
	defer db.Close()
}

func UpdateStudents(res http.ResponseWriter, req *http.Request) {
	SetupResponse(&res, req)
	db := db.DbConn()
	id := req.URL.Query().Get("id")
	var stud models.Student
	req.ParseForm()

	if err := json.NewDecoder(req.Body).Decode(&stud); err != nil {
		fmt.Println("err", err)
	}

	// Stud := models.Student{Fname: stud.Fname, Lname: stud.Lname, Email: stud.Email, Age: stud.Age, Mobile: stud.Mobile}
	var Stud = models.Student{}
	// db.Model(&Stud).Updates(models.Student{Fname: stud.Fname, Lname: stud.Lname, Email: stud.Email, Age: stud.Age, Mobile: stud.Mobile})
	db.Model(&Stud).Where("ID=?", id).Updates(models.Student{Fname: stud.Fname, Lname: stud.Lname, Email: stud.Email, Age: stud.Age, Mobile: stud.Mobile})
	json.NewEncoder(res).Encode(Stud)
	log.Println("Data Deleted")
	defer db.Close()
}

func FindStudents(res http.ResponseWriter, req *http.Request) {
	SetupResponse(&res, req)
	db := db.DbConn()
	id := req.URL.Query().Get("id")
	var stud models.Student
	req.ParseForm()

	if err := json.NewDecoder(req.Body).Decode(&stud); err != nil {
		fmt.Println("err", err)
	}
	output := models.Student{}
	Stud := db.Where("ID = ?", id).Find(&output)
	json.NewEncoder(res).Encode(Stud.Value)
	log.Println("Data OF Single Student")
	defer db.Close()
}

func Search(res http.ResponseWriter, req *http.Request) {
	SetupResponse(&res, req)
	db := db.DbConn()
	search := req.URL.Query().Get("search")
	req.ParseForm()

	output := []models.Student{}
	Stud := db.Where("Fname LIKE ?", "%"+search+"%").Find(&output)
	json.NewEncoder(res).Encode(Stud.Value)
	log.Println("Data OF Single Student")
	defer db.Close()
}

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

func Login(res http.ResponseWriter, req *http.Request) {
	validToken, err := GenerateJWT()
	fmt.Println("1")
	if err != nil {
		fmt.Println("Failed to generate token")
	}
	fmt.Println("2", validToken)
	req.Header.Set("Token", validToken)
}
