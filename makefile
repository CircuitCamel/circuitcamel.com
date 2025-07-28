run:
	go run ./cmd/site/main.go

build:
	go build -o bin/site ./cmd/site/main.go

clean:
	rm bin/site && rm -d bin/

full: build
	./bin/site

submodule:
	git submodule update --recursive --remote --init