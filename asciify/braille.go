package asciify

import "github.com/brylleee/asciifx/asciifx"

type Braille struct{}

func UseBraille() asciifx.AsciifyMethod {
	return asciifx.AsciifyMethod{
		SupportsColor: false,
		ColorRange:    asciifx.RGBI{R: 0, G: 0, B: 0, I: 2},
	}
}

func (braille Braille) Asciify(asciifx *asciifx.AsciiFx) [][]rune {
	var result [][]rune = make([][]rune, asciifx.Height)
	var line []rune = make([]rune, 0)

	braille_unicode_offset := 0x2800
	corresponding_char_offset := 0

	for i, ix := 0, 0; i+4 < asciifx.Height; func() { i += 4; ix++ }() {
		for j := 0; j+2 < asciifx.Width; j += 2 {
			corresponding_char_offset += int(asciifx.Space[i][j].I>>7) + int((asciifx.Space[i][j+1].I>>7)<<3) +
				int((asciifx.Space[+1][j].I>>7)<<1) + int((asciifx.Space[+1][j+1].I>>7)<<4) +
				int((asciifx.Space[+2][j].I>>7)<<2) + int((asciifx.Space[+2][j+1].I>>7)<<5) +
				int((asciifx.Space[+3][j].I>>7)<<6) + int((asciifx.Space[+3][j+1].I>>7)<<7)

			line = append(line, rune(braille_unicode_offset+corresponding_char_offset))
			corresponding_char_offset = 0
		}

		result[ix] = line
		line = []rune{}
	}

	return result
}
