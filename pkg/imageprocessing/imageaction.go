package imageprocessing

import "image"

type ImageAction interface {
	Transform(image.Image) (image.Image, error)
}
