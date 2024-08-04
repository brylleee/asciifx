package dithering

import (
	"math"

	"github.com/brylleee/asciifx/asciifx"
)

type FloydSteinberg struct{}

func UseFloydSteinberg() FloydSteinberg {
	return FloydSteinberg{}
}

func getColorDistance(originalColor asciifx.RGBI, RGBColor []uint8) int {
	redDifference := int(originalColor.R - RGBColor[0])
	greenDifference := int(originalColor.G - RGBColor[1])
	blueDifference := int(originalColor.B - RGBColor[2])

	return redDifference*redDifference + greenDifference*greenDifference + blueDifference*blueDifference
}

func getClosestColor(color asciifx.RGBI, rgbPalette [][]uint8, grayPalette []uint8) asciifx.RGBI {
	closestMatch := 255*255 + 255*255 + 255*255 + 1
	rgbMatchIndex, grayMatchIndex := 0, 0

	for idx, i := range rgbPalette {
		if getColorDistance(color, i) < closestMatch {
			closestMatch = getColorDistance(color, i)
			rgbMatchIndex = idx
		}
	}

	grayMatchIndex = int(color.I) / int(255/(len(grayPalette)))

	return asciifx.RGBI{
		R: rgbPalette[rgbMatchIndex][0],
		G: rgbPalette[rgbMatchIndex][1],
		B: rgbPalette[rgbMatchIndex][2],
		I: grayPalette[grayMatchIndex],
	}
}

func (floydSteinberg FloydSteinberg) Dither(asciifxObj *asciifx.AsciiFx) {
	for i := 0; i < asciifxObj.Height; i++ {
		for j := 0; j < asciifxObj.Width; j++ {
			oldValues := asciifxObj.Space[i][j]
			newValues := getClosestColor(asciifxObj.Space[i][j], asciifxObj.AsciifyChoice.Self().RGBColor, asciifxObj.AsciifyChoice.Self().GrayColors)

			errors := asciifx.RGBI{
				R: oldValues.R - newValues.R,
				G: oldValues.G - newValues.G,
				B: oldValues.B - newValues.B,
				I: oldValues.I - newValues.I,
			}

			if j < asciifxObj.Width-1 {
				asciifxObj.Space[i][j+1].R += uint8(math.Round(float64(errors.R) * 7.0 / 16.0))
				asciifxObj.Space[i][j+1].G += uint8(math.Round(float64(errors.G) * 7.0 / 16.0))
				asciifxObj.Space[i][j+1].B += uint8(math.Round(float64(errors.B) * 7.0 / 16.0))
				asciifxObj.Space[i][j+1].I += uint8(math.Round(float64(errors.I) * 7.0 / 16.0))
			}

			if i < asciifxObj.Height-1 {
				if j > 0 {
					asciifxObj.Space[i+1][j-1].R += uint8(math.Round(float64(errors.R) * 3.0 / 16.0))
					asciifxObj.Space[i+1][j-1].G += uint8(math.Round(float64(errors.G) * 3.0 / 16.0))
					asciifxObj.Space[i+1][j-1].B += uint8(math.Round(float64(errors.B) * 3.0 / 16.0))
					asciifxObj.Space[i+1][j-1].I += uint8(math.Round(float64(errors.I) * 3.0 / 16.0))
				}

				asciifxObj.Space[i+1][j].R += uint8(math.Round(float64(errors.R) * 5.0 / 16.0))
				asciifxObj.Space[i+1][j].G += uint8(math.Round(float64(errors.G) * 5.0 / 16.0))
				asciifxObj.Space[i+1][j].B += uint8(math.Round(float64(errors.B) * 5.0 / 16.0))
				asciifxObj.Space[i+1][j].I += uint8(math.Round(float64(errors.I) * 5.0 / 16.0))

				if j < asciifxObj.Width-1 {
					asciifxObj.Space[i+1][j+1].R += uint8(math.Round(float64(errors.R) * 1.0 / 16.0))
					asciifxObj.Space[i+1][j+1].G += uint8(math.Round(float64(errors.G) * 1.0 / 16.0))
					asciifxObj.Space[i+1][j+1].B += uint8(math.Round(float64(errors.B) * 1.0 / 16.0))
					asciifxObj.Space[i+1][j+1].I += uint8(math.Round(float64(errors.I) * 1.0 / 16.0))
				}
			}
		}
	}
}
