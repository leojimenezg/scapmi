package utils

import (
	"bytes"
	"fmt"
	"image/png"
	"log"

	"gioui.org/op/paint"
	"github.com/leojimenezg/scapmi/gui/assets"
	"github.com/pkg/browser"
)

func LoadPNG(name string) paint.ImageOp {
	fileName := fmt.Sprintf("public/%s", name)
	file, err := assets.PngImgs.ReadFile(fileName)
	if err != nil {
		return paint.ImageOp{}
	}

	image, err := png.Decode(bytes.NewReader(file))
	if err != nil {
		return paint.ImageOp{}
	}
	return paint.NewImageOp(image)
}

func OpenURL(url string) {
	err := browser.OpenURL(url)
	if err != nil {
		log.Printf("failed to open %s URL: %v", url, err)
	}
}
