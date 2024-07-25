// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package userstore

import (
	"context"
)

type Querier interface {
	CreateUser(ctx context.Context, arg CreateUserParams) error
	DeletUserByEmail(ctx context.Context, email string) error
	GetUserByEmail(ctx context.Context, email string) (User, error)
	GetUserById(ctx context.Context, id int32) (User, error)
	GetUsers(ctx context.Context) ([]User, error)
	UpdatePasswordByEmail(ctx context.Context, arg UpdatePasswordByEmailParams) error
	UpdateRefreshTokenByEmail(ctx context.Context, arg UpdateRefreshTokenByEmailParams) error
}

var _ Querier = (*Queries)(nil)