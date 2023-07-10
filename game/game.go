package game

import (
	_ "image/png"

	"github.com/alivedise/tsuruki/boss"
	"github.com/alivedise/tsuruki/input"
	"github.com/alivedise/tsuruki/interfaces"
	"github.com/alivedise/tsuruki/player"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Player1 *player.Player
	Player2 *player.Player
	Player3 *player.Player
	Player4 *player.Player
	Player5 *player.Player
	Boss    *boss.Boss
	input   *input.Input
}

func NewGame() *Game {
	g := &Game{
		input:   input.NewInput(),
		Boss:    boss.NewBoss("images/boss.png", 100, 100),
		Player1: player.NewPlayer("images/luin.png", 150.0, 240.0),
		Player2: player.NewPlayer("images/yuna.png", 240.0, 150.0),
		Player3: player.NewPlayer("images/kaito.png", 40.0, 250.0),
		Player4: player.NewPlayer("images/oruta.png", 250.0, 40.0),
		Player5: player.NewPlayer("images/namalie.png", 250.0, 250.0),
	}
	return g
}

func (g *Game) GetCreatureList() []interfaces.Creature {
	return []interfaces.Creature{
		g.Player1,
		g.Player2,
		g.Player3,
		g.Player4,
		g.Player5,
	}
}

func (g *Game) Update() error {
	g.input.Update()
	g.Player1.Control(g.input)
	g.Player2.RandomMove()
	g.Player3.RandomMove()
	g.Player4.RandomMove()
	g.Player5.RandomMove()

	g.Boss.Update(g)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	//g.Boss.DrawSkill(screen)
	g.Boss.Draw(screen)
	g.Player1.Draw(screen)
	g.Player2.Draw(screen)
	g.Player3.Draw(screen)
	g.Player4.Draw(screen)
	g.Player5.Draw(screen)
	//ebitenutil.DebugPrint(screen, "Hello, World!\n"+g.Boss.Skill.State+"\n"+strconv.FormatFloat(g.Boss.X, 'f', -1, 64)+","+strconv.FormatFloat(g.Boss.Y, 'f', -1, 64))
	//ebitenutil.DebugPrint(screen, g.Boss.Skill.State)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	// The unit of outsideWidth/Height is device-independent pixels.
	// By multiplying them by the device scale factor, we can get a hi-DPI screen size.
	s := ebiten.DeviceScaleFactor()
	return int(float64(outsideWidth) * s), int(float64(outsideHeight) * s)
}
