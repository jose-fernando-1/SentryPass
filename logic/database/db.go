package database

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3" // Importing SQLite driver
)

// Initializing SQLite database
func InitDB() *sql.DB {
	// Creating the database file if it doesn't exist
	db, err := sql.Open("sqlite3", "sentrypass.db")
	if err != nil {
		log.Fatal("Failed to open the database: ", err)
	}

	// Verify the database connection
	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}

	// Enable foreign key constraints
	_, err = db.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		log.Fatal("Failed to enable foreign keys: ", err)
	}

	log.Println("Database connected and foreign keys enabled.")
	return db
}

func RunMigrations(db *sql.DB) {
	files, err := os.ReadDir("migrations")
	if err != nil {
		log.Fatal("Failed to read migrations directory: ", err)
	}

	for _, file := range files {
		script, err := os.ReadFile(filepath.Join("migrations", file.Name()))
		if err != nil {
			log.Fatalf("Failed to read file %s: %v", file.Name(), err)
		}

		_, err = db.Exec(string(script))
		if err != nil {
			log.Fatalf("Failed to execute migration %s: %v", file.Name(), err)
		}

		log.Printf("Migration %s applied successfully.", file.Name())
	}
}
