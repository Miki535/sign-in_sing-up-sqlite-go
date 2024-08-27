package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func database(dbName, email, password string) {
	// Відкриваємо базу даних
	database, err := sql.Open("sqlite3", dbName)
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	// Створюємо таблицю, якщо вона не існує
	statement, err := database.Prepare("CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, email TEXT, password TEXT)")
	if err != nil {
		log.Fatal(err)
	}
	defer statement.Close()

	_, err = statement.Exec()
	if err != nil {
		log.Fatal(err)
	}

	// Вставляємо дані
	statement, err = database.Prepare("INSERT INTO people (email, password) VALUES (? ,?)")
	if err != nil {
		log.Fatal(err)
	}
	defer statement.Close()

	_, err = statement.Exec(email, password)
	if err != nil {
		log.Fatal(err)
	}
}
