add:
	git add . && git commit -m 'updated(${task}): ${commit}'
push:
	git push origin ${ORIGIN}
run:
	docker compose up --build -d 
up:
	docker compose up -d 
stop:
	docker compose down 
logs:
	docker logs blog-app 

build:
	@go build -o bin/blog-app ./Delivery/main.go

test:
	@go test $(shell go list ./Tests/...) -v

test-coverage:
	@go test -coverprofile=coverage.out $(shell go list ./Tests...')
	@go tool cover -func=coverage.out
