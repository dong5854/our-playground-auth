package config

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"entgo.io/ent/examples/fs/ent"

	"github.com/Team-OurPlayground/our-playground-auth/internal/util/customerror"
)

var db *sql.DB
var entClient *ent.Client

func CreateConnectionString(dbName string, host string, port string, userName string, password string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		userName,
		password,
		host,
		port,
		dbName)
}

func MysqlInitialize() error {
	connStr := CreateConnectionString(
		GetEnv("DB_NAME"),
		GetEnv("DB_HOST"),
		GetEnv("DB_PORT"),
		GetEnv("DB_USER"),
		GetEnv("DB_PW"))

	var err error
	db, err = sql.Open("mysql", connStr)
	if err != nil {
		return customerror.Wrap(err, customerror.DBConnection, "failed to connect to database")
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetConnMaxIdleTime(time.Minute * 3)
	err = db.Ping()
	if err != nil {
		return customerror.Wrap(err, customerror.DBConnection, "failed to ping the database")
	}

	return nil
}

func GetDBInstance() *sql.DB {
	if db == nil {
		log.Panic(customerror.New(customerror.DBConnection, "DB has not been initialized"))
	}
	return db
}

func EntMysqlInitialize() {
	connStr := CreateConnectionString(
		GetEnv("DB_NAME"),
		GetEnv("DB_HOST"),
		GetEnv("DB_PORT"),
		GetEnv("DB_USER"),
		GetEnv("DB_PW"))

	var err error
	entClient, err = ent.Open("mysql", connStr)
	if err != nil {
		log.Fatal(customerror.Wrap(err, customerror.DBConnection, "failed to connect to database"))
	}

	if err = entClient.Schema.Create(context.Background()); err != nil {
		log.Fatal(customerror.Wrap(err, customerror.DBConnection, "failed creating schema resources"))
	}
}

func GetEntClient() *ent.Client {
	if entClient == nil {
		log.Panic(customerror.New(customerror.DBConnection, "entClient has not been initialized"))
	}
	return entClient
}
