package http

import (
	"net/http"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/usecases"
)

func (hc *HttpController) getChallengeRoute(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	address := ctx.Value(common.ContextKeyAddress).(string)

	challenge, err := hc.getChallenge(ctx, &usecases.GetChallengeInput{
		Address: address,
	})

	if err != nil {
		hc.presenter.PresentBadRequest(w, r, err)
		return
	}

	hc.presenter.PresentChallenge(w, r, challenge)
}
