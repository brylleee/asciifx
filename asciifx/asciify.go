package asciifx

type Asciify interface {
	GetRGBColors() [][]uint8
	GetGrayColors() []uint8
	Asciify(asciifx *AsciiFx) [][]rune
}
