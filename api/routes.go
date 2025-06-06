package api

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/symball/go-gin-boilerplate/index"
	"github.com/symball/go-gin-boilerplate/todos"
	"net/http"
)

var router *gin.Engine

// Initiate a GIN router with site routes

func NewRouter(authMiddleware *jwt.GinJWTMiddleware, corsMiddleware gin.HandlerFunc) *gin.Engine {

	var routes = Routes{
		{
			"Index",
			http.MethodGet,
			"/",
			index.Index,
			false,
		},
		{
			"TodosGet",
			http.MethodGet,
			"/todos",
			todos.TodosGet,
			true,
		},
		{
			"TodosIdGet",
			http.MethodGet,
			"/todos/:TodoId",
			todos.TodosGetById,
			true,
		},
		{
			"TodosIdDelete",
			http.MethodDelete,
			"/todos/:TodoId",
			todos.TodosDeleteById,
			true,
		},
		{
			"TodosPost",
			http.MethodPost,
			"/todos",
			todos.TodosPost,
			true,
		},
		{
			"TodosPut",
			http.MethodPut,
			"/todos/:TodoId",
			todos.TodosPutById,
			true,
		},
		{
			"LoginPost",
			http.MethodPost,
			"/login",
			authMiddleware.LoginHandler,
			false,
		},
		{
			"LogoutGet",
			http.MethodGet,
			"/logout",
			authMiddleware.LogoutHandler,
			true,
		},
		{
			"authRefresh",
			http.MethodGet,
			"/refresh",
			authMiddleware.RefreshHandler,
			true,
		},
	}

	router := gin.Default()
	router.Use(corsMiddleware)
	secureRouter := router.Group("/", authMiddleware.MiddlewareFunc())
	for _, route := range routes {

		if route.Secure {
			switch route.Method {
			case http.MethodGet:
				secureRouter.GET(route.Pattern, route.HandlerFunc)
			case http.MethodPost:
				secureRouter.POST(route.Pattern, route.HandlerFunc)
			case http.MethodPut:
				secureRouter.PUT(route.Pattern, route.HandlerFunc)
			case http.MethodPatch:
				secureRouter.PATCH(route.Pattern, route.HandlerFunc)
			case http.MethodDelete:
				secureRouter.DELETE(route.Pattern, route.HandlerFunc)
			}
		} else {
			switch route.Method {
			case http.MethodGet:
				router.GET(route.Pattern, route.HandlerFunc)
			case http.MethodPost:
				router.POST(route.Pattern, route.HandlerFunc)
			case http.MethodPut:
				router.PUT(route.Pattern, route.HandlerFunc)
			case http.MethodPatch:
				router.PATCH(route.Pattern, route.HandlerFunc)
			case http.MethodDelete:
				router.DELETE(route.Pattern, route.HandlerFunc)
			}
		}

	}

	return router
}

// Endpoint definition for the router to initate with
type Route struct {
	// Name is the name of this Route.
	Name string
	// Method is the string for the HTTP method. ex) GET, POST etc..
	Method string
	// Pattern is the pattern of the URI.
	Pattern string
	// HandlerFunc is the handler function of this route.
	HandlerFunc gin.HandlerFunc
	// Boolean is flag to determine whether route is protected
	Secure bool
}

// Routes is the list of the generated Route.
type Routes []Route
