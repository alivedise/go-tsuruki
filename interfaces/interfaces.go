package interfaces

import "github.com/hajimehoshi/ebiten/v2"

type World interface {
	GetCreatureList() []Creature
}

type SkillRotation interface {
	Next() Skill
	Current() Skill
}

type SkillState interface {
	GetProgress() float64
	Is(string) bool
	Next(World, Creature)
	Set(string)
	Update(World, Creature)
}

type SkillIndicator interface {
	Draw(*ebiten.Image, Creature)
	SetRectangleData(float64, float64, float64, float64)
	GetSkillSize() float64
	Clear()
}

type Castbar interface {
	Draw(Creature)
}

type Creature interface {
	GetRotation() SkillRotation
	GetInfo() CreatureInfo
	GetCastbar() Castbar
	CastNext()
}

type CreatureInfo interface {
	GetSpeed() float64
	GetPosition() (float64, float64)
	SetTargetPosition(float64, float64)
	GetElapsedTime() float64
	GetTargetPosition() (float64, float64)
	GetImage() *ebiten.Image
	Move(float64, float64)
	Draw(*ebiten.Image)
	IncreaseElapsedTime(float64)
	ClearElapsedTime()
}

type Skill interface {
	EnterCast(g World, c Creature)
	Execute(g World, c Creature)
	State() SkillState
	Name() string
	GetIndicator() SkillIndicator
	TimeInfo() float64
}
