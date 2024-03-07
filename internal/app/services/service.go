package services

import (
	"context"
	"fmt"
	"github.com/abaykerimov/test_kmf/internal/domain/entity"
	"github.com/abaykerimov/test_kmf/internal/domain/entity/dto"
	"github.com/abaykerimov/test_kmf/internal/infrastructure/clients/nat_api"
	"time"
)

type NatClient interface {
	GetExchangeRates(ctx context.Context, date string) (*nat_api.GetExchangeRatesResponse, error)
}

type Service struct {
	repository entity.Repository
	client     NatClient
}

func NewService(repository entity.Repository, client NatClient) *Service {
	return &Service{
		repository: repository,
		client:     client,
	}
}

func (s *Service) GetByDate(ctx context.Context, date, code string) ([]*entity.Rate, error) {
	return s.repository.GetByDate(ctx, date, code)
}

func (s *Service) Create(ctx context.Context, date string) error {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	chErr := make(chan error)

	go func(ctx context.Context) {
		response, cliErr := s.client.GetExchangeRates(ctx, date)
		if cliErr != nil {
			chErr <- cliErr
			return
		}

		for _, v := range response.Rates {
			repErr := s.repository.Create(ctx, entity.CreateRate(dto.NewCreateDTO(v.Name, v.Title, date, v.Description)))
			if repErr != nil {
				chErr <- repErr
				return
			}
		}

		cancel()
	}(ctx)

	select {
	case <-ctx.Done():
		switch ctx.Err() {
		case context.DeadlineExceeded:
			fmt.Println("context timeout exceeded")
		case context.Canceled:
			fmt.Println("context cancelled by force. whole process is complete")
		}
	case err := <-chErr:
		fmt.Println("process fail causing by some error:", err.Error())
	}

	return nil
}
