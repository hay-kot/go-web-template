
api:
	go run ./app/api/

cli:
	go build -o manage ./app/cli 

test:
	go test ./app/... ./internal/... ./pkgs/... -v -cover