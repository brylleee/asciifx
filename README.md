# asciifx
An image to text-art command-line and library converter!

![Preview Image](https://github.com/brylleee/asciifx/blob/main/preview.png?raw=true)

### How to use
You can use asciifx as a library in your already existing **Go** projects or as a command-line application.

#### Library way
1. Import the asciifx headers in your main code:
```go
import (
	"github.com/brylleee/asciifx/asciifx"
	"github.com/brylleee/asciifx/asciify"
	"github.com/brylleee/asciifx/dithering"
	"github.com/brylleee/asciifx/downsampler"
)
```
2. You can then create an asciifx object, load any image, then convert it
```go
func main() {
  // ...
  asciifxobj := &asciifx.AsciiFx{}             // create asciifx object
  err = asciifxobj.Load("images/maki.png")     // load any image
	if err != nil {                              // ensure no errors
		log.Fatal(err)
	}

  ditheringMethod = dithering.UseFloydSteinberg()             // Choose dithering algorithm
  downsamplerMethod = downsampler.UseNearestNeighbor(5)       // Choose downsampling algorithm
  asciifyMethod = asciify.UseBraille()                        // Choose output style

  result := asciifxobj.Convert(ditheringMethod, downsamplerMethod, asciifyMethod)      // Convert and store result

  // Looping over every lines in the result
	for _, v := range result {
		fmt.Println(v)
	}
}
```

#### Command-line interface
You can download the binary (if I haven't provided one yet, build it yourself please)\
You must have [Golang](https://go.dev/) installed in your system to build it.
1. Clone the repository
```sh
$ git clone https://github.com/brylleee/asciifx && cd asciifx
```
2. Simply build the main source code
```sh
$ go build main.go && ./main
```

#### Dedicated website
I am hosting a website that you can use to convert any images to asciifx conveniently.\
You can also see examples there.\
Please check it out [here](https://kairooo.online/asciifx)!

---

## Documentation
Link to more documentation is [here](https://pkg.go.dev/github.com/brylleee/asciifx@v0.0.0-20241117152057-392af225bc4b).
