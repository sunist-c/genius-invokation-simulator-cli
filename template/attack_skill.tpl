package {{.Skill.CharacterName}}

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
	def "github.com/sunist-c/genius-invokation-simulator-backend/mod/definition"
    impl "github.com/sunist-c/genius-invokation-simulator-backend/mod/implement"

	. "{{.Paths.Mod}}"
)

type {{.Skill.SkillName}}Type def.AttackSkill

var (
    {{.Skill.SkillName}}Entity {{.Skill.SkillName}}Type = impl.NewAttackSkillWithOpts(
        impl.WithAttackSkillID({{.Skill.SkillID}}),
		impl.WithAttackSkillType({{.Skill.SkillType}}),
		impl.WithAttackSkillCost(map[enum.ElementType]uint{
            // todo: modify me
        }),
		impl.WithAttackSkillActiveDamageHandler(func(ctx def.Context) (elementType enum.ElementType, damageAmount uint) {
            // todo: implement me
			panic("not implement yet")
        }),
		impl.WithAttackSkillBackgroundDamageHandler(func(ctx def.Context) (damageAmount uint) {
            // todo: implement me
            panic("not implement yet")
        }),
    )
)

func Get{{.Skill.SkillName}}Entity() def.AttackSkill {
	return {{.Skill.SkillName}}Entity
}

func init() {
    var modEntity def.Mod = GetMod()
	modEntity.RegisterSkill(Get{{.Skill.SkillName}}Entity())
}