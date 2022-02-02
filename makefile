
build:
	cd backend && go build ./app/api/

api:
	cd backend && go run ./app/api/

cli:
	cd backend && go build -o manage ./app/cli 

coverage:
	cd backend && go test -race -coverprofile=coverage.out -covermode=atomic ./app/... ./internal/... ./pkgs/... -v -cover

test:
	cd backend && go test ./app/... ./internal/... ./pkgs/... -v


test-client:
	cd backend && go run ./app/api/ &
	sleep 5
	cd client && npm run test
	