include ./config/dev.env

build:
	@GOARCH=amd64 \
		GOOS=linux \
		CGO_ENABLED=0 \
		go build -o ./bin/user ./cmd/user/main.go

compose-up:
	@docker compose -f ./deployments/docker-compose.yaml --env-file ./config/dev.env up --build -d

compose-down:
	@docker compose -f ./deployments/docker-compose.yaml down

protogen:
	@docker run --rm -v "$(PWD):/defs" namely/protoc-all -f ./pkg/pb/user.proto -o ./ -l go

test:
	mkdir -p ./tmp
	go clean -testcache
	go test -v ./... -coverprofile=./tmp/coverage.out

coverage:
	gotestsum --format dots --packages="./..." -- -coverprofile=./tmp/coverage.out
	go tool cover -html=./tmp/coverage.out -o ./tmp/coverage.html
