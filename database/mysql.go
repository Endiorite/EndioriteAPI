package database

import (
	"EndioriteAPI/config"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectMySQL() {
	dsn := config.GetEnv("MYSQL_DSN", "root:password@tcp(127.0.0.1:3306)/mydb")

	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Erreur d'ouverture MySQL: %v", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("Impossible de se connecter à MySQL: %v", err)
	}

	log.Println("Connexion MySQL établie")
}
