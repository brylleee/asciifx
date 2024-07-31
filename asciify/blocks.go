package asciify

import "github.com/brylleee/asciifx/asciifx"

type Blocks struct{}

func UseBlocks() asciifx.AsciifyMethod {
	return asciifx.AsciifyMethod{
		SupportsColor: false,
		ColorRange:    asciifx.RGBI{R: 0, G: 0, B: 0, I: 4},
	}
}

func (blocks Blocks) Asciify(asciifx *asciifx.AsciiFx) [][]rune {
	var result [][]rune = make([][]rune, asciifx.Height)
	var line []rune = make([]rune, 0)
	blockValues := []rune{'░', '▒', '▓', '█'}

	for idxi, i := range asciifx.Space {
		for idxj := range i {
			line = append(line, []rune{blockValues[asciifx.Space[idxi][idxj].I>>6], blockValues[asciifx.Space[idxi][idxj].I>>6]}...)
		}

		result[idxi] = line
		line = []rune{}
	}

	return result
}
