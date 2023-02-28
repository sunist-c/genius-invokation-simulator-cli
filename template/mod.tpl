package {{.context.Paths.Mod}}

import (
    def "github.com/sunist-c/genius-invokation-simulator-backend/mod/definition"
    impl "github.com/sunist-c/genius-invokation-simulator-backend/mod/implement"
)

type {{.context.Mod.ModName}}Type def.Mod

var (
	{{.context.Mod.ModName}} {{.context.Mod.ModName}}Type = impl.NewMod()
)

func GetMod() def.Mod {
	return {{.context.Mod.ModName}}
}