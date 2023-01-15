helloworldserver: clean
	go install github.com/valyala/fasthttp@latest
	go build

clean:
	rm -f helloworldserver
