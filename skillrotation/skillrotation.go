package skillrotation

import (
	"github.com/alivedise/tsuruki/interfaces"
	"github.com/alivedise/tsuruki/skill"
	"github.com/alivedise/tsuruki/skills"
)

type SkillRotation struct {
	skillList []interfaces.Skill
	current   int64
	loop      bool
}

func (sr *SkillRotation) Next(c interfaces.Creature) interfaces.Skill {
	size := len(sr.skillList)
	c.CastDone(sr.skillList[sr.current])
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
	if sr.current == int64(len(sr.skillList)) && !sr.loop {
		return nil
	}
	return sr.skillList[sr.current]
}

func (sr *SkillRotation) ShouldLoop() bool {
	return sr.loop
}

func NewPhase2SkillRotation() interfaces.SkillRotation {
	melee := &skills.HateSword{
		SkillState: &skill.SkillState{},
		SkillConfig: &skill.SkillConfig{
			Precast:  1.0,
			Cast:     1.5,
			Backwing: 1.0,
		},
		Indicator: skill.NewSkillIndicator(
			"rectangle",
			15,
			250,
		),
	}
	rush := &skills.Rush{
		SkillState: &skill.SkillState{},
		SkillConfig: &skill.SkillConfig{
			Precast:  1.5,
			Cast:     1.5,
			Backwing: 1.0,
		},
		Indicator: skill.NewSkillIndicator(
			"rectangle",
			35,
			180,
		),
	}
	teleport := &skills.Teleport{
		SkillState: &skill.SkillState{},
		SkillConfig: &skill.SkillConfig{
			Precast:  1.0,
			Cast:     1.0,
			Backwing: 0.1,
		},
		Indicator: skill.NewSkillIndicator(
			"rectangle",
			35,
			250,
		),
	}
	ranges := &skills.RangeSword{
		SkillState: &skill.SkillState{},
		SkillConfig: &skill.SkillConfig{
			Precast:  0.0,
			Cast:     1.0,
			Backwing: 0.0,
		},
		Indicator: skill.NewSkillIndicator(
			"rectangle",
			35,
			250,
		),
	}
	hugesector := &skills.HugeSector{
		SkillState: &skill.SkillState{},
		SkillConfig: &skill.SkillConfig{
			Precast:  0.5,
			Cast:     2.5,
			Backwing: 5.0,
		},
		Indicator: skill.NewSkillIndicator(
			"sector",
			35,
			250,
		),
	}
	list := []interfaces.Skill{melee, rush, teleport, ranges, rush, teleport, hugesector}
	return &SkillRotation{
		skillList: list,
		current:   0,
		loop:      true,
	}
}

func NewPhase1SkillRotation() interfaces.SkillRotation {
	melee1 := &skills.MeleeSword{
		SkillState: &skill.SkillState{},
		SkillConfig: &skill.SkillConfig{
			Precast:  1.0,
			Cast:     1.5,
			Backwing: 1.0,
		},
		Indicator: skill.NewSkillIndicator(
			"rectangle",
			15,
			250,
		),
	}
	melee2 := &skills.MeleeSword{
		SkillState: &skill.SkillState{},
		SkillConfig: &skill.SkillConfig{
			Precast:  1.0,
			Cast:     1.5,
			Backwing: 1.0,
		},
		Indicator: skill.NewSkillIndicator(
			"rectangle",
			15,
			250,
		),
	}
	range1 := &skills.RangeSword{
		SkillState: &skill.SkillState{},
		SkillConfig: &skill.SkillConfig{
			Precast:  1.0,
			Cast:     1.5,
			Backwing: 1.0,
		},
		Indicator: skill.NewSkillIndicator(
			"rectangle",
			35,
			250,
		),
	}
	rush := &skills.Rush{
		SkillState: &skill.SkillState{},
		SkillConfig: &skill.SkillConfig{
			Precast:  1.0,
			Cast:     1.5,
			Backwing: 1.0,
		},
		Indicator: skill.NewSkillIndicator(
			"rectangle",
			35,
			180,
		),
	}
	list := []interfaces.Skill{melee1, melee2, range1, rush}
	return &SkillRotation{
		skillList: list,
		current:   0,
		loop:      true,
	}
}

func NewDummySkillRotationSector() interfaces.SkillRotation {
	hugesector := &skills.HugeSector{
		SkillState: &skill.SkillState{},
		SkillConfig: &skill.SkillConfig{
			Precast:  3.5,
			Cast:     1.5,
			Backwing: 0.1,
		},
		Indicator: skill.NewSkillIndicator(
			"sector",
			35,
			250,
		),
	}
	list := []interfaces.Skill{hugesector}
	return &SkillRotation{
		skillList: list,
		current:   0,
		loop:      false,
	}
}

func NewDummySkillRotationMelee() interfaces.SkillRotation {
	melee := &skills.MeleeSword{
		SkillState: &skill.SkillState{},
		SkillConfig: &skill.SkillConfig{
			Precast:  3.0,
			Cast:     1.5,
			Backwing: 0,
		},
		Indicator: skill.NewSkillIndicator(
			"rectangle",
			15,
			250,
		),
	}
	list := []interfaces.Skill{melee}
	return &SkillRotation{
		skillList: list,
		current:   0,
		loop:      false,
	}
}
