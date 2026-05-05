package initializers

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	var err error
	if err = godotenv.Load(); err != nil {
		log.Println("No .env file is exist there in our code!")

	}

	// dsn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%s sslmode=%s",

	// 	os.Getenv("localhost"),
	// 	os.Getenv("user"),
	// 	os.Getenv("dbname"),
	// 	os.Getenv("password"),
	// 	os.Getenv("port"),
	// 	os.Getenv("sslmode"),
	// )

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)

	db, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	CreateTable()

	if err := db.Ping(); err != nil {
		log.Fatal("fail to ping the anotherdb:", err)
	}
	fmt.Println("successfully connected to the database")

}
func CreateTable() {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS tasks (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		assign_to VARCHAR(255)[]
	);`)
	if err != nil {
		log.Fatalf("Error creating users table: %v", err)
	} else {
		fmt.Println("Tasks table created successfully")
	}
}

func GetDB() *sql.DB {
	return db
}
