package postgres

import (
	pb "comment-service/genproto/comment_service"
	"database/sql"

	"github.com/google/uuid"
)

type commentRepo struct {
	db *sql.DB
}

func NewCommentRepo(db *sql.DB) *commentRepo {
	return &commentRepo{
		db: db,
	}
}

//rpc CreateComment(Comment) returns (Comment);
//rpc GetCommentById(GetCommentId) returns (Comment);
//rpc GetAllCommentsByPostId(GetPostID) returns (AllComments);
//rpc GetAllCommentsByOwnerId(GetOwnerID) returns (AllComments);

func (c *commentRepo) CreateComment(comment *pb.Comment) (*pb.Comment, error) {
	if comment.Id == "" {
		comment.Id = uuid.NewString()
	}
	query := `INSERT INTO comments(id, content, owner_id, post_id) VALUES($1, $2, $3, $4) RETURNING id, content, owner_id, post_id, created_at`

	rowComment := c.db.QueryRow(query,
		comment.Id,
		comment.Content,
		comment.OwnerId,
		comment.PostId)
	if err := rowComment.Scan(&comment.Id,
		&comment.Content,
		&comment.OwnerId,
		&comment.PostId,
		&comment.CreatedAt); err != nil {
		return nil, err
	}

	return comment, nil
}

func (c *commentRepo) GetCommentById(commentId *pb.GetCommentId) (*pb.Comment, error) {
	query := `SELECT id, content, owner_id, post_id, created_at FROM comments WHERE id = $1`

	rowComment := c.db.QueryRow(query, commentId.Id)
	respComment := pb.Comment{}

	if err := rowComment.Scan(&respComment.Id,
		&respComment.Content,
		&respComment.OwnerId,
		&respComment.PostId,
		&respComment.CreatedAt); err != nil {
		return nil, err
	}

	return &respComment, nil
}

func (c *commentRepo) GetAllCommentsByPostId(postId *pb.GetPostID) (*pb.AllComments, error) {
	query := `SELECT id, content, owner_id, post_id, created_at FROM comments WHERE deleted_at IS NULL AND post_id = $1`

	var respComments pb.AllComments
	rows, err := c.db.Query(query, postId.PostId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var respComment = pb.Comment{}
		if err := rows.Scan(&respComment.Id,
			&respComment.Content,
			&respComment.OwnerId,
			&respComment.PostId,
			&respComment.CreatedAt); err != nil {
			return nil, err
		}

		respComments.Comments = append(respComments.Comments, &respComment)
	}

	return &respComments, nil
}

func (c *commentRepo) GetAllCommentsByOwnerId(ownerId *pb.GetOwnerID) (*pb.AllComments, error) {
	query := `SELECT id, content, owner_id, post_id, created_at FROM comments WHERE deleted_at IS NULL AND owner_id = $1`

	var respComments pb.AllComments
	rows, err := c.db.Query(query, ownerId.OwnerId)
	if err != nil {
		return nil, err
	}

	if respComments.Comments == nil {
		respComments.Comments = []*pb.Comment{}
	}
	for rows.Next() {
		var respComment = pb.Comment{}
		if err := rows.Scan(&respComment.Id,
			&respComment.Content,
			&respComment.OwnerId,
			&respComment.PostId,
			&respComment.CreatedAt); err != nil {
			return nil, err
		}

		respComments.Comments = append(respComments.Comments, &respComment)
	}

	return &respComments, nil
}
