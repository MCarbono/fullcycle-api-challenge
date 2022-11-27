package usecase

import (
	"context"
	"fullcycle-api-challenge/server/entity"
	"fullcycle-api-challenge/server/gateway"
	"fullcycle-api-challenge/server/repository"
	"fullcycle-api-challenge/server/usecase/dto"

	"github.com/google/uuid"
)

type DollarCotationUseCase struct {
	DollarCotationRepository repository.DollarCotationRepository
}

func (uc *DollarCotationUseCase) DollarCotation(ctx context.Context) (*dto.DollarCotationDTO, error) {
	dollarCotationGatewayResponse, err := gateway.DollarCotation("https://economia.awesomeapi.com.br/json/last/USD-BRL")
	if err != nil {
		return nil, err
	}
	id := uuid.New().String()
	dollarCotation := entity.NewDollarCotation(id, dollarCotationGatewayResponse)
	err = uc.DollarCotationRepository.Save(dollarCotation)
	if err != nil {
		return nil, err
	}
	return dto.NewDollarCotationDTO(dollarCotation), nil
}
