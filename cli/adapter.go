package cli

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/model/adapter"
	"github.com/sunist-c/genius-invokation-simulator-cli/data"
)

type NewSkillContextToCreatSkillArgumentAdapter struct{}

func (adapter NewSkillContextToCreatSkillArgumentAdapter) Convert(source NewSkillContext) (success bool, result data.CreateSkillArguments) {
	return true, data.CreateSkillArguments{
		OwnCharacterName: source.CharacterName,
		SkillName:        source.SkillName,
		SkillType:        source.SkillType,
	}
}

func NewNewSkillContextToCreatSkillArgumentAdapter() adapter.Adapter[NewSkillContext, data.CreateSkillArguments] {
	return NewSkillContextToCreatSkillArgumentAdapter{}
}
