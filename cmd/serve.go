package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/symball/go-gin-boilerplate/api"
	"github.com/symball/go-gin-boilerplate/config"
	"github.com/symball/go-gin-boilerplate/storage"
	"log"
	"time"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{

	Use:   "serve",
	Short: "Serve the API",
	Long:  `A Gin based API server`,
	Run: func(cmd *cobra.Command, args []string) {
		storage.DBInit()

		// Prepare auth middleware
		authMiddleware := api.MiddlewareAuthInit()
		corsMiddleware := api.MiddlewareCorsInit()

		// Prepare router
		router := api.NewRouter(authMiddleware, corsMiddleware)
		log.Fatal(router.Run(":" + config.AppConfig.Port))
	},
}

func init() {
	serveCmd.Flags().StringP("port", "p", "8080", "API port if not using env.PORT")
	serveCmd.Flags().String("auth_realm", "development zone", "Security Realm to use with JWT")
	serveCmd.Flags().Bool("auth_cookie_secure", false, "Whether the auth cookie cannot be read by JS")
	serveCmd.Flags().String("auth_key", "Secret Key", "Secret key to use as salt as part of JWT auth")
	serveCmd.Flags().String("auth_identity_key", "username", "Identity key to use as part of JWT auth")
	serveCmd.Flags().String("auth_header_key", "Bearer", "Prefix to use with the Authorization header")
	serveCmd.Flags().Duration("auth_session_length", time.Duration(3), "Session Length in hours")
	serveCmd.Flags().StringSliceP("cors_allowed_origins", "o", []string{"http://localhost:5173"}, "List of allowed origins for CORS requests")

	viper.BindPFlag("port", serveCmd.Flags().Lookup("port"))
	viper.BindPFlag("auth_cookie_secure", serveCmd.Flags().Lookup("auth_cookie_secure"))
	viper.BindPFlag("auth_realm", serveCmd.Flags().Lookup("auth_realm"))
	viper.BindPFlag("auth_key", serveCmd.Flags().Lookup("auth_key"))
	viper.BindPFlag("auth_identity_key", serveCmd.Flags().Lookup("auth_identity_key"))
	viper.BindPFlag("auth_header_key", serveCmd.Flags().Lookup("auth_header_key"))
	viper.BindPFlag("auth_header_key", serveCmd.Flags().Lookup("auth_header_key"))
	viper.BindPFlag("auth_session_length", serveCmd.Flags().Lookup("auth_session_length"))
	viper.BindPFlag("cors_allowed_origins", serveCmd.Flags().Lookup("cors_allowed_origins"))

	rootCmd.AddCommand(serveCmd)

}
