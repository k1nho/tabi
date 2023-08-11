.PHONY: run
run:
	go run main.go


.PHONY: build
build:
	go build -o bin/tabi main.go


.PHONY: lauch
lauch:
	./bin/tabi

.PHONY: test
test:
	go test -v ./...
