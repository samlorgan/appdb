package main

import (
	"dbapp/ent"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Running cmd/populate main")
	dbuser := "appuser"
	dbpass := "pass"
	dbhost := "localhost"
	dbport := "3306"
	dbname := "appdb"
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True", dbuser, dbpass, dbhost, dbport, dbname)
	client, err := ent.Open("mysql", connectionString)
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()

}
