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

	return (redDifference * redDifference) + (greenDifference * greenDifference) + (blueDifference * blueDifference)
}

func getClosestColor(color asciifx.RGBI, rgbPalette [][]uint8, grayPalette []uint8) asciifx.RGBI {
	closestMatch := (255 * 255) + (255 * 255) + (255 * 255) + 1
	rgbMatchIndex, grayMatchIndex := 0, 0

	grayStepSize := float64(255) / float64(len(grayPalette)-1)
	grayMatchIndex = int(math.Round(float64(color.I) / grayStepSize))

	if rgbPalette == nil {
		return asciifx.RGBI{
			R: color.R,
			G: color.G,
			B: color.B,
			I: grayPalette[grayMatchIndex],
		}
	} else {
		for idx, i := range rgbPalette {
			if getColorDistance(color, i) < closestMatch {
				closestMatch = getColorDistance(color, i)
				rgbMatchIndex = idx
			}
		}

		return asciifx.RGBI{
			R: rgbPalette[rgbMatchIndex][0],
			G: rgbPalette[rgbMatchIndex][1],
			B: rgbPalette[rgbMatchIndex][2],
			I: grayPalette[grayMatchIndex],
		}
	}
}

func (floydSteinberg FloydSteinberg) Dither(asciifxObj *asciifx.AsciiFx) {
	for i := 0; i < asciifxObj.Height; i++ {
		for j := 0; j < asciifxObj.Width; j++ {
			oldValues := asciifxObj.Space[i][j]
			newValues := getClosestColor(oldValues, asciifxObj.AsciifyChoice.GetRGBColors(), asciifxObj.AsciifyChoice.GetGrayColors())

			asciifxObj.Space[i][j] = newValues

			errors := []int{
				int(oldValues.R) - int(newValues.R),
				int(oldValues.G) - int(newValues.G),
				int(oldValues.B) - int(newValues.B),
				int(oldValues.I) - int(newValues.I),
			}

			if j+1 < asciifxObj.Width {
				asciifxObj.Space[i][j+1].R = asciifx.Clamp(int(asciifxObj.Space[i][j+1].R) + int(math.Round(float64(errors[0])*7.0/16.0)))
				asciifxObj.Space[i][j+1].G = asciifx.Clamp(int(asciifxObj.Space[i][j+1].G) + int(math.Round(float64(errors[1])*7.0/16.0)))
				asciifxObj.Space[i][j+1].B = asciifx.Clamp(int(asciifxObj.Space[i][j+1].B) + int(math.Round(float64(errors[2])*7.0/16.0)))
				asciifxObj.Space[i][j+1].I = asciifx.Clamp(int(asciifxObj.Space[i][j+1].I) + int(math.Round(float64(errors[3])*7.0/16.0)))
			}

			if i+1 < asciifxObj.Height {
				if j > 0 {
					asciifxObj.Space[i+1][j-1].R = asciifx.Clamp(int(asciifxObj.Space[i+1][j-1].R) + int(math.Round(float64(errors[0])*3.0/16.0)))
					asciifxObj.Space[i+1][j-1].G = asciifx.Clamp(int(asciifxObj.Space[i+1][j-1].G) + int(math.Round(float64(errors[1])*3.0/16.0)))
					asciifxObj.Space[i+1][j-1].B = asciifx.Clamp(int(asciifxObj.Space[i+1][j-1].B) + int(math.Round(float64(errors[2])*3.0/16.0)))
					asciifxObj.Space[i+1][j-1].I = asciifx.Clamp(int(asciifxObj.Space[i+1][j-1].I) + int(math.Round(float64(errors[3])*3.0/16.0)))
				}

				asciifxObj.Space[i+1][j].R = asciifx.Clamp(int(asciifxObj.Space[i+1][j].R) + int(math.Round(float64(errors[0])*5.0/16.0)))
				asciifxObj.Space[i+1][j].G = asciifx.Clamp(int(asciifxObj.Space[i+1][j].G) + int(math.Round(float64(errors[1])*5.0/16.0)))
				asciifxObj.Space[i+1][j].B = asciifx.Clamp(int(asciifxObj.Space[i+1][j].B) + int(math.Round(float64(errors[2])*5.0/16.0)))
				asciifxObj.Space[i+1][j].I = asciifx.Clamp(int(asciifxObj.Space[i+1][j].I) + int(math.Round(float64(errors[3])*5.0/16.0)))

				if j+1 < asciifxObj.Width {
					asciifxObj.Space[i+1][j+1].R = asciifx.Clamp(int(asciifxObj.Space[i+1][j+1].R) + int(math.Round(float64(errors[0])*1.0/16.0)))
					asciifxObj.Space[i+1][j+1].G = asciifx.Clamp(int(asciifxObj.Space[i+1][j+1].G) + int(math.Round(float64(errors[1])*1.0/16.0)))
					asciifxObj.Space[i+1][j+1].B = asciifx.Clamp(int(asciifxObj.Space[i+1][j+1].B) + int(math.Round(float64(errors[2])*1.0/16.0)))
					asciifxObj.Space[i+1][j+1].I = asciifx.Clamp(int(asciifxObj.Space[i+1][j+1].I) + int(math.Round(float64(errors[3])*1.0/16.0)))
				}
			}
		}
	}
}
