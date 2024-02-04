package storage

import (
	"comment-service/storage/postgres"
	"comment-service/storage/repo"
	"database/sql"
)

type IStorage interface {
	Comment() repo.CommentStorageI
}

type storagePg struct {
	db          *sql.DB
	commentRepo repo.CommentStorageI
}

func NewStoragePg(db *sql.DB) *storagePg {
	return &storagePg{
		db:          db,
		commentRepo: postgres.NewCommentRepo(db),
	}
}

func (s storagePg) Comment() repo.CommentStorageI {
	return s.commentRepo
}
