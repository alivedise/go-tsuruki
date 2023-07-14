package game

import (
	_ "image/png"
	"math/rand"
	"strconv"

	"github.com/alivedise/tsuruki/ai"
	"github.com/alivedise/tsuruki/boss"
	"github.com/alivedise/tsuruki/input"
	"github.com/alivedise/tsuruki/interfaces"
	"github.com/alivedise/tsuruki/player"
	"github.com/alivedise/tsuruki/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	Player1   *player.Player
	Player2   *player.Player
	Player3   *player.Player
	Player4   *player.Player
	Player5   *player.Player
	Boss      *boss.Boss
	BossDummy *boss.BossDummy
	input     *input.Input
	width     float64
	height    float64
}

func NewGame() *Game {
	boss := boss.NewBoss("images/boss.png", 200, 300)
	g := &Game{
		input:   input.NewInput(),
		Boss:    boss,
		Player1: player.NewPlayer("images/luin.png", 250.0, 240.0, nil),
		Player2: player.NewPlayer("images/yuna.png", 240.0, 150.0, &ai.RangeInducerAI{}),
		Player3: player.NewPlayer("images/kaito.png", 40.0, 250.0, &ai.MeleeInducerAI{}),
		Player4: player.NewPlayer("images/oruta.png", 250.0, 40.0, &ai.RangeDealerAI{}),
		Player5: player.NewPlayer("images/namalie.png", 250.0, 250.0, &ai.RangeDealerAI{}),
		width:   400.0,
		height:  600.0,
	}
	return g
}

func (g *Game) GetWidth() float64 {
	return g.width
}

func (g *Game) GetHeight() float64 {
	return g.height
}

func (g *Game) CheckOverBoundry(x, y float64) (float64, float64) {
	x1 := x
	y1 := y
	if x < 0 {
		x1 = 0
	}
	if x > g.width {
		x1 = g.width
	}
	if y < 0 {
		y1 = 0
	}
	if y > g.height {
		y1 = g.height
	}
	return x1, y1
}

func (g *Game) GetRangeInducer() interfaces.Creature {
	x0, y0 := g.Boss.GetInfo().GetPosition()
	x1, y1 := g.Player2.GetInfo().GetPosition()
	x2, y2 := g.Player3.GetInfo().GetPosition()
	d1 := utils.Distance(x0, y0, x1, y1)
	d2 := utils.Distance(x0, y0, x2, y2)
	if d1 > d2 {
		return g.Player2
	} else {
		return g.Player3
	}
}

func (g *Game) GetMeleeInducer() interfaces.Creature {
	x0, y0 := g.Boss.CreatureInfo.GetPosition()
	x1, y1 := g.Player2.CreatureInfo.GetPosition()
	x2, y2 := g.Player3.CreatureInfo.GetPosition()
	d1 := utils.Distance(x0, y0, x1, y1)
	d2 := utils.Distance(x0, y0, x2, y2)
	if d1 < d2 {
		return g.Player2
	} else {
		return g.Player3
	}
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

func (g *Game) GetEnemyList() []interfaces.Creature {
	return []interfaces.Creature{
		g.Boss,
	}
}

func (g *Game) GetFarthestCreature(source interfaces.Creature) interfaces.Creature {
	players := g.GetCreatureList()

	target := players[0]
	x1, y1 := source.GetInfo().GetPosition()
	for _, char := range players {
		if char != source {
			x2, y2 := target.GetInfo().GetPosition()
			x3, y3 := char.GetInfo().GetPosition()
			if utils.Distance(x1, y1, x2, y2) < utils.Distance(x1, y1, x3, y3) {
				target = char
			}
		}
	}
	return target
}

func (g *Game) GetNearestCreature(source interfaces.Creature) interfaces.Creature {
	players := g.GetCreatureList()

	target := players[0]
	x1, y1 := source.GetInfo().GetPosition()
	for _, char := range players {
		if char != source {
			x2, y2 := target.GetInfo().GetPosition()
			x3, y3 := char.GetInfo().GetPosition()
			if utils.Distance(x1, y1, x2, y2) > utils.Distance(x1, y1, x3, y3) {
				target = char
			}
		}
	}
	return target
}

func (g *Game) GetHighestHateCreature(source interfaces.Creature) interfaces.Creature {
	players := g.GetCreatureList()
	randomInt := rand.Intn(len(players))
	return players[randomInt]
}

func (g *Game) Notify() {
	list := g.GetCreatureList()
	for _, p := range list {
		p.Notify(g)
	}
}

func (g *Game) Update() error {
	g.input.Update()
	g.Player1.Control(g.input)
	g.Player2.Update(g)
	g.Player3.Update(g)
	g.Player4.Update(g)
	g.Player5.Update(g)

	g.Boss.Update(g)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Boss.Draw(screen)
	g.Player1.Draw(screen)
	g.Player2.Draw(screen)
	g.Player3.Draw(screen)
	g.Player4.Draw(screen)
	g.Player5.Draw(screen)
	x1, y1 := g.GetRangeInducer().GetInfo().GetPosition()
	ebitenutil.DebugPrint(screen, strconv.FormatFloat(x1, 'f', -1, 64)+","+strconv.FormatFloat(y1, 'f', -1, 64))
	//x1, y1 := g.Player3.GetInfo().GetPosition()
	//x2, y2 := g.Player3.GetInfo().GetTargetPosition()
	//ebitenutil.DebugPrint(screen, strconv.FormatFloat(x1, 'f', -1, 64)+","+strconv.FormatFloat(y1, 'f', -1, 64)+
	//		"\n"+strconv.FormatFloat(x2, 'f', -1, 64)+","+strconv.FormatFloat(y2, 'f', -1, 64))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	// The unit of outsideWidth/Height is device-independent pixels.
	// By multiplying them by the device scale factor, we can get a hi-DPI screen size.
	s := ebiten.DeviceScaleFactor()
	return int(float64(outsideWidth) * s), int(float64(outsideHeight) * s)
}
