.PHONY: install build watch

install:
	@echo "Downloading dependencies..."
	go mod download

build: 
	@echo "Building application..."
	rm -f bin/main
	go build -o bin/main main.go

watch:
	@echo "Watching..."
	reflex -c reflex.conf

start:
	@echo "Starting..."
	./bin/main