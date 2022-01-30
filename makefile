
build:
	go build ./app/api/

api:
	go run ./app/api/

cli:
	go build -o manage ./app/cli 

test:
	go test -race -coverprofile=coverage.out -covermode=atomic ./app/... ./internal/... ./pkgs/... -v -cover