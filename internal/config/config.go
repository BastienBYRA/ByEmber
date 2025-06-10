package config

import (
	"database/sql"

	"bastienbyra.fr/bastienbyra/ByEmber/utils"
)

type Config struct {
	DB                *sql.DB
	EncryptionService *utils.EncryptionService
}

func NewConfig(db *sql.DB, encryptionService *utils.EncryptionService) *Config {
	return &Config{
		DB:                db,
		EncryptionService: encryptionService,
	}
}
