package utils

import (
	"path/filepath"

	"github.com/qeesung/image2ascii/convert"
)

func GetAsciiString() string {
	convertOptions := convert.DefaultOptions
	convertOptions.FixedWidth = 60
	convertOptions.FixedHeight = 30

	converter := convert.NewImageConverter()
	return converter.ImageFile2ASCIIString(filepath.Join(ImagesDir, "PwnLogo.png"), &convertOptions)
}
