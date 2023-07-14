package boss

import (
	"fmt"
	"image/color"

	"github.com/alivedise/tsuruki/creature"
	"github.com/alivedise/tsuruki/input"
	"github.com/alivedise/tsuruki/interfaces"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type BossDummy struct {
	CreatureInfo interfaces.CreatureInfo
	Castbar      interfaces.Castbar
	Rotation     interfaces.SkillRotation
	name         string
}

func NewBossDummy(path string, x, y float64) *BossDummy {
	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		fmt.Println(err)
	}
	img.Fill(color.RGBA{R: 255, G: 192, B: 203, A: 128})
	bar := ebiten.NewImage(img.Bounds().Dx(), 15)
	boss := &BossDummy{
		name: "dummy",
		CreatureInfo: &creature.CreatureInfo{
			X:     x,
			Y:     y,
			Image: img,
			Speed: 2.0,
		},
		Castbar: &creature.Castbar{
			Image: bar,
		},
	}
	return boss
}

func (c *BossDummy) Notify(g interfaces.World) {}

func (c *BossDummy) CastNext() {
	c.Rotation.Next(c)
}

func (c *BossDummy) CastDone(s interfaces.Skill) {
	c.GetInfo().Hide()
}

func (c *BossDummy) SetRotation(sr interfaces.SkillRotation) {
	c.Rotation = sr
}

func (c *BossDummy) GetInfo() interfaces.CreatureInfo {
	return c.CreatureInfo
}

func (c *BossDummy) RandomMove() {
}

func (c *BossDummy) Draw(screen *ebiten.Image) {
	if c.GetInfo().IsHidden() {
		return
	}
	c.Castbar.Draw(c)
	if c.Rotation != nil && c.Rotation.Current() != nil {
		c.Rotation.Current().GetIndicator().Draw(screen, c)
	}
	c.GetInfo().Draw(screen)
}

func (c *BossDummy) GetRotation() interfaces.SkillRotation {
	return c.Rotation
}

func (c *BossDummy) Move(dir input.Dir) {
	vx, vy := dir.Vector()
	c.GetInfo().Move(float64(vx), float64(vy))
}

func (c *BossDummy) Control(input *input.Input) {
	if dir, ok := input.Dir(); ok {
		c.Move(dir)
	}
}

func (c *BossDummy) GetCastbar() interfaces.Castbar {
	return c.Castbar
}

// Boss update; decide what action(skill) to do when being invoked
func (c *BossDummy) Update(g interfaces.World) {
	if c.Rotation != nil && c.Rotation.Current() != nil {
		c.Rotation.Current().State().Update(g, c)
	}
}
