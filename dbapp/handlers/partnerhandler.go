package handlers

import (
	"dbapp/ent"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (app *App) GetPartner(w http.ResponseWriter, r *http.Request) {
	partnerTitle := mux.Vars(r)["title"]
	// curl -v -X GET localhost:8080/partner{name}
	type GetPartnerResponse struct {
		Partner ent.Partner
	}
	w.Header().Set("Content-Type", "application/json")
	response := GetPartnerResponse{}
	// s := []utils.SalesItem{}
	p := app.Database.GetAllPartners()
	for _, partner := range p {
		if partner.WhamTitle == partnerTitle {
			response.Partner = *partner
			err := json.NewEncoder(w).Encode(response)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}
	http.Error(w, fmt.Sprintf("unable to find parner %s", partnerTitle), http.StatusNotFound)

}
func (app *App) GetAllPartners(w http.ResponseWriter, r *http.Request) {
	// curl -v -X GET localhost:8080/allpartners
	type GetAllPartnersResponse struct {
		Partners []ent.Partner
	}
	w.Header().Set("Content-Type", "application/json")
	response := GetAllPartnersResponse{}
	// s := []utils.SalesItem{}
	p := app.Database.GetAllPartners()
	for _, partner := range p {
		response.Partners = append(response.Partners, *partner)
	}

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
