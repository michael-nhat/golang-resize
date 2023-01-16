package main

import (
	"flag"
	"fmt"

	// "io/ioutil"
	"log"
	"os"

	// "unsafe"
	// "math/rand"
	// "github.com/h2non/bimg"
	"time"

	"github.com/michael-nhat/golang-resize/src/draft"
	"github.com/michael-nhat/golang-resize/src/resize"
	"github.com/michael-nhat/golang-resize/src/utils"
	"github.com/valyala/fasthttp"
)

func requestHandler(ctx *fasthttp.RequestCtx) {
	// prevent for browser auto job
	switch string(ctx.Path()) {
	case
		"/",
		"/favicon.ico":
		return
	}
	// draft.TestRequest(ctx)
	// return

	var params = getArgF(ctx.QueryArgs().Peek)

	var fileShortPath = params("img")
	var filePath = currentDir + "/src/jiangzi_tupian/" + string(fileShortPath)
	// fmt.Printf("wtf: %s\n",string(params("wft")))
	// fmt.Printf(filePath)
	
	if string(params("x-oss-process")) == "" {
		resize.GetFile(ctx, filePath)
	} else {
		resize.Resize(filePath, ctx, string(params("x-oss-process")))
	}
	// fmt.Printf("Requested path is %q\n", ctx.Path())

	// var oss_params = utils.GetParams(xossp)

	// ctx.Write(newImage)
	// ctx.Response.Write(fileBytes)
	// ctx.Response.SendFile(filePath)
	// ctx.Response.SetBody(buffer)
	// ctx.SetContentType("application/octet-stream")
	// ctx.SetContentType("image/jpeg")
	// ctx.Response.Header.Set("Content-Type", "image/jpeg")
	

	utils.UNUSED(
		// xossp,
		// fileBytes,

		// oss_params,
		draft.TestRequest,
		utils.Aa)
}

var (
	addr          = flag.String("addr", ":8089", "TCP address to listen to")
	compress      = flag.Bool("compress", false, "Whether to enable transparent response compression")
	currentDir, _ = os.Getwd()
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
