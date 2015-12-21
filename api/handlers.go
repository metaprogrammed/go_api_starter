package api

import (
	"encoding/json"
	"fmt"
	"go-rest-json-boilerplate/api/utils"
	"go-rest-json-boilerplate/resources"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func UserIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resources.AllUsers); err != nil {
		log.Println(err)
	}
}

func UserShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var userId int
	var err error
	if userId, err = strconv.Atoi(vars["userId"]); err != nil {
		log.Println(err)
	}
	user := resources.FindUser(userId)
	if user.Id > 0 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(user); err != nil {
			log.Println(err)
		}
		return
	}
	// If not found, return a 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(utils.JsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		log.Println(err)
	}

}

/*
Curl command for testing:
curl -H "Content-Type: application/json" -d '{"username":"user3", "name":"David", "email":"david@example.com"}' http://localhost:8080/users
*/
func UserCreate(w http.ResponseWriter, r *http.Request) {
	var user resources.User
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Println(err)
	}
	if err := r.Body.Close(); err != nil {
		log.Println(err)
	}
	if err := json.Unmarshal(body, &user); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Println(err)
		}
	}
	new_user := resources.CreateUser(user)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(new_user); err != nil {
		log.Println(err)
	}
}
