package main

import (
	"log"
	"path/filepath"
	"sentrypass/logic/database"
	"sentrypass/ui/screens"

	"gioui.org/app"
	"github.com/sqweek/dialog"
)

func main() {
	// Show a directory selection dialog
	dbDir, err := dialog.Directory().Title("Select Database Directory").Browse()
	if err != nil {
		log.Fatalf("Failed to select directory: %v", err)
	}

	dbPath := filepath.Join(dbDir, "database.db")
	log.Printf("Selected database path: %s", dbPath)

	db := database.InitDB(dbPath)
	defer db.Close()

	database.RunSchemaMigration(db)

	// Prompt the user for the master password (mocked for now)
	masterPassword := "example_password" // Replace with actual password input logic
	masterPasswordHash := "hashed_" + masterPassword
	masterPasswordSalt := "salt_" + masterPassword

	database.PopulateInitialData(db, masterPasswordHash, masterPasswordSalt)

	log.Println("Application started")
	go func() {
		window := new(app.Window)
		err := screens.Run_GUI(window)
		if err != nil {
			log.Fatal(err)
		}
	}()
	app.Main()
}
