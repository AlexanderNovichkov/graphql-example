.PHONY: generate migrateup migratedown

generate:
	go run github.com/99designs/gqlgen generate

migrateup:
	migrate -path migration -database "postgresql://user:password@localhost:5432/graphql?sslmode=disable" -verbose up

migratedown:
	migrate -path migration -database "postgresql://user:password@localhost:5432/graphql?sslmode=disable" -verbose down