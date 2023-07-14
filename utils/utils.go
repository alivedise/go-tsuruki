package utils

import (
	"fmt"
	"math"

	"github.com/alivedise/tsuruki/point"
)

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

func InterpolateByRatio(x1, y1, x2, y2, ratio float64) (float64, float64) {
	dx := x2 - x1
	dy := y2 - y1
	offsetX := dx * ratio
	offsetY := dy * ratio

	newX := x1 + offsetX
	newY := y1 + offsetY
	return newX, newY
}

func CalculateSlope(p1, p2 point.Point) (float64, error) {
	if p1.X == p2.X {
		return 0, fmt.Errorf("无法计算斜率，两点的 x 值相同")
	}

	slope := (p2.Y - p1.Y) / (p2.X - p1.X)
	return slope, nil
}

func CalculateRadian(x1, y1, x2, y2 float64) float64 {

	return math.Atan2(x2-x1, y2-y1)
}

func CalculateRectanglePoints2(x, y, l, w, s float64) [5]point.Point {
	radians := math.Pi * s / 180.0
	halfLength := l / 2.0
	halfWidth := w / 2.0

	// 计算矩形的四个角点坐标
	// 第一个角点（左上角）
	x1 := x - halfLength*math.Cos(radians) + halfWidth*math.Sin(radians)
	y1 := y - halfLength*math.Sin(radians) - halfWidth*math.Cos(radians)

	// 第二个角点（右上角）
	x2 := x + halfLength*math.Cos(radians) + halfWidth*math.Sin(radians)
	y2 := y + halfLength*math.Sin(radians) - halfWidth*math.Cos(radians)

	// 第三个角点（右下角）
	x3 := x + halfLength*math.Cos(radians) - halfWidth*math.Sin(radians)
	y3 := y + halfLength*math.Sin(radians) + halfWidth*math.Cos(radians)

	// 第四个角点（左下角）
	x4 := x - halfLength*math.Cos(radians) - halfWidth*math.Sin(radians)
	y4 := y - halfLength*math.Sin(radians) + halfWidth*math.Cos(radians)

	return [5]point.Point{{X: x1, Y: y1}, {X: x2, Y: y2}, {X: x3, Y: y3}, {X: x4, Y: y4}, {X: x, Y: y}}
}

func CalculateRectanglePoints(start, end point.Point, width float64, height float64) [5]point.Point {
	centerX := (start.X + end.X) / 2
	centerY := (start.Y + end.Y) / 2

	slope, _ := CalculateSlope(start, end)
	angle := math.Atan(slope) * 180 / math.Pi

	return CalculateRectanglePoints2(centerX, centerY, height, width, angle)
}
