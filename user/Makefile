build:
	GOARCH=amd64 \
	GOOS=linux \
	CGO_ENABLED=0 \
	APP_ENV=development \
	go build -o ./bin/user ./cmd/user/main.go

protogen:
	@echo "Running generation protofiles..."
	docker run --rm -v "$(PWD):/defs" namely/protoc-all -f ./pkg/pb/user.proto -o ./ -l go

