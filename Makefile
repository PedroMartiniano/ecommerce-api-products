default:
	cd cmd/api && go run .

docs:
	cd internal && swag init -g ../cmd/api/main.go

new_migration:
	migrate create -ext sql -dir internal/migrations -seq $(name)