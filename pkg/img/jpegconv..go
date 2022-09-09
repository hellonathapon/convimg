package img

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"os"
)

func Jpegconv(pngSrc image.Image) error {
	//* create new bacnkground data
	bg := image.NewRGBA(pngSrc.Bounds())

	//* draw color to bg
	draw.Draw(bg, bg.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)

	//* place png data on top of bg data to create an new JPEG image
	draw.Draw(bg, bg.Bounds(), pngSrc, pngSrc.Bounds().Min, draw.Over)

	//* create new JPEG file
	// dt := time.Now()
	// fmt.Sprintf("%v", dt.Format())
	jpegFile, err := os.Create("./myJPEG.jpg")

	if err != nil {
		return errors.New("Unable to create JPEG file")
	}

	defer jpegFile.Close()

	//* set options of new image
	var opt jpeg.Options
    opt.Quality = 80

	//* finally, bound both JPEG data and JPEG file together
	err = jpeg.Encode(jpegFile, bg, &opt)

	if err != nil {
		return errors.New("Unable to encode data to JPEG file")
	}

	fmt.Println("Converted PNG to JPEG file")
	return nil
}