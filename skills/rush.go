package skills

import (
	"github.com/alivedise/tsuruki/interfaces"
	"github.com/alivedise/tsuruki/skill"
	"github.com/alivedise/tsuruki/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

type Rush struct {
	SkillState *skill.SkillState
	Indicator  *skill.SkillIndicator
}

func (s *Rush) State() interfaces.SkillState {
	return s.SkillState
}

func (s *Rush) EnterCast(g interfaces.World, c interfaces.Creature) {
	players := g.GetCreatureList()

	target := players[0]
	x1, y1 := c.GetInfo().GetPosition()
	for _, char := range players {
		if char != c {
			x2, y2 := target.GetInfo().GetPosition()
			x3, y3 := char.GetInfo().GetPosition()
			if utils.Distance(x1, y1, x2, y2) < utils.Distance(x1, y1, x3, y3) {
				target = char
			}
		}
	}
	x2, y2 := target.GetInfo().GetPosition()
	newX, newY := utils.Interpolate(x1, y1, x2, y2, c.GetRotation().Current().GetIndicator().GetSkillSize())
	c.GetInfo().SetTargetPosition(newX, newY)
	s.GetIndicator().SetRectangleData(x1, y1, newX, newY)
}

func (s *Rush) Execute(g interfaces.World, c interfaces.Creature) {
	s.Indicator.Clear()
	dt := 1.0 / ebiten.ActualTPS()
	c.GetInfo().IncreaseElapsedTime(dt)

	if c.GetInfo().GetElapsedTime() >= 3.0 {
		c.GetInfo().ClearElapsedTime()
		s.State().Set("executed")
		s.SkillState.Next(g, c)
	} else {
		x1, y1 := c.GetInfo().GetPosition()
		x2, y2 := c.GetInfo().GetTargetPosition()
		c.GetInfo().Move((x2-x1)*c.GetInfo().GetSpeed()*dt, (y2-y1)*c.GetInfo().GetSpeed()*dt)
	}
}

func (s *Rush) Name() string {
	return "Rush"
}

func (s *Rush) GetIndicator() interfaces.SkillIndicator {
	return s.Indicator
}

func (s *Rush) TimeInfo() float64 {
	if s.State().Is("precast") {
		return 1000.0
	} else if s.State().Is("casting") {
		return 1500.0
	} else {
		return 500.0
	}
}
