run:
	docker-compose --env-file ./.env up

stop:
	docker-compose down

db:
	docker-compose up db

test:
	go get ./...
	go test -p 1 -v -cover -race -timeout 30s -coverprofile=coverage.out -a ./... 
	go tool cover -func=coverage.out

