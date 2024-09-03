package repository

import (
	"context"
	"database/sql"
	"errors"
	"learn-golang-database/entity"
	"strconv"
)

type CommentRepositoryImpl struct {
	Db *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &CommentRepositoryImpl{Db: db}

}

func (repository *CommentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	script := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	result, err := repository.Db.ExecContext(ctx, script, comment.Email, comment.Comment)
	if err != nil {
		return comment, err
	}
	id, err := result.LastInsertId()
	comment.Id = int32(id)
	return comment, nil
}

func (repository *CommentRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Comment, error) {
	script := "SELECT id, email, comment From comments WHERE id = ?"
	comment := entity.Comment{}
	rows, err := repository.Db.QueryContext(ctx, script, id)
	if err != nil {
		return comment, err
	}
	if rows.Next() {
		// ada
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		return comment, nil
	} else {
		// tidak ada
		return comment, errors.New("Id" + strconv.Itoa(int(id)) + "not foung")
	}
}

func (repository *CommentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	script := "SELECT id, email, comment FROM comments"
	rows, err := repository.Db.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var comments []entity.Comment
	for rows.Next() {
		comment := entity.Comment{}
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		comments = append(comments, comment)
	}
	return comments, nil
}
