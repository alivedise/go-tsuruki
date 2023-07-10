package creature

import (
	"image/color"

	"github.com/alivedise/tsuruki/interfaces"
	"github.com/hajimehoshi/ebiten/v2"
)

type Castbar struct {
	Image *ebiten.Image
}

func (cb *Castbar) Draw(c interfaces.Creature) {
	rotation := c.GetRotation()
	if rotation == nil {
		return
	}
	current := rotation.Current()
	if current == nil {
		return
	}
	if current.State() == nil {
		return
	}
	op := &ebiten.DrawImageOptions{}
	progress := current.State().GetProgress()

	cb.Image.Fill(color.Gray{Y: 128})
	c.GetInfo().GetImage().DrawImage(cb.Image, op) // 绘制进度条
	if progress != 0 {
		barWidth := int(float64(cb.Image.Bounds().Dx()) * progress)
		if barWidth <= 0 {
			return
		}
		bar := ebiten.NewImage(barWidth, 15)
		bar.Fill(color.RGBA{4, 59, 92, 1})
		c.GetInfo().GetImage().DrawImage(bar, op)
	}
}
