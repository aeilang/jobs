// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: users.sql

package userStore

import (
	"context"
)

const createUser = `-- name: CreateUser :exec
insert into users 
(name, email, password_hash)
values ($1, $2, $3)
`

type CreateUserParams struct {
	Name         string `db:"name" json:"name"`
	Email        string `db:"email" json:"email"`
	PasswordHash string `db:"password_hash" json:"password_hash"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.exec(ctx, q.createUserStmt, createUser, arg.Name, arg.Email, arg.PasswordHash)
	return err
}

const deleteUserByEmail = `-- name: DeleteUserByEmail :exec
update users set
is_deleted = true
where id = $1
`

func (q *Queries) DeleteUserByEmail(ctx context.Context, id int64) error {
	_, err := q.exec(ctx, q.deleteUserByEmailStmt, deleteUserByEmail, id)
	return err
}

const getUserById = `-- name: GetUserById :one
select name, email, password_hash, refresh_token
from users
where is_deleted = false
and id = $1
`

type GetUserByIdRow struct {
	Name         string `db:"name" json:"name"`
	Email        string `db:"email" json:"email"`
	PasswordHash string `db:"password_hash" json:"password_hash"`
	RefreshToken string `db:"refresh_token" json:"refresh_token"`
}

func (q *Queries) GetUserById(ctx context.Context, id int64) (GetUserByIdRow, error) {
	row := q.queryRow(ctx, q.getUserByIdStmt, getUserById, id)
	var i GetUserByIdRow
	err := row.Scan(
		&i.Name,
		&i.Email,
		&i.PasswordHash,
		&i.RefreshToken,
	)
	return i, err
}

const updateUserById = `-- name: UpdateUserById :exec
update users set
name = $1, email = $2, password_hash = $3, refresh_token=$4, version = version +1
where id = $5 and version = $6 and is_deleted = false
`

type UpdateUserByIdParams struct {
	Name         string `db:"name" json:"name"`
	Email        string `db:"email" json:"email"`
	PasswordHash string `db:"password_hash" json:"password_hash"`
	RefreshToken string `db:"refresh_token" json:"refresh_token"`
	ID           int64  `db:"id" json:"id"`
	Version      int32  `db:"version" json:"version"`
}

func (q *Queries) UpdateUserById(ctx context.Context, arg UpdateUserByIdParams) error {
	_, err := q.exec(ctx, q.updateUserByIdStmt, updateUserById,
		arg.Name,
		arg.Email,
		arg.PasswordHash,
		arg.RefreshToken,
		arg.ID,
		arg.Version,
	)
	return err
}

const updateUserRefreshTokenById = `-- name: UpdateUserRefreshTokenById :exec
update users set
refresh_token = $1
where id = $2
`

type UpdateUserRefreshTokenByIdParams struct {
	RefreshToken string `db:"refresh_token" json:"refresh_token"`
	ID           int64  `db:"id" json:"id"`
}

func (q *Queries) UpdateUserRefreshTokenById(ctx context.Context, arg UpdateUserRefreshTokenByIdParams) error {
	_, err := q.exec(ctx, q.updateUserRefreshTokenByIdStmt, updateUserRefreshTokenById, arg.RefreshToken, arg.ID)
	return err
}