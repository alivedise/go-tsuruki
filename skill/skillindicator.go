package skill

import (
	"fmt"
	"image"
	"image/color"
	"math"

	"github.com/alivedise/tsuruki/interfaces"
	"github.com/alivedise/tsuruki/point"
	"github.com/alivedise/tsuruki/utils"
	"github.com/fogleman/gg"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type SkillIndicator struct {
	image       *ebiten.Image
	renderType  string
	width       float64
	height      float64
	radius      float64
	startAngle  float64
	endAngle    float64
	source      point.Point
	destination point.Point
	points      [5]point.Point
	color       color.Color
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

	si.points = utils.CalculateRectanglePoints(point.Point{X: si.source.X, Y: si.source.Y}, point.Point{X: si.destination.X, Y: si.destination.Y}, si.width, si.height)
}

func (si *SkillIndicator) SetSectorData(x, y, r, startAngle, endAngle float64) {
	si.renderType = "sector"
	si.source = point.Point{X: x, Y: y}
	si.radius = r
	si.startAngle = startAngle
	si.endAngle = endAngle
}

func (si *SkillIndicator) Clear() {
	si.source = point.Point{X: -1000, Y: -1000}
	si.destination = point.Point{}
	si.points = [5]point.Point{}
	si.startAngle = 0
	si.endAngle = 0
	si.radius = 0
}

func (si *SkillIndicator) DrawRectangle(screen *ebiten.Image, c interfaces.Creature) {
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

func (si *SkillIndicator) DrawSector(screen *ebiten.Image, c interfaces.Creature) {
	// TODO
	op := &ebiten.DrawImageOptions{}
	width := 250
	height := 250
	radius := 100.0
	centerX, centerY := float64(width/2), float64(height/2)
	dc := gg.NewContext(width, height)
	startRad := gg.Radians(si.startAngle)
	endRad := gg.Radians(si.endAngle)

	// 绘制扇形
	dc.MoveTo(centerX, centerY)
	dc.LineTo(centerX+math.Cos(startRad)*radius, centerY+math.Sin(startRad)*radius)
	dc.DrawArc(centerX, centerY, radius, startRad, endRad)
	dc.LineTo(centerX, centerY)
	if si.color != nil {
		r, g, b, _ := si.color.RGBA()
		dc.SetRGB(float64(r), float64(g), float64(b))
	}
	dc.SetRGB(55, 120, 0)
	dc.FillPreserve()
	dc.Stroke()

	// 显示结果
	dc.SavePNG("sector.png")
	img, _, err := ebitenutil.NewImageFromFile("sector.png")
	if err != nil {
		fmt.Println(err)
	}

	dw := img.Bounds().Size().X
	dh := img.Bounds().Size().Y
	op.GeoM.Translate(-float64(dw)/2, -float64(dh)/2)
	op.GeoM.Rotate(-c.GetInfo().GetFaceAngle() + (135)*math.Pi/180)
	op.GeoM.Translate(si.source.X, si.source.Y)
	scale := ebiten.DeviceScaleFactor()
	op.GeoM.Scale(scale, scale)
	op.Filter = ebiten.FilterLinear
	screen.DrawImage(img, op)
}

func (si *SkillIndicator) Draw(screen *ebiten.Image, c interfaces.Creature) {
	current := c.GetRotation().Current().State()
	if current.Is("precast") || current.Is("backwing") {
		return
	}
	if si.renderType == "sector" {
		si.DrawSector(screen, c)
	} else {
		si.DrawRectangle(screen, c)
	}
}
