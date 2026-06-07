package config

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
type Config struct {
	ServerPort         string   `mapstructure:"PORT"`
	MongoURI           string   `mapstructure:"MONGO_URI"`
	DBName             string   `mapstructure:"DB_NAME"`
	JWTSecretKey       string   `mapstructure:"JWT_SECRET_KEY"`
	JWTExpirationHours int      `mapstructure:"JWT_EXPIRATION_HOURS"`
	EnableCache        bool     `mapstructure:"ENABLE_CACHE"`
	RedisAddr          string   `mapstructure:"REDIS_ADDR"`
	RedisPassword      string   `mapstructure:"REDIS_PASSWORD"`
	LogLevel           string   `mapstructure:"LOG_LEVEL"`
	LogFormat          string   `mapstructure:"LOG_FORMAT"`
	CookieDomains      []string `mapstructure:"COOKIE_DOMAINS"`
	SecureCookie       bool     `mapstructure:"SECURE_COOKIE"`
	AllowedOrigins     []string `mapstructure:"ALLOWED_ORIGINS"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AutomaticEnv()

	// defaults
	viper.SetDefault("PORT", "8080")
	viper.SetDefault("ENABLE_CACHE", false)
	viper.SetDefault("JWT_EXPIRATION_HOURS", 72)
	viper.SetDefault("ALLOWED_ORIGINS", "http://localhost:5173")
	viper.SetDefault("COOKIE_DOMAINS", "localhost")

	// explicit env binding (IMPORTANT)
	viper.BindEnv("MONGO_URI")
	viper.BindEnv("DB_NAME")
	viper.BindEnv("JWT_SECRET_KEY")
	viper.BindEnv("JWT_EXPIRATION_HOURS")
	viper.BindEnv("ENABLE_CACHE")
	viper.BindEnv("REDIS_ADDR")
	viper.BindEnv("REDIS_PASSWORD")
	viper.BindEnv("LOG_LEVEL")
	viper.BindEnv("LOG_FORMAT")
	viper.BindEnv("COOKIE_DOMAINS")
	viper.BindEnv("SECURE_COOKIE")
	viper.BindEnv("ALLOWED_ORIGINS")

	// only load local file in dev (optional)
	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return
	}

	// FIXED: parsing AFTER unmarshal (inside function!)
	if allowedOrigins := viper.GetString("ALLOWED_ORIGINS"); allowedOrigins != "" {
		parts := strings.Split(allowedOrigins, ",")

		var cleaned []string
		for _, p := range parts {
			p = strings.TrimSpace(p)
			p = strings.Trim(p, "\"'")
			if p != "" {
				cleaned = append(cleaned, p)
			}
		}

		config.AllowedOrigins = cleaned
	}

	if cookieDomains := viper.GetString("COOKIE_DOMAINS"); cookieDomains != "" {
		parts := strings.Split(cookieDomains, ",")

		var cleaned []string
		for _, p := range parts {
			p = strings.TrimSpace(p)
			p = strings.Trim(p, "\"'")
			if p != "" {
				cleaned = append(cleaned, p)
			}
		}

		config.CookieDomains = cleaned
	}
	return
}
