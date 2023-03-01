package advisor

import (
	"github.com/c-bata/go-prompt"
	"github.com/sunist-c/genius-invokation-simulator-cli/localization"
)

func InvalidAdvisor() Advisor {
	return NewAdvisorWithOpts(
		WithAdvisorDepth(0),
		WithAdvisorSuggesterFunctions(
			func(ctx *SuggesterContext) {
				ctx.Result = []prompt.Suggest{
					{
						Text:        "",
						Description: localization.LanguagePack.GetTranslation(localization.GetLocalLanguage(), "invalid_command_desc"),
					},
				}
				ctx.Abort()
			},
		),
	)
}

func NoMatchedAdvisor() Advisor {
	return NewAdvisorWithOpts(
		WithAdvisorDepth(0),
		WithAdvisorSuggesterFunctions(
			func(ctx *SuggesterContext) {
				ctx.Result = append([]prompt.Suggest{
					{
						Text:        "",
						Description: localization.LanguagePack.GetTranslation(localization.GetLocalLanguage(), "no_matched_command_desc"),
					},
				}, ctx.Result...)
			}),
	)
}
