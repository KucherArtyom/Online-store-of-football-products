package repository

import (
	"Go-Kurs/config"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func NewPostgresDB() (*sql.DB, error) {
	log := logrus.New()

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.AppConfig.Database.Host,
		config.AppConfig.Database.Port,
		config.AppConfig.Database.User,
		config.AppConfig.Database.Password,
		config.AppConfig.Database.Name)

	log.WithFields(logrus.Fields{
		"host":   config.AppConfig.Database.Host,
		"port":   config.AppConfig.Database.Port,
		"user":   config.AppConfig.Database.User,
		"dbname": config.AppConfig.Database.Name,
	}).Info("Connecting to database")

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.WithError(err).Error("Failed to open database connection")
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.WithError(err).Error("Failed to ping database")
		return nil, err
	}

	log.Info("Successfully connected to database")
	return db, nil
}
