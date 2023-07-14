package player

import (
	"fmt"
	"math/rand"

	"github.com/alivedise/tsuruki/creature"
	"github.com/alivedise/tsuruki/input"
	"github.com/alivedise/tsuruki/interfaces"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Player struct {
	CreatureInfo interfaces.CreatureInfo
	Rotation     interfaces.SkillRotation
	Castbar      interfaces.Castbar
	ai           interfaces.PlayerAI
}

func NewPlayer(path string, x float64, y float64, ai interfaces.PlayerAI) *Player {
	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		fmt.Println(err)
	}
	bar := ebiten.NewImage(img.Bounds().Dx(), 15)
	return &Player{
		Castbar: &creature.Castbar{
			Image: bar,
		},
		CreatureInfo: &creature.CreatureInfo{
			Image: img,
			X:     x,
			Y:     y,
			Speed: 0.1,
		},
		ai: ai,
	}
}

func (c *Player) Notify(g interfaces.World) {
	c.ai.Detect(g, c)
}

func (c *Player) GetInfo() interfaces.CreatureInfo {
	return c.CreatureInfo
}

func (c *Player) CastNext() {
	c.Rotation.Next(c)
}

func (c *Player) SetRotation(s interfaces.SkillRotation) {

}

func (c *Player) CastDone(s interfaces.Skill) {

}

func (c *Player) GetCastbar() interfaces.Castbar {
	return c.Castbar
}

func (c *Player) GetRotation() interfaces.SkillRotation {
	return c.Rotation
}

func (c *Player) Update(g interfaces.World) {
	if c.ai == nil {
		return
	}
	c.ai.Detect(g, c)
	x2, y2 := c.GetInfo().GetTargetPosition()
	if x2 == 0 && y2 == 0 {
		return
	}
	if x2 > 400 || y2 > 600 {
		return
	}

	x1, y1 := c.GetInfo().GetPosition()
	dx := x2 - x1
	dy := y2 - y1
	c.GetInfo().Move(dx*c.CreatureInfo.GetSpeed(), dy*c.CreatureInfo.GetSpeed())
}

func (c *Player) RandomMove() {
	randomInt := rand.Intn(4)
	move := [4][]int{{2, 0}, {0, 2}, {-2, 0}, {0, -2}}
	r := move[randomInt]
	c.GetInfo().Move(float64(r[0]), float64(r[1]))
}

func (c *Player) Draw(screen *ebiten.Image) {
	c.Castbar.Draw(c)
	if c.Rotation != nil && c.Rotation.Current() != nil {
		c.Rotation.Current().GetIndicator().Draw(screen, c)
	}
	c.GetInfo().Draw(screen)
}

func (c *Player) Move(dir input.Dir) {
	vx, vy := dir.Vector()
	c.GetInfo().Move(10*float64(vx), 10*float64(vy))
}

func (c *Player) Control(input *input.Input) {
	if dir, ok := input.Dir(); ok {
		c.Move(dir)
	}
}
