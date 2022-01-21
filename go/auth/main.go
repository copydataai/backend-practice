package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type UserLogin struct {
	Email    string
	Password string
}

type UserCreate struct {
	Name       string
	Country    string
	Speciality string
	Role       string
	Email      string
	Password   string
}

func insertUser(db sql.DB, newUser *UserCreate) string {
	db.QueryRow("INSERT INTO users(name, country, speciality, role, created_at, email, password) VALUES($1, $2, $3, $4, $5, $6)", newUser.Name, newUser.Country, newUser.Speciality, newUser.Role, time.Now().String(), newUser.Email, newUser.Password)
	return "User Created"
}

func verifyUser(db sql.DB, user *UserLogin) string {
	var userVerify UserLogin
	userQuery := db.QueryRow("SELECT password from users WHERE email = $1", user.Email)
	userQuery.Scan(&userVerify.Email, &userVerify.Password)
	return "is user"
}

func main() {
	configDB := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s", user, password, dbname, host, port)
	db, err := sql.Open("postgres", configDB)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	} else {
		fmt.Println("DB Connected...")
	}
}
