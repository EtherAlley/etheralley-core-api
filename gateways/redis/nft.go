package redis

import (
	"context"
	"encoding/json"
	"time"

	"github.com/etheralley/etheralley-core-api/entities"
)

const NFTNamespace = "non_fungible_token"

func (g *gateway) GetNonFungibleMetadata(ctx context.Context, contract *entities.Contract, tokenId string) (*entities.NonFungibleMetadata, error) {

	metadataString, err := g.client.Get(ctx, getFullKey(NFTNamespace, contract.Address, tokenId, contract.Blockchain)).Result()

	if err != nil {
		return nil, err
	}

	metadataJson := &nonFungibleMetadataJson{}
	err = json.Unmarshal([]byte(metadataString), metadataJson)

	if err != nil {
		return nil, err
	}

	metadata := fromNonFungibleMetadataJson(metadataJson)

	return metadata, nil
}

func (g *gateway) SaveNonFungibleMetadata(ctx context.Context, contract *entities.Contract, tokenId string, metadata *entities.NonFungibleMetadata) error {
	metadataJson := toNonFungibleMetadataJson(metadata)
	bytes, err := json.Marshal(metadataJson)

	if err != nil {
		return err
	}

	_, err = g.client.Set(ctx, getFullKey(NFTNamespace, contract.Address, tokenId, contract.Blockchain), bytes, time.Hour).Result()

	return err
}
