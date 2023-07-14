package skills

import (
	"github.com/alivedise/tsuruki/interfaces"
	"github.com/alivedise/tsuruki/skill"
	"github.com/alivedise/tsuruki/utils"
)

type RangeSword struct {
	SkillState  *skill.SkillState
	Indicator   *skill.SkillIndicator
	SkillConfig *skill.SkillConfig
}

func (s *RangeSword) EnterCast(g interfaces.World, c interfaces.Creature) {
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

func (s *RangeSword) Execute(g interfaces.World, c interfaces.Creature) {
	c.GetInfo().SetTargetPosition(-1, -1)
	s.State().Set("executed")
	s.State().Next(g, c)
}

func (s *RangeSword) Name() string {
	return "RangeSword"
}

func (s *RangeSword) State() interfaces.SkillState {
	return s.SkillState
}

func (s *RangeSword) GetIndicator() interfaces.SkillIndicator {
	return s.Indicator
}

func (s *RangeSword) GetConfig() interfaces.SkillConfig {
	return s.SkillConfig
}
