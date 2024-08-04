package asciify

import (
	"fmt"
	"math"

	"github.com/brylleee/asciifx/asciifx"
)

type HTMLCode struct {
	SupportsColor bool
	GrayColors    []uint8
	RGBColors     [][]uint8
}

func UseHTMLCode() HTMLCode {
	const SUPPORTS_COLOR bool = true
	const GRAY_COLORS_SIZE int = 256
	var RGB_COLORS [][]uint8 = nil // nil means full support color

	grayColors := make([]uint8, GRAY_COLORS_SIZE)

	for i := 0; i < GRAY_COLORS_SIZE; i++ {
		grayColors[i] = uint8(math.Round(float64(255)/float64(GRAY_COLORS_SIZE-1))) * uint8(i)
	}

	return HTMLCode{
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

func (htmlCode HTMLCode) Asciify(asciifxObj *asciifx.AsciiFx) []string {
	var result []string = make([]string, asciifxObj.Height)
	var line string

	for i := 0; i < asciifxObj.Height; i++ {
		for j := 0; j < asciifxObj.Width; j++ {
			line += fmt.Sprintf("<span style=\"color:#%x%x%x;\">██</span>", asciifxObj.Space[i][j].R, asciifxObj.Space[i][j].G, asciifxObj.Space[i][j].B)
		}

		result[i] = line
		line = ""
	}

	return result
}
