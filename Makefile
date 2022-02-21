.PHONY: build
build:
	docker build -t poncheska/sa-auth -f builds/Dockerfile .
	docker push poncheska/sa-auth

.PHONY: run
run:
	go run ./main.go