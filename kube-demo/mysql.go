package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log/slog"
	"os"
	"time"
)

type DatabaseConfig struct {
	username string
	password string
	hostname string
	port     string
	dbName   string
}

var (
	databaseConn *sql.DB
)

func connectDatabase() error {
	slog.Info("Trying to connect to DB")
	db, err := sql.Open("mysql", createDSN(true))
	if err != nil {
		return fmt.Errorf("failed to open mysql connection: %w", err)
	}

	databaseConn = db

	if err := createDatabase(os.Getenv("DATABASE_NAME")); err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return fmt.Errorf("failed to ping db: %w", err)
	}

	db, err = sql.Open("mysql", createDSN(false))
	if err != nil {
		return fmt.Errorf("failed to open mysql connection using database name: %w", err)
	}

	slog.Info("connected to database")
	databaseConn = db

	return nil
}

func createDatabase(dbName string) error {
	slog.Info("Creating database")
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	tx, err := databaseConn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	res, err := tx.ExecContext(ctx, fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbName))
	if err != nil {
		return err
	}

	no, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if no == 0 {
		return errors.New("failed to create database, no rows affected")
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit tx: %w", err)
	}
	return nil
}

func createDSN(skipDB bool) string {
	dbCfg := getDatabaseConfig()
	if skipDB {
		return fmt.Sprintf("%s:%s@tcp(%s)/%s", dbCfg.username, dbCfg.password, dbCfg.hostname, "")
	}
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", dbCfg.username, dbCfg.password, dbCfg.hostname, dbCfg.dbName)
}

func getDatabaseConfig() DatabaseConfig {
	return DatabaseConfig{
		username: os.Getenv("DATABASE_USERNAME"),
		password: os.Getenv("DATABASE_PASSWORD"),
		hostname: os.Getenv("MYSQL_SERVICE_HOST"),
		port:     os.Getenv("MYSQL_SERVICE_PORT"),
		dbName:   os.Getenv("DATABASE_NAME"),
	}
}
