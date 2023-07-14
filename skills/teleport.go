package skills

import (
	"github.com/alivedise/tsuruki/interfaces"
	"github.com/alivedise/tsuruki/skill"
	"github.com/alivedise/tsuruki/utils"
)

type Teleport struct {
	SkillState  *skill.SkillState
	Indicator   *skill.SkillIndicator
	SkillConfig *skill.SkillConfig
}

func (s *Teleport) EnterCast(g interfaces.World, c interfaces.Creature) {
	x1, y1 := c.GetInfo().GetPosition()
	target := g.GetFarthestCreature(c)
	x2, y2 := target.GetInfo().GetPosition()
	d := utils.Distance(x1, y1, x2, y2)
	newX, newY := utils.Interpolate(x1, y1, x2, y2, d-50)
	c.GetInfo().SetTargetPosition(newX, newY)
}

func (s *Teleport) Execute(g interfaces.World, c interfaces.Creature) {
	x1, y1 := c.GetInfo().GetPosition()
	x, y := c.GetInfo().GetTargetPosition()
	c.GetInfo().MoveTo(x, y)
	c.GetInfo().SetFaceAngle(utils.CalculateRadian(x, y, x1, y1))
	c.GetInfo().SetTargetPosition(-1, -1)
	s.State().Set("executed")
	s.State().Next(g, c)
}

func (s *Teleport) Name() string {
	return "Teleport"
}

func (s *Teleport) State() interfaces.SkillState {
	return s.SkillState
}

func (s *Teleport) GetIndicator() interfaces.SkillIndicator {
	return s.Indicator
}

func (s *Teleport) GetConfig() interfaces.SkillConfig {
	return s.SkillConfig
}
