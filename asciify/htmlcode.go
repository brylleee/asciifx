package asciify

import (
	"math"

	"github.com/brylleee/asciifx/asciifx"
)

type HTMLCode struct {
	SupportsColor bool
	GrayColors    []uint8
	RGBColors     [][]uint8
}

func UseHTMLCode() Braille {
	const SUPPORTS_COLOR bool = true
	const GRAY_COLORS_SIZE int = 256
	var RGB_COLORS [][]uint8 = [][]uint8{{0, 0, 0}}

	grayColors := make([]uint8, GRAY_COLORS_SIZE)

	for i := 0; i < GRAY_COLORS_SIZE; i++ {
		grayColors[i] = uint8(math.Round(float64(255)/float64(GRAY_COLORS_SIZE-1))) * uint8(i)
	}

	return Braille{
		SupportsColor: SUPPORTS_COLOR,
		GrayColors:    grayColors,
		RGBColors:     RGB_COLORS,
	}
}

func (htmlCode HTMLCode) GetGrayColors() []uint8 {
	return htmlCode.GrayColors
}

func (htmlCode HTMLCode) GetRGBColors() [][]uint8 {
	return htmlCode.RGBColors
}

func (htmlCode HTMLCode) Asciify(asciifx *asciifx.AsciiFx) [][]rune {
	var result [][]rune = make([][]rune, asciifx.Height/4)
	var line []rune = make([]rune, 0)

	for i := 0; i < asciifx.Height; i++ {
		for j := 0; j < asciifx.Width; j++ {

		}

		result[i] = line
		line = []rune{}
	}

	return result
}
