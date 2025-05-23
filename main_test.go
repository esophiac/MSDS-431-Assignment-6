package main

import (
	"fmt"
	imageprocessing "goroutines_pipeline/image_processing"
	"strings"
	"testing"
)

// test the loadImage function in main.go
//func TestLoadImage(t *testing.T) {
//
//}

// slice of string paths for testing
var imagePaths []string = []string{"images/image1.jpeg",
	"images/image2.jpeg",
	"images/image3.jpeg",
	"images/image4.jpeg",
}

// test the loadImage function in main.go
// testing the ReadImage function is handled in the imageprocessing tests, so this is
// just making sure the channels are acting appropriately
func TestLoadImage(t *testing.T) {

	expected := 0 // unbuffered channel

	result := loadImage(imagePaths)

	if len(result) != expected {
		t.Errorf("Expected %v, got %v", expected, len(result))
	}

}

// testing the resize function in main.go
func TestResize(t *testing.T) {

	input := loadImage(imagePaths)

	expected := 0 // unbuffered channel

	result := resize(input)

	if len(result) != expected {
		t.Errorf("Expected %v, got %v", expected, len(result))
	}
}

// testing the convertToGrayscale function in main.go
func TestConvertToGrayScale(t *testing.T) {

	input := loadImage(imagePaths)

	expected := 0 // unbuffered channel

	result := convertToGrayscale(input)

	if len(result) != expected {
		t.Errorf("Expected %v, got %v", expected, len(result))
	}
}

// testing the saveImage function in main.go
func TestSaveImage(t *testing.T) {

	input := loadImage(imagePaths)

	expected := 0 // unbuffered channel

	result := saveImage(input)

	if len(result) != expected {
		t.Errorf("Expected %v, got %v", expected, len(result))
	}
}

// benchmark to get time with goroutines
func BenchmarkGoRoutine(b *testing.B) {
	// list of file paths
	imagePaths := []string{"images/image1.jpeg",
		"images/image2.jpeg",
		"images/image3.jpeg",
		"images/image4.jpeg",
	}

	channel1 := loadImage(imagePaths)        // read files and returns first channel
	channel2 := resize(channel1)             // resize the image and return another channel
	channel3 := convertToGrayscale(channel2) // convert to grayscale and return another channel
	writeResults := saveImage(channel3)      // save processed files

	for success := range writeResults {
		if success {
			fmt.Println("Success!")
		} else {
			fmt.Println("Failed!")
		}
	}
}

// benchmark to get time of program without goroutines
func BenchmarkWithout(b *testing.B) {

	imagePaths := []string{"images/image1.jpeg",
		"images/image2.jpeg",
		"images/image3.jpeg",
		"images/image4.jpeg",
	}

	for _, value := range imagePaths {

		outPath := strings.Replace(value, "images/", "images/output/", 1)

		imgPath := imageprocessing.ReadImage(value)
		imgSize := imageprocessing.Resize(imgPath)
		imgGray := imageprocessing.Grayscale(imgSize)
		imageprocessing.WriteImage(outPath, imgGray)

		//fmt.Println("image processing complete")
	}

}
