package config

import (
	"Golang-crud/helper"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" //postgress
	"github.com/rs/zerolog/log"
)

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "postgress"
	dbName = "test"
)
func DatabaseConnection() *sql.DB{
	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	db, err := sql.Open("postgres", sqlInfo)
	helper.PanicIfError(err)

	err = db.Ping()
	helper.PanicIfError(err)

	log.Info().Msg("Connected to database!!")

	return db
}