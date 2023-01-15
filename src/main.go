package main

import (
	"flag"
	"fmt"
	"github.com/valyala/fasthttp"
	"log"
	"github.com/michael-nhat/golang-resize/src/utils"
)

var (
	addr     = flag.String("addr", ":8089", "TCP address to listen to")
	compress = flag.Bool("compress", false, "Whether to enable transparent response compression")
)

func main() {
	flag.Parse()

	h := requestHandler
	if *compress {
		h = fasthttp.CompressHandler(h)
	}

	if err := fasthttp.ListenAndServe(*addr, h); err != nil {
		log.Fatalf("Error in ListenAndServe: %v", err)
	}
}

type PeekFun func(string) []byte

func getArgF(peekf PeekFun) PeekFun {
	return func(targetArg string) []byte {
		return peekf(targetArg)
	}
}

func requestHandler(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "He!\n")

	fmt.Fprintf(ctx, "RequestURI is %q\n", ctx.RequestURI())
	fmt.Fprintf(ctx, "Requested path is %q\n", ctx.Path())
	fmt.Fprintf(ctx, "Query string is %q\n", ctx.QueryArgs().Peek("img"))
	fmt.Fprintf(ctx, "\n")

	var params = getArgF(ctx.QueryArgs().Peek)

	fmt.Fprintf(ctx, "Requested path is %q\n", params("img"))
	var xossp = params("x-oss-process")
	
	ctx.SetContentType("text/plain; charset=utf8")
}
