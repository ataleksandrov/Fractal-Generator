package frctl

import (
	"fmt"
	"math/cmplx"
	"strconv"
	"strings"
)

const MAX_ITERATIONS = 800

// predefined func
func applyFunc(c, z complex128) complex128 {
	return cmplx.Exp(cmplx.Cos(c * z))
}

func ParseImageBounds(bounds string) (int, int, error) {
	buff := strings.Split(bounds, "x")
	h, err := strconv.Atoi(buff[0])
	if err == nil {
		w, err := strconv.Atoi(buff[1])
		return h, w, err
	}

	return 0, 0, err
}

func ParseAreaBounds(bounds string) ([4]float64, error) {
	var b [4]float64
	buff := strings.Split(bounds, ":")
	var err error
	for i := 0; i < 4 && err == nil; i++ {
		b[i], err = strconv.ParseFloat(buff[i], 64)
	}

	return b, err
}

func printAdditionalInfo(o *options, format string, args ...interface{}) {
	if !o.quietMode {
		fmt.Printf(format, args...)
	}
}
