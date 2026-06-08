package database

import (
	"fmt"
	"log"

	"server/pkg/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgres(cfg *config.Config) *gorm.DB {

	pg := cfg.Postgres

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		pg.Host,
		pg.Port,
		config.GetEnv("POSTGRES_USER"),
		config.GetEnv("POSTGRES_PASSWORD"),
		pg.Database,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("postgres connect failed:", err)
	}

	log.Println("postgres connected")

	return db
}
