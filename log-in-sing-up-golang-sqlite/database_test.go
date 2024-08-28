package main

import (
	"database/sql"
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func setupTestDB(t *testing.T) *sql.DB {
	// Створюємо тимчасову базу даних для тестування
	dbFile := "./test.db"
	_ = os.Remove(dbFile) // Видаляємо базу даних, якщо вона вже існує

	db := initDatabase(dbFile)
	return db
}

func TestDatabase(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	// Хешування пароля
	password := "testpassword"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	assert.NoError(t, err, "Не вдалося хешувати пароль")

	// Виклик функції database
	database := func(db *sql.DB, email, hashedPassword string) {
		// Вставляємо дані
		statement, err := db.Prepare("INSERT INTO people (email, password) VALUES (?, ?)")
		if err != nil {
			t.Fatal(err)
		}
		defer statement.Close()

		_, err = statement.Exec(email, hashedPassword)
		if err != nil {
			t.Fatal(err)
		}
	}
	database(db, "test@example.com", string(hashedPassword))

	// Перевіряємо, чи були вставлені дані
	var email, passwordFromDB string
	row := db.QueryRow("SELECT email, password FROM people WHERE email = ?", "test@example.com")
	err = row.Scan(&email, &passwordFromDB)
	assert.NoError(t, err, "Не вдалося виконати запит для отримання даних")

	if assert.Equal(t, "test@example.com", email, "Email повинен бути 'test@example.com'") {
		t.Log("Email test passed")
	} else {
		t.Log("Email test failed")
	}

	// Перевірка пароля
	err = bcrypt.CompareHashAndPassword([]byte(passwordFromDB), []byte(password))
	if assert.NoError(t, err, "Пароль не співпадає") {
		t.Log("Password test passed")
	} else {
		t.Log("Password test failed")
	}
}
