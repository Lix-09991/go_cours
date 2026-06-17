package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./users.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL
		);
	`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`DELETE FROM users`)
	if err != nil {
		log.Fatal(err)
	}

	email := "test@example.com"
	password := "qwerty123"

	_, err = db.Exec(`INSERT INTO users (email, password) VALUES (?, ?)`, email, password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(auth(db, "no-user@example.com", "123"))
	fmt.Println(auth(db, email, "wrongpass"))
	fmt.Println(auth(db, email, password))
}

func auth(db *sql.DB, login, pass string) string {
	var savedPass string
	err := db.QueryRow(`SELECT password FROM users WHERE email = ?`, login).Scan(&savedPass)
	if err != nil {
		if err == sql.ErrNoRows {
			return "пользователь не найден"
		}
		return "ошибка БД"
	}

	if savedPass != pass {
		return "пароль неверный"
	}

	return "пароль верный"
}
