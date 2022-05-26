package http

import (
	"context"
	"fmt"
	"net/http"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/presenters"
	"github.com/etheralley/etheralley-core-api/usecases"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type HttpController struct {
	settings            common.ISettings
	logger              common.ILogger
	presenter           presenters.IPresenter
	getProfile          usecases.IGetProfileUseCase
	saveProfile         usecases.ISaveProfileUseCase
	getChallenge        usecases.IGetChallengeUseCase
	verifyChallenge     usecases.IVerifyChallengeUseCase
	getNonFungibleToken usecases.IGetNonFungibleTokenUseCase
	resolveAddress      usecases.IResolveAddressUseCase
	getFungibleToken    usecases.IGetFungibleTokenUseCase
	getStatistic        usecases.IGetStatisticUseCase
	getInteraction      usecases.IGetInteractionUseCase
	recordProfileView   usecases.IRecordProfileViewUseCase
	getTopProfiles      usecases.IGetTopProfilesUseCase
	getListingMetadata  usecases.IGetListingMetadataUseCase
	getListings         usecases.IGetListingsUseCase
	refreshProfile      usecases.IRefreshProfileUseCase
	getCurrency         usecases.IGetCurrencyUseCase
	getStoreMetadata    usecases.IGetStoreMetadataUseCase
	verifyRateLimit     usecases.IVerifyRateLimitUseCase
}

func NewHttpController(
	settings common.ISettings,
	logger common.ILogger,
	presenter presenters.IPresenter,
	getProfile usecases.IGetProfileUseCase,
	saveProfile usecases.ISaveProfileUseCase,
	getChallenge usecases.IGetChallengeUseCase,
	verifyChallenge usecases.IVerifyChallengeUseCase,
	getNonFungibleToken usecases.IGetNonFungibleTokenUseCase,
	resolveAddress usecases.IResolveAddressUseCase,
	getFungibleToken usecases.IGetFungibleTokenUseCase,
	getStatistic usecases.IGetStatisticUseCase,
	getInteraction usecases.IGetInteractionUseCase,
	recordProfileView usecases.IRecordProfileViewUseCase,
	getTopProfiles usecases.IGetTopProfilesUseCase,
	getListingMetadata usecases.IGetListingMetadataUseCase,
	getListings usecases.IGetListingsUseCase,
	refreshProfile usecases.IRefreshProfileUseCase,
	getCurrency usecases.IGetCurrencyUseCase,
	getStoreMetadata usecases.IGetStoreMetadataUseCase,
	verifyRateLimit usecases.IVerifyRateLimitUseCase,
) *HttpController {
	return &HttpController{
		settings,
		logger,
		presenter,
		getProfile,
		saveProfile,
		getChallenge,
		verifyChallenge,
		getNonFungibleToken,
		resolveAddress,
		getFungibleToken,
		getStatistic,
		getInteraction,
		recordProfileView,
		getTopProfiles,
		getListingMetadata,
		getListings,
		refreshProfile,
		getCurrency,
		getStoreMetadata,
		verifyRateLimit,
	}
}

func (hc *HttpController) Start() error {
	ctx := context.Background()

	hc.logger.Info(ctx).Msg("starting http controller")

	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://etheralley.io", "http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	r.Use(middleware.NoCache)
	r.Use(hc.timer)
	r.Use(hc.realIP)
	r.Use(hc.requestId)
	r.Use(hc.recoverer)
	r.Use(hc.timeout)
	r.Use(hc.rateLimit)

	r.Get("/", hc.healthRoute)

	r.Route("/profiles", func(r chi.Router) {
		r.Get("/top", hc.getTopProfilesRoute)

		r.Route("/{address}", func(r chi.Router) {
			r.Use(hc.resolveAddressRoute)
			r.With(hc.recordProfileViewMiddleware).Get("/", hc.getProfileByAddressRoute)
			r.With(hc.authenticate).Put("/", hc.saveProfileRoute)
			r.Get("/refresh", hc.refreshProfileRoute)
		})
	})

	r.Route("/challenges/{address}", func(r chi.Router) {
		r.Use(hc.resolveAddressRoute)
		r.Get("/", hc.getChallengeRoute)
	})

	r.Route("/contracts", func(r chi.Router) {
		r.Use(hc.parseContract)
		r.Get("/token", hc.getTokenRoute)
		r.Get("/nft", hc.getNFTRoute)
		r.Get("/statistic", hc.getStatisticRoute)
	})

	r.Get("/currency", hc.getCurrencyRoute)

	r.Route("/transactions", func(r chi.Router) {
		r.Use(hc.parseTransaction)
		r.Get("/interaction", hc.getInteractionRoute)
	})

	r.Route("/listings", func(r chi.Router) {
		r.Post("/", hc.getListingsRoute)
		r.Get("/metadata", hc.getStoreMetadataRoute)
		r.Get("/metadata/{tokenid}", hc.getMetadataByIdRoute)
	})

	port := hc.settings.Port()

	hc.logger.Info(ctx).Msgf("listening on port %v", port)

	err := http.ListenAndServe(fmt.Sprintf(":%v", port), r)

	hc.logger.Error(ctx).Err(err).Msg("error in http controller")

	return err
}

func (hc *HttpController) Exit() {
	ctx := context.Background()
	hc.logger.Error(ctx).Msg("detected exit in http controller")
}
