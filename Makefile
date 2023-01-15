resize: clean
	# go get -u -v github.com/valyala/fasthttp
	go build ./src/main.go 

clean:
	rm -f resize

build:
	go build ./src/main.go 