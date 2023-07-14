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

type SkillConfig struct {
	Precast  float64
	Cast     float64
	Backwing float64
}

func (sc *SkillConfig) GetCastTime(state string) float64 {
	if state == "precast" {
		return sc.Precast * 1000
	} else if state == "casting" {
		return sc.Cast * 1000
	} else if state == "backwing" {
		return sc.Backwing * 1000
	}
	return 0
}

func (s *SkillState) Is(state string) bool {
	return s.State == state
}

func (s *SkillState) Get() string {
	return s.State
}

func (s *SkillState) GetProgress() float64 {
	return float64(s.current.Sub(s.start).Milliseconds()) / float64(s.end.Sub(s.start).Milliseconds())
}

func (s *SkillState) GetCastTime(sk interfaces.Skill) float64 {
	return sk.GetConfig().GetCastTime(s.State)
}

func (s *SkillState) Set(state string) {
	s.State = state
}

func (s *SkillState) Update(g interfaces.World, c interfaces.Creature) {
	if s == nil {
		return
	}
	// Whenever Update is called, update the time in "current" to Now
	s.current = time.Now()
	if s.State == "executing" {
		c.GetRotation().Current().Execute(g, c)
		return
	}
	if s.end.Before(s.current) {
		s.Next(g, c)
	}
}

func (ss *SkillState) Next(g interfaces.World, c interfaces.Creature) {
	if ss.State == "init" || ss.State == "" {
		ss.State = "precast"
		ss.start = time.Now()
		ss.end = ss.start.Add(time.Millisecond * time.Duration(ss.GetCastTime(c.GetRotation().Current())))
	} else if ss.State == "precast" {
		ss.State = "casting"
		ss.start = time.Now()
		ss.end = ss.start.Add(time.Millisecond * time.Duration(ss.GetCastTime(c.GetRotation().Current())))
		c.GetRotation().Current().EnterCast(g, c)
	} else if ss.State == "casting" {
		ss.State = "executing"
		// Execute should decide when to enter executed
	} else if ss.State == "executing" || ss.State == "executed" {
		ss.State = "backwing"
		ss.start = time.Now()
		ss.end = ss.start.Add(time.Millisecond * time.Duration(ss.GetCastTime(c.GetRotation().Current())))
	} else if ss.State == "backwing" {
		c.GetRotation().Current().GetIndicator().Clear()
		ss.State = "init"
		c.CastNext()
	}
}
