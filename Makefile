start: 	clean deps build-run

clean:
	rm -f main

build-run:
	go build ./src/main.go && ./main

build:
	go build ./src/main.go
	
install: 
	go get -u -v github.com/valyala/fasthttp
	go get -u github.com/h2non/bimg
	go get -u "io/ioutil@latest"
	go get -u "github.com/go-redis/redis"

deps: 
	go mod tidy
