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
	fmt.Printf("path is %s\n", string(ctx.Path()))
	fmt.Printf("RdequestURI is %q\n", ctx.RequestURI())
	a,b,found := strings.Cut(string(ctx.Path()), "/");
	if b != "" && found == true {
		fmt.Printf("Requested path is %q\n", a)

	}
	fmt.Printf("Query string is %q\n", ctx.QueryArgs().Peek("img"))
	ctx.SetContentType("text/plain; charset=utf8")
}
