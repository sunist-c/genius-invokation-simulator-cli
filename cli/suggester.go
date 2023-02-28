package cli

import (
	"github.com/c-bata/go-prompt"
	"github.com/sunist-c/genius-invokation-simulator-cli/data"
)

type SuggesterFunc func(d *prompt.Document) []prompt.Suggest

func ConvertSuggesterFuncToCompleter(f SuggesterFunc) func(d prompt.Document) []prompt.Suggest {
	return func(d prompt.Document) []prompt.Suggest {
		adapterLayer := &d
		return f(adapterLayer)
	}
}

func ListCharactersSuggesterFunc(ctx *data.Context) SuggesterFunc {
	return func(d *prompt.Document) []prompt.Suggest {
		return nil
	}
}

func PrintLegalSuggesterFunc(legalJudge func(argument string) (isLegal bool)) SuggesterFunc {
	return func(d *prompt.Document) []prompt.Suggest {
		input := d.GetWordBeforeCursor()
		if legalJudge(input) {
			return []prompt.Suggest{
				{
					Text:        input,
					Description: languagePack.GetTranslation(LocalLanguage(), "correct_input_format"),
				},
			}
		} else {
			return []prompt.Suggest{
				{
					Text:        input,
					Description: languagePack.GetTranslation(LocalLanguage(), "incorrect_input_format"),
				},
			}
		}
	}
}
