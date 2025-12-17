package utils

import (
	"bytes"
	"embed"
	"fmt"
	"image/png"

	"gioui.org/op/paint"
)

func LoadImage(fs *embed.FS, name string) paint.ImageOp {
	fileName := fmt.Sprintf("public/%s", name)
	file, err := fs.ReadFile(fileName)
	if err != nil {
		return paint.ImageOp{}
	}
	image, err := png.Decode(bytes.NewReader(file))
	if err != nil {
		return paint.ImageOp{}
	}
	return paint.NewImageOp(image)
}
