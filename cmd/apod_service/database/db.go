package database

import (
	"context"
	"fmt"
	"time"

	"github.com/Marsredskies/apod_service/envconfig"
	"github.com/jmoiron/sqlx"
)

type DB struct {
	db *sqlx.DB
}

func New(dbconfig envconfig.Database) (*DB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dbEngine, err := ConnectDB(ctx, dbconfig)
	if err != nil {
		return nil, err
	}

	return &DB{
		db: dbEngine,
	}, nil
}

func ConnectDB(ctx context.Context, dbconfig envconfig.Database) (*sqlx.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbconfig.Host, dbconfig.Port, dbconfig.Username, dbconfig.Password, dbconfig.DBName)

	db, err := sqlx.Open(dbconfig.DBName, connStr)
	if err != nil {
		return nil, err
	}

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxIdleTime(10 * time.Second)
	db.SetConnMaxLifetime(10 * time.Second)
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(5)

	return db, nil
}

func RequireNewDBClient(ctx context.Context, dbconfig envconfig.Database) *DB {
	db, err := New(dbconfig)
	if err != nil {
		panic(fmt.Errorf("failed to initialise db client: %v", err))
	}
	return db
}
