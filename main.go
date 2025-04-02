package main

import (
	"log"
	"os"
	"sentrypass/logic/database"
	"sentrypass/ui/screens"

	"gioui.org/app"
)

func main() {
	//Initializing database
	db := database.InitDB()
	defer db.Close()
	database.RunMigrations(db)

	log.Println("Application started")
	go func() {
		window := new(app.Window)
		err := screens.Run_GUI(window)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}
