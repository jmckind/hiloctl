.PHONY: build
build:
	go build -o hiloctl main.go

.PHONY: clean
clean:
	rm	*.out \
		hiloctl

.PHONY: run
run:
	go run main.go

.PHONY: test
test:
	go test ./... -coverprofile cover.out
