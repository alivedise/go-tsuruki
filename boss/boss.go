package boss

import (
	"fmt"

	"github.com/alivedise/tsuruki/creature"
	"github.com/alivedise/tsuruki/input"
	"github.com/alivedise/tsuruki/interfaces"
	"github.com/alivedise/tsuruki/skillrotation"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Boss struct {
	CreatureInfo interfaces.CreatureInfo
	Castbar      interfaces.Castbar
	Rotation     interfaces.SkillRotation
}

func NewBoss(path string, x, y float64) *Boss {
	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		fmt.Println(err)
	}
	bar := ebiten.NewImage(img.Bounds().Dx(), 15)
	boss := &Boss{
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
	boss.Rotation = skillrotation.NewPhase1SkillRotation()
	return boss
}

func (c *Boss) CastNext() {
	c.Rotation.Next()
}

func (c *Boss) GetInfo() interfaces.CreatureInfo {
	return c.CreatureInfo
}

func (c *Boss) RandomMove() {
}

func (c *Boss) Draw(screen *ebiten.Image) {
	c.Castbar.Draw(c)
	if c.Rotation.Current() != nil {
		c.Rotation.Current().GetIndicator().Draw(screen, c)
	}
	c.GetInfo().Draw(screen)
}

func (c *Boss) GetRotation() interfaces.SkillRotation {
	return c.Rotation
}

func (c *Boss) Move(dir input.Dir) {
	vx, vy := dir.Vector()
	c.GetInfo().Move(float64(vx), float64(vy))
}

func (c *Boss) Control(input *input.Input) {
	if dir, ok := input.Dir(); ok {
		c.Move(dir)
	}
}

func (c *Boss) GetCastbar() interfaces.Castbar {
	return c.Castbar
}

// Boss update; decide what action(skill) to do when being invoked
func (c *Boss) Update(g interfaces.World) {
	if c.Rotation != nil && c.Rotation.Current() != nil {
		c.Rotation.Current().State().Update(g, c)
	}
}
