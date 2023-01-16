package draft

import (
	// "flag"
	"fmt"
	// "log"

	// "math/rand"
	// "time"

	// "github.com/h2non/bimg"
	// "github.com/michael-nhat/golang-resize/src/utils"
	"github.com/valyala/fasthttp"
)



func TestRequest(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "Hdedd!\n")
	fmt.Fprintf(ctx, "RdequestURI is %q\n", ctx.RequestURI())
	fmt.Fprintf(ctx, "Requested path is %q\n", ctx.Path())
	fmt.Fprintf(ctx, "Query string is %q\n", ctx.QueryArgs().Peek("img"))
	fmt.Fprintf(ctx, "\n")
	ctx.SetContentType("text/plain; charset=utf8")
}
