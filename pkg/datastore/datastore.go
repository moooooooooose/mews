package datastore

import "image"

type Datastore interface {
	Save(image.Image) error
}
