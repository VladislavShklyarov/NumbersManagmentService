package config

import "os"

func LoadDBConfig() PostgresConfig {
	return PostgresConfig{
		Master: DBConfig{
			Host:     os.Getenv("NUMBERS_STORAGE_POSTGRES_MASTER_HOST"),
			Port:     5432,
			User:     os.Getenv("NUMBERS_STORAGE_POSTGRES_MASTER_USER"),
			Password: os.Getenv("NUMBERS_STORAGE_POSTGRES_MASTER_PASSWORD"),
			DBName:   os.Getenv("NUMBERS_STORAGE_POSTGRES_MASTER_DBNAME"),
			SSLMode:  "disable",
		},
		Slave: DBConfig{
			Host:     os.Getenv("NUMBERS_STORAGE_POSTGRES_SLAVE_HOST"),
			Port:     5432,
			User:     os.Getenv("NUMBERS_STORAGE_POSTGRES_SLAVE_USER"),
			Password: os.Getenv("NUMBERS_STORAGE_POSTGRES_SLAVE_PASSWORD"),
			DBName:   os.Getenv("NUMBERS_STORAGE_POSTGRES_SLAVE_DBNAME"),
			SSLMode:  "disable",
		},
	}
}
