package presenter

import (
	"net/http"

	"github.com/etheralley/etheralley-backend/profiles-api/entities"
)

func (p *presenter) PresentCurrency(w http.ResponseWriter, r *http.Request, currency *entities.Currency) {
	json := toCurrencyJson(currency)

	p.presentJSON(w, r, http.StatusOK, json)
}
