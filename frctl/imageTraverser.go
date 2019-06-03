package frctl

import (
	"image"
	"image/color"
	"math/cmplx"
)

type imageTraverser struct {
	imgWidth, imgHeight int
	xa, xb, ya, yb      float64
}

func NewImageTraverser(imgWidth, imgHeight int, rect [4]float64) *imageTraverser {
	return &imageTraverser{
		imgWidth:  imgWidth,
		imgHeight: imgHeight,
		xa:        rect[0],
		xb:        rect[1],
		ya:        rect[2],
		yb:        rect[3],
	}
}

func (imgTr *imageTraverser) traverseImg(sRow, eRow float64, img *image.RGBA) {
	var cReal, cImag float64

	var iterator int
	for row := sRow; row < eRow; row++ {
		cImag = (row - float64(img.Bounds().Dy())/imgTr.yb) * ((imgTr.yb - imgTr.ya) / float64(img.Bounds().Dy()))

		for col := 0.0; col < float64(img.Bounds().Dx()); col++ {
			cReal = (col - float64(img.Bounds().Dx())/imgTr.xb) * ((imgTr.xb - imgTr.xa) / float64(img.Bounds().Dx()))

			c := complex(cReal, cImag)
			z := complex(0.0, 0.0)

			for iterator = 0; iterator < MAX_ITERATIONS; iterator++ {
				if cmplx.Abs(z) > (imgTr.xb - imgTr.xa) {
					break
				}
				z = applyFunc(c, z)
			}

			if iterator < MAX_ITERATIONS {
				img.SetRGBA(int(col), int(row), color.RGBA{uint8(iterator * 10), uint8(iterator * 2), 0, 255})
			} else {
				img.SetRGBA(int(col), int(row), color.RGBA{0, 0, 0, 255})
			}
		}
	}
}
