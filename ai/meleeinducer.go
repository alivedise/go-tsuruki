package ai

import (
	"github.com/alivedise/tsuruki/interfaces"
	"github.com/alivedise/tsuruki/utils"
)

type MeleeInducerAI struct {
}

func (ai *MeleeInducerAI) Detect(g interfaces.World, c interfaces.Creature) {
	c.GetInfo().SetTargetPosition(g.GetWidth()/2-50, g.GetHeight()/2)
	currentSkill := g.GetEnemyList()[0].GetRotation().Current().Name()
	x1, y1 := c.GetInfo().GetPosition()
	if currentSkill == "MeleeSword" {
		// move to other side arround boss
	} else if currentSkill == "Teleport" {
		// move to area center around 9 o'clock side
	} else if currentSkill == "HugeSector" {
		// go to opponent side of boss
		x2, y2 := g.GetEnemyList()[0].GetInfo().GetPosition()
		x3, y3 := utils.InterpolateByRatio(x1, y1, x2, y2, 2.0)
		c.GetInfo().SetTargetPosition(x3, y3)
	}
}
