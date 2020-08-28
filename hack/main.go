package main

import (
	"fmt"
	"github.com/moooooooooose/mews/pkg/imageprocessing"
	"image/jpeg"
	"log"
	"os"
)

// hack main to run through actions on beeg test image
func main() {
	imageToTransform := "beeg.jpg"

	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	imageFile, err := os.Open(fmt.Sprintf("%s/hack/images/%s", currentDir, imageToTransform))
	if err != nil {
		log.Fatal(err)
	}

	image, err := jpeg.Decode(imageFile)
	if err != nil {
		log.Fatal(err)
	}
	imageFile.Close()

	processorPipeline := imageprocessing.NewProcessorPipeline()
	boundResize, err := imageprocessing.NewActionBoundResize()
	if err != nil {
		log.Fatal(err)
	}
	processorPipeline.AddAction(boundResize)

	processedImage, err := processorPipeline.Transform(image)
	if err != nil {
		log.Fatal(err)
	}

	createOutput, err := os.Create(fmt.Sprintf("%s/hack/images/new_beeg.jpg", currentDir))
	if err != nil {
		log.Fatal(err)
	}
	defer createOutput.Close()

	jpeg.Encode(createOutput, processedImage, nil)
}
