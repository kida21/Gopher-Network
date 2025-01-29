package store

import (
	"context"
	"database/sql"
)
type User struct {
	ID int64  `json:"id"`
	UserName string `json:"username"`
	Email string  `json:"email"`
	Password string `json:"-"`
	CreatedAt string `json:"created_At"`
}
type UserStore struct {
	db *sql.DB
}

func (u*UserStore)Create(ctx context.Context, user*User)error{
query := `
INSERT INTO users(username,email,password)
VALUES($1,$2,$3) RETURNING id,created_At
`
if err:= u.db.QueryRowContext(
	ctx,
	query,
	user.UserName,
	user.Email,
	user.Password,
).Scan(
	&user.ID,
	&user.CreatedAt,
);err != nil{
	return err
}
  return nil
}