package boss

import (
	"fmt"
	"math"

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
	dummyList    []interfaces.Creature
}

func (b *Boss) Notify(g interfaces.World) {}

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
	boss.Rotation = skillrotation.NewPhase2SkillRotation()
	return boss
}

func (c *Boss) CastNext() {
	c.Rotation.Next(c)
	// decide to call dummy or not
	current := c.GetRotation().Current().Name()

	if current == "Rush" {
		x, y := c.GetInfo().GetPosition()
		dummy := NewBossDummy("images/boss.png", x, y)
		dummy.SetRotation(skillrotation.NewDummySkillRotationMelee())
		c.dummyList = append(c.dummyList, dummy)
	} else if current == "HugeSector" {
		x, y := c.GetInfo().GetPosition()
		dummy := NewBossDummy("images/boss.png", x, y)
		dummy.SetRotation(skillrotation.NewDummySkillRotationSector())
		dummy.GetInfo().SetFaceAngle(c.GetInfo().GetFaceAngle() + math.Pi)
		c.dummyList = append(c.dummyList, dummy)
	}

	toRemove := []int{}

	for i, d := range c.dummyList {
		if d.GetInfo().IsHidden() {
			toRemove = append(toRemove, i)
		}
	}
	for _, i := range toRemove {
		c.dummyList = append(c.dummyList[:i], c.dummyList[i+1:]...)
	}
}

func (c *Boss) CastDone(s interfaces.Skill) {

}

func (c *Boss) SetRotation(r interfaces.SkillRotation) {
	c.Rotation = r
}

func (c *Boss) GetInfo() interfaces.CreatureInfo {
	return c.CreatureInfo
}

func (c *Boss) RandomMove() {
}

func (c *Boss) Draw(screen *ebiten.Image) {
	for _, dummy := range c.dummyList {
		dummy.Draw(screen)
	}
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
	for _, dummy := range c.dummyList {
		dummy.Update(g)
	}
}
