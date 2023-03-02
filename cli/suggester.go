package cli

import (
	"github.com/c-bata/go-prompt"
	"github.com/sunist-c/genius-invokation-simulator-cli/advisor"
	"github.com/sunist-c/genius-invokation-simulator-cli/data"
	"github.com/sunist-c/genius-invokation-simulator-cli/localization"
)

func InitializeJudgeSuggesterFunc(context *data.Context) advisor.SuggesterFunc {
	return func(ctx *advisor.SuggesterContext) {
		if !context.Initialized {
			ctx.AppendSuggestBefore(prompt.Suggest{
				Text:        ctx.Document.GetWordBeforeCursor(),
				Description: localization.LanguagePack.GetTranslation(localization.GetLocalLanguage(), "mod_not_init"),
			})
		}
	}
}

func LegalJudgeSuggesterFunc(isLegal func(argument string) bool) advisor.SuggesterFunc {
	return func(ctx *advisor.SuggesterContext) {
		ctx.Next()
		input := ctx.Document.GetWordBeforeCursor()
		if isLegal(input) {
			ctx.AppendSuggestBefore(prompt.Suggest{
				Text:        input,
				Description: localization.LanguagePack.GetTranslation(localization.GetLocalLanguage(), "correct_input_format"),
			})
		} else {
			ctx.AppendSuggestBefore(prompt.Suggest{
				Text:        input,
				Description: localization.LanguagePack.GetTranslation(localization.GetLocalLanguage(), "incorrect_input_format"),
			})
		}
	}
}

func FirstLetterUpperCaseLegalSuggesterFunc() advisor.SuggesterFunc {
	return LegalJudgeSuggesterFunc(func(argument string) bool {
		if argument == "" {
			return false
		} else if runes := []rune(argument); len(runes) == 0 {
			return false
		} else if runes[0] >= 97 && runes[0] <= 122 {
			return false
		} else {
			return runes[0] >= 65 && runes[0] <= 90
		}
	})
}

func PackagePathLegalSuggesterFunc() advisor.SuggesterFunc {
	return LegalJudgeSuggesterFunc(func(argument string) bool {
		if argument == "" {
			return false
		} else if runes := []rune(argument); len(runes) == 0 {
			return false
		} else if !((97 <= runes[0] && runes[0] <= 122) || 65 <= runes[0] && runes[0] <= 90) {
			return false
		} else {
			for _, r := range runes {
				if !((97 <= r && r <= 122) ||
					(65 <= r && r <= 90) ||
					(48 <= r && r <= 57) ||
					r == '/' || r == '.' ||
					r == '-' || r == '_') {
					return false
				}
			}

			return true
		}
	})
}

func StaticSuggesterFunc(staticSuggestions ...prompt.Suggest) advisor.SuggesterFunc {
	return func(ctx *advisor.SuggesterContext) {
		ctx.AppendSuggestAfter(staticSuggestions...)
	}
}
