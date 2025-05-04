package main

import (
	"fmt"
	imageprocessing "goroutines_pipeline/image_processing"
	"image"
	"strings"
)

// struct holding data like input path, output file path, and image object that is
// used to perform operations
type Job struct {
	InputPath string
	Image     image.Image
	OutPath   string
}

// takes slice of image paths and returns a channel that will output
// job structs
func loadImage(paths []string) <-chan Job {
	out := make(chan Job)
	go func() {
		// For each file path p, create a Job and add it to the channel
		for _, p := range paths {
			job := Job{InputPath: p,
				OutPath: strings.Replace(p, "images/", "images/output/", 1)}
			job.Image = imageprocessing.ReadImage(p)
			out <- job
		}
		close(out)
	}()
	return out
}

// accepts a channel of job structs, resizes the image from the function in the
// imageprocessing package, and then sends the modified job to the next stage
func resize(input <-chan Job) <-chan Job {
	out := make(chan Job)
	go func() {
		// For each input job, create a new job after resize and add it to
		// the out channel
		for job := range input { // Read from the channel
			job.Image = imageprocessing.Resize(job.Image)
			out <- job
		}
		close(out)
	}()
	return out
}

// accepts a channel of job structs, converts to grayscale, and the returns to
// modified job
func convertToGrayscale(input <-chan Job) <-chan Job {
	out := make(chan Job)
	go func() {
		for job := range input { // Read from the channel
			job.Image = imageprocessing.Grayscale(job.Image)
			out <- job
		}
		close(out)
	}()
	return out
}

// accepts a channel of Job structs and saves the images in the Job to the
// output directory, then returns an output of boolean values
func saveImage(input <-chan Job) <-chan bool {
	out := make(chan bool)
	go func() {
		for job := range input { // Read from the channel
			imageprocessing.WriteImage(job.OutPath, job.Image)
			// include error check to determine if WriteImage succeeded or failed
			out <- true
		}
		close(out)
	}()
	return out
}

func main() {

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
