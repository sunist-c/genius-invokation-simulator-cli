package {{.generateContext.PackagePath}}

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/model/adapter"

	{{range .generateContext.ImportPackages}} {{.Alias}} "{{.Full}}"
	{{end}}
)

type {{.generateContext.AdapterName}} struct {
	dictionary map[{{.generateContext.SourceType}}]{{.generateContext.DestType}}
}

func (adapter *{{.generateContext.AdapterName}}) Convert(source {{.generateContext.SourceType}}) (success bool, result {{.generateContext.DestType}}) {
	dictionaryResult, exist := adapter.dictionary[source]
	return exist, dictionaryResult
}

func New{{.generateContext.AdapterName}}() adapter.Adapter[{{.generateContext.SourceType}}, {{.generateContext.DestType}}] {
	return &{{.generateContext.AdapterName}} {
		dictionary: map[{{.generateContext.SourceType}}]{{.generateContext.DestType}}{
			{{range .generateContext.Enums}} {{.SourceValue}}: {{.DestValue}},
			{{end}}
		},
	}
}