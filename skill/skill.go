package skill

import (
	"time"

	"github.com/alivedise/tsuruki/interfaces"
)

type SkillState struct {
	State   string
	start   time.Time
	end     time.Time
	current time.Time
}

func (s *SkillState) Is(state string) bool {
	return s.State == state
}

func (s *SkillState) GetProgress() float64 {
	return float64(s.current.Sub(s.start).Milliseconds()) / float64(s.end.Sub(s.start).Milliseconds())
}

func (s *SkillState) GetCastTime(sk interfaces.Skill) float64 {
	return sk.TimeInfo()
}

func (s *SkillState) Set(state string) {
	s.State = state
}

func (s *SkillState) Update(g interfaces.World, c interfaces.Creature) {
	if s == nil {
		return
	}
	s.current = time.Now()
	if s.State == "executing" {
		c.GetRotation().Current().Execute(g, c)
	}
	if s.end.Before(s.current) {
		s.Next(g, c)
		s.start = time.Now()
		s.end = s.start.Add(time.Millisecond * time.Duration(s.GetCastTime(c.GetRotation().Current())))
	}
}

func (ss *SkillState) Next(g interfaces.World, c interfaces.Creature) {
	if ss.State == "precast" {
		ss.State = "casting"
		c.GetRotation().Current().EnterCast(g, c)
	} else if ss.State == "casting" {
		ss.State = "executing"
		// Execute should decide when to enter executed
	} else if ss.State == "executed" {
		ss.State = "backwing"
	} else if ss.State == "backwing" {
		c.GetRotation().Current().GetIndicator().Clear()
		ss.State = "precast"
		c.CastNext()
	}
}
