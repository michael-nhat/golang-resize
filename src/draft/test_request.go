package draft

import (
	// "flag"
	"fmt"
	"strings"
	// "log"

	// "math/rand"
	// "time"

	// "github.com/h2non/bimg"
	// "github.com/michael-nhat/golang-resize/src/utils"
	"github.com/valyala/fasthttp"
)

func TestRequest(ctx *fasthttp.RequestCtx) {
	fmt.Printf("Hdedd!\n")
	fmt.Printf("RdequestURI is %q\n", ctx.RequestURI())
	a := strings.Split(string(ctx.Path()),"/");
	fmt.Printf("Requested path is %q\n", a[len(a) - 1])
	fmt.Printf("Query string is %q\n", ctx.QueryArgs().Peek("img"))
	ctx.SetContentType("text/plain; charset=utf8")
}
