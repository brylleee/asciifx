package asciify

import "github.com/brylleee/asciifx/asciifx"

type Braille struct {
	SupportsColor bool
	GrayColors    []uint8
	RGBColors     [][]uint8
}

func UseBraille() Braille {
	const SUPPORTS_COLOR bool = false
	const GRAY_COLORS_SIZE int = 2
	var RGB_COLORS [][]uint8 = [][]uint8{}

	grayColors := make([]uint8, GRAY_COLORS_SIZE)

	for i := 0; i < GRAY_COLORS_SIZE; i++ {
		grayColors[i] = uint8(255/GRAY_COLORS_SIZE) * uint8(i+1)
	}

	return Braille{
		SupportsColor: SUPPORTS_COLOR,
		GrayColors:    grayColors,
		RGBColors:     RGB_COLORS,
	}
}

func (braille Braille) Self() Braille {
	return braille
}

func (braille Braille) Asciify(asciifx *asciifx.AsciiFx) [][]rune {
	var result [][]rune = make([][]rune, asciifx.Height/4)
	var line []rune = make([]rune, 0)

	braille_unicode_offset := 0x2800
	corresponding_char_offset := 0

	for i, ix := 0, 0; i+4 < asciifx.Height; func() { i += 4; ix++ }() {
		for j := 0; j+2 < asciifx.Width; j += 2 {
			corresponding_char_offset += int(asciifx.Space[i][j].I>>7) + int((asciifx.Space[i][j+1].I>>7)<<3) +
				int((asciifx.Space[i+1][j].I>>7)<<1) + int((asciifx.Space[i+1][j+1].I>>7)<<4) +
				int((asciifx.Space[i+2][j].I>>7)<<2) + int((asciifx.Space[i+2][j+1].I>>7)<<5) +
				int((asciifx.Space[i+3][j].I>>7)<<6) + int((asciifx.Space[i+3][j+1].I>>7)<<7)

			line = append(line, rune(braille_unicode_offset+corresponding_char_offset))
			corresponding_char_offset = 0
		}

		result[ix] = line
		line = []rune{}
	}

	return result
}
