package service

import (
	pb "comment-service/genproto/comment_service"
	l "comment-service/pkg/logger"
	"comment-service/storage"
	"context"
	"database/sql"
)

type CommentService struct {
	storage storage.IStorage
	logger  l.Logger
	pb.UnimplementedCommentServiceServer
}

func NewCommentService(db *sql.DB, log l.Logger) *CommentService {
	return &CommentService{
		storage: storage.NewStoragePg(db),
		logger:  log, UnimplementedCommentServiceServer: pb.UnimplementedCommentServiceServer{},
	}
}

//rpc CreateComment(Comment) returns (Comment);
//rpc GetAllCommentsByPostId(GetPostID) returns (AllComments);
//rpc GetAllCommentsByOwnerId(GetOwnerID) returns (AllComments);

func (c *CommentService) CreateComment(ctx context.Context, comment *pb.Comment) (*pb.Comment, error) {
	return c.storage.Comment().CreateComment(comment)
}

func (c *CommentService) GetCommentById(ctx context.Context, req *pb.GetCommentId) (*pb.Comment, error) {
	return c.storage.Comment().GetCommentById(req)
}

func (c *CommentService) GetAllCommentsByPostId(ctx context.Context, postId *pb.GetPostID) (*pb.AllComments, error) {
	return c.storage.Comment().GetAllCommentsByPostId(postId)
}

func (c *CommentService) GetAllCommentsByOwnerId(ctx context.Context, ownerId *pb.GetOwnerID) (*pb.AllComments, error) {
	return c.storage.Comment().GetAllCommentsByOwnerId(ownerId)
}
