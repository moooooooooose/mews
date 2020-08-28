package imageprocessing

import (
	"fmt"
	size "github.com/nfnt/resize"
	"image"
)

const (
	DefaultWidth = 100
	DefaultHeight = 100
)

type ActionBoundResize struct{
	Width uint
	Height uint
}

type BoundResizeOptions func(*ActionBoundResize) error

var _ ImageAction = ActionBoundResize{}

func NewActionBoundResize(options ...BoundResizeOptions) (ImageAction, error) {
	actionBoundResize := &ActionBoundResize{}

	// get those defaults
	actionBoundResize.Height = DefaultHeight
	actionBoundResize.Width = DefaultWidth

	// get any optional params
	for _, option := range options {
		err := option(actionBoundResize)
		if err != nil {
			return nil, fmt.Errorf("error happened while setting bound resize options %w", err)
		}
	}
	return actionBoundResize, nil
}

func SetSizeWidth(size uint) func(boundResize *ActionBoundResize) error {
	return func(boundResize *ActionBoundResize) error {
		boundResize.Width = size
		return nil
	}
}

func SetSizeHeight(size uint) func(boundResize *ActionBoundResize) error {
	return func(boundResize *ActionBoundResize) error {
		boundResize.Height = size
		return nil
	}
}

func (a ActionBoundResize) Transform(image image.Image) (image.Image, error) {
	imageBounds := image.Bounds()

	if imageBounds.Dx() > imageBounds.Dy() {
		return resize(image, a.Width, 0)
	}
	return resize(image, 0, a.Height)
}

func resize(image image.Image, width, height uint) (image.Image, error) {
	return size.Resize(width, height, image, size.Lanczos2), nil
}


