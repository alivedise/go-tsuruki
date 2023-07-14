package skills

import (
	"github.com/alivedise/tsuruki/interfaces"
	"github.com/alivedise/tsuruki/skill"
	"github.com/alivedise/tsuruki/utils"
)

type HateSword struct {
	SkillState  *skill.SkillState
	Indicator   *skill.SkillIndicator
	SkillConfig *skill.SkillConfig
}

func (s *HateSword) GetConfig() interfaces.SkillConfig {
	return s.SkillConfig
}

func (s *HateSword) EnterCast(g interfaces.World, c interfaces.Creature) {
	x2, y2 := g.GetHighestHateCreature(c).GetInfo().GetPosition()
	x1, y1 := c.GetInfo().GetPosition()
	newX, newY := utils.Interpolate(x1, y1, x2, y2, c.GetRotation().Current().GetIndicator().GetSkillSize())
	c.GetInfo().SetTargetPosition(newX, newY)
	s.GetIndicator().SetRectangleData(x1, y1, newX, newY)
}

func (s *HateSword) Execute(g interfaces.World, c interfaces.Creature) {
	c.GetInfo().SetTargetPosition(-1, -1)
	s.State().Set("executed")
	s.State().Next(g, c)
}

func (s *HateSword) State() interfaces.SkillState {
	return s.SkillState
}

func (s *HateSword) GetIndicator() interfaces.SkillIndicator {
	return s.Indicator
}

func (s *HateSword) Name() string {
	return "HateSword"
}
