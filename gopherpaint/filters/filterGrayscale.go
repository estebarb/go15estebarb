package filters

import (
	"context"
	"image"

	"github.com/disintegration/imaging"
)

func FilterGrayscale(_ context.Context, m image.Image) image.Image {
	res := imaging.Grayscale(m)
	return res
}
