package main

import (
	"math"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func main() {
	p := message.NewPrinter(language.English)
	p.Printf("%f\n", math.MaxFloat32)
}
