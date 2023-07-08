package main

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Skill interface {
	Precast()
	Cast()
	Backwing()
}

type SkillCaster struct {
	Skill
	Casting       time.Time
	CastStartTime time.Time
	CastEndTime   time.Time
	CastbarImage  *ebiten.Image
	casted        bool
	rushing       bool
	TargetX       float64
	TargetY       float64
	Speed         float64
	ElapsedTime   float64
}

func (skill *SkillCaster) Precast() {

}

type MeleeSword struct{}

func (ms *MeleeSword) Cast() {

}
