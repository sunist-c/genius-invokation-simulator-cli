package generator

import (
	"github.com/sunist-c/genius-invokation-simulator-cli/log"

	"fmt"
	"os"
	"path"
	"text/template"
)

type Generator[Context any] struct {
	workingDirectory string
	textTemplate     *template.Template
	logger           log.Logger
}

func (g *Generator[Context]) GenerateFile(fileName string, ctx Context) (err error) {
	filePath := path.Join(g.workingDirectory, fileName)
	g.logger.Logf("start generate file [%v]", filePath)
	defer func() {
		if err != nil {
			g.logger.Errorf("generate file [%v] with error: %v", filePath, err)
		} else {
			g.logger.Logf("generate file [%v] success", filePath)
		}
	}()

	if g.textTemplate == nil {
		return fmt.Errorf("nil text template")
	}

	file, openFileErr := os.OpenFile(path.Join(g.workingDirectory, fileName), os.O_RDWR|os.O_CREATE, 0666)
	if openFileErr != nil {
		return openFileErr
	} else {
		defer func() {
			closeErr := file.Close()
			if closeErr != nil && err == nil {
				err = fmt.Errorf("cannot close file %v: %v", path.Join(g.workingDirectory, fileName), closeErr)
			}
		}()
	}

	if generateErr := g.textTemplate.Execute(file, ctx); generateErr != nil {
		return generateErr
	}

	return nil
}

type Options[Context any] func(option *Generator[Context])

func WithGeneratorWorkingDirectory[Context any](directory string) Options[Context] {
	return func(option *Generator[Context]) {
		option.workingDirectory = directory
	}
}

func WithGeneratorTemplate[Context any](tpl *template.Template) Options[Context] {
	return func(option *Generator[Context]) {
		option.textTemplate = tpl
	}
}

func NewGeneratorWithOpts[Context any](options ...Options[Context]) *Generator[Context] {
	generator := &Generator[Context]{
		logger: log.Default(),
	}
	for _, option := range options {
		option(generator)
	}

	return generator
}
