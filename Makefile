include .env
export

# make proto file=delivery.proto
proto:
	protoc --go_out=. --go_opt=paths=source_relative \
        --go-grpc_out=. --go-grpc_opt=paths=source_relative \
        internal/model/proto/$(file)

migrate-create:  ### create new migration: make migrate-create name=add_table_xyz
	migrate create -ext sql -dir migrations $(name)

# migrate-up num=1
migrate-up: ### migration up
	migrate -path migrations -database '$(PG_URL)?sslmode=disable' up $(num)

sql-gen:
	go generate

run:
	go run cmd/app/main.go