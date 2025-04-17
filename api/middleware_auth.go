package api

import (
	"fmt"
	"github.com/symball/go-gin-boilerplate/auth"
	. "github.com/symball/go-gin-boilerplate/config"
	"github.com/symball/go-gin-boilerplate/users"
	"log"
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

var authMiddleware *jwt.GinJWTMiddleware

// Initiate a middleware instance and return reference for use in router
func MiddlewareAuthInit() *jwt.GinJWTMiddleware {

	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       AppConfig.AuthRealm,
		Key:         []byte(AppConfig.AuthKey),
		Timeout:     time.Hour * AppConfig.AuthSessionLength,
		MaxRefresh:  time.Hour * 24,
		IdentityKey: AppConfig.AuthIdentityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*users.User); ok {
				return jwt.MapClaims{
					AppConfig.AuthIdentityKey: v.Username,
					"roles":                   v.Roles,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &users.User{
				Username: claims[AppConfig.AuthIdentityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals auth.LoginRequest
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			username := loginVals.Username
			password := loginVals.Password

			user, err := users.GetOneByUsername(c, username)
			if err != nil {
				log.Print(err.Error())
				return nil, jwt.ErrFailedAuthentication
			}

			if user.Status != users.Active {
				return nil, jwt.ErrFailedAuthentication
			}

			fmt.Printf("Testing user: %s", user.Username)
			pass, err := auth.CheckPassword(password, user.Password)
			if err != nil {
				log.Print(err.Error())
				return nil, jwt.ErrFailedAuthentication
			}
			if pass {
				return &users.User{
					Id:       user.Id,
					Username: username,
					Roles:    user.Roles,
				}, nil
			}
			return nil, jwt.ErrFailedAuthentication
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		LoginResponse: func(c *gin.Context, code int, token string, responseTime time.Time) {
			c.JSON(http.StatusOK, gin.H{
				"code":   http.StatusOK,
				"token":  token,
				"expire": responseTime.Format(time.RFC3339),
			})
		},
		LogoutResponse: func(c *gin.Context, code int) {
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  "logout success",
			})
		},
		RefreshResponse: func(c *gin.Context, code int, token string, responseTime time.Time) {
			c.JSON(http.StatusOK, gin.H{
				"code":   http.StatusOK,
				"token":  token,
				"expire": responseTime.Format(time.RFC3339),
			})
		},
		SendCookie:     true,
		SecureCookie:   AppConfig.AuthCookieSecure,
		CookieHTTPOnly: AppConfig.AuthCookieSecure,
		CookieName:     "token",
		CookieSameSite: http.SameSiteStrictMode, //SameSiteDefaultMode, SameSiteLaxMode, SameSiteStrictMode, SameSiteNoneMode
		// Define places in request authentication accepted
		TokenLookup:   "cookie:token, header:Authorization, query:token",
		TokenHeadName: AppConfig.AuthHeaderKey,
		TimeFunc:      time.Now,
	})
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
	errInit := authMiddleware.MiddlewareInit()
	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}
	return authMiddleware
}
