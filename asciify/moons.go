package asciify

import "github.com/brylleee/asciifx/asciifx"

type Moons struct {
	SupportsColor bool
	GrayColors    []uint8
	RGBColors     [][]uint8
}

func UseMoons() Moons {
	const SUPPORTS_COLOR bool = false
	const GRAY_COLORS_SIZE int = 2
	var RGB_COLORS [][]uint8 = [][]uint8{}

	grayColors := make([]uint8, GRAY_COLORS_SIZE)

	for i := 0; i < GRAY_COLORS_SIZE; i++ {
		grayColors[i] = uint8(255/GRAY_COLORS_SIZE) * uint8(i+1)
	}

	return Moons{
		SupportsColor: SUPPORTS_COLOR,
		GrayColors:    grayColors,
		RGBColors:     RGB_COLORS,
	}
}

func (moons Moons) Self() Moons {
	return moons
}

func (moons Moons) Asciify(asciifx *asciifx.AsciiFx) [][]rune {
	var result [][]rune = make([][]rune, asciifx.Height/4)
	var line []rune = make([]rune, 0)

	moonValues := map[[4]uint]rune{
		{0, 0, 0, 0}: '🌑',
		{1, 0, 0, 1}: '🌑',
		{0, 0, 0, 1}: '🌒',
		{0, 0, 1, 0}: '🌒',
		{0, 0, 1, 1}: '🌓',
		{0, 1, 0, 1}: '🌓',
		{0, 1, 1, 1}: '🌔',
		{1, 0, 1, 1}: '🌔',
		{0, 1, 1, 0}: '🌕',
		{1, 1, 1, 1}: '🌕',
		{1, 1, 0, 1}: '🌖',
		{1, 1, 1, 0}: '🌖',
		{1, 1, 0, 0}: '🌗',
		{1, 0, 1, 0}: '🌗',
		{1, 0, 0, 0}: '🌘',
		{0, 1, 0, 0}: '🌘',
	}

	for i, ix := 0, 0; i+4 < asciifx.Height; func() { i += 4; ix++ }() {
		for j := 0; j < asciifx.Width-4; j += 4 {
			moonValue := [4]uint{
				uint(asciifx.Space[i][j].I >> 7),
				uint(asciifx.Space[i][j+1].I >> 7),
				uint(asciifx.Space[i][j+2].I >> 7),
				uint(asciifx.Space[i][j+3].I >> 7),
			}

			line = append(line, moonValues[moonValue])
		}

		result[ix] = line
		line = []rune{}
	}

	return result
}
