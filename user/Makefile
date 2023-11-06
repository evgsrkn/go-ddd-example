include ./config/dev.env

POSTGRESQL_URL="postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable"
MIGRATE=migrate -database ${POSTGRESQL_URL} -path ./migrations

build:
	@GOARCH=amd64 \
		GOOS=linux \
		CGO_ENABLED=0 \
		go build -o ./bin/user ./cmd/user/main.go

compose-up:
	@docker compose -f ./deployments/docker-compose.dev.yml --env-file ./config/dev.env up --build -d

compose-down:
	@docker compose -f ./deployments/docker-compose.dev.yml down

migrate-up:
	@$(MIGRATE) up

migrate-new:
	@read -p "Enter the name of new migration: " name; \
	$(MIGRATE) create -ext sql -dir migrations $$name

migrate-down:
	@$(MIGRATE) down

migrate-force:
	@read -p "Enter version migration: " name; \
	$(MIGRATE) force $$name

migrate-drop:
	@$(MIGRATE) drop

protogen:
	@docker run --rm -v "$(PWD):/defs" namely/protoc-all -f ./pkg/pb/user.proto -o ./ -l go

test:
	mkdir -p ./tmp
	go clean -testcache
	go test -v ./... -coverprofile=./tmp/coverage.out

coverage:
	gotestsum --format dots --packages="./..." -- -coverprofile=./tmp/coverage.out
	go tool cover -html=./tmp/coverage.out -o ./tmp/coverage.html
