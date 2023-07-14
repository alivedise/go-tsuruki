package ai

import (
	"github.com/alivedise/tsuruki/interfaces"
)

type RangeInducerAI struct {
}

func (ai *RangeInducerAI) Detect(g interfaces.World, c interfaces.Creature) {
	c.GetInfo().SetTargetPosition(400-50, 300)
	currentSkill := g.GetEnemyList()[0].GetRotation().Current().Name()
	if currentSkill == "MeleeSword" {
		// move to other side arround boss
	} else if currentSkill == "Teleport" {
	} else if currentSkill == "HugeSector" {
		// go to opponent side of boss
	} else {
		c.GetInfo().SetTargetPosition(400-50, 300)
	}
}
