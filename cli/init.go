package cli

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/sunist-c/genius-invokation-simulator-backend/mod/implement"
	"github.com/sunist-c/genius-invokation-simulator-cli/generator"
	"github.com/sunist-c/genius-invokation-simulator-cli/log"
	"os"
	"path"
	"text/template"
)

var (
	workingDirectory string

	advisor Advisor

	logger = log.Default()

	affiliationSelector = &indexSelector{indexes: map[int]string{}}
	elementTypeSelector = &indexSelector{indexes: map[int]string{}}
	weaponTypeSelector  = &indexSelector{indexes: map[int]string{}}

	characterTemplate *template.Template
	skillTemplate     *template.Template
	cardTemplate      *template.Template
	ruleTemplate      *template.Template
	modTemplate       *template.Template

	characterGenerator *generator.Generator[NewCharacterContext]
	skillGenerator     *generator.Generator[NewSkillContext]
	cardGenerator      *generator.Generator[NewCardContext]
	ruleGenerator      *generator.Generator[NewRuleContext]
	modGenerator       *generator.Generator[InitModContext]
)

func initSelectors() {
	// init affiliationSelector
	affiliationSelector.loadTranslation(
		"affiliation_mondstadt_desc",
		"affiliation_liyue_desc",
		"affiliation_inazuma_desc",
		"affiliation_sumeru_desc",
		"affiliation_fatui_desc",
		"affiliation_hilichurl_desc",
		"affiliation_monster_desc",
		"affiliation_undefined_desc",
	)

	// init elementTypeSelector
	elementTypeSelector.loadTranslation(
		"element_currency_desc",
		"element_anemo_desc",
		"element_cryo_desc",
		"element_dendro_desc",
		"element_electro_desc",
		"element_geo_desc",
		"element_hydro_desc",
		"element_pyro_desc",
		"element_none_desc",
		"element_undifined_desc",
	)

	// init weaponTypeSelector
	weaponTypeSelector.loadTranslation(
		"weapon_sword_desc",
		"weapon_claymore_desc",
		"weapon_bow_desc",
		"weapon_catalyst_desc",
		"weapon_polearm_desc",
		"weapon_others_desc",
		"",
	)
}

func initAdvisors() {
	// main commands
	rootAdvisor := NewAdvisorWithOpts(
		WithAdvisorDepth(0),
		WithAdvisorEntrance(""),
		WithAdvisorBuiltinSuggestions(AllCommand...),
		WithAdvisorFunctionChain(
			ExitCliHandler(),
		),
	)

	{
		// new entity command
		newAdvisor := rootAdvisor.PluginChildWithOpts(
			WithAdvisorEntrance("new"),
			WithAdvisorBuiltinSuggestions(AllEntityType...),
			WithAdvisorFunctionChain(func(argument string) {
				fmt.Println(argument)
			}),
		)

		{
			// new character command
			newAdvisor.PluginChildWithOpts(
				WithAdvisorEntrance("character"),
				WithAdvisorSuggestionFunction(func(d *prompt.Document) []prompt.Suggest {
					return []prompt.Suggest{
						{
							Text:        "character_name",
							Description: languagePack.GetTranslation(LocalLanguage(), "new_character_help"),
						},
						{
							Text:        "Keqing",
							Description: languagePack.GetTranslation(LocalLanguage(), "new_character_example"),
						},
					}
				}),
				WithAdvisorFunctionChain(
					NewCharacterHandler(),
				),
			)
		}
	}

	{
		// list entities command
		rootAdvisor.PluginChildWithOpts(
			WithAdvisorEntrance("list"),
			WithAdvisorBuiltinSuggestions(AllEntityType...),
			WithAdvisorFunctionChain(func(argument string) {
				fmt.Println(argument)
			}),
		)
	}

	advisor = rootAdvisor
}

func initGenerators() {
	characterGenerator = generator.NewGeneratorWithOpts[NewCharacterContext](
		generator.WithGeneratorTemplate[NewCharacterContext](characterTemplate),
		generator.WithGeneratorWorkingDirectory[NewCharacterContext](path.Join(workingDirectory, "./character")),
	)
}

func init() {
	pwd, err := os.Getwd()
	if err != nil {
		logger.Panicf("get working directory failed: %v", err)
	} else {
		workingDirectory = pwd
	}

	logger.Logf("initializing metadata in [%v]", workingDirectory)
	implement.InitMetaData()
	logger.Logf("current mod id: %v", implement.ModID())

	logger.Logf("initializing io-util.selectors")
	initSelectors()

	logger.Logf("initializing advisors")
	initAdvisors()
}
