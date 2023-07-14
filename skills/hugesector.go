package skills

import (
	"github.com/alivedise/tsuruki/interfaces"
	"github.com/alivedise/tsuruki/skill"
	"github.com/alivedise/tsuruki/utils"
)

type HugeSector struct {
	SkillState  *skill.SkillState
	Indicator   *skill.SkillIndicator
	SkillConfig *skill.SkillConfig
}

func (s *HugeSector) EnterCast(g interfaces.World, c interfaces.Creature) {
	x1, y1 := c.GetInfo().GetPosition()
	target := g.GetFarthestCreature(c)
	x2, y2 := target.GetInfo().GetPosition()
	d := utils.Distance(x1, y1, x2, y2)
	newX, newY := utils.Interpolate(x1, y1, x2, y2, d-50)
	c.GetInfo().SetTargetPosition(newX, newY)
	c.GetRotation().Current().GetIndicator().SetSectorData(x1, y1, 100, 0, 270)
}

func (s *HugeSector) Execute(g interfaces.World, c interfaces.Creature) {
	c.GetInfo().SetTargetPosition(-1, -1)
	s.State().Set("executed")
	s.State().Next(g, c)
}

func (s *HugeSector) Name() string {
	return "HugeSector"
}

func (s *HugeSector) State() interfaces.SkillState {
	return s.SkillState
}

func (s *HugeSector) GetIndicator() interfaces.SkillIndicator {
	return s.Indicator
}

func (s *HugeSector) GetConfig() interfaces.SkillConfig {
	return s.SkillConfig
}
