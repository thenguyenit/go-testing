package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thenguyenit/testing/rest-api-testing/models"
)

var app App

func TestMain(m *testing.M) {
	fmt.Println("Start TestMain")
	app = App{}
	app.Initialize("root", "123456", "go_rest_api_testing")

	ensureTableExists()

	code := m.Run()

	clearTestData()

	fmt.Println("Exit TestMain")
	os.Exit(code)
}

func TestCreateUser(t *testing.T) {

	userFaked := map[string]interface{}{
		"name": "NTN",
		"age":  30.0,
	}
	u, _ := json.Marshal(userFaked)

	req, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(u))
	response := executeRequest(req)

	assert.Equal(t, http.StatusCreated, response.Code)
	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)
	assert.Equal(t, userFaked["name"], m["name"])
	assert.Equal(t, userFaked["age"], m["age"])
	assert.Equal(t, 1.0, m["id"])

	clearTable()
}

func TestEmptyTable(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest(http.MethodGet, "/users", nil)
	response := executeRequest(req)

	assert.Equal(t, http.StatusOK, response.Code)
	body := response.Body.String()

	assert.Equal(t, "[]", body)
}

func TestGetNonExistenUser(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest(http.MethodGet, "/user/15", nil)
	response := executeRequest(req)

	assert.Equal(t, http.StatusNotFound, response.Code)
	body := response.Body.String()

	assert.Equal(t, "404 page not found\n", body)
}

func TestGetUsers(t *testing.T) {
	//Insert some faked users
	statement := `INSERT INTO users(name, age)
	VALUES ('Johnny Nguyen', 39), ('Herry Ford', 60)
	`
	_, err := app.DB.Exec(statement)
	if err != nil {
		log.Fatal(err)
	}

	//Test query
	req, _ := http.NewRequest(http.MethodGet, "/users", nil)
	res := executeRequest(req)

	assert.Equal(t, http.StatusOK, res.Code)

	userList := []models.User{}
	err = json.Unmarshal(res.Body.Bytes(), &userList)
	assert.NoError(t, err)

	for _, u := range userList {
		assert.IsType(t, models.User{}, u)
	}

	clearTable()
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	app.Router.ServeHTTP(rr, req)

	return rr
}

func clearTable() {
	app.DB.Exec("DELETE FROM users")
	app.DB.Exec("ALTER TABLE users AUTO_INCREMENT = 1")
}

const tableCreationQuery = `
CREATE TABLE IF NOT EXISTS users
(
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    age INT NOT NULL
)`

func ensureTableExists() {
	if _, err := app.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTestData() {

}
