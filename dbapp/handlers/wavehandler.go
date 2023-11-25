package handlers

import (
	"fmt"
	"net/http"
	"os"
)

func WaveHandler(w http.ResponseWriter, r *http.Request) {
	// curl -v -X GET localhost:8080/wave
	appName := os.Getenv("APP_NAME")
	w.Header().Set("Content-Type", "application/json")
	message := fmt.Sprintf(`{"message": "Waves! Hello from the %s wave page!"}`, appName)
	w.Write([]byte(message))

}
