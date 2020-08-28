package main

import (
	"flag"
	"fmt"
	"github.com/moooooooooose/mews/pkg/imageprocessing"
	"image/jpeg"
	"io"
	"log"
	"net/http"
	"os"
)

const (
	imageLocation = "/tmp/mews.jpg"
	imageResizeLocation = "/tmp/mews_resize.jpg"
)

var imageUrl string
var bound bool
var t string
var id string

func main() {
	flag.StringVar(&t, "token", "", "google auth token")
	flag.StringVar(&id, "sheet", "", "google sheet id")
	flag.StringVar(&imageUrl, "imageUrl", "", "image url")
	flag.BoolVar(&bound, "bound", false, "bound image resize")
	flag.Parse()

	if t == "" || id == "" || imageUrl == "" {
		flag.Usage()
		os.Exit(1)
	}

	resp, err := http.Get(imageUrl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	img, err := os.Create(imageLocation)
	if err != nil {
		log.Fatal(err)
	}
	defer img.Close()

	_, err = io.Copy(img, resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	imageFile, err := os.Open(imageLocation)
	if err != nil {
		log.Fatal(err)
	}

	image, err := jpeg.Decode(imageFile)
	if err != nil {
		log.Fatal(err)
	}
	imageFile.Close()

	processorPipeline := imageprocessing.NewProcessorPipeline()

	if bound {
		boundResize, err := imageprocessing.NewActionBoundResize()
		if err != nil {
			log.Fatal(err)
		}
		processorPipeline.AddAction(boundResize)
	} else {
		processorPipeline.AddAction(imageprocessing.NewGoogleSheetResize(id, t))
	}

	processedImage, err := processorPipeline.Transform(image)
	if err != nil {
		log.Fatal(err)
	}

	createOutput, err := os.Create(fmt.Sprintf(imageResizeLocation))
	if err != nil {
		log.Fatal(err)
	}
	defer createOutput.Close()

	jpeg.Encode(createOutput, processedImage, nil)
}
