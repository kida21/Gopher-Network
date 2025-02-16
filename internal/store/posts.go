package store

import (
	"context"
	"database/sql"
	"errors"
	"time"

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
	User User `json:"user"`
}

type PostWithMetaData struct{
 Post
 CommentCount int64 `json:"comment_count"`
}
func(s*PostStore)GetUserFeed(ctx context.Context,userId int64)([]PostWithMetaData,error){
	query:=`
	 SELECT p.id,p.user_id,p.title,p.content,p.created_At,p.tags,u.username
	 COUNT(C.id) AS comments_count FROM posts p
	 LEFT JOIN comments c ON c.post_id = p.id
	 LEFT JOIN users u on p.user_id = u.id
	 JOIN followers f ON f.follower_id = p.user_id OR p.user_id=$1
	 WHERE f.user_id = $1 OR p.user_id = $1
	 GROUP BY p.id,p.username
	 ORDER BY p.created_At DESC
	`
	ctx,cancel := context.WithTimeout(ctx,time.Second*5)
	defer cancel()
	rows,err:= s.db.QueryContext(ctx,query,userId)
	if err!=nil{
		return nil,err
	}
	defer rows.Close()
	var feed []PostWithMetaData
	for rows.Next(){
		var p PostWithMetaData
		err:=rows.Scan(
			&p.ID,
			&p.UserID,
			&p.Title,
			&p.Content,
			&p.CreatedAt,
			pq.Array(&p.Tags),
			&p.User.UserName,
			&p.CommentCount,
		)
		if err!=nil{
			return nil,err
		}
		feed = append(feed,p)
	
	}
	return feed,nil
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

