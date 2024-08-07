package asciify

import (
	"math"

	"github.com/brylleee/asciifx/asciifx"
)

type Moons struct {
	SupportsColor bool
	GrayColors    []uint8
	RGBColors     [][]uint8
}

func UseMoons() Moons {
	const SUPPORTS_COLOR bool = false
	const GRAY_COLORS_SIZE int = 2
	var RGB_COLORS [][]uint8 = [][]uint8{{0, 0, 0}}

	grayColors := make([]uint8, GRAY_COLORS_SIZE)

	for i := 0; i < GRAY_COLORS_SIZE; i++ {
		grayColors[i] = uint8(math.Round(float64(255)/float64(GRAY_COLORS_SIZE-1))) * uint8(i)
	}

	return Moons{
		SupportsColor: SUPPORTS_COLOR,
		GrayColors:    grayColors,
		RGBColors:     RGB_COLORS,
	}
}

func (moons Moons) GetGrayColors() []uint8 {
	return moons.GrayColors
}

func (moons Moons) GetRGBColors() [][]uint8 {
	return moons.RGBColors
}

func (moons Moons) Asciify(asciifx *asciifx.AsciiFx) []string {
	var result []string = make([]string, asciifx.Height/4)
	var line string

	moonValues := map[[4]uint]string{
		{0, 0, 0, 0}: "🌑",
		{1, 0, 0, 1}: "🌑",
		{0, 0, 0, 1}: "🌒",
		{0, 0, 1, 0}: "🌒",
		{0, 0, 1, 1}: "🌓",
		{0, 1, 0, 1}: "🌓",
		{0, 1, 1, 1}: "🌔",
		{1, 0, 1, 1}: "🌔",
		{0, 1, 1, 0}: "🌕",
		{1, 1, 1, 1}: "🌕",
		{1, 1, 0, 1}: "🌖",
		{1, 1, 1, 0}: "🌖",
		{1, 1, 0, 0}: "🌗",
		{1, 0, 1, 0}: "🌗",
		{1, 0, 0, 0}: "🌘",
		{0, 1, 0, 0}: "🌘",
	}

	for i, ix := 0, 0; i+4 < asciifx.Height; func() { i += 4; ix++ }() {
		for j := 0; j < asciifx.Width-4; j += 4 {
			moonValue := [4]uint{
				uint(asciifx.Space[i][j].I >> 7),
				uint(asciifx.Space[i][j+1].I >> 7),
				uint(asciifx.Space[i][j+2].I >> 7),
				uint(asciifx.Space[i][j+3].I >> 7),
			}

			line += moonValues[moonValue]
		}

		result[ix] = line
		line = ""
	}

	return result
}
