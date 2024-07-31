package asciify

import "asciifx/asciifx"

type moons struct {
	SupportsColor bool
	ColorRange    asciifx.RGBI
}

func UseMoons() moons {
	return moons{
		SupportsColor: false,
		ColorRange:    asciifx.RGBI{R: 0, G: 0, B: 0, I: 2},
	}
}

func (moons moons) Asciify(asciifx *asciifx.AsciiFx) [][]rune {
	var result [][]rune = make([][]rune, asciifx.Height)
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
