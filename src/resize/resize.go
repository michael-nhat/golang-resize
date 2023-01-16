package resize

import (
	// "flag"
	"fmt"
	// "path/filepath"
	"strings"

	// "log"
	"os"
	"strconv"

	// "math/rand"
	// "github.com/clarketm/json"
	// "time"

	"github.com/h2non/bimg"
	"github.com/michael-nhat/golang-resize/src/utils"
	"github.com/valyala/fasthttp"
	// "github.com/clarketm/json"
)

// Exporting interface instead of struct
// type Oss_params interface {
// }
// Struct is not exported
type Oss_params struct {
	w int
	h int    // Defaults to 0
	m string // default to ""
	// not implement
	limit string // default to "0"
	l     int
	s     int
	color string
}

// We are forced to call the constructor to get an instance of candidate
// input is partial update
func New_oss_params(raw Oss_params) bool {

	// return Oss_params{name, 0}  // enforce the default value here
	return true
}

func Res(ctx *fasthttp.RequestCtx, buffer []byte) {
	ctx.Response.SetBody(buffer)
	ctx.Response.Header.Set("Content-Type", "image/jpeg")
}

func GetFile(ctx *fasthttp.RequestCtx, filePath string) {
	buffer, err := bimg.Read(filePath)
	if err != nil {
		// fmt.Fprintln(os.Stderr, err)
		utils.ErrRes(ctx, 404, "read file")
		return
	}
	Res(ctx, buffer)
}

func Resize(filePath string, ctx *fasthttp.RequestCtx, xossString string) {
	// aa := Oss_params{m: "skjd", w: 33, h: 33, limit: "skdjf"}
	// utils.ErrRes(ctx, 501, "fdsf")
	// return

	xossMap := make(map[string]string)
	xossArr := strings.Split(xossString, ",")
	for i := 1; i < len(xossArr); i++ {
		par := strings.Split(xossArr[i], "_")
		xossMap[par[0]] = par[1]
	}
	oss_params := Oss_params{limit: xossMap["limit"], m: xossMap["m"]}
	if xossMap["w"] != "" {
		w, err := strconv.Atoi(xossMap["w"])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		oss_params.w = w
	}
	if xossMap["h"] != "" {
		h, err := strconv.Atoi(xossMap["h"])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		oss_params.h = h
	}

	buffer, err := bimg.Read(filePath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	size, err := bimg.NewImage(buffer).Size()
	imgW := size.Width
	imgH := size.Height
	tarH := oss_params.h
	tarW := oss_params.w
	wRatio := float32(tarW) / float32(imgW)
	hRatio := float32(tarH) / float32(imgH)

	// fmt.Printf("%#v", oss_params)

	if oss_params.m == "fill" {
		if tarW > 0 && tarH > 0 {
			buffer, err := bimg.NewImage(buffer).Resize(tarW, tarH)
			if err != nil {
				utils.ErrRes(ctx, 500, "Resize")
			}
			Res(ctx, buffer)
		}
	} else if oss_params.m == "fixed" {
		if tarW > 0 && tarH > 0 {
			buffer, err := bimg.NewImage(buffer).ForceResize(int(tarW), int(tarH))
			if err != nil {
				utils.ErrRes(ctx, 500, "ForceResize")
			}
			Res(ctx, buffer)
		}
	} else {
		if hRatio > wRatio {
			buffer, err := bimg.NewImage(buffer).Resize(tarW, 0)
			if err != nil {
				utils.ErrRes(ctx, 500, "Resize")
			}
			Res(ctx, buffer)
		} else {
			buffer, err := bimg.NewImage(buffer).Resize(0, tarH)
			if err != nil {
				utils.ErrRes(ctx, 500, "Resize")
			}
			Res(ctx, buffer)
		}
	}
	// Res(ctx, buffer)

	utils.UNUSED(buffer,
		oss_params,
		err,
		// fileBytes,
	)
}
