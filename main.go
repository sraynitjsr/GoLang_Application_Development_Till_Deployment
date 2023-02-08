package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Employee struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

var employees []Employee = []Employee{
	{
		Name: "A",
		Id:   10,
	},
	{
		Name: "B",
		Id:   20,
	},
}

func homePage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the Employee API"))
}

func getEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	id, _ := strconv.ParseInt(vars["id"], 10, 64)
	emp := Employee{
		Name: name,
		Id:   int(id),
	}

	flag := false

	for _, e := range employees {
		if e.Name == emp.Name && e.Id == emp.Id {
			flag = true
			break
		}
	}

	if flag {
		json.NewEncoder(w).Encode(emp)
	} else {
		w.Write([]byte("No Such Employee Exists"))
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", homePage)
	router.HandleFunc("/employee/{name}/{id}", getEmployee)
	http.ListenAndServe(":8080", router)
}
