package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	. "github.com/symball/go-gin-boilerplate/config"
)

// Initiate a middleware instance and return reference for use in router
func MiddlewareCorsInit() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowCredentials = true
	config.AllowMethods = []string{"OPTIONS", "DELETE", "HEAD", "GET", "POST", "PUT"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Credentials", "Access-Control-Allow-Headers"}
	config.ExposeHeaders = []string{"Content-Length", "Content-Type", "Set-Cookie", "Access-Control-Allow-Origin", "Access-Control-Allow-Credentials", "Access-Control-Allow-Headers"}

	if AppConfig.CorsAllowedOrigins[0] == "*" {
		config.AllowAllOrigins = true
	} else {
		config.AllowOrigins = AppConfig.CorsAllowedOrigins
	}
	corsMiddleware := cors.New(config)
	return corsMiddleware
}
