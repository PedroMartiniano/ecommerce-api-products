default:
	cd cmd/api && go run .

docs:
	cd internal && swag init -g ../cmd/api/main.go

all:
	@make docs && make

new_migration:
	migrate create -ext sql -dir internal/infra/migrations -seq $(name)