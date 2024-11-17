package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jessevdk/go-flags"

	"github.com/brylleee/asciifx/asciifx"
	"github.com/brylleee/asciifx/asciify"
	"github.com/brylleee/asciifx/dithering"
	"github.com/brylleee/asciifx/downsampler"
)

var banner string = `
 █████╗ ███████╗ ██████╗██╗██╗███████╗██╗  ██╗
██╔══██╗██╔════╝██╔════╝██║██║██╔════╝╚██╗██╔╝
███████║███████╗██║     ██║██║█████╗   ╚███╔╝ 
██╔══██║╚════██║██║     ██║██║██╔══╝   ██╔██╗ 
██║  ██║███████║╚██████╗██║██║██║     ██╔╝ ██╗
╚═╝  ╚═╝╚══════╝ ╚═════╝╚═╝╚═╝╚═╝     ╚═╝  ╚═╝
by ka1ro (brylleee) v1.0 -- A1SBERG
`

var options struct {
	Asciify     string `long:"asciify" short:"a" value-name:"braille" description:"Asciify method to use (ansicode, blocks, braille, htmlcode, moons)"`
	Dithering   string `long:"dithering" short:"d" value-name:"floydsteinberg" description:"Dithering method to employ (bayermatrix, floydsteinberg, random)"`
	Downsampler string `long:"downsampler" short:"s" value-name:"nearestneighbor" description:"Downsampling method to apply (default: nearestneighbor)"`
	Ratio       int    `long:"ratio" short:"r" default:"3" value-name:"3" description:"Ratio on how much to shrink the output"`
	// Raw         bool   `long:"raw" short:"w" description:"Enable raw mode (do not format)"`
	Help bool `long:"help" short:"h" description:"Show the help manual"`
}

func main() {
	asciifxobj := &asciifx.AsciiFx{}
	parser := flags.NewParser(&options, flags.Default)
	parser.Usage = os.Args[0] + " <image> [OPTIONS]"

	// Overriding default help flag to add banner first
	if len(os.Args) == 1 || (len(os.Args) > 1 && (os.Args[1] == "-h" || os.Args[1] == "--help")) {
		fmt.Println(banner)
		parser.WriteHelp(os.Stdout)
		return
	}
	_, err := parser.Parse()
	if err != nil {
		log.Fatal(err)
	}

	err = asciifxobj.Load(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	var asciifyMethod asciifx.Asciify
	var ditheringMethod asciifx.Dithering
	var downsamplerMethod asciifx.Downsampler

	switch options.Asciify {
	case "ansicode":
		asciifyMethod = asciify.UseANSICode()
	case "blocks":
		asciifyMethod = asciify.UseBlocks()
	case "braille":
		asciifyMethod = asciify.UseBraille()
	case "htmlcode":
		asciifyMethod = asciify.UseHTMLCode()
	case "moons":
		asciifyMethod = asciify.UseMoons()
	default:
		asciifyMethod = asciify.UseBraille()
	}

	switch options.Dithering {
	case "bayermatrix":
		ditheringMethod = dithering.UseBayerMatrix()
	case "floydsteinberg":
		ditheringMethod = dithering.UseFloydSteinberg()
	case "random":
		ditheringMethod = dithering.UseRandom()
	default:
		ditheringMethod = dithering.UseFloydSteinberg()
	}

	switch options.Downsampler {
	case "nearestneighbor":
		downsamplerMethod = downsampler.UseNearestNeighbor(options.Ratio)
	default:
		downsamplerMethod = downsampler.UseNearestNeighbor(options.Ratio)
	}

	result := asciifxobj.Convert(ditheringMethod, downsamplerMethod, asciifyMethod)

	for _, v := range result {
		fmt.Println(v)
	}
}
