package {{.Paths.Mod}}

import (
    def "github.com/sunist-c/genius-invokation-simulator-backend/mod/definition"
    impl "github.com/sunist-c/genius-invokation-simulator-backend/mod/implement"
)

type {{.Mod.ModName}}Type def.Mod

var (
	{{.Mod.ModName}}Entity {{.Mod.ModName}}Type = impl.NewMod()
)

func GetMod() def.Mod {
	return {{.Mod.ModName}}Entity
}