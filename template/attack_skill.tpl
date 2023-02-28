package {{.context.Skills.CharacterName}}

import (
	. "github.com/sunist-c/genius-invokation-simulator-backend/enum"
	def "github.com/sunist-c/genius-invokation-simulator-backend/mod/definition"
    impl "github.com/sunist-c/genius-invokation-simulator-backend/mod/implement"

	mod "{{.context.Paths.Mod}}"
)

type {{.context.Skills.SkillName}}Type def.AttackSkill

var (
    {{.context.Skills.SkillName}} {{.context.Skills.SkillName}}Type = impl.NewAttackSkillWithOpts(
        impl.WithAttackSkillID(uint16({{.context.Skills.SkillID}})),
		impl.WithAttackSkillType({{.context.Skills.SkillType}}),
		impl.WithAttackSkillCost(map[ElementType]uint{
            // todo: modify me
        }),
		impl.WithAttackSkillActiveDamageHandler(func(ctx def.Context) (elementType ElementType, damageAmount uint) {
            // todo: implement me
			panic("not implement yet")
        }),
		impl.WithAttackSkillBackgroundDamageHandler(func(ctx def.Context) (damageAmount uint) {
            // todo: implement me
            panic("not implement yet")
        }),
    )
)

func Get{{.context.Skills.SkillName}}Entity() def.AttackSkill {
	var entity = {{.context.Skills.SkillName}}
	return entity
}

func init() {
    var modEntity def.Mod = GetMod()
	modEntity.RegisterSkill(Get{{.context.Skills.SkillName}}Entity())
}