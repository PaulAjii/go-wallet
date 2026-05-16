package config

import (
	"log"
	"os"

	"github.com/PaulAjii/go-wallet/pkg/sysmsg"
	"github.com/joho/godotenv"
)

type Config struct {
	App      AppConfig
	Database DatabaseConfig
	Supabase SupabaseConfig
	Resend   ResendConfig
}

type AppConfig struct {
	Env            string
	Port           string
	AllowedOrigins string
	FrontendURL    string
}

type DatabaseConfig struct {
	URL string
}

type SupabaseConfig struct {
	ProjectRef     string
	URL            string
	AnonKey        string
	ServiceRoleKey string
	JWTSecret      string
}

type ResendConfig struct {
	APIKey    string
	FromEmail string
}

var ApplicationConfig *Config

func Load() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("%s: %v", sysmsg.NoEnvFile, err)
	}

	ApplicationConfig = &Config{
		AppConfig{
			Env:            os.Getenv("APP_ENV"),
			Port:           os.Getenv("APP_PORT"),
			AllowedOrigins: os.Getenv("APP_ALLOWED_ORIGINS"),
			FrontendURL:    os.Getenv("APP_FRONTEND_URL"),
		},
		DatabaseConfig{
			URL: os.Getenv("DATABASE_URI"),
		},
		SupabaseConfig{
			ProjectRef:     os.Getenv("SUPABASE_PROJECT_REF"),
			URL:            os.Getenv("SUPABASE_URL"),
			AnonKey:        os.Getenv("SUPABASE_ANON_KEY"),
			ServiceRoleKey: os.Getenv("SUPABASE_SERVICE_ROLE_KEY"),
			JWTSecret:      os.Getenv("SUPABASE_JWT_SECRET"),
		},
		ResendConfig{
			APIKey:    os.Getenv("RESEND_API_KEY"),
			FromEmail: os.Getenv("RESEND_FROM_EMAIL"),
		},
	}
}
