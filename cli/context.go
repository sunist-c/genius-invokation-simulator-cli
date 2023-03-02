package cli

import "github.com/sunist-c/genius-invokation-simulator-backend/enum"

type context = map[string]interface{}

type NewCharacterContext struct {
	CharacterName string
}

type NewSkillContext struct {
	SkillName string
	SkillID   uint16
	SkillType enum.SkillType
}

type NewCardContext struct{}

type NewRuleContext struct{}

type InitModContext struct{}
