package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// variable for database
var DB *gorm.DB
var err error

const DNS = "root:MySQLRoot@tcp(localhost:3306)/goapi?charset=utf8mb4&parseTime=True&loc=Local"

type User struct {
	gorm.Model
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}

func InitialMigration() {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}
	DB.AutoMigrate(&User{})
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	// get the list of users available
	w.Header().Set("Content-Type", "application/json")
	var users []User
	DB.Find(&users)
	// once get all data, encode it back to send client
	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	// get a particular user
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user User
	// want to get id from the parameters
	DB.First(&user, params["id"])
	json.NewEncoder(w).Encode(user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	// use json module & decode data getting from the request body
	json.NewDecoder(r.Body).Decode(&user)
	// save data
	DB.Create(&user)
	// pass data back to the browser
	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user User
	DB.First(&user, params["id"])
	// after gettting data; 3 values from params, update with new values
	// decode information
	json.NewDecoder(r.Body).Decode(&user)
	// update values of the same user
	DB.Save(&user)
	// pass data back to the browser
	json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user User
	// should delete user reference & params whatever we're passing
	DB.Delete(&user, params["id"])
	json.NewEncoder(w).Encode("The User is Deleted successfully!")
}
