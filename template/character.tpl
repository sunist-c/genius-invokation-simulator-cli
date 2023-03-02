package character

import (
	def "github.com/sunist-c/genius-invokation-simulator-backend/mod/definition"
	impl "github.com/sunist-c/genius-invokation-simulator-backend/mod/implement"

	. "{{.Paths.Mod}}"{{if .Conditions.HasSkills}}
	skill "{{.Conditions.SkillsPath}}"{{end}}
)

type {{.Character.CharacterName}}Type def.Character

var (
	{{.Character.CharacterName}}Entity {{.Character.CharacterName}}Type = impl.NewCharacterWithOpts(
		impl.WithCharacterID({{.Character.CharacterID}}),
		impl.WithCharacterName({{.Character.CharacterName}}),
		impl.WithCharacterAffiliation({{.Character.CharacterAffiliation}}),
		impl.WithCharacterVision({{.Character.CharacterVision}}),
		impl.WithCharacterWeapon({{.Character.CharacterWeapon}}),
		impl.WithCharacterHP({{.Character.CharacterHP}}),
		impl.WithCharacterMP({{.Character.CharacterMP}}),
		impl.WithCharacterSkills(
			{{range .Character.CharacterSkills}}skill.Get{{.SkillName}}Entity(),
		{{end}}),
	)
)

func Get{{.Character.CharacterName}}Entity() def.Character {
	return {{.Character.CharacterName}}Entity
}

func init() {
	var modEntity def.Mod = GetMod()
	modEntity.RegisterCharacter(Get{{.Character.CharacterName}}Entity())
}