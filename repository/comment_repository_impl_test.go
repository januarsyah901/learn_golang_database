package repository

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	learn_golang_database "learn-golang-database"
	"learn-golang-database/entity"
	"testing"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(learn_golang_database.GetConnection())
	ctx := context.Background()
	comment := entity.Comment{Email: "masjanu@gmail.com", Comment: "apkaha benar ada ikan berkepala lele?"}
	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
func TestFindById(t *testing.T) {
	commentRepository := NewCommentRepository(learn_golang_database.GetConnection())
	comment, err := commentRepository.FindById(context.Background(), 222)
	if err != nil {
		panic(err)
	}
	fmt.Println(comment)
}
func TestFindByAll(t *testing.T) {
	commentRepository := NewCommentRepository(learn_golang_database.GetConnection())
	comments, err := commentRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}
	for _, comment := range comments {
		fmt.Println(comment)
	}
}
