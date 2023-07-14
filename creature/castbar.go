package creature

import (
	"image/color"

	"github.com/alivedise/tsuruki/interfaces"
	"github.com/hajimehoshi/bitmapfont/v2"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
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
		bar := ebiten.NewImage(barWidth, 25)
		clr := color.RGBA{4, 59, 92, 1}
		if current.State().Is("precast") {
			clr = color.RGBA{R: 255, G: 99, B: 71}
		} else if current.State().Is("backwing") {
			clr = color.RGBA{R: 238, G: 130, B: 238}
		}
		bar.Fill(clr)
		c.GetInfo().GetImage().DrawImage(bar, op)
		text.Draw(c.GetInfo().GetImage(), current.State().Get(), bitmapfont.Face, 8, 12, color.White)
	}
}
