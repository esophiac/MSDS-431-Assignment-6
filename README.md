# Assignment 6: Data Pipelines with Concurrency

The goal of this assignment is to explore data throughput times using concurrency in Go. This assignment began with cloning a repository that processes a list of images (full details about the original repository in the Background section) and then making adjustments to verify how goroutines and channels can imapct run times.

Made the following changes to the repository:
- Add error checking for image file input and output.
    - In the imageprocessing packages, changed the ReadImage function to print out error message on nil and when getting a err for decoding the image
    - Also changed the output function to output errors with the log function in the os package
    - include error check to determine if WriteImage succeeded or failed
    - include an error check to determine if the file input is valid
- Replaced four images from original repository with image5, image6, image7, and image8
    - Kept the initial images as well for testing
    - used image1.jpeg for testing in imageprocessing as well
- Added unit tests in image_processing subdirectory for image_processing.go and main_test in the main directory
    - testing in the image_processing directory tests the resize and grayscale functions to verify they work
    - testing in the main directory tests that the channels function appropriately
- Added benchmark methods for capturing pipeline throughput times.
    - The benchmarks for running the code with goroutines and running the code without go routines are in the main_test.go file.
- Additional code modifications
    - Added more comments to explain how the functions worked
    - other readbility changes
    - changes to the main.go file to demonstrate functionality with goroutines and without goroutines
- Build, test, and run the pipeline program with and without goroutines.
    - You can run the program in both methods from the main.go file.
- Created README.md to document work
- OPTIONAL: changed the resize function in image_processing so that images will not be distorted

This assignment found that the version made with goroutines was faster (0.1238 ns/op) than the version without goroutines (0.1715 ns/op). 

## Background
This respository started off as a cloned repository from a project by [Amrit Singh](https://www.codeheim.io/) to demonstrate Go image processing pipeline with concurrency. The initial repository can be found [here](https://github.com/code-heim/go_21_goroutines_pipeline). He also put together a video tutorial [here](https://www.youtube.com/watch?v=8Rn8yOQH62k).

### Image Sources
For this assignment, we found other images to use. These are the sources for the images I've found.

image5.jpg: "[Milky Way Galaxy of Joshua Tree](https://www.flickr.com/photos/115357548@N08/18202006112)" by [Joshua Tree National Park](https://www.flickr.com/photos/115357548@N08) is marked with [Public Domain Mark 1.0](https://creativecommons.org/publicdomain/mark/1.0/?ref=openverse).

image6.jpg: "[Hawai‘i Volcanoes National Park](https://www.flickr.com/photos/42600860@N02/32852214577)" by [National Park Service](https://www.flickr.com/photos/42600860@N02) is marked with [Public Domain Mark 1.0](https://creativecommons.org/publicdomain/mark/1.0/?ref=openverse).

image7.jpg: "[Glacier National Park](https://www.flickr.com/photos/42600860@N02/51971764978)" by [National Park Service](https://www.flickr.com/photos/42600860@N02) is marked with [Public Domain Mark 1.0](https://creativecommons.org/publicdomain/mark/1.0/?ref=openverse).

image8.jpg: "[Bison feeding near Mud Volcano](https://www.flickr.com/photos/80223459@N05/51916511191)" by [YellowstoneNPS](https://www.flickr.com/photos/80223459@N05) is marked with [Public Domain Mark 1.0](https://creativecommons.org/publicdomain/mark/1.0/?ref=openverse).

## Roles of Programs and Data
These are the programs in the repository. Data information is in the background section.

- image_processing
    - images: images used for testing the image_processing package
        - image1.jpeg: from original repository
    - image_processesing.go: adjusted resize function
    - image_processing_test.go: testing the grayscale and resize functions. 
- images
    - output
        - image5.jpeg: processed image
        - image6.jpeg: processed image
        - image7.jpeg: processed image
        - image8.jpeg: processed image
    - image1.jpeg: from original repository, for testing
    - image2.jpeg: from original repository, for testing
    - image3.jpeg: from original repository, for testing
    - image4.jpeg: from original repository, for testing
    - image5.jpeg: new image
    - image6.jpeg: new image
    - image7.jpeg: new image
    - image8.jpeg: new image
- gitignore: from original repository, removed .exe so that the original and final build are included in this repository.
- go.mod: defines the module's properties
- go.sum: record of the library the project depends on
- LICENSE: from the original repository
- main_test.go: tests the functions in main.go to make sure the channels are working
- main.go: runs the program with goroutines and without goroutines
- README.md


## Application
An executable for this project was created using Windows. To create your own executable, run **go build** in the same directory as the go program. For more information, see the Gopher documentation on creating an executable [here](https://go.dev/doc/tutorial/compile-install).

## Use of AI
AI was not used for this assingment.

