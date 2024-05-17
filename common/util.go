package common

import (
	"strconv"
	"strings"
)

func Convert2XY(value string) (int, int) {
	scale := strings.Split(value, ",")
	x, _ := strconv.Atoi(scale[0])
	y, _ := strconv.Atoi(scale[1])

	return x, y
}
