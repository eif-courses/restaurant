package config

import "os"

type Config struct {
	DatabaseUrl           string
	BattleNetClientID     string
	BattleNetClientSecret string
	BattleNetRegion       string
	BattleNetBaseUrl      string
}

func NewConfig() *Config {
	return &Config{
		DatabaseUrl:           os.Getenv("DATABASE_URL"),
		BattleNetClientID:     os.Getenv("BATTLENET_CLIENT_ID"),
		BattleNetClientSecret: os.Getenv("BATTLENET_CLIENT_SECRET"),
		BattleNetRegion:       os.Getenv("BATTLENET_REGION"),
		BattleNetBaseUrl:      os.Getenv("BATTLENET_BASE_URL"),
	}
}
