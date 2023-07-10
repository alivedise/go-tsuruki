package skillrotation

import (
	"github.com/alivedise/tsuruki/interfaces"
	"github.com/alivedise/tsuruki/skill"
	"github.com/alivedise/tsuruki/skills"
)

type SkillRotation struct {
	skillList []interfaces.Skill
	current   int64
}

func (sr *SkillRotation) Next() interfaces.Skill {
	size := len(sr.skillList)
	sr.current = sr.current + 1
	if sr.current >= int64(size) {
		sr.current = 0
	}
	return sr.skillList[sr.current]
}

func (sr *SkillRotation) Current() interfaces.Skill {
	if sr == nil {
		return nil
	}
	if len(sr.skillList) == 0 {
		return nil
	}
	return sr.skillList[sr.current]
}

type Phase1SkillRotation struct{}

func NewPhase1SkillRotation() interfaces.SkillRotation {
	melee1 := &skills.MeleeSword{
		SkillState: &skill.SkillState{
			State: "precast",
		},
		Indicator: skill.NewSkillIndicator(
			"rectangle",
			15,
			250,
		),
	}
	melee2 := &skills.MeleeSword{
		SkillState: &skill.SkillState{
			State: "precast",
		},
		Indicator: skill.NewSkillIndicator(
			"rectangle",
			15,
			250,
		),
	}
	range1 := &skills.RangeSword{
		SkillState: &skill.SkillState{
			State: "precast",
		},
		Indicator: skill.NewSkillIndicator(
			"rectangle",
			35,
			250,
		),
	}
	rush := &skills.Rush{
		SkillState: &skill.SkillState{
			State: "precast",
		},
		Indicator: skill.NewSkillIndicator(
			"rectangle",
			35,
			250,
		),
	}
	list := []interfaces.Skill{melee1, melee2, range1, rush}
	return &SkillRotation{
		skillList: list,
		current:   0,
	}
}
