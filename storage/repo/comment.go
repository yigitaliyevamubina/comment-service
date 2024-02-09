package repo

import pb "comment-service/genproto/comment_service"

//rpc CreateComment(Comment) returns (Comment);
//rpc GetCommentById(GetCommentId) returns (Comment);
//rpc GetAllCommentsByPostId(GetPostID) returns (AllComments);
//rpc GetAllCommentsByOwnerId(GetOwnerID) returns (AllComments);

type CommentStorageI interface {
	CreateComment(*pb.Comment) (*pb.Comment, error)
	GetCommentById(*pb.GetCommentId) (*pb.Comment, error)
	GetAllCommentsByPostId(*pb.GetPostID) (*pb.AllComments, error)
	GetAllCommentsByOwnerId(*pb.GetOwnerID) (*pb.AllComments, error)
}
