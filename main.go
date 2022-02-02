package main

import (
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/controllers"
	"github.com/etheralley/etheralley-core-api/controllers/http"
	"github.com/etheralley/etheralley-core-api/gateways/ethereum"
	"github.com/etheralley/etheralley-core-api/gateways/mongo"
	"github.com/etheralley/etheralley-core-api/gateways/redis"
	"github.com/etheralley/etheralley-core-api/gateways/thegraph"
	"github.com/etheralley/etheralley-core-api/usecases"
	"go.uber.org/dig"
)

func awaitSigterm(logger common.ILogger) {
	logger.Info("awaiting sigterm")

	cancelChan := make(chan os.Signal, 1)
	signal.Notify(cancelChan, syscall.SIGTERM, syscall.SIGINT)
	sig := <-cancelChan

	logger.Infof("caught sigterm %v", sig)
}

func setRandSeed() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	container := dig.New()
	container.Provide(common.NewSettings)
	container.Provide(common.NewLogger)
	container.Provide(common.NewHttpClient)
	container.Provide(common.NewGraphQLClient)
	container.Provide(redis.NewGateway)
	container.Provide(mongo.NewGateway)
	container.Provide(ethereum.NewGateway)
	container.Provide(thegraph.NewGateway)
	container.Provide(usecases.NewGetChallenge)
	container.Provide(usecases.NewGetProfile)
	container.Provide(usecases.NewGetDefaultProfile)
	container.Provide(usecases.NewSaveProfile)
	container.Provide(usecases.NewVerifyChallenge)
	container.Provide(usecases.NewGetNonFungibleToken)
	container.Provide(usecases.NewGetAllNonFungibleTokens)
	container.Provide(usecases.NewResolveAddress)
	container.Provide(usecases.NewGetFungibleToken)
	container.Provide(usecases.NewGetAllFungibleTokens)
	container.Provide(usecases.NewGetStatistic)
	container.Provide(usecases.NewGetAllStatistics)
	container.Provide(http.NewHttpController)

	setRandSeed()

	// start controllers in concurrent go routines
	err := container.Invoke(func(controller *http.HttpController) {
		go controllers.StartController(controller)
	})

	if err != nil {
		panic(err)
	}

	// blocking call in main routine to await sigterm
	err = container.Invoke(awaitSigterm)

	if err != nil {
		panic(err)
	}

	// TODO: Shutdown gracefully below
}
