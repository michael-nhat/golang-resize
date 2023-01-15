package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	// "unsafe"

	// "math/rand"
	"time"

	"github.com/h2non/bimg"
	"github.com/michael-nhat/golang-resize/src/draft"
	"github.com/michael-nhat/golang-resize/src/utils"
	"github.com/valyala/fasthttp"
)

func requestHandler(ctx *fasthttp.RequestCtx) {
	draft.TestRequest(ctx)
	var params = getArgF(ctx.QueryArgs().Peek)

	var fileShortPath = params("img")
	var filePathx, _ = os.Getwd()
	var filePath = filePathx + "/src/jiangzi_tupian/" + string(fileShortPath)

	fmt.Printf(filePath)

	buffer, err := bimg.Read(filePath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	newImage, err := bimg.NewImage(buffer).Resize(400, 600)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	bimg.Write("new.jpg", newImage)

	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(ctx, "Requested path is %q\n", params("img"))
	var xossp = params("x-oss-process")
	utils.Aa()

	// ctx.Write(newImage)
	// ctx.Response.Write(fileBytes)
	// ctx.Response.SendFile(filePath)
	ctx.Response.SetBody(buffer)
	// ctx.SetContentType("application/octet-stream")
	// ctx.SetContentType("image/jpeg")
	ctx.Response.Header.Set("Content-Type", "image/jpeg")
	// ctx.SetContentType("text/plain; charset=utf8")

	UNUSED(xossp, fileBytes)
}

func UNUSED(x ...interface{}) {}

var (
	addr     = flag.String("addr", ":8089", "TCP address to listen to")
	compress = flag.Bool("compress", false, "Whether to enable transparent response compression")
)

func main() {
	fmt.Printf("Hi %d\n", time.Now().UnixNano()%100)

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
