package utils

import (
	"fmt"
	"strings"

	// "flag"
	// "fmt"
	// "log"
	// "math/rand"
	// "time"
	// "github.com/h2non/bimg"
	// "github.com/michael-nhat/golang-resize/src/utils"
	// "encoding/json"
	// "time"

	// "github.com/clarketm/json"
	"github.com/valyala/fasthttp"
)

// wtf?? export with first capital??
func Aa() int16 {
	return 2
}

func GetParams(str string) interface{} {
	s := strings.Split(str, ",")
	fmt.Printf("%v", s)
	return "ksdjf"
}

func UNUSED(x ...interface{}) {}

var (
	strContentType     = "Content-Type"
	strApplicationJSON = "application/json"
)

func ErrRes(ctx *fasthttp.RequestCtx, code int, err string) {
	ctx.Response.Header.Set(strContentType, strApplicationJSON)
	ctx.Response.SetStatusCode(code)
	ctx.Response.SetBodyString("Error: "+err)
	// if err := json.NewEncoder(ctx).Encode(obj); err != nil {
	// 	ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
	// }	else {
	// 	ctx.Error("Unknow Error", fasthttp.StatusInternalServerError)
	// }
}

func GetInfoResizeFile(filePath string,xossp string) {
	
}