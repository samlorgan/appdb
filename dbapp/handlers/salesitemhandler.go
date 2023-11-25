package handlers

import (
	"dbapp/types"
	"encoding/json"

	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (app *App) CountSalesItemsHandler(w http.ResponseWriter, r *http.Request) {
	// curl -v -X GET localhost:8080/countsalesitems
	type CountSalesItemsResponse struct {
		Count int
	}
	w.Header().Set("Content-Type", "application/json")
	response := CountSalesItemsResponse{}
	s := []types.SalesItem{}
	response.Count = len(s)
	// GetSalesItems(db, "code", []string{"P01", "P02", "P03"})
	// Use json.NewEncoder to write the response struct to the ResponseWriter as JSON
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (app *App) GetSalesItems(field string, values []string) (*[]types.SalesItem, error) {
	s := []types.SalesItem{}

	return &s, nil
}

func (app *App) GetSalesItemsHandler(w http.ResponseWriter, r *http.Request) {
	salesItemName := mux.Vars(r)["name"]
	// curl -v -X GET localhost:8080/salesitem/{name}
	type GetSalesItemsResponse struct {
		SalesItems []types.SalesItem
	}
	w.Header().Set("Content-Type", "application/json")
	response := GetSalesItemsResponse{}
	// s := []utils.SalesItem{}
	s, err := app.GetSalesItems("name", []string{salesItemName})
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to find %s - %v", salesItemName, err), http.StatusBadRequest)
		return
	}
	response.SalesItems = *s
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
