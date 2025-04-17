package users

import (
	"github.com/gin-gonic/gin"
	"github.com/symball/go-gin-boilerplate/storage"
)

func GetOneByUsername(ctx *gin.Context, username string) (*User, error) {
	user := new(User)
	err := storage.DBGet().NewSelect().Model(user).Where("username = ?", username).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}
