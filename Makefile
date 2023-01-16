resize: clean
	# go get -u -v github.com/valyala/fasthttp
	go build ./src/main.go 

clean:
	rm -f resize

build:
	go build ./src/main.go 

install: 
	go get -u -v github.com/valyala/fasthttp
	go get -u github.com/h2non/bimg
	go get -u "io/ioutil@latest"