install:
	brew install entr
	go get ./...
start: build-local 
	./build/migration
	AWS_REGION=us-east-1 GIN_MODE=test ./build/api
dev :
	echo "STARTING HOT RELOAD ENV"
	find . -type f -name "*.go" | entr -r make start
stop-db:
	docker-compose -f ./docker/local/docker-compose.yml down
start-db:	
	docker-compose -f ./docker/local/docker-compose.yml up -d
copy-documentation:
	cp ./docs/rpost-it-golang.json ./build/docs.json
copy-web :
	cp -r ./web ./build
build-local: build-folder copy-documentation copy-web
	go build -o ./build/api ./cmd/local/main.go
	go build -o ./build/migration ./cmd/migration/main.go
	cp env.local.json ./build/env.json
deploy: build-folder copy-documentation copy-web
	GOOS=linux go build -o ./build/api ./cmd/lambda/main.go
	GOOS=linux go build -o ./build/migration ./cmd/migration/main.go
	cp env.json ./build/env.json
	serverless deploy
build-folder:
	mkdir -p build