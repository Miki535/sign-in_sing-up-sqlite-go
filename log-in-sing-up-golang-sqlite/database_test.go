package main

import (
	"database/sql"
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func setupTestDB(t *testing.T) *sql.DB {
	// Створюємо тимчасову базу даних для тестування
	dbFile := "./test.db"
	_ = os.Remove(dbFile) // Видаляємо базу даних, якщо вона вже існує

	// Відкриваємо базу даних
	db, err := sql.Open("sqlite3", dbFile)
	assert.NoError(t, err, "Не вдалося відкрити базу даних")

	// Create table 'people'
	_, err = db.Exec(`CREATE TABLE people (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL,
		password TEXT NOT NULL
	)`)
	assert.NoError(t, err, "Не вдалося створити таблицю 'people'")

	return db
}

func TestDatabase(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()
	// Вказуємо назву тестової бази даних
	dbName := "./test.db"

	database(dbName, "test@example.com", "testpassword")
	// Перевіряємо, чи були вставлені дані
	var email, password string
	row := db.QueryRow("SELECT email, password FROM people WHERE email = ?", "test@example.com")
	err := row.Scan(&email, &password)
	assert.NoError(t, err, "Не вдалося виконати запит для отримання даних")

	if assert.Equal(t, "test@example.com", email, "Email should be 'test@example.com'") {
		t.Log("Email test passed")
	} else {
		t.Log("Email test failed")
	}

	// Перевірка password
	if assert.Equal(t, "testpassword", password, "Password should be 'testpassword'") {
		t.Log("Password test passed")
	} else {
		t.Log("Password test failed")
	}

}
