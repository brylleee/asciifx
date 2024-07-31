package dithering

import "github.com/brylleee/asciifx/asciifx"

type BayerMatrix struct{}

func (bayerMatrix *BayerMatrix) Dither(asciifx *asciifx.AsciiFx) {
	for i := 0; i < asciifx.Height; i++ {
		for j := 0; j < asciifx.Width; j++ {

		}
	}
}
