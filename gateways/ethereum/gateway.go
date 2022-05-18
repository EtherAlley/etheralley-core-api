package ethereum

import (
	"context"
	"errors"
	"fmt"

	cmn "github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/gateways"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

type gateway struct {
	settings cmn.ISettings
	logger   cmn.ILogger
	http     cmn.IHttpClient
}

func NewGateway(logger cmn.ILogger, settings cmn.ISettings, http cmn.IHttpClient) gateways.IBlockchainGateway {
	return &gateway{
		settings,
		logger,
		http,
	}
}

// Use the main uris for the core blockchain workflows
func (gw *gateway) getClient(ctx context.Context, blockchain cmn.Blockchain) (*ethclient.Client, error) {
	switch blockchain {
	case cmn.ETHEREUM:
		return ethclient.DialContext(ctx, gw.settings.EthereumMainURI())
	case cmn.POLYGON:
		return ethclient.DialContext(ctx, gw.settings.PolygonMainURI())
	case cmn.OPTIMISM:
		return ethclient.DialContext(ctx, gw.settings.OptimismMainURI())
	case cmn.ARBITRUM:
		return ethclient.DialContext(ctx, gw.settings.ArbitrumMainURI())
	}
	return nil, fmt.Errorf("eth get client %v", blockchain)
}

// Parse for go-ethereum http error to determine if its retryable.
// Wrap in common.ErrRetryable if status code is 429 and include msg
func (gw *gateway) tryWrapRetryable(ctx context.Context, msg string, err error) error {
	if err == nil {
		return nil
	}

	var e rpc.HTTPError
	if errors.As(err, &e) && e.StatusCode == 429 {
		gw.logger.Warn(ctx).Err(err).Msgf("eth rate limit: %v", msg)
		return fmt.Errorf("%v %v %w", msg, err, cmn.ErrRetryable)
	}

	return err
}
