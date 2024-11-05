run:
	go run cmd/main.go

build-img:
	docker build -f Dockerfile -t vini65599/go-zip-code-temperature:latest .

run-img:
	docker run -it --rm vini65599/go-zip-code-temperature:latest

run-from-hub:
	docker run -it --rm --pull always vini65599/go-zip-code-temperature:latest

.PHONY: build-img run-img run run-from-hub