// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package userstore

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.createUserStmt, err = db.PrepareContext(ctx, createUser); err != nil {
		return nil, fmt.Errorf("error preparing query CreateUser: %w", err)
	}
	if q.deletUserByEmailStmt, err = db.PrepareContext(ctx, deletUserByEmail); err != nil {
		return nil, fmt.Errorf("error preparing query DeletUserByEmail: %w", err)
	}
	if q.getUserByEmailStmt, err = db.PrepareContext(ctx, getUserByEmail); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserByEmail: %w", err)
	}
	if q.getUserByIdStmt, err = db.PrepareContext(ctx, getUserById); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserById: %w", err)
	}
	if q.getUsersStmt, err = db.PrepareContext(ctx, getUsers); err != nil {
		return nil, fmt.Errorf("error preparing query GetUsers: %w", err)
	}
	if q.updatePasswordByEmailStmt, err = db.PrepareContext(ctx, updatePasswordByEmail); err != nil {
		return nil, fmt.Errorf("error preparing query UpdatePasswordByEmail: %w", err)
	}
	if q.updateRefreshTokenByEmailStmt, err = db.PrepareContext(ctx, updateRefreshTokenByEmail); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateRefreshTokenByEmail: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.createUserStmt != nil {
		if cerr := q.createUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createUserStmt: %w", cerr)
		}
	}
	if q.deletUserByEmailStmt != nil {
		if cerr := q.deletUserByEmailStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deletUserByEmailStmt: %w", cerr)
		}
	}
	if q.getUserByEmailStmt != nil {
		if cerr := q.getUserByEmailStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserByEmailStmt: %w", cerr)
		}
	}
	if q.getUserByIdStmt != nil {
		if cerr := q.getUserByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserByIdStmt: %w", cerr)
		}
	}
	if q.getUsersStmt != nil {
		if cerr := q.getUsersStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUsersStmt: %w", cerr)
		}
	}
	if q.updatePasswordByEmailStmt != nil {
		if cerr := q.updatePasswordByEmailStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updatePasswordByEmailStmt: %w", cerr)
		}
	}
	if q.updateRefreshTokenByEmailStmt != nil {
		if cerr := q.updateRefreshTokenByEmailStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateRefreshTokenByEmailStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                            DBTX
	tx                            *sql.Tx
	createUserStmt                *sql.Stmt
	deletUserByEmailStmt          *sql.Stmt
	getUserByEmailStmt            *sql.Stmt
	getUserByIdStmt               *sql.Stmt
	getUsersStmt                  *sql.Stmt
	updatePasswordByEmailStmt     *sql.Stmt
	updateRefreshTokenByEmailStmt *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                            tx,
		tx:                            tx,
		createUserStmt:                q.createUserStmt,
		deletUserByEmailStmt:          q.deletUserByEmailStmt,
		getUserByEmailStmt:            q.getUserByEmailStmt,
		getUserByIdStmt:               q.getUserByIdStmt,
		getUsersStmt:                  q.getUsersStmt,
		updatePasswordByEmailStmt:     q.updatePasswordByEmailStmt,
		updateRefreshTokenByEmailStmt: q.updateRefreshTokenByEmailStmt,
	}
}