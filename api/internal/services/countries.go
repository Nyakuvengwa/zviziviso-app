package services

import (
	"context"
	"errors"
	"fmt"
	repository "zviziviso-app/api/internal/db"
)

var (
	ErrCountryNotFound = errors.New("country not found")
)

type CountryService interface {
	ListCountries(ctx context.Context) ([]repository.Country, error)
	GetCountryById(ctx context.Context, countryId int32) (repository.Country, error)
}

type countryService struct {
	queries repository.Queries
}

// GetCountryById implements CountryService.
func (c *countryService) GetCountryById(ctx context.Context, countryId int32) (repository.Country, error) {
	if countryId > 0 {
		return c.queries.GetCountry(ctx, countryId)
	}

	return repository.Country{}, fmt.Errorf("%w (id %d)", ErrCountryNotFound, countryId)
}

// ListCountries implements CountryService.
func (c *countryService) ListCountries(ctx context.Context) ([]repository.Country, error) {
	return c.queries.ListCountries(ctx)
}

func NewCountryService(queries repository.Queries) CountryService {
	return &countryService{queries: queries}
}
