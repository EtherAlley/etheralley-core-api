package usecases

import (
	"context"
	"strings"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/gateways"
	"github.com/etheralley/etheralley-core-api/gateways/ethereum"
	"github.com/etheralley/etheralley-core-api/gateways/redis"
)

func NewResolveAddressUseCase(logger *common.Logger, blockchainGateway *ethereum.Gateway, cacheGateway *redis.Gateway) IResolveAddressUseCase {
	return ResolveAddress(logger, blockchainGateway, cacheGateway)
}

// if the address contains a ".", we assume its an attempted ens address and try to resolve
// attempt to cache resolved addresses
// if no "." we check for valid hex address and return if valid
func ResolveAddress(logger *common.Logger, blockchainGateway gateways.IBlockchainGateway, cacheGateway gateways.ICacheGateway) IResolveAddressUseCase {
	return func(ctx context.Context, input string) (string, error) {
		normalized := strings.ToLower(input)

		if strings.Contains(normalized, ".") {

			address, err := cacheGateway.GetENSAddressFromName(ctx, normalized)

			if err == nil {
				logger.Debugf("cache hit for ens name %v -> address %v", input, address)
				return address, err
			}

			logger.Debugf("cache miss for ens name %v", normalized)

			address, err = blockchainGateway.GetENSAddressFromName(normalized)

			if err != nil {
				logger.Debugf("chain miss for ens name %v err: %v", normalized, err)
				return address, err
			}

			logger.Debugf("chain hit for ens name %v -> address %v", normalized, address)

			cacheGateway.SaveENSAddress(ctx, normalized, address)

			return address, err
		}

		if err := common.ValidateField(normalized, `required,eth_addr`); err != nil {
			return normalized, err
		}

		return normalized, nil
	}
}
