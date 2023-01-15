resize: clean
	# go get -u -v github.com/valyala/fasthttp
	go build ./src/resize.go 

clean:
	rm -f resize

build:
	go build ./src/resize.go