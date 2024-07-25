// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package sadstore

import (
	"context"
)

type Querier interface {
	CreateSad(ctx context.Context, arg CreateSadParams) error
	DeleteSadById(ctx context.Context, id int32) error
	GetSadById(ctx context.Context, id int32) (Sad, error)
	GetSadByUserId(ctx context.Context, id int32) ([]Sad, error)
	GetSads(ctx context.Context) ([]Sad, error)
	UpdateSad(ctx context.Context, arg UpdateSadParams) error
}

var _ Querier = (*Queries)(nil)
