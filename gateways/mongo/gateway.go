package mongo

import (
	"context"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/gateways"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Gateway struct {
	logger   common.ILogger
	profiles *mongo.Collection
}

func NewGateway(settings common.ISettings, logger common.ILogger) gateways.IDatabaseGateway {
	ctx := context.Background()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(settings.DatabaseURI()).SetMaxConnecting(100))

	if err != nil {
		logger.Err(ctx, err, "mongo connection error")
		panic(err)
	}

	db := client.Database(settings.Database())
	profiles := db.Collection("profiles")

	return &Gateway{
		logger,
		profiles,
	}
}

type profileBson struct {
	Address           string                  `bson:"_id"`
	NonFungibleTokens *[]nonFungibleTokenBson `bson:"non_fungible_tokens"`
	FungibleTokens    *[]fungibleTokenBson    `bson:"fungible_tokens"`
	Statistics        *[]statisticBson        `bson:"statistics"`
	Interactions      *[]interactionBson      `bson:"interactions"`
}

type contractBson struct {
	Blockchain common.Blockchain `bson:"blockchain"`
	Address    string            `bson:"address"`
	Interface  common.Interface  `bson:"interface"`
}

type nonFungibleTokenBson struct {
	Contract *contractBson `bson:"contract"`
	TokenId  string        `bson:"token_id"`
}

type fungibleTokenBson struct {
	Contract *contractBson `bson:"contract"`
}

type statisticBson struct {
	Contract *contractBson        `bson:"contract"`
	Type     common.StatisticType `bson:"type"`
}

type transactionBson struct {
	Id         string            `bson:"id"`
	Blockchain common.Blockchain `bson:"blockchain"`
}

type interactionBson struct {
	Transaction *transactionBson   `bson:"transaction"`
	Type        common.Interaction `bson:"type"`
	Timestamp   uint64             `bson:"timestamp"`
}
