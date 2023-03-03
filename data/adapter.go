package data

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
	"github.com/sunist-c/genius-invokation-simulator-backend/mod/implement"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/adapter"
	"github.com/sunist-c/genius-invokation-simulator-cli/parser"
)

var (
	skillTypeToStringAdapter       = parser.NewSkillTypeToStringAdapter()
	stringToSkillTypeAdapter       = parser.NewStringToSkillTypeAdapter()
	languageToShortStringAdapter   = parser.NewLanguageToShortStringAdapter()
	ShortStringToLanguageAdapter   = parser.NewShortStringToLanguageAdapter()
	descriptionTypeToStringAdapter = parser.NewDescriptionTypeToStringAdapter()
	stringToDescriptionTypeAdapter = parser.NewStringToDescriptionTypeAdapter()
)

func generateLongID(shortID uint16) uint64 {
	return implement.NewEntityWithOpts(implement.WithEntityID(shortID)).TypeID()
}

type SkillContextToExportedSkillContext struct{}

func (adapter SkillContextToExportedSkillContext) Convert(source SkillContext) (success bool, result ExportedSkillContext) {
	exported := ExportedSkillContext{
		ShortID:             source.Index.ShortID,
		OwnCharacter:        source.CreateArguments.OwnCharacterName,
		OwnCharacterShortID: source.CreateArguments.OwnCharacterShortID,
		SkillName:           source.CreateArguments.SkillName,
		Descriptions:        map[string]ExportedSkillDescription{},
	}

	_, exported.SkillType = skillTypeToStringAdapter.Convert(source.CreateArguments.SkillType)
	for language, description := range source.Descriptions {
		_, exportedLanguage := languageToShortStringAdapter.Convert(language)
		_, descriptionType := descriptionTypeToStringAdapter.Convert(description.Type)
		exported.Descriptions[exportedLanguage] = ExportedSkillDescription{
			DescriptionType:  descriptionType,
			SkillName:        description.Name,
			SkillDescription: description.Description,
		}
	}

	return true, exported
}

func NewSkillContextToExportedSkillContextAdapter() adapter.Adapter[SkillContext, ExportedSkillContext] {
	return SkillContextToExportedSkillContext{}
}

type ExportedSkillContextToSkillContextAdapter struct{}

func (adapter ExportedSkillContextToSkillContextAdapter) Convert(source ExportedSkillContext) (success bool, result SkillContext) {
	unexported := SkillContext{
		Index: EntityIndex{
			ShortID: source.ShortID,
			LongID:  generateLongID(source.ShortID),
		},
		Descriptions: map[enum.Language]SkillDescription{},
		CreateArguments: CreateSkillArguments{
			OwnCharacterName:    source.OwnCharacter,
			OwnCharacterShortID: source.OwnCharacterShortID,
			SkillName:           source.SkillName,
		},
	}

	_, unexported.CreateArguments.SkillType = stringToSkillTypeAdapter.Convert(source.SkillType)
	for language, description := range source.Descriptions {
		_, enumLanguage := ShortStringToLanguageAdapter.Convert(language)
		_, enumDescriptionType := stringToDescriptionTypeAdapter.Convert(description.DescriptionType)
		unexported.Descriptions[enumLanguage] = SkillDescription{
			Type:        enumDescriptionType,
			Name:        description.SkillName,
			Description: description.SkillDescription,
		}
	}

	return true, unexported
}

func NewExportedSkillContextToSkillContextAdapter() adapter.Adapter[ExportedSkillContext, SkillContext] {
	return ExportedSkillContextToSkillContextAdapter{}
}
