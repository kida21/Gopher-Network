package store

import (
	"context"
	"database/sql"
)
type Follower struct{
	UserId int64 `json:"user_id"`
	FollowId int64 `json:"follow_id"`
	created_At string `json:"created_At"`
}

type FollowStore struct {
	db *sql.DB
}
func(F*FollowStore) Follow(ctx context.Context,followId,userId int64)error{
  query:=`INSERT  INTO followers(user_id,follower_id) VALUES ($1,$2)`
  _,err:=F.db.ExecContext(ctx,query,userId,followId)
  return err
}
func(F*FollowStore) Unfollow(ctx context.Context,followId,userId int64)error{
 query:=`DELETE FROM followers WHERE user_id=$1 AND follower_id=$2`
  _,err:=F.db.ExecContext(ctx,query,userId,followId)
  return err
}