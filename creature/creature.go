package creature

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type CreatureInfo struct {
	X           float64
	Y           float64
	TargetX     float64
	TargetY     float64
	Image       *ebiten.Image
	Speed       float64
	ElapsedTime float64
}

var TARGET_W = 40.0
var TARGET_H = 40.0

func (ci *CreatureInfo) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	dw := ci.Image.Bounds().Size().X
	dh := ci.Image.Bounds().Size().Y
	scaleX := TARGET_W / float64(dw)
	scaleY := TARGET_H / float64(dh)
	op.GeoM.Translate((ci.X-20)/scaleX, (ci.Y-20)/scaleY)
	op.GeoM.Scale(scaleX, scaleY)
	scale := ebiten.DeviceScaleFactor()
	op.GeoM.Scale(scale, scale)
	op.Filter = ebiten.FilterLinear
	screen.DrawImage(ci.Image, op)
}

func (ci *CreatureInfo) IncreaseElapsedTime(dt float64) {
	ci.ElapsedTime += dt
}

func (ci *CreatureInfo) SetTargetPosition(x, y float64) {
	ci.TargetX = x
	ci.TargetY = y
}

func (ci *CreatureInfo) Move(x, y float64) {
	ci.X = ci.X + float64(x)
	if ci.X < 0 {
		ci.X = 0
	}
	ci.Y = ci.Y + float64(y)
	if ci.Y < 0 {
		ci.Y = 0
	}
}

func (ci *CreatureInfo) GetImage() *ebiten.Image {
	return ci.Image
}

func (ci *CreatureInfo) GetSpeed() float64 {
	return ci.Speed
}

func (ci *CreatureInfo) GetPosition() (float64, float64) {
	return ci.X, ci.Y
}

func (ci *CreatureInfo) GetElapsedTime() float64 {
	return ci.ElapsedTime
}

func (ci *CreatureInfo) ClearElapsedTime() {
	ci.ElapsedTime = 0
}

func (ci *CreatureInfo) GetTargetPosition() (float64, float64) {
	return ci.TargetX, ci.TargetY
}
