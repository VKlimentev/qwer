up:
	docker-compose up -d

down:
	docker-compose down

build:
	docker-compose build

test:
	go test -v ./...