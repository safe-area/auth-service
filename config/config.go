package config

import (
	"encoding/json"
	"os"
	"path"
)

type Config struct {
	JWTSecret string `json:"jwt_secret"`
	PgConfig  struct {
		Host string `json:"host"`
		Port string `json:"port"`
		User string `json:"user"`
		Pass string
		DB   string `json:"db"`
	} `json:"pg_config"`
}

// ParseConfig of service
func ParseConfig(configPath, secretPath string) (*Config, error) {
	fileBody, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = json.Unmarshal(fileBody, &cfg)
	if err != nil {
		return nil, err
	}

	jwtSecret, err := os.ReadFile(path.Join(secretPath, "jwt-secret"))
	if err != nil {
		return nil, err
	}
	cfg.JWTSecret = string(jwtSecret)

	pgPass, err := os.ReadFile(path.Join(secretPath, "postgres-password"))
	if err != nil {
		return nil, err
	}
	cfg.PgConfig.Pass = string(pgPass)

	return &cfg, nil
}
