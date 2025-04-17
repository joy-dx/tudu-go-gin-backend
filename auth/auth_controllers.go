package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/symball/go-gin-boilerplate/lib"
	"net/http"
)

func RegisterPost(c *gin.Context) {
	var registerForm RegisterRequest
	if errors := c.ShouldBindJSON(&registerForm); errors != nil {
		outPutData := lib.HandleValidationErrors(errors)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": outPutData})
	}

}
