package main

import (
	"image/color"
	"log"
	"os"

	"github.com/hellonathapon/convimg/pkg/img"
)

func main() {
	pngFile, err := os.Open("mypng.png")

	if err != nil {
		log.Println("Unable to open Png source image")
		os.Exit(1)
	}

	defer pngFile.Close()

	err = img.Jpegconv(pngFile, color.White, "./pkg/converted-image")

	if err != nil {
		log.Println("Error to convert JPEG")
		os.Exit(1)
	}

	log.Println("Converted PNG to JPEG")
}