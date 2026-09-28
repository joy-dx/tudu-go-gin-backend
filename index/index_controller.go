package index

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Default endpoint for the service to be used with health checks
func Index(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}
