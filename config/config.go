package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
)

type AppConfigSpec struct {
	AuthArgonSaltLength  uint32        `mapstructure:"auth_argon_salt_length"`
	AuthArgonMemory      uint32        `mapstructure:"auth_argon_memory"`
	AuthArgonIterations  uint32        `mapstructure:"auth_argon_iterations"`
	AuthArgonParallelism uint8         `mapstructure:"auth_argon_parallelism"`
	AuthArgonKeyLength   uint32        `mapstructure:"auth_argon_key_length"`
	AuthCookieSecure     bool          `mapstructure:"auth_cookie_secure"`
	AuthRealm            string        `mapstructure:"auth_realm"`
	AuthKey              string        `mapstructure:"auth_key"`
	AuthIdentityKey      string        `mapstructure:"auth_identity_key"`
	AuthHeaderKey        string        `mapstructure:"auth_header_key"`
	AuthSessionLength    time.Duration `mapstructure:"auth_session_length"`
	CorsAllowedOrigins   []string      `mapstructure:"cors_allowed_origins"`
	PostgresDSN          string        `mapstructure:"postgres_dsn"`
	Port                 string        `mapstructure:"port"`
}

var AppConfig AppConfigSpec

func ConfigInit(cfgFile string) {

	fmt.Print("Loading Viper config")
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		if home, err := os.UserHomeDir(); err == nil {
			// Search config in home directory with name ".cobra" (without extension).
			viper.AddConfigPath(home)
			viper.SetConfigType("yaml")
			viper.SetConfigName(".cobra")
		}
	}

	viper.AutomaticEnv()
	viper.SetEnvPrefix("tudu")

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	err := viper.Unmarshal(&AppConfig)
	if err != nil {
		log.Fatal("unable to decode into struct, %v", err)
	}
}
