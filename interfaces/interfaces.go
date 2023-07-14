package interfaces

import "github.com/hajimehoshi/ebiten/v2"

type World interface {
	GetCreatureList() []Creature
	GetEnemyList() []Creature
	GetNearestCreature(Creature) Creature
	GetFarthestCreature(Creature) Creature
	GetHighestHateCreature(Creature) Creature
	GetRangeInducer() Creature
	GetMeleeInducer() Creature
	Notify()
	GetWidth() float64
	GetHeight() float64
}

type SkillRotation interface {
	Next(Creature) Skill
	Current() Skill
	ShouldLoop() bool
}

type SkillConfig interface {
	GetCastTime(key string) float64
}

type SkillState interface {
	GetProgress() float64
	Is(string) bool
	Next(World, Creature)
	Set(string)
	Get() string
	Update(World, Creature)
}

type SkillIndicator interface {
	Draw(*ebiten.Image, Creature)
	SetRectangleData(float64, float64, float64, float64)
	SetSectorData(float64, float64, float64, float64, float64)
	GetSkillSize() float64
	Clear()
}

type Castbar interface {
	Draw(Creature)
}

type Creature interface {
	GetRotation() SkillRotation
	SetRotation(SkillRotation)
	GetInfo() CreatureInfo
	GetCastbar() Castbar
	CastNext()
	CastDone(Skill)
	Draw(*ebiten.Image)
	Update(World)
	Notify(World)
}

type CreatureInfo interface {
	GetSpeed() float64
	GetPosition() (float64, float64)
	SetTargetPosition(float64, float64)
	GetElapsedTime() float64
	GetTargetPosition() (float64, float64)
	GetImage() *ebiten.Image
	Move(float64, float64)
	MoveTo(float64, float64)
	Update()
	Draw(*ebiten.Image)
	IncreaseElapsedTime(float64)
	ClearElapsedTime()
	SetFaceAngle(float64)
	GetFaceAngle() float64
	Show()
	Hide()
	IsHidden() bool
}

type Skill interface {
	EnterCast(g World, c Creature)
	Execute(g World, c Creature)
	State() SkillState
	GetConfig() SkillConfig
	Name() string
	GetIndicator() SkillIndicator
}

type PlayerAI interface {
	Detect(World, Creature)
}
