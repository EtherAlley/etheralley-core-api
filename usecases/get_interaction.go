package usecases

import (
	"context"
	"errors"
	"strings"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/entities"
	"github.com/etheralley/etheralley-core-api/gateways"
)

type GetInteractionInput struct {
	Address     string            `validate:"required,eth_addr"`
	Interaction *InteractionInput `validate:"required,dive"`
}

// Get the transaction data for a given interaction input
//
// Validate transaction against interaction type
type IGetInteractionUseCase func(ctx context.Context, input *GetInteractionInput) (*entities.Interaction, error)

func NewGetInteractionUseCase(
	logger common.ILogger,
	blockchainGateway gateways.IBlockchainGateway,
) IGetInteractionUseCase {
	return func(ctx context.Context, input *GetInteractionInput) (*entities.Interaction, error) {
		if err := common.ValidateStruct(input); err != nil {
			return nil, err
		}

		address := input.Address
		tx := &entities.Transaction{
			Id:         input.Interaction.Transaction.Id,
			Blockchain: input.Interaction.Transaction.Blockchain,
		}

		data, err := blockchainGateway.GetTransactionData(ctx, tx)

		if err != nil {
			return nil, err
		}

		err = validateTransaction(address, data)

		if err != nil {
			return nil, err
		}

		var validationErr error
		switch input.Interaction.Type {
		case common.CONTRACT_CREATION:
			validationErr = validateContractCreation(data)
		case common.SEND_ETHER:
			validationErr = validateSendEther(data)
		default:
			validationErr = errors.New("unsupported interaction type")
		}

		if validationErr != nil {
			return nil, validationErr
		}

		interaction := &entities.Interaction{
			Transaction:     tx,
			Type:            input.Interaction.Type,
			Timestamp:       data.Timestamp,
			TransactionData: data,
		}

		return interaction, nil
	}
}

func validateTransaction(address string, data *entities.TransactionData) error {
	if !strings.EqualFold(data.From, address) {
		return errors.New("transaction - invalid from address")
	}
	return nil
}

func validateContractCreation(data *entities.TransactionData) error {
	if data.To != nil {
		return errors.New("contract creation - found to address")
	}

	return nil
}

func validateSendEther(data *entities.TransactionData) error {
	if data.Value == "0" {
		return errors.New("send ether - zero value")
	}

	return nil
}
