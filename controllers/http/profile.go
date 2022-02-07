package http

import (
	"encoding/json"
	"net/http"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/entities"
)

func (hc *HttpController) getProfileByAddressRoute(w http.ResponseWriter, r *http.Request) {
	address := r.Context().Value(common.ContextKeyAddress).(string)

	profile, err := hc.getProfile(r.Context(), address)

	if err != nil {
		RenderError(w, http.StatusBadRequest, "bad request")
		return
	}

	Render(w, http.StatusOK, profile)
}

func (hc *HttpController) saveProfileRoute(w http.ResponseWriter, r *http.Request) {
	address := r.Context().Value(common.ContextKeyAddress).(string)

	profile := &entities.Profile{}
	err := json.NewDecoder(r.Body).Decode(profile)

	if err != nil {
		RenderError(w, http.StatusBadRequest, "bad request")
		return
	}

	err = hc.saveProfile(r.Context(), address, profile)

	if err != nil {
		RenderError(w, http.StatusBadRequest, "bad request")
		return
	}

	RenderNoBody(w, http.StatusCreated)
}
