package utils

import (
	"bytes"
	"fmt"
	"image/png"

	"gioui.org/op/paint"
	// "gioui.org/widget"
	"github.com/leojimenezg/scapmi/gui/assets"
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

// func LoadSVG(name string) *widget.Icon {
// 	fileName := fmt.Sprintf("public/%s", name)
// 	file, err := assets.SvgIcons.ReadFile(fileName)
// 	if err != nil {
// 		return &widget.Icon{}
// 	}
// 	icon, err := widget.NewIcon(file)
// 	if err != nil {
// 		return &widget.Icon{}
// 	}
// 	return icon
// }
