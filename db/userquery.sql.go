// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: userquery.sql

package grpcservice

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  email,
  password, 
  date_created
) VALUES ($1, $2, $3)
RETURNING id, email, password, date_created
`

type CreateUserParams struct {
	Email       string
	Password    string
	DateCreated string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Email, arg.Password, arg.DateCreated)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.DateCreated,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT id, email, password, date_created FROM users WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.DateCreated,
	)
	return i, err
}

const removeUser = `-- name: RemoveUser :one
DELETE FROM users WHERE id = $1
RETURNING id, email, password, date_created
`

func (q *Queries) RemoveUser(ctx context.Context, id string) (User, error) {
	row := q.db.QueryRowContext(ctx, removeUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.DateCreated,
	)
	return i, err
}
