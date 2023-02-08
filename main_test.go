package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestHomePage(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/", homePage)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "Welcome to the Employee API", rr.Body.String())
}

func TestGetEmployee(t *testing.T) {
	req, _ := http.NewRequest("GET", "/employee/A/10", nil)
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/employee/{name}/{id}", getEmployee)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var emp Employee
	err := json.Unmarshal(rr.Body.Bytes(), &emp)
	assert.Nil(t, err)

	assert.Equal(t, "A", emp.Name)
	assert.Equal(t, 10, emp.Id)
}

func TestGetNonExistentEmployee(t *testing.T) {
	req, _ := http.NewRequest("GET", "/employee/C/30", nil)
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/employee/{name}/{id}", getEmployee)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "No Such Employee Exists", rr.Body.String())
}
