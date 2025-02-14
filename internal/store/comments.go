package store

import (
	"context"
	"database/sql"
)

type Comment struct{
	Id int64 `json:"id"`
	PostId int64 `json:"post_id"`
	UserId int64 `json:"user_id"`
	Content string `json:"content"`
	CreatedAt string `json:"created_At"`
	User User `json:"user"`
 }
type CommentStore struct {
	db *sql.DB
}

func(c*CommentStore) GetPostById(ctx context.Context,postId int64)([]Comment,error){
  query:=`
    SELECT c.id,c.post_id,c.user_id,c.content,c.created_At,users.username,users.id FROM comments c
	JOIN users on users.id = c.user_id
	WHERE c.post_id=$1
	ORDER BY c.created_At DESC; 
  `
  row,err:= c.db.QueryContext(ctx,query,postId)
  if err!=nil{
	return nil,err
  }
  defer row.Close()
  comments:=[]Comment{}
  for row.Next(){
	var c Comment
	c.User=User{}

	err:=row.Scan(&c.Id,&c.PostId,&c.UserId,&c.Content,&c.CreatedAt,&c.User.UserName,&c.User.ID)
	if err!=nil{
		return nil,err
	}
	comments = append(comments, c)
 }
 return comments,nil
}