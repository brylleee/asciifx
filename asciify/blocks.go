package asciify

import "github.com/brylleee/asciifx/asciifx"

type Blocks struct {
	SupportsColor bool
	GrayColors    []uint8
	RGBColors     [][]uint8
}

func UseBlocks() Blocks {
	const SUPPORTS_COLOR bool = false
	const GRAY_COLORS_SIZE int = 4
	var RGB_COLORS [][]uint8 = [][]uint8{}

	grayColors := make([]uint8, GRAY_COLORS_SIZE)

	for i := 0; i < GRAY_COLORS_SIZE; i++ {
		grayColors[i] = uint8(255/GRAY_COLORS_SIZE) * uint8(i+1)
	}

	return Blocks{
		SupportsColor: SUPPORTS_COLOR,
		GrayColors:    grayColors,
		RGBColors:     RGB_COLORS,
	}
}

func (blocks Blocks) Self() Blocks {
	return blocks
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
