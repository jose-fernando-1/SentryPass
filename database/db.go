package database

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3" // Importa o driver SQLite
)

// Inicializa o banco de dados SQLite
func InitDB() *sql.DB {
	// Abre ou cria o arquivo do database
	db, err := sql.Open("sqlite3", "sentrypass.db")
	if err != nil {
		log.Fatal("Erro ao abrir o banco de dados: ", err)
	}
	log.Println("Banco de dados conectado com sucesso.")
	return db
}

// Executa todas as migrações no folder migrations
func RunMigrations(db *sql.DB) {
	files, err := os.ReadDir("migrations")
	if err != nil {
		log.Fatal("Erro ao ler diretório migrations: ", err)
	}

	for _, file := range files {
		script, err := os.ReadFile(filepath.Join("migrations", file.Name()))
		if err != nil {
			log.Fatalf("Erro ao ler o arquivo %s: %v", file.Name(), err)
		}

		_, err = db.Exec(string(script))
		if err != nil {
			log.Fatalf("Erro ao executar a migração %s: %v", file.Name(), err)
		}

		log.Printf("Migração %s aplicada com sucesso.", file.Name())
	}
}
