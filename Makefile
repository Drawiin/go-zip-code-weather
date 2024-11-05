run:
	@echo "Running application with .env file"
	@set -a && source .env && go run cmd/main.go

build-img:
	docker build -f Dockerfile -t vini65599/go-zip-code-temperature:latest .

run-img:
	docker run -it --rm -p 8080:8080 vini65599/go-zip-code-temperature:latest

push-img:
	docker push vini65599/go-zip-code-temperature:latest

run-from-hub:
	docker run -it --rm --pull always -p 8080:8080 vini65599/go-zip-code-temperature:latest

.PHONY: build-img run-img run run-from-hub