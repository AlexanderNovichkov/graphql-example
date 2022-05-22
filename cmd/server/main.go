package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/AlexanderNovichkov/graphql-example/internal/db/repository"
	"github.com/AlexanderNovichkov/graphql-example/internal/graph"
	"github.com/AlexanderNovichkov/graphql-example/internal/graph/gqlgenerated"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func GetEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func migrateDatabase(db *sqlx.DB) {
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://migration",
		"postgres", driver)
	failOnError(err, "Failed to create migration")

	err = m.Up()
	if err != migrate.ErrNoChange {
		failOnError(err, "Failed to migrate database")
	}
}

var graphqlPort = GetEnvOrDefault("PORT", "8080")
var database = GetEnvOrDefault("DATABASE",
	"host=localhost port=5432 user=user password=password dbname=graphql sslmode=disable")

func main() {
	db, err := sqlx.Open("postgres", database)
	failOnError(err, "Failed to connect to database: "+database)
	failOnError(db.Ping(), "Failed to ping database: "+database)
	defer db.Close()

	migrateDatabase(db)

	repo := repository.NewRepository(db)

	srv := handler.NewDefaultServer(
		gqlgenerated.NewExecutableSchema(
			gqlgenerated.Config{
				Resolvers: &graph.Resolver{Repository: repo},
			},
		),
	)

	http.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	http.Handle("/graphql", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", graphqlPort)
	log.Fatal(http.ListenAndServe(":"+graphqlPort, nil))
}
