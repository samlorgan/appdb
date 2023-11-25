package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func connect(dbuser, dbpass, dbhost, dbport, dbname string) (*sql.DB, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbuser, dbpass, dbhost, dbport, dbname)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func getTables(db *sql.DB, schema string) ([]string, error) {
	rows, err := db.Query("SELECT TABLE_NAME FROM information_schema.tables WHERE table_schema = ?", schema)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tables []string
	for rows.Next() {
		var table string
		if err := rows.Scan(&table); err != nil {
			return nil, err
		}
		tables = append(tables, table)
	}
	return tables, rows.Err()
}

func dropTable(db *sql.DB, table string) error {
	_, err := db.Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", table))
	return err
}

func main() {
	fmt.Println("Running cmd/teardown main")
	godotenv.Load()
	appName := os.Getenv("APP_NAME")
	wd, _ := os.Getwd()
	fmt.Println("Project:", appName, "Working Dir:", wd)
	dbuser := os.Getenv("DB_USER")
	dbpass := os.Getenv("DB_PASS")
	dbhost := os.Getenv("DB_HOST")
	dbport := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	db, err := connect(dbuser, dbpass, dbhost, dbport, dbname)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tables, err := getTables(db, dbname)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found %d tables\n", len(tables))

	for _, table := range tables {
		fmt.Printf("Processing table %s\n", table)

		if err := dropTable(db, table); err != nil {
			fmt.Printf("Error dropping table %s %s \n", table, err.Error())
		}
		if err == nil {
			fmt.Printf("Dropped table %s\n", table)
		}

	}
}
