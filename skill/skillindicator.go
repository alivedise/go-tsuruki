package skill

import (
	"fmt"
	"image"
	"image/color"
	"math"

	"github.com/alivedise/tsuruki/interfaces"
	"github.com/alivedise/tsuruki/point"
	"github.com/hajimehoshi/ebiten/v2"
)

func calculateSlope(p1, p2 point.Point) (float64, error) {
	if p1.X == p2.X {
		return 0, fmt.Errorf("无法计算斜率，两点的 x 值相同")
	}

	slope := (p2.Y - p1.Y) / (p2.X - p1.X)
	return slope, nil
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

func calculateRectanglePoints(start, end point.Point, width float64, height float64) [5]point.Point {
	centerX := (start.X + end.X) / 2
	centerY := (start.Y + end.Y) / 2

	slope, _ := calculateSlope(start, end)
	angle := math.Atan(slope) * 180 / math.Pi

	return CalculateRectanglePoints2(centerX, centerY, height, width, angle)
}

type SkillIndicator struct {
	image       *ebiten.Image
	renderType  string
	width       float64
	height      float64
	source      point.Point
	destination point.Point
	points      [5]point.Point
}

func NewSkillIndicator(renderType string, width float64, height float64) *SkillIndicator {
	image := ebiten.NewImage(3, 3)
	image.Fill(color.RGBA{189, 22, 64, 1})
	return &SkillIndicator{
		renderType: renderType,
		width:      width,
		height:     height,
		image:      image,
	}
}

func (si *SkillIndicator) GetSkillSize() float64 {
	return si.height
}

func (si *SkillIndicator) SetRectangleData(x1, y1, x2, y2 float64) {
	si.source = point.Point{X: x1, Y: y1}
	si.destination = point.Point{X: x2, Y: y2}

	si.points = calculateRectanglePoints(point.Point{X: si.source.X, Y: si.source.Y}, point.Point{X: si.destination.X, Y: si.destination.Y}, si.width, si.height)
}

func (si *SkillIndicator) Clear() {
	si.source = point.Point{}
	si.destination = point.Point{}
	si.points = [5]point.Point{}
}

func (si *SkillIndicator) Draw(screen *ebiten.Image, c interfaces.Creature) {
	if si.source.X == 0 {
		return
	}
	op := &ebiten.DrawTrianglesOptions{}
	op.Address = ebiten.AddressUnsafe
	scale := ebiten.DeviceScaleFactor()
	indices := []uint16{}
	for i := 0; i < 4; i++ {
		indices = append(indices, uint16(i), uint16(i+1)%uint16(4), uint16(4))
	}
	screen.DrawTriangles([]ebiten.Vertex{
		{DstX: float32(si.points[0].X * scale), DstY: float32(si.points[0].Y * scale), SrcX: 0, SrcY: 0, ColorR: 1, ColorG: 1, ColorB: 1, ColorA: 1},
		{DstX: float32(si.points[1].X * scale), DstY: float32(si.points[1].Y * scale), SrcX: 0, SrcY: 0, ColorR: 1, ColorG: 1, ColorB: 1, ColorA: 1},
		{DstX: float32(si.points[2].X * scale), DstY: float32(si.points[2].Y * scale), SrcX: 0, SrcY: 0, ColorR: 1, ColorG: 1, ColorB: 1, ColorA: 1},
		{DstX: float32(si.points[3].X * scale), DstY: float32(si.points[3].Y * scale), SrcX: 0, SrcY: 0, ColorR: 1, ColorG: 1, ColorB: 1, ColorA: 1},
		{DstX: float32(si.points[4].X * scale), DstY: float32(si.points[4].Y * scale), SrcX: 0, SrcY: 0, ColorR: 1, ColorG: 1, ColorB: 1, ColorA: 1},
	}, indices, si.image.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image), op)
}
