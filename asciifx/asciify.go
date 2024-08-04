package asciifx

type Asciify interface {
	GetColorRange() RGBI
	Asciify(asciifx *AsciiFx) [][]rune
}
