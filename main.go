package main

import (
	"log"

	"github.com/ataleksand/Fractal-Generator/frctl"
	"github.com/spf13/pflag"
)

var numberOfGoroutines int
var quiteMode bool
var imgBounds string
var areaBounds string
var outputFile string

func init() {
	pflag.IntVarP(&numberOfGoroutines, "tasks", "t", 1, "Number of goroutine should be positive integer")
	pflag.BoolVarP(&quiteMode, "quiet", "q", false, "True/False value")
	pflag.StringVarP(&imgBounds, "size", "s", "640x480", "Bounds of the png image")
	pflag.StringVarP(&areaBounds, "rect", "r", "-2.0:2.0:-2.0:2.0", "Bounds of the searching area")
	pflag.StringVarP(&outputFile, "output", "o", "fractal.png", "Result file name")

	pflag.Parse()
}

func main() {
	imgWidth, imgHeight, err := frctl.ParseImageBounds(imgBounds)
	if err != nil {
		log.Fatal("Could not parse image bounds.")
	}

	areaRect, err := frctl.ParseAreaBounds(areaBounds)
	if err != nil {
		log.Fatal("Could not parse area bounds.")
	}

	opt := frctl.NewOptions(numberOfGoroutines, outputFile, quiteMode)
	it := frctl.NewImageTraverser(imgWidth, imgHeight, areaRect)

	fg := frctl.NewFractalGenerator(opt, it)
	fg.Generate()
}
