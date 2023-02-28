package data

import "github.com/sunist-c/genius-invokation-simulator-backend/enum"

type EntityIndex struct {
	ShortID uint16
	LongID  uint64
}

type CharacterContext struct {
	Index        EntityIndex
	Descriptions map[enum.Language]CharacterDescription
}

type SkillContext struct {
	Index        EntityIndex
	Descriptions map[enum.Language]SkillDescription
}

type EventContext struct {
	Index        EntityIndex
	Descriptions map[enum.Language]EventDescription
}

type CardContext struct {
	Index        EntityIndex
	Descriptions map[enum.Language]CardDescription
}

type SummonContext struct {
	Index        EntityIndex
	Descriptions map[enum.Language]SummonDescription
}

type ModifierContext struct {
	Index        EntityIndex
	Descriptions map[enum.Language]ModifierDescription
}

type Context struct {
	Mod        ModDescription
	Characters map[uint64]CharacterContext
	Skills     map[uint64]SkillContext
	Events     map[uint64]EventContext
	Cards      map[uint64]CardContext
	Summons    map[uint64]SummonContext
	Modifiers  map[uint64]ModifierContext
}
