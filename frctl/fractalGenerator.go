package frctl

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
	"sync"
	"time"
)

type fractalGenerator struct {
	opt   *options
	imgTr *imageTraverser
}

func NewFractalGenerator(o *options, i *imageTraverser) *fractalGenerator {
	return &fractalGenerator{
		opt:   o,
		imgTr: i,
	}
}

func (fg *fractalGenerator) Generate() {
	totalTime := time.Now()

	img := image.NewRGBA(image.Rect(0, 0, fg.imgTr.imgWidth, fg.imgTr.imgHeight))
	rowsPerGoroutine := float64(img.Bounds().Dy()) / float64(fg.opt.numberOfGoroutines)

	wg := &sync.WaitGroup{}
	wg.Add(fg.opt.numberOfGoroutines)

	var sRow float64
	for i := 0; i < fg.opt.numberOfGoroutines; i++ {
		sRow = float64(i) * rowsPerGoroutine
		printAdditionalInfo(fg.opt, "Started: goroutine %d \n", int(i))

		go func(sRow float64, i int) {
			t := time.Now()
			fg.imgTr.traverseImg(sRow, sRow+rowsPerGoroutine, img)
			printAdditionalInfo(fg.opt, "Finished: goroutine %d with time %v \n", i, time.Since(t))
			wg.Done()
		}(sRow, int(i))
	}
	wg.Wait()

	f, err := os.Create(fg.opt.outputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if err := png.Encode(f, img); err != nil {
		log.Fatal(err)
	}

	printAdditionalInfo(fg.opt, "Number of goroutines: %d \n", fg.opt.numberOfGoroutines)
	fmt.Printf("Total execution time: %v", time.Since(totalTime))
}
