start: 	clean deps build

clean:
	rm -f resize

build-run:
	go build ./src/main.go && ./main

install: 
	go get -u -v github.com/valyala/fasthttp
	go get -u github.com/h2non/bimg
	go get -u "io/ioutil@latest"
	go get -u "github.com/go-redis/redis"

deps: 
	go mod tidy

build:
	go build ./src/main.go