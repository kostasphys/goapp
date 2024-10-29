.DEFAULT_GOAL := goapp

.PHONY: all
all: clean goapp

.PHONY: goapp
goapp:
	mkdir -p bin
	go build -o bin ./...
	make -C client

.PHONY: clean
clean:
	go clean
	make -C client clean
	rm -f bin/*
