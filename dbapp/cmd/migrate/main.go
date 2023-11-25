package main

import (
	"context"
	"dbapp/ent"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Running cmd/migrate main")
	godotenv.Load()
	appName := os.Getenv("APP_NAME")
	wd, _ := os.Getwd()
	fmt.Println("Project:", appName, "Working Dir:", wd)
	dbuser := os.Getenv("DB_USER")
	dbpass := os.Getenv("DB_PASS")
	dbhost := os.Getenv("DB_HOST")
	dbport := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True", dbuser, dbpass, dbhost, dbport, dbname)
	client, err := ent.Open("mysql", connectionString)
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	fmt.Println("Completed migrate successfully!")
}
