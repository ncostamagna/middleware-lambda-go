export GOSUMDB=off

.PHONY: build
build:
	@echo "=> Building service"
	env GOOS=linux go build -ldflags="-s -w" -o bin/hello 		    		cmd/hello/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/authorizer 		    	cmd/authorizer/main.go
																			    
.PHONY: start
start:
	@echo "=> Starting service"
	@make format
	@make build
	@sls offline

.PHONY: format
format:
	@gofmt -w internal/
	@gofmt -w cmd/

.PHONY: lint
lint:
	@echo "=> Running linter"
	@${GOPATH}/bin/golangci-lint run ./internal/... ./cmd/...

.PHONY: deploy
deploy:
	@echo "=> Running deployment"
	@make build
	@sls deploy
