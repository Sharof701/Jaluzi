
tidy:
	@go mod tidy
	@go mod vendor
 
run:
	@go run cmd/main.go

migration:
	@migrate create -ext sql -dir ./migrations -seq $(name)

migrateup:
	@migrate -path ./migrations -database "$(DB_URL)" -verbose up

migratedown:
	@migrate -path ./migrations -database "$(DB_URL)" -verbose down

migrateforce:
	@migrate -path ./migrations -database "$(DB_URL)" -verbose force 1

swag_init:
	swag init -g api/api.go -o api/docs