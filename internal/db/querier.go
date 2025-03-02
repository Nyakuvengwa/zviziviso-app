// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package repository

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateUserDetails(ctx context.Context, arg CreateUserDetailsParams) (uuid.UUID, error)
	GetCountry(ctx context.Context, id int32) (Country, error)
	GetProvincesByCountryId(ctx context.Context, countryID int32) ([]Province, error)
	GetProvincesById(ctx context.Context, id int32) (Province, error)
	GetUserByEmailOrUsername(ctx context.Context, arg GetUserByEmailOrUsernameParams) ([]User, error)
	GetUserSummaryDetails(ctx context.Context, userID uuid.UUID) (GetUserSummaryDetailsRow, error)
	ListCountries(ctx context.Context) ([]Country, error)
	UpdateUserDetails(ctx context.Context, arg UpdateUserDetailsParams) error
	UpdateUserPassword(ctx context.Context, arg UpdateUserPasswordParams) error
}

var _ Querier = (*Queries)(nil)
