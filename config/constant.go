package config

import "os"

var (
	JWT_SECRET_KEY = os.Getenv("SECRET_KEY")
	JWT_ISSUER     = "dating-apps.com"
	PORT           = ":8080"
	DATABASE       = os.Getenv("DATABASE_URL")
	SALT_PASSWORD  = os.Getenv("SALT_PASSWORD")
)
