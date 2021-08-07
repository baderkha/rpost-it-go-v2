install:
	brew install entr
	go get ./...
start: build-local 
	cd build && ./migration
	cd build && AWS_REGION=us-east-1 GIN_MODE=test ./api
dev :
	echo "STARTING HOT RELOAD ENV"
	find . -type f -name "*.go" | entr -r make start
stop-db:
	docker-compose -f ./docker/local/docker-compose.yml down
start-db:	
	docker-compose -f ./docker/local/docker-compose.yml up -d
build-local: build-folder
	go build -o ./build/api ./cmd/api/main.go
	go build -o ./build/migration ./cmd/migration/main.go
	cp env.local.json ./build/env.json
build-folder:
	mkdir -p build