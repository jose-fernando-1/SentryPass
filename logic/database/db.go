package database

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	"sentrypass/logic/utils"

	_ "github.com/mattn/go-sqlite3" // Importing SQLite driver
)

// InitDB initializes the SQLite database at the specified path.
func InitDB(dbPath string) *sql.DB {
	// Ensure the directory exists
	err := os.MkdirAll(filepath.Dir(dbPath), os.ModePerm)
	if err != nil {
		log.Fatalf("Failed to create database directory: %v", err)
	}

	log.Printf("Database path: %s", dbPath)

	// Open the database file
	db, err := sql.Open("sqlite3", dbPath)
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

// RunSchemaMigration runs the schema initialization script.
func RunSchemaMigration(db *sql.DB) {
	scriptPath := filepath.Join("migrations", "01_schema.sql")
	script, err := os.ReadFile(scriptPath)
	if err != nil {
		log.Fatalf("Failed to read schema migration file: %v", err)
	}

	_, err = db.Exec(string(script))
	if err != nil {
		log.Fatalf("Failed to execute schema migration: %v", err)
	}

	log.Println("Schema migration applied successfully.")
}

// PopulateInitialData populates the database with initial data.
func PopulateInitialData(db *sql.DB, masterPasswordHash, masterPasswordSalt string) {
	scriptPath := filepath.Join("migrations", "02_migration.sql")
	script, err := os.ReadFile(scriptPath)
	if err != nil {
		log.Fatalf("Failed to read data population file: %v", err)
	}

	// Split the script into individual SQL statements
	statements := utils.SplitSQLStatements(string(script))

	// Execute each statement
	for _, stmt := range statements {
		_, err := db.Exec(stmt, masterPasswordHash, masterPasswordSalt)
		if err != nil {
			log.Fatalf("Failed to execute data population statement: %v", err)
		}
	}

	log.Println("Initial data populated successfully.")
}
