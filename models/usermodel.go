package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID        int            `json:"id"`
	Email     string         `json:"email"`
	Name      sql.NullString `json:"name"`
	AvatarURL sql.NullString `json:"avatarURL"`
}
type DB struct {
	*sql.DB
}

// ConnectDB initializes a new in-memory SQLite database and returns a DB instance.
func ConnectDB() (*DB, error) {
	db, err := sql.Open("sqlite3", ":memory:") // In-memory database
	if err != nil {
		return nil, err
	}
	return &DB{db}, nil
}

// CreateTable creates the users table if it does not already exist.
func CreateTable(db *sql.DB) error {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY AUTOINCREMENT, email TEXT UNIQUE, name TEXT, avatarURL TEXT)")
	return err
}

// InsertUser inserts a new user into the database with the given email.
// If a user with the same email already exists, no insertion is performed.
func InsertUser(db *sql.DB, email string) error {
	// Check if user already exists
	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)", email).Scan(&exists)
	if err != nil {
		return fmt.Errorf("failed to check user existence: %v", err)
	}

	if exists {
		// User already exists, no need to insert
		log.Printf("User already exists - Email: %s", email)
		return nil
	}

	// Insert new user
	_, err = db.Exec("INSERT INTO users (email) VALUES (?)", email)
	if err != nil {
		return fmt.Errorf("failed to insert user: %v", err)
	}

	log.Printf("User inserted - Email: %s", email)
	return nil
}

// StoreUserProfile inserts or updates a user profile based on the given email, displayName, and avatarURL.
// If the user already exists, their profile information will be updated.
func StoreUserProfile(db *sql.DB, email string, displayName string, avatarURL string) error {
	// Prepare the upsert statement
	query := `
	 INSERT INTO users (email, name, avatarURL) 
	 VALUES (?, ?, ?)
	 ON CONFLICT(email) DO UPDATE SET
		 name = excluded.name,
		 avatarURL = excluded.avatarURL
 `

	// Execute the upsert
	result, err := db.Exec(query, email, displayName, avatarURL)
	if err != nil {
		return fmt.Errorf("failed to upsert user: %v", err)
	}

	// Check if a new row was inserted or an existing one was updated
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %v", err)
	}

	if rowsAffected == 0 {
		log.Printf("User information unchanged - Email: %s", email)
	} else {
		//log.Printf("User upserted - Email: %s, Name: %s, AvatarURL: %s", email, displayName, avatarURL)
	}

	return nil
}

// GetUserByEmail retrieves a user by their email address.
// Returns a User object and an error if the user is not found or if there is a database error.
func GetUserByEmail(db *sql.DB, email string) (*User, error) {
	var user User
	err := db.QueryRow("SELECT * FROM users WHERE email = ?", email).Scan(&user.ID, &user.Email, &user.Name, &user.AvatarURL)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// DeleteUserByEmail removes a user from the database based on their email address.
func DeleteUserByEmail(db *sql.DB, email string) error {
	_, err := db.Exec("DELETE FROM users WHERE email = ?", email)
	return err
}

// ListUsers retrieves all users from the database and returns them as a slice of User objects.
func ListUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query("SELECT id, email, name, avatarURL FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		err := rows.Scan(&u.ID, &u.Email, &u.Name, &u.AvatarURL)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}
