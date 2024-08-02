package asciifx

import (
	"log"
	"math"
	"os"

	"image"
	_ "image/jpeg"
	_ "image/png"
)

// RGBI represents a single pixel's color values
// Red, Green, Blue and Intensity (the grayscale value of the RGB)
type RGBI struct {
	R uint8
	G uint8
	B uint8
	I uint8
}

// AsciiFx represents a single image to be processed to ascii art
// It has information about image path, its RGBI values, etc.
type AsciiFx struct {
	Image   image.Image
	ImgPath string
	Space   [][]RGBI

	asciifychoice     Asciify
	ditheralgochoice  Dithering
	downsamplerchoice Downsampler

	Width  int
	Height int
}

// Load loads an image supplied as a string image path to an AsciiFx struct
// It reads the image then initializes the width and height as well as
// allocating the space it will work on
func (asciifx *AsciiFx) Load(path string) error {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	// img, fmt, err
	image, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	asciifx.Image = image
	asciifx.Width = image.Bounds().Size().X
	asciifx.Height = image.Bounds().Size().Y

	asciifx.allocateSpace()

	return nil
}

// Conver converts the image loaded in AsciiFx to ascii arts employing user chose dithering alogrithm
// and asciify method. It returns the result as a 2D array of runes
func (asciifx *AsciiFx) Convert(ditherAlgorithm Dithering, downsampler Downsampler, asciify Asciify) [][]rune {
	asciifx.extractColors()
	asciifx.ditheralgochoice = ditherAlgorithm
	asciifx.downsamplerchoice = downsampler
	asciifx.asciifychoice = asciify

	asciifx.ditheralgochoice.Dither(asciifx)
	asciifx.downsamplerchoice.Downsample(asciifx)
	return asciifx.asciifychoice.Asciify(asciifx)
}

// extractColors extracts every color in the image loaded in AsciiFx and stores it in the Space property of AsciiFx
func (asciifx *AsciiFx) extractColors() {
	for i := 0; i < asciifx.Height; i++ {
		for j := 0; j < asciifx.Width; j++ {
			r, g, b, _ := asciifx.Image.At(j, i).RGBA()
			asciifx.Space[i][j] = RGBI{R: uint8(r >> 8), G: uint8(g >> 8), B: uint8(b >> 8),
				I: uint8(math.Round(0.299*float64(r))) + uint8(math.Round(0.587*float64(g))) + uint8(math.Round(0.114*float64(b)))}
		}
	}
}

// allocateSpace allocates Space values according to the width and height of the image loaded in AsciiFx
func (asciifx *AsciiFx) allocateSpace() {
	asciifx.Space = make([][]RGBI, asciifx.Height)
	for i := 0; i < asciifx.Height; i++ {
		asciifx.Space[i] = make([]RGBI, asciifx.Width)
	}
}
