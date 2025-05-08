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
