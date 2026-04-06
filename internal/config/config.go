package config

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

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
	Type     string         `mapstructure:"type"`
	Postgres PostgresConfig `mapstructure:"postgres"`
}

type PostgresConfig struct {
	Migrations MigrationsConfig `mapstructure:"migrations"`
	Master     DBConfig         `mapstructure:"master"`
	Slave      DBConfig         `mapstructure:"slave"`
}

type MigrationsConfig struct {
	Enabled bool   `mapstructure:"enabled"`
	Path    string `mapstructure:"path"`
}

type DBConfig struct {
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

func Load() (*Config, error) {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AddConfigPath("./configs")
	viper.AddConfigPath(".")

	// ENV
	viper.SetEnvPrefix("NUMBERS")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.SetDefault("env", "local")

	viper.SetDefault("http.host", "0.0.0.0")
	viper.SetDefault("http.port", 8080)

	viper.SetDefault("storage.type", "postgres")

	viper.SetDefault("storage.postgres.migrations.enabled", true)
	viper.SetDefault("storage.postgres.migrations.path", "./migrations")

	viper.SetDefault("storage.postgres.master.port", 5432)
	viper.SetDefault("storage.postgres.slave.port", 5432)

	viper.SetDefault("storage.postgres.master.sslmode", "disable")
	viper.SetDefault("storage.postgres.slave.sslmode", "disable")

	_ = viper.ReadInConfig()

	cfg := &Config{
		Env: viper.GetString("env"),

		HTTP: HTTPConfig{
			Host: viper.GetString("http.host"),
			Port: viper.GetInt("http.port"),
		},

		Storage: StorageConfig{
			Type: viper.GetString("storage.type"),

			Postgres: PostgresConfig{
				Migrations: MigrationsConfig{
					Enabled: viper.GetBool("storage.postgres.migrations.enabled"),
					Path:    viper.GetString("storage.postgres.migrations.path"),
				},

				Master: DBConfig{
					Host:     viper.GetString("storage.postgres.master.host"),
					Port:     viper.GetInt("storage.postgres.master.port"),
					User:     viper.GetString("storage.postgres.master.user"),
					Password: viper.GetString("storage.postgres.master.password"),
					DBName:   viper.GetString("storage.postgres.master.dbname"),
					SSLMode:  viper.GetString("storage.postgres.master.sslmode"),
				},

				Slave: DBConfig{
					Host:     viper.GetString("storage.postgres.slave.host"),
					Port:     viper.GetInt("storage.postgres.slave.port"),
					User:     viper.GetString("storage.postgres.slave.user"),
					Password: viper.GetString("storage.postgres.slave.password"),
					DBName:   viper.GetString("storage.postgres.slave.dbname"),
					SSLMode:  viper.GetString("storage.postgres.slave.sslmode"),
				},
			},
		},

		Phone: PhoneConfig{
			DefaultCountryCode: viper.GetString("phone.default_country_code"),
		},
	}

	if cfg.Storage.Postgres.Master.Host == "" {
		return nil, fmt.Errorf("postgres master host is empty")
	}

	return cfg, nil
}
