package asciify

import (
	"fmt"
	"math"
	"sync"

	"github.com/brylleee/asciifx/asciifx"
)

type ANSICode struct {
	SupportsColor bool
	GrayColors    []uint8
	RGBColors     [][]uint8
}

func UseANSICode() ANSICode {
	const SUPPORTS_COLOR bool = true
	const GRAY_COLORS_SIZE int = 256
	var RGB_COLORS [][]uint8 = nil // nil means full support color

	grayColors := make([]uint8, GRAY_COLORS_SIZE)

	for i := 0; i < GRAY_COLORS_SIZE; i++ {
		grayColors[i] = uint8(math.Round(float64(255)/float64(GRAY_COLORS_SIZE-1))) * uint8(i)
	}

	return ANSICode{
		SupportsColor: SUPPORTS_COLOR,
		GrayColors:    grayColors,
		RGBColors:     RGB_COLORS,
	}
}

func (ansiCode ANSICode) GetGrayColors() []uint8 {
	return ansiCode.GrayColors
}

func (ansiCode ANSICode) GetRGBColors() [][]uint8 {
	return ansiCode.RGBColors
}

func (ansiCode ANSICode) Asciify(asciifxObj *asciifx.AsciiFx) []string {
	var result []string = make([]string, asciifxObj.Height)
	var line string

	var wg sync.WaitGroup

	for i := 0; i < asciifxObj.Height; i++ {
		for j := 0; j < asciifxObj.Width; j++ {
			wg.Add(1)

			go func() {
				defer wg.Done()
				line += fmt.Sprintf("\033[38;2;%d;%d;%dm██", asciifxObj.Space[i][j].R, asciifxObj.Space[i][j].G, asciifxObj.Space[i][j].B)
			}()
		}

		result[i] = line
		line = ""
	}

	wg.Wait()
	return result
}
