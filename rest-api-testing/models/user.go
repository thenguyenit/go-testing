package models

import (
	"database/sql"
	"errors"
	"fmt"
)

//User represent a User Model
type User struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

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
	statement := fmt.Sprintf("INSERT INTO users(name, age) VALUES('%s', %d)", u.Name, u.Age)
	_, err := db.Exec(statement)
	if err != nil {
		return err
	}
	err = db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&u.ID)

	if err != nil {
		return err
	}
	return nil
}

func GetUsers(db *sql.DB, start, count int) ([]User, error) {
	statement := fmt.Sprintf("SELECT id, name, age FROM users LIMIT %d OFFSET %d", count, start)
	rows, err := db.Query(statement)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	users := []User{}

	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Age); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}
