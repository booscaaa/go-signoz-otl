package postgres

import (
	"context"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"github.com/uptrace/opentelemetry-go-extra/otelsql"
	"github.com/uptrace/opentelemetry-go-extra/otelsqlx"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

// GetConnection return connection pool from postgres drive SQLX
func GetConnection(context context.Context, provider *sdktrace.TracerProvider) *sqlx.DB {
	databaseURL := viper.GetString("database.url")

	db, err := otelsqlx.ConnectContext(
		context,
		"postgres",
		databaseURL,
		otelsql.WithAttributes(semconv.DBSystemPostgreSQL),
		otelsql.WithTracerProvider(provider),
	)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

// RunMigrations run scripts on path database/migrations
func RunMigrations() {
	databaseURL := viper.GetString("database.url")
	m, err := migrate.New("file://database/migrations", databaseURL)
	if err != nil {
		log.Println(err)
	}

	if err := m.Up(); err != nil {
		log.Println(err)
	}
}
