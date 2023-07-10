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
}

func NewPlayer(path string, x float64, y float64) *Player {
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
			Speed: 2.0,
		},
	}
}

func (c *Player) GetInfo() interfaces.CreatureInfo {
	return c.CreatureInfo
}

func (c *Player) CastNext() {
	c.Rotation.Next()
}

func (c *Player) GetCastbar() interfaces.Castbar {
	return c.Castbar
}

func (c *Player) GetRotation() interfaces.SkillRotation {
	return c.Rotation
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
	c.GetInfo().Move(float64(vx), float64(vy))
}

func (c *Player) Control(input *input.Input) {
	if dir, ok := input.Dir(); ok {
		c.Move(dir)
	}
}

// Boss update; decide what action(skill) to do when being invoked
func (c *Player) Update(g interfaces.World) {
	if c.Rotation.Current() != nil {
		//c.rotation.Current().State().Update(g, c)
	}
}
