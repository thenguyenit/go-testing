## API Specification
* Create a new user in response to a valid POST request at /user,
* Update a user in response to a valid PUT request at /user/{id},
* Delete a user in response to a valid DELETE request at /user/{id},
* Fetch a user in response to a valid GET request at /user/{id}, and
* Fetch a list of users in response to a valid GET request at /users.

## Use govendor to manage dependency
```
govendor init
$ govendor fetch github.com/gorilla/mux
$ govendor fetch github.com/go-sql-driver/mysql

```

## Create a new database

```mysql
mysql -uroot -p
create database rest_api_testing;
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    age INT NOT NULL
);
```

## Write User Model with the basic code and return error 

```go
//GetUser to get a User
func (u *User) GetUser(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (u *User) UpdateUser(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (u *User) DeleteUser(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (u *User) CreateUser(db *sql.DB) error {
  	return errors.New("Not implemented")
}  

func GetUsers(db *sql.DB, start, count int) ([]User, error) {
  	return errors.New("Not implemented")
}
```

## Write main.go to initial the application

```go
//App represent the App initial information
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func main() {
	app := App{}
	app.Initialize("root", "123456", "go_rest_api_testing")
	app.Run("APP_ADDR")
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
}

func (app *App) CreateUser(w http.ResponseWriter, r *http.Request) {
}
```

## Write main_test.go

## Implement the request handler

## Implement User model
