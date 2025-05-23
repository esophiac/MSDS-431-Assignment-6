package imageprocessing

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"

	"github.com/nfnt/resize"
)

// read an image file and return an image object
func ReadImage(path string) image.Image {
	inputFile, err := os.Open(path)
	if err != nil {
		log.Fatal("Unable to read input file "+path, err)
	}
	defer inputFile.Close()

	// Decode the image
	img, _, err := image.Decode(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	return img
}

// write the image object to a file
func WriteImage(path string, img image.Image) {
	outputFile, err := os.Create(path)
	if err != nil {
		log.Fatal("Unable to output to file "+path, err)
	}
	defer outputFile.Close()

	// Encode the image to the new file
	err = jpeg.Encode(outputFile, img, nil)
	if err != nil {
		log.Fatal(err)
	}
}

// convert image object to grayscale
func Grayscale(img image.Image) image.Image {
	// Create a new grayscale image
	bounds := img.Bounds()
	grayImg := image.NewGray(bounds)

	// Convert each pixel to grayscale
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			originalPixel := img.At(x, y)
			grayPixel := color.GrayModel.Convert(originalPixel)
			grayImg.Set(x, y, grayPixel)
		}
	}
	return grayImg
}

// shrinks image to so that the longest axis is 500 px
// keeps the rest of the image in ratio
func Resize(img image.Image) image.Image {
	dimensions := img.Bounds()

	//width := dimensions.Max.X - dimensions.Min.X
	//height := dimensions.Max.Y - dimensions.Min.Y

	width := dimensions.Max.X
	height := dimensions.Max.Y

	var newWidth float64
	var newHeight float64
	var ratio float64

	if width > height {
		ratio = 500 / float64(width)

		fmt.Println(ratio)

		newWidth = ratio * float64(width)
		newHeight = ratio * float64(height)

	} else if height < width {
		ratio = 500 / float64(height)

		newWidth = ratio * float64(width)
		newHeight = ratio * float64(height)

	} else {
		newWidth = 500
		newHeight = 500
	}

	finalWidth := uint(newWidth)
	finalHeight := uint(newHeight)

	//newWidth := uint(500)
	//newHeight := uint(500)
	resizedImg := resize.Resize(finalWidth, finalHeight, img, resize.Lanczos3)
	return resizedImg
}
