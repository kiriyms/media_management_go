run: build
	@./backend/bin/app

build:
	@cd backend && go build -o ./bin/app ./cmd
# 	@cd backend && go build -ldflags="-X main.commit=local" -o ./bin/app ./cmd

test:
	@cd backend && go test -v ./...

push:
	@git push origin main
