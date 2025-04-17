package index

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Default endpoint for the service to be used with health checks
func Index(c *gin.Context) {
	c.String(http.StatusOK, "OK!")
}
