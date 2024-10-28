package initializers

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	var err error
	dsn := "host=localhost user=root dbname=taskmgmtdb password=secret port=5432 sslmode=disable"
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	CreateTable()

	if err := db.Ping(); err != nil {
		log.Fatal("fail to ping the anotherdb:", err)
	}
	fmt.Println("successfully connected to the database")
	// // Alter the users table to add email and address columns
	// err = alterUsersTable()
	// if err != nil {
	// 	log.Fatalf("Failed to alter users table: %v", err)
	// }

	// fmt.Println("Database connection successful and users table altered!")
}

func CreateTable() {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		email VARCHAR(320) UNIQUE NOT NULL,
		address TEXT,
		password VARCHAR(255) NOT NULL
	);`)
	if err != nil {
		log.Fatalf("Error creating users table: %v", err)
	} else {
		fmt.Println("User table created successfully")
	}
}

// alterUsersTable alters the users table by adding email and address columns if they don't exist
// func alterUsersTable() error {
// 	// Step 1: Add the columns as nullable
// 	query := `
//     ALTER TABLE users
//     ADD COLUMN IF NOT EXISTS email VARCHAR(255),
//     ADD COLUMN IF NOT EXISTS address VARCHAR(255),
// 	ADD COLUMN IF NOT EXISTS password VARCHAR(255);
//     `
// 	_, err := db.Exec(query)
// 	if err != nil {
// 		return fmt.Errorf("error executing alter query: %w", err)
// 	}

// 	// Step 2: Check existing records and ensure no NULL values before applying NOT NULL constraints
// 	checkQuery := `
//     SELECT COUNT(*) FROM users WHERE email IS NULL OR address IS NULL;
//     `
// 	var count int
// 	err = db.QueryRow(checkQuery).Scan(&count)
// 	if err != nil {
// 		return fmt.Errorf("error checking NULL values: %w", err)
// 	}

// 	if count > 0 {
// 		return fmt.Errorf("cannot set NOT NULL constraints because there are existing NULL values in the users table")
// 	}

// 	// Step 3: Set the columns to NOT NULL
// 	alterQuery := `
//     ALTER TABLE users
//     ALTER COLUMN email SET NOT NULL,
//     ALTER COLUMN address SET NOT NULL,
// 	ALTER COLUMN address SET NOT NULL;
//     `
// 	_, err = db.Exec(alterQuery)
// 	if err != nil {
// 		return fmt.Errorf("error setting not null constraint: %w", err)
// 	}

// 	log.Println("Users table altered successfully: email and address and password columns added and updated.")
// 	return nil
// }

// GetDB returns the database connection
func GetDB() *sql.DB {
	return db
}
