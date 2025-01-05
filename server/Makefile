build:
	@go build -o bin/api
run: build
	@./bin/api
seed:
	@go run scripts/seed.go scripts/seed_hotel_room.go
test:
	go test -v ./...
