package dithering

import (
	"math/rand/v2"

	"github.com/brylleee/asciifx/asciifx"
)

type Random struct {
	RandomThreshold uint8
}

func UseRandom() Random {
	return Random{
		RandomThreshold: uint8(rand.IntN(255)),
	}
}

func (random Random) Dither(asciifxObj *asciifx.AsciiFx) {
	for i := 0; i < asciifxObj.Height; i++ {
		for j := 0; j < asciifxObj.Width; j++ {
			asciifxObj.Space[i][j] = asciifx.RGBI{
				R: func() uint8 {
					if asciifxObj.Space[i][j].R > random.RandomThreshold {
						return uint8(255)
					} else {
						return uint8(0)
					}
				}(),
				G: func() uint8 {
					if asciifxObj.Space[i][j].G > random.RandomThreshold {
						return uint8(255)
					} else {
						return uint8(0)
					}
				}(),
				B: func() uint8 {
					if asciifxObj.Space[i][j].B > random.RandomThreshold {
						return uint8(255)
					} else {
						return uint8(0)
					}
				}(),
				I: func() uint8 {
					if asciifxObj.Space[i][j].I > random.RandomThreshold {
						return uint8(255)
					} else {
						return uint8(0)
					}
				}(),
			}
		}
	}
}
