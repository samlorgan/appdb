package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

func CreateRoutes(app App) http.Handler {
	standardMiddleware := alice.New(LogRequest, PanicRecoveryMiddleware)
	r := mux.NewRouter()
	r.HandleFunc("/wave", WaveHandler).Methods("GET")
	r.HandleFunc("/countsalesitems", app.CountSalesItemsHandler).Methods("GET")
	r.HandleFunc("/salesitem/{name}", app.GetSalesItemsHandler).Methods("GET")
	return standardMiddleware.Then(r)

}
