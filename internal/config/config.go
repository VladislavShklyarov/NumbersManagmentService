package config

import (
	"github.com/spf13/viper"
	"strings"
)

func Load() (*Config, error) {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AddConfigPath("./configs")
	viper.AddConfigPath(".")

	// ENV
	viper.AutomaticEnv()
	viper.SetEnvPrefix("PHONE")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// defaults
	viper.SetDefault("env", "local")
	viper.SetDefault("http.port", 8080)
	viper.SetDefault("http.host", "0.0.0.0")
	viper.SetDefault("storage.type", "memory")

	if err := viper.ReadInConfig(); err != nil {
		// не критично — можем жить на env
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

type Config struct {
	Env     string        `mapstructure:"env"`
	HTTP    HTTPConfig    `mapstructure:"http"`
	Storage StorageConfig `mapstructure:"storage"`
	Phone   PhoneConfig   `mapstructure:"phone"`
}

type HTTPConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type StorageConfig struct {
	Type     string         `mapstructure:"type"` // memory | postgres
	Postgres PostgresConfig `mapstructure:"postgres"`
}

type PostgresConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
	SSLMode  string `mapstructure:"sslmode"`
}

type PhoneConfig struct {
	DefaultCountryCode string `mapstructure:"default_country_code"`
}
