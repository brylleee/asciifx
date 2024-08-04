package asciify

import (
	"math"

	"github.com/brylleee/asciifx/asciifx"
)

type Blocks struct {
	SupportsColor bool
	GrayColors    []uint8
	RGBColors     [][]uint8
}

func UseBlocks() Blocks {
	const SUPPORTS_COLOR bool = false
	const GRAY_COLORS_SIZE int = 4
	var RGB_COLORS [][]uint8 = [][]uint8{{0, 0, 0}}

	grayColors := make([]uint8, GRAY_COLORS_SIZE)

	for i := 0; i < GRAY_COLORS_SIZE; i++ {
		grayColors[i] = uint8(math.Round(float64(255)/float64(GRAY_COLORS_SIZE-1))) * uint8(i)
	}

	return Blocks{
		SupportsColor: SUPPORTS_COLOR,
		GrayColors:    grayColors,
		RGBColors:     RGB_COLORS,
	}
}

func (blocks Blocks) GetGrayColors() []uint8 {
	return blocks.GrayColors
}

func (blocks Blocks) GetRGBColors() [][]uint8 {
	return blocks.RGBColors
}

func (blocks Blocks) Asciify(asciifx *asciifx.AsciiFx) []string {
	var result []string = make([]string, asciifx.Height)
	var line string
	blockValues := []string{"░", "▒", "▓", "█"}

	for idxi, i := range asciifx.Space {
		for idxj := range i {
			line += blockValues[asciifx.Space[idxi][idxj].I>>6] + blockValues[asciifx.Space[idxi][idxj].I>>6]
		}

		result[idxi] = line
		line = ""
	}

	return result
}
