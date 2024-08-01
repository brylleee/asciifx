package downsampler

import (
	"github.com/brylleee/asciifx/asciifx"
)

type NearestNeighbor struct {
	Ratio int
}

func UseNearestNeighbor(ratio int) NearestNeighbor {
	return NearestNeighbor{
		Ratio: ratio,
	}
}

func (nearestneighbor NearestNeighbor) Downsample(asciifxObj *asciifx.AsciiFx) {
	new_width := asciifxObj.Width / nearestneighbor.Ratio
	new_height := asciifxObj.Height / nearestneighbor.Ratio

	new_space := make([][]asciifx.RGBI, new_height)

	for i := 0; i < new_height; i++ {
		new_space[i] = make([]asciifx.RGBI, new_width)
		for j := 0; j < new_width; j++ {
			x, y := j*nearestneighbor.Ratio, i*nearestneighbor.Ratio
			new_space[i][j] = asciifxObj.Space[y][x]
		}
	}

	asciifxObj.Height = new_height
	asciifxObj.Width = new_width
	asciifxObj.Space = new_space
}
