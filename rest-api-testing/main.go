package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/thenguyenit/testing/rest-api-testing/models"
)

//App represent the App initial information
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

//Initialize to create DB connection and Router
func (app *App) Initialize(user, pass, dbname string) {
	connectionString := fmt.Sprintf("%s:%s@/%s", user, pass, dbname)

	var err error
	app.DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	app.Router = mux.NewRouter()
	app.InitializeRoutes()

}

func (app *App) InitializeRoutes() {
	app.Router.HandleFunc("/users", app.CreateUser).Methods(http.MethodPost)
	app.Router.HandleFunc("/users", app.GetUsers).Methods(http.MethodGet)
}

//Run to start the application
func (app *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, app.Router))
}

func (app *App) GetUsers(w http.ResponseWriter, r *http.Request) {

	users, err := models.GetUsers(app.DB, 0, 10)
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, models.Error{err.Error()})
	}
	respondWithJSON(w, http.StatusOK, users)
}

func (app *App) CreateUser(w http.ResponseWriter, r *http.Request) {
	var u models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&u); err != nil {
		respondWithJSON(w, http.StatusInternalServerError, models.Error{err.Error()})
	}

	err := u.CreateUser(app.DB)
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, models.Error{err.Error()})
	} else {
		respondWithJSON(w, http.StatusCreated, u)
	}

}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func main() {
	app := App{}
	app.Initialize("root", "123456", "go_rest_api_testing")
	app.Run("APP_ADDR")
}
