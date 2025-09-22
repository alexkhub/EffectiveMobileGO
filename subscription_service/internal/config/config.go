package config

import (
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	Postgres PostgresCFG `mapstructure:"postgres"`
	Port     int         `mapstructure:"server_port"`
	LogLevel string      `mapstructure:"log_level"`
	LogFile  string      `mapstructure:"log_file"`
}

type PostgresCFG struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Dbname   string `mapstructure:"dbname"`
	Sslmode  string `mapstructure:"sslmode"`
}

func Load() (*Config, error) {
	_ = godotenv.Load(".env")
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigFile("configs/config.yaml")
	_ = v.ReadInConfig()

	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	_ = v.BindEnv("postgres.user", "POSTGRES_USER")
	_ = v.BindEnv("postgres.password", "POSTGRES_PASSWORD")
	_ = v.BindEnv("postgres.dbname", "POSTGRES_DB")

	cfg := &Config{}
	if err := v.Unmarshal(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
