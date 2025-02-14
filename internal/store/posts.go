package store

import (
	"context"
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

type PostStore struct {
	db *sql.DB
}
type Post struct{
	ID int64   `json:"id"`
	Content string  `json:"content"`
	Title string  `json:"title"`
	UserID int64   `json:"user_id"`
	Tags []string  `json:"tags"`   
	CreatedAt string `json:"created_At"`
	UpdatedAt string  `json:"updated_At"`
	Comments []Comment  `json:"comments"`
}

func (s*PostStore) Create(ctx context.Context,post *Post)error{
 query := `
 INSERT INTO posts (content,title,user_id,tags)
 VALUES ($1,$2,$3,$4) RETURNING id,created_At,updated_At
 `
 err:= s.db.QueryRowContext(
	ctx,
	query,
	post.Content,
	post.Title,
	post.UserID,
	pq.Array(post.Tags),
).Scan(
	&post.ID,
	&post.CreatedAt,
	&post.UpdatedAt,
)
if err != nil{
	return err
}
return nil
}

func(s*PostStore) GetPostById(ctx context.Context,id int64)(*Post,error)  {
	var post Post
	query := `
	 SELECT id,content,title,user_id,tags,created_At,Updated_At FROM posts 
	 WHERE id = $1 
	`
	err :=s.db.QueryRowContext(
		ctx,
		query,
		id,
	).Scan(
     &post.ID,
	 &post.Content,
	 &post.Title,
	 &post.UserID,
	 pq.Array(&post.Tags),
	 &post.CreatedAt,
	 &post.UpdatedAt,
	 
	)
	if err !=nil{
		switch{
		case errors.Is(err,sql.ErrNoRows):
			return nil,ErrNotFound;
		default:
			return nil,err
		}
		
	}
	return &post,nil
}
func (s*PostStore) Delete(ctx context.Context,id int64)error {
	query := `DELETE from posts where id =$1`
	res,err:=s.db.ExecContext(ctx,query,id)
	if err!=nil{
		return nil
	}
	row,err:= res.RowsAffected()
	if err!=nil{
		return nil
	}
	if row ==0{
		return ErrNotFound
	}
	return nil
}
func(S*PostStore)Update(ctx context.Context,post *Post)(error){
   
	query := `UPDATE posts SET content=$1,title=$2 WHERE id=$3`
	
	_,err:=S.db.ExecContext(ctx,query,post.Content,post.Title,post.ID)
	if err!=nil{
		return err
	}
	return nil
}

