package handlers

import (
    "net/http"

    "github.com/kaminski-pawel/goth_tut1/components"
    "github.com/kaminski-pawel/goth_tut1/models"
)

func CountryList(w http.ResponseWriter, r *http.Request) {
    component := components.CountryList(models.Countries)
    component.Render(r.Context(), w)
}

func CountryDetail(w http.ResponseWriter, r *http.Request) {
    http.NotFound(w, r)
}

func SearchCountries(w http.ResponseWriter, r *http.Request) {
    http.NotFound(w, r)
}
