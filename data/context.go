package data

import "github.com/sunist-c/genius-invokation-simulator-backend/enum"

type EntityIndex struct {
	ShortID uint16
	LongID  uint64
}

type CreateCharacterArguments struct {
	CharacterName          string
	CharacterAffiliation   enum.Affiliation
	CharacterVision        enum.ElementType
	CharacterWeapon        enum.WeaponType
	CharacterHP            uint
	CharacterMP            uint
	CharacterSkillsShortID []uint16
}

type CreateSkillArguments struct {
	OwnCharacterName    string
	OwnCharacterShortID uint16
	SkillName           string
	SkillType           enum.SkillType
}

type CreateCardArguments struct{}

type CreateRuleArguments struct{}

type InitModArguments struct {
	ID          uint64
	PackagePath string
}

type CharacterContext struct {
	Index          EntityIndex
	Descriptions   map[enum.Language]CharacterDescription
	CreatArguments CreateCharacterArguments
}

type SkillContext struct {
	Index           EntityIndex
	Descriptions    map[enum.Language]SkillDescription
	CreateArguments CreateSkillArguments
}

type EventContext struct {
	Index        EntityIndex
	Descriptions map[enum.Language]EventDescription
}

type CardContext struct {
	Index           EntityIndex
	Descriptions    map[enum.Language]CardDescription
	CreateArguments CreateCardArguments
}

type SummonContext struct {
	Index        EntityIndex
	Descriptions map[enum.Language]SummonDescription
}

type ModifierContext struct {
	Index        EntityIndex
	Descriptions map[enum.Language]ModifierDescription
}

type ModContext struct {
	Descriptions  map[enum.Language]ModDescription
	InitArguments InitModArguments
}

type Context struct {
	Initialized bool
	Mod         ModContext
	Paths       map[string]string
	Characters  map[uint64]CharacterContext
	Skills      map[uint64]SkillContext
	Events      map[uint64]EventContext
	Cards       map[uint64]CardContext
	Summons     map[uint64]SummonContext
	Modifiers   map[uint64]ModifierContext
}

func (ctx *Context) RegisterSkill(index EntityIndex, arguments CreateSkillArguments) {
	if ctx.Skills == nil {
		ctx.Skills = map[uint64]SkillContext{}
	}

	ctx.Skills[index.LongID] = SkillContext{
		Index:           index,
		Descriptions:    map[enum.Language]SkillDescription{},
		CreateArguments: arguments,
	}
}

func (ctx *Context) RegisterCharacter(index EntityIndex, arguments CreateCharacterArguments) {
	if ctx.Characters == nil {
		ctx.Characters = map[uint64]CharacterContext{}
	}

	ctx.Characters[index.LongID] = CharacterContext{
		Index:          index,
		Descriptions:   map[enum.Language]CharacterDescription{},
		CreatArguments: arguments,
	}
}
