package asciifx

type Dithering interface {
	Dither(asciifx *AsciiFx)
}

func Threshold(value uint8, thresholdValue uint8) uint8 {
	if value > thresholdValue {
		return 255
	} else {
		return 0
	}
}
