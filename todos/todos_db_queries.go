package todos

import (
	"github.com/gin-gonic/gin"
	"github.com/symball/go-gin-boilerplate/storage"
)

func Add(ctx *gin.Context, todo *Todo) (*Todo, error) {
	_, err := storage.DBGet().NewInsert().Model(todo).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

// Remove todo record from persistent storage
func Delete(ctx *gin.Context, todo *Todo) error {
	_, err := storage.DBGet().NewDelete().Model(todo).WherePK().Exec(ctx)
	return err
}

// Retrieve a single Todo record from the DB
func GetOneById(ctx *gin.Context, TodoId int64) (*Todo, error) {
	todo := new(Todo)
	err := storage.DBGet().NewSelect().Model(todo).Where("id = ?", TodoId).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

// Retrieve all Todo records from the DB
func GetAll(ctx *gin.Context) ([]Todo, error) {
	var todos []Todo
	err := storage.DBGet().NewSelect().Model(&todos).Order("id ASC").Scan(ctx)
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func Update(ctx *gin.Context, todo *Todo) (*Todo, error) {
	_, err := storage.DBGet().NewUpdate().Model(todo).WherePK().Exec(ctx)
	if err != nil {
		return nil, err
	}
	return todo, nil
}
