package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func Connect() {

	//establish connection with heroku postgres
	fmt.Println("logging into the database")
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(db)
}
