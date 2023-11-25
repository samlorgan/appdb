package main

import (
	"dbapp/db"
	"dbapp/handlers"
	"dbapp/utils"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	fmt.Println("Running cmd/api main")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Panic occurred:", err)
			fmt.Printf("Stack Trace: %v", spew.Sdump(err))

		}
	}()

	go backgroundLogger()
	app := handlers.App{Database: db.Init()}
	r := handlers.CreateRoutes(app)
	port := fmt.Sprintf(":%s", os.Getenv("PORT_NUMBER"))
	server := &http.Server{Addr: port, Handler: r}
	fmt.Println("Server is running on port", port)
	log.Fatal(server.ListenAndServe())
}

func backgroundLogger() {
	appName := os.Getenv("APP_NAME")
	fmt.Printf("Project %s Main\n", appName)

	for {
		utils.PrintTime(appName + " Heartbeat")
		time.Sleep(30 * time.Second)
	}

}
