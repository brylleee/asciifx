package dithering

import "github.com/brylleee/asciifx/asciifx"

type BayerMatrix struct {
	BayerSize   int
	BayerMatrix [8][8]uint8
}

func UseBayerMatrix() BayerMatrix {
	return BayerMatrix{
		BayerSize: 8,
		BayerMatrix: [8][8]uint8{
			{0, 32, 8, 40, 2, 34, 10, 42},
			{48, 16, 56, 24, 50, 18, 58, 26},
			{12, 44, 4, 36, 14, 46, 6, 38},
			{60, 28, 52, 20, 62, 30, 54, 22},
			{3, 35, 11, 43, 1, 33, 9, 41},
			{51, 19, 59, 27, 49, 17, 57, 25},
			{15, 47, 7, 39, 13, 45, 5, 37},
			{63, 31, 55, 23, 61, 29, 53, 21},
		},
	}
}

func (bayerMatrix BayerMatrix) Dither(asciifxObj *asciifx.AsciiFx) {
	for i := 0; i < asciifxObj.Height; i++ {
		for j := 0; j < asciifxObj.Width; j++ {
			asciifxObj.Space[i][j] = asciifx.RGBI{
				R: asciifx.Threshold(asciifxObj.Space[i][j].R, bayerMatrix.BayerMatrix[i%(bayerMatrix.BayerSize)][j%(bayerMatrix.BayerSize)]<<2),
				G: asciifx.Threshold(asciifxObj.Space[i][j].G, bayerMatrix.BayerMatrix[i%(bayerMatrix.BayerSize)][j%(bayerMatrix.BayerSize)]<<2),
				B: asciifx.Threshold(asciifxObj.Space[i][j].B, bayerMatrix.BayerMatrix[i%(bayerMatrix.BayerSize)][j%(bayerMatrix.BayerSize)]<<2),
				I: asciifx.Threshold(asciifxObj.Space[i][j].I, bayerMatrix.BayerMatrix[i%(bayerMatrix.BayerSize)][j%(bayerMatrix.BayerSize)]<<2),
			}
		}
	}
}
