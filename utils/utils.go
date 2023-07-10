package utils

import "math"

func Distance(x1, y1, x2, y2 float64) float64 {
	dx := float64(x2 - x1)
	dy := float64(y2 - y1)
	return math.Sqrt(dx*dx + dy*dy)
}

func Interpolate(x1, y1, x2, y2, length float64) (float64, float64) {
	dx := x2 - x1
	dy := y2 - y1
	olength := Distance(x1, y1, x2, y2)

	offsetX := (dx / olength) * length
	offsetY := (dy / olength) * length

	newX := x1 + offsetX
	newY := y1 + offsetY
	return newX, newY
}
