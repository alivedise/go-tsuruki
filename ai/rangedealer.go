package ai

import (
	"math"
	"math/rand"

	"github.com/alivedise/tsuruki/interfaces"
	"github.com/alivedise/tsuruki/utils"
)

type RangeDealerAI struct {
	angle float64
}

func (ai *RangeDealerAI) findThirdPoint(x1, y1, x2, y2, l float64) (float64, float64) {
	// 计算从 P1 到 P2 的向量
	dx := x2 - x1
	dy := y2 - y1

	// 计算从 P1 到 P2 的长度
	distance := utils.Distance(x1, y1, x2, y2)

	// 计算新坐标系的 x 轴向量
	axisX := dx / distance
	axisY := dy / distance

	if ai.angle == 0.0 {
		// 随机选择角度范围
		minAngle := 30.0 // 30 度
		maxAngle := 80.0 // 80 度
		ai.angle = rand.Float64()*(maxAngle-minAngle) + minAngle
	}

	// 计算角度的弧度值
	radian1 := ai.angle * math.Pi / 180.0
	// 计算角度的弧度值
	radian2 := -ai.angle * math.Pi / 180.0

	// 计算从 P1 到 P3 的长度
	l3 := l

	// 计算 P3 在新坐标系中的位置
	x3 := x1 + l3*math.Cos(radian1)*axisX - l3*math.Sin(radian1)*axisY
	y3 := y1 + l3*math.Sin(radian1)*axisX + l3*math.Cos(radian1)*axisY

	x4 := x1 + l3*math.Cos(radian2)*axisX - l3*math.Sin(radian2)*axisY
	y4 := y1 + l3*math.Sin(radian2)*axisX + l3*math.Cos(radian2)*axisY

	if y4 > y3 {
		return x4, y4
	}
	if math.Abs(y3-y4) <= 1 && x4 > x3 {
		return x4, y4
	}
	return x3, y3
}

func (ai *RangeDealerAI) Detect(g interfaces.World, c interfaces.Creature) {
	//x0, y0 := c.GetInfo().GetPosition()
	x1, y1 := g.GetRangeInducer().GetInfo().GetPosition()
	x2, y2 := g.GetEnemyList()[0].GetInfo().GetPosition()
	x3, y3 := ai.findThirdPoint(x2, y2, x1, y1, 80.0)

	c.GetInfo().SetTargetPosition(x3, y3)
	currentSkill := g.GetEnemyList()[0].GetRotation().Current().Name()
	if currentSkill == "MeleeSword" {
		// move to other side arround boss
	} else if currentSkill == "Teleport" {
		// move to area center around 9 o'clock side
	} else if currentSkill == "HugeSector" {
		// go to opponent side of boss
		x3, y3 := ai.findThirdPoint(x2, y2, x1, y1, 130.0)
		c.GetInfo().SetTargetPosition(x3, y3)
	} else {
	}
}
