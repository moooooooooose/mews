package imageprocessing

import (
	"bufio"
	"errors"
	"image"
	_ "image/png"
	"os"
)

var testError = errors.New("test")

var mockImageFiles = []string{
	"../../testdata/images/sad-pumpkin-spice-latte-unicorn.png",
}

func BuildMockProcessorPipeline(modifyFn func(*ProcessorPipelineMock)) ProcessorPipeline {
	mock := &ProcessorPipelineMock{}
	if modifyFn != nil {
		modifyFn(mock)
	}
	return mock
}

func BuildMockImageAction(modifyFn func(*ImageActionMock)) ImageAction {
	mock := &ImageActionMock{}
	if modifyFn != nil {
		modifyFn(mock)
	}
	return mock
}

func BuildMockImage() image.Image {
	if len(mockImageFiles) < 1 {
		panic(errors.New("no mock image files found"))
	}
	file, err := os.Open(mockImageFiles[0])
	if err != nil {
		panic(err)
	}
	defer file.Close()
	decodedImage, _, err := image.Decode(bufio.NewReader(file))
	if err != nil {
		panic(err)
	}
	return decodedImage
}
