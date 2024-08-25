// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package userStore

import (
	"context"
)

type Querier interface {
	CreateUser(ctx context.Context, arg CreateUserParams) error
	DeleteUserByEmail(ctx context.Context, id int64) error
	GetUserById(ctx context.Context, id int64) (GetUserByIdRow, error)
	UpdateUserById(ctx context.Context, arg UpdateUserByIdParams) error
	UpdateUserRefreshTokenById(ctx context.Context, arg UpdateUserRefreshTokenByIdParams) error
}

var _ Querier = (*Queries)(nil)
