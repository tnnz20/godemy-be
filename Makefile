run:
	go run cmd/main.go

migrateup:
	migrate -path db/migrations -database "postgres://postgres:qwe456zxc@localhost:5432/godemy_test?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgres://postgres:qwe456zxc@localhost:5432/godemy_test?sslmode=disable" -verbose down

.PHONY: run, migrateup, migratedown