package mongo

import (
	"context"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (g *Gateway) GetProfileByAddress(ctx context.Context, address string) (*entities.Profile, error) {
	profileBson := &profileBson{}

	err := g.profiles.FindOne(ctx, bson.D{primitive.E{Key: "_id", Value: address}}).Decode(profileBson)

	if err == mongo.ErrNoDocuments {
		return nil, common.ErrNotFound
	}

	profile := fromProfileBson(profileBson)

	return profile, err
}

func (g *Gateway) SaveProfile(ctx context.Context, profile *entities.Profile) error {
	profileBson := toProfileBson(profile)

	_, err := g.profiles.UpdateOne(ctx, bson.D{primitive.E{Key: "_id", Value: profile.Address}}, bson.D{primitive.E{Key: "$set", Value: profileBson}}, options.Update().SetUpsert(true))

	return err
}

func fromProfileBson(profileBson *profileBson) *entities.Profile {
	profile := &entities.Profile{
		Address:           profileBson.Address,
		NonFungibleTokens: &[]entities.NonFungibleToken{},
		FungibleTokens:    &[]entities.FungibleToken{},
		Statistics:        &[]entities.Statistic{},
		Interactions:      &[]entities.Interaction{},
	}
	return profile
}

func toProfileBson(profile *entities.Profile) *profileBson {
	profileBson := &profileBson{
		Address:           profile.Address,
		NonFungibleTokens: &[]nonFungibleTokenBson{},
		FungibleTokens:    &[]fungibleTokenBson{},
		Statistics:        &[]statisticBson{},
		Interactions:      &[]interactionBson{},
	}
	return profileBson
}
