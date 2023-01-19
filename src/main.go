package main

import (
	"flag"
	"fmt"
	// "go/printer"
	"strings"

	// "io/ioutil"
	"log"
	"os"

	// "unsafe"
	// "math/rand"
	// "github.com/h2non/bimg"
	"time"

	"github.com/go-redis/redis"
	"github.com/michael-nhat/golang-resize/src/draft"
	"github.com/michael-nhat/golang-resize/src/resize"
	"github.com/michael-nhat/golang-resize/src/utils"
	"github.com/valyala/fasthttp"
)

func requestHandler(ctx *fasthttp.RequestCtx) {
	// draft.TestRequest(ctx)
	switch string(ctx.Path()) {
	case
		"/",
		"/favicon.ico":
		return
	}

	params := getArgF(ctx.QueryArgs().Peek)
	reqPath := string(ctx.Path())
	strTrimPath := strings.TrimPrefix(reqPath, "/")
	_,fileShortPath,_ := strings.Cut(strTrimPath, "/");
	filePath := "/data/jiangzi_tupian/" + fileShortPath
	xossp := string(params("x-oss-process"))

	fileStat, err := os.Stat(filePath)
	if err != nil {
		utils.ErrRes(ctx, 500, "wrong file " + filePath)
		return
	}
	if xossp == "" {
		resize.GetFile(ctx, filePath)
	} else {
		resizedFileInfoKey := filePath + fmt.Sprint(fileStat.ModTime().Unix()) + xossp
		resizedFile := clientRedis.Get(resizedFileInfoKey).Val()
		if resizedFile != "" {
			resize.Res(ctx, []byte(resizedFile))
		} else {
			resize.Resize(filePath, ctx, xossp, clientRedis, resizedFileInfoKey)
		}
	}

	utils.UNUSED(
		// xossp,
		// fileBytes,

		// oss_params,
		draft.TestRequest,
		utils.Aa)
}

var (
	clientRedis = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	addr     = flag.String("addr", ":8089", "TCP address to listen to")
	compress = flag.Bool("compress", false, "Whether to enable transparent response compression")
	// currentDir, err = os.Getwd()
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
