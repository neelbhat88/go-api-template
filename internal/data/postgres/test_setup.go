package postgres

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"sync"
)

const (
	TestDBName     = "testdb"
	TestDBPassword = "postgres"
	TestDBUser     = "postgres"
)

var (
	containerHost string
	containerPort int64

	initDBOnce sync.Once
)

func CreateTestDatabase() (*sqlx.DB, func()) {
	initDBOnce.Do(setupTestDatabaseContainer)

	// Update config to connect to default postgres DB first
	cfg := DatabaseConfig{
		Port:     containerPort,
		Host:     containerHost,
		Name:     "postgres",
		User:     TestDBUser,
		Password: TestDBPassword,
		SSLMode:  false,
	}
	db, err := ConnectPostgres(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to postgres DB")
	}
	defer db.Close()

	// Small hack according to blogpost above to close existing connections that might be hanging, because otherwise we can't create a new database from the template
	dropConnections(db, "template1")
	dropConnections(db, TestDBName)

	_, err = db.Exec(`
		drop database if exists testdb		
	`)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to drop testdb")
	}

	_, err = db.Exec(`create database testdb`)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create testdb")
	}

	// Now connect back to the testdb
	cfg.Name = TestDBName
	db, err = ConnectPostgres(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to testdb")
	}
	cleanup := func() {
		if err := db.Close(); err != nil {
			log.Error().Err(err).Msg("Failed to close testdb connection")
		}
	}

	return db, cleanup
}

func setupTestDatabaseContainer() {
	// 1. Create PostgreSQL container request
	containerReq := testcontainers.ContainerRequest{
		Image:        "postgres:16.3-alpine",
		ExposedPorts: []string{"5432/tcp"},
		WaitingFor:   wait.ForListeningPort("5432/tcp"),
		Env: map[string]string{
			"POSTGRES_DB":       TestDBName,
			"POSTGRES_PASSWORD": TestDBPassword,
			"POSTGRES_USER":     TestDBUser,
		},
	}

	// 2. Start PostgreSQL container
	dbContainer, _ := testcontainers.GenericContainer(
		context.Background(),
		testcontainers.GenericContainerRequest{
			ContainerRequest: containerReq,
			Started:          true,
		})

	// 3.1 Get host and port of PostgreSQL container
	host, _ := dbContainer.Host(context.Background())
	port, _ := dbContainer.MappedPort(context.Background(), "5432")

	// Do the trick found here: https://www.maragu.dk/blog/speeding-up-postgres-integration-tests-in-go
	// This will leverage Postgres's template database feature to speed up getting a clean new DB on each test
	templateCfg := DatabaseConfig{
		Port:     int64(port.Int()),
		Host:     host,
		Name:     "template1",
		User:     TestDBUser,
		Password: TestDBPassword,
		SSLMode:  false,
	}

	// Connect to template1 DB and run migrations
	_, cleanup, err := InitializeDB(templateCfg)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to initialize template1 DB")
	}
	defer cleanup()

	// Store to package level variables for use in CreateTestDatabase
	containerHost = host
	containerPort = int64(port.Int())
}

func dropConnections(db *sqlx.DB, name string) {
	query := `
		select pg_terminate_backend(pg_stat_activity.pid)
		from pg_stat_activity
		where pg_stat_activity.datname = $1 and pid <> pg_backend_pid()`
	_, err := db.Exec(query, name)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to drop connections")
	}
}
