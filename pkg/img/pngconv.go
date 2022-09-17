package img

import (
	"errors"
	"image/jpeg"
	"os"
)

func Pngconv(jpegFile *os.File) error {
	jpegSrc, err := jpeg.Decode(jpegFile)

	if err != nil {
		return errors.New("Unable to decode JPEG file")
	}

	return nil
}