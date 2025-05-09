package imageprocessing

import (
	//"fmt"
	"image"
	"testing"
	//"image/color"
)

// variable creation here
var goodIn string = "images/image1.jpeg"

//var badIn string = "/images/image100.jpeg"

var imgTest image.Image = ReadImage(goodIn)

//var badOut string = "test/uotput/images/image1.jpeg"

// test that the Grayscale image object functions
func TestGrayscale(t *testing.T) {

	grayImg := Grayscale(imgTest)

	bounds := grayImg.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := grayImg.At(x, y).RGBA()
			if r != g || r != b || b != g {
				t.Errorf("Greyscale conversion test failed.")
			}
		}
	}
}

func TestResize(t *testing.T) {

	resizeImg := Resize(imgTest)
	dimensions := resizeImg.Bounds()

	width := dimensions.Max.X
	height := dimensions.Max.Y

	if width != 500 && height != 500 {
		t.Errorf("Resize test failed.")
	}
}
