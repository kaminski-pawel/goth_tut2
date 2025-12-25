package handlers

import (
    "net/http"

    "github.com/kaminski-pawel/goth_tut2/components"
    "github.com/kaminski-pawel/goth_tut2/models"
)

func CountryList(w http.ResponseWriter, r *http.Request) {
    component := components.CountryList(models.Countries)
    component.Render(r.Context(), w)
}

func CountryDetail(w http.ResponseWriter, r *http.Request) {
    name := r.PathValue("name")
    for _, country := range models.Countries {
	if country.Name == name {
	    component := components.CountryDetail(country)
	    component.Render(r.Context(), w)
	    return
	}
    }
    http.NotFound(w, r)
}

func SearchCountries(w http.ResponseWriter, r *http.Request) {
    http.NotFound(w, r)
}
