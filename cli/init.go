package cli

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
	"github.com/sunist-c/genius-invokation-simulator-backend/mod/implement"
	"github.com/sunist-c/genius-invokation-simulator-cli/advisor"
	"github.com/sunist-c/genius-invokation-simulator-cli/data"
	"github.com/sunist-c/genius-invokation-simulator-cli/generator"
	"github.com/sunist-c/genius-invokation-simulator-cli/localization"
	"github.com/sunist-c/genius-invokation-simulator-cli/log"
	"os"
	"path"
	"path/filepath"
	"text/template"
)

var (
	workingDirectory string

	rootAdvisor advisor.Advisor

	logger = log.Default()

	affiliationSelector = newIndexSelector[enum.Affiliation]()
	elementTypeSelector = newIndexSelector[enum.ElementType]()
	weaponTypeSelector  = newIndexSelector[enum.WeaponType]()
	skillTypeSelector   = newIndexSelector[enum.SkillType]()

	characterTemplate        *template.Template
	attackSkillTemplate      *template.Template
	cooperativeSkillTemplate *template.Template
	passiveSkillTemplate     *template.Template
	currencySkillTemplate    *template.Template
	cardTemplate             *template.Template
	ruleTemplate             *template.Template
	modTemplate              *template.Template

	characterGenerator *generator.Generator[NewCharacterContext]
	cardGenerator      *generator.Generator[NewCardContext]
	ruleGenerator      *generator.Generator[NewRuleContext]
	modGenerator       *generator.Generator[InitModContext]
)

func initSelectors() {
	// init affiliationSelector
	affiliationSelector.loadOptions([]indexOption[enum.Affiliation]{
		{
			key:       "affiliation_mondstadt_desc",
			reference: enum.AffiliationMondstadt,
		},
		{
			key:       "affiliation_liyue_desc",
			reference: enum.AffiliationLiyue,
		},
		{
			key:       "affiliation_inazuma_desc",
			reference: enum.AffiliationInazuma,
		},
		{
			key:       "affiliation_sumeru_desc",
			reference: enum.AffiliationSumeru,
		},
		{
			key:       "affiliation_fatui_desc",
			reference: enum.AffiliationFatui,
		},
		{
			key:       "affiliation_hilichurl_desc",
			reference: enum.AffiliationHilichurl,
		},
		{
			key:       "affiliation_monster_desc",
			reference: enum.AffiliationMonster,
		},
		{
			key:       "affiliation_undefined_desc",
			reference: enum.AffiliationUndefined,
		},
	}...)

	// init elementTypeSelector
	elementTypeSelector.loadOptions([]indexOption[enum.ElementType]{
		{
			key:       "element_currency_desc",
			reference: enum.ElementCurrency,
		},
		{
			key:       "element_anemo_desc",
			reference: enum.ElementAnemo,
		},
		{
			key:       "element_cryo_desc",
			reference: enum.ElementCryo,
		},
		{
			key:       "element_dendro_desc",
			reference: enum.ElementDendro,
		},
		{
			key:       "element_electro_desc",
			reference: enum.ElementElectro,
		},
		{
			key:       "element_geo_desc",
			reference: enum.ElementGeo,
		},
		{
			key:       "element_hydro_desc",
			reference: enum.ElementHydro,
		},
		{
			key:       "element_pyro_desc",
			reference: enum.ElementPyro,
		},
		{
			key:       "element_none_desc",
			reference: enum.ElementNone,
		},
		{
			key:       "element_undifined_desc",
			reference: enum.ElementUndefined,
		},
	}...)

	// init weaponTypeSelector
	weaponTypeSelector.loadOptions([]indexOption[enum.WeaponType]{
		{
			key:       "weapon_sword_desc",
			reference: enum.WeaponSword,
		},
		{
			key:       "weapon_claymore_desc",
			reference: enum.WeaponClaymore,
		},
		{
			key:       "weapon_bow_desc",
			reference: enum.WeaponBow,
		},
		{
			key:       "weapon_catalyst_desc",
			reference: enum.WeaponCatalyst,
		},
		{
			key:       "weapon_polearm_desc",
			reference: enum.WeaponPolearm,
		},
		{
			key:       "weapon_others_desc",
			reference: enum.WeaponOthers,
		},
	}...)

	// init skillTypeSelector
	skillTypeSelector.loadOptions([]indexOption[enum.SkillType]{
		{
			key:       "skill_passive_desc",
			reference: enum.SkillPassive,
		},
		{
			key:       "skill_normal_attack_desc",
			reference: enum.SkillNormalAttack,
		},
		{
			key:       "skill_elemental_skill_desc",
			reference: enum.SkillElementalSkill,
		},
		{
			key:       "skill_elemental_burst_desc",
			reference: enum.SkillElementalBurst,
		},
		{
			key:       "skill_cooperative_desc",
			reference: enum.SkillCooperative,
		},
	}...)
}

func initAdvisors() {
	// main commands
	initializingAdvisor := advisor.NewAdvisorWithOpts(
		advisor.WithAdvisorDepth(0),
		advisor.WithAdvisorEntrance(""),
		advisor.WithAdvisorSuggesterFunctions(
			InitializeJudgeSuggesterFunc(data.GetContext())),
		advisor.WithAdvisorBuiltinSuggestions(AllCommand...),
		advisor.WithAdvisorFunctionChain(
			ExitCliHandler(),
		),
	)

	{
		// new entity command
		newAdvisor := initializingAdvisor.PluginChildWithOpts(
			advisor.WithAdvisorEntrance("new"),
			advisor.WithAdvisorSuggesterFunctions(
				InitializeJudgeSuggesterFunc(data.GetContext()),
			),
			advisor.WithAdvisorBuiltinSuggestions(AllEntityType...),
			advisor.WithAdvisorFunctionChain(func(argument string) {
				fmt.Println(argument)
			}),
		)

		{
			// new character command
			newAdvisor.PluginChildWithOpts(
				advisor.WithAdvisorEntrance("character"),
				advisor.WithAdvisorSuggesterFunctions(
					InitializeJudgeSuggesterFunc(data.GetContext()),
					FirstLetterUpperCaseLegalSuggesterFunc(),
					StaticSuggesterFunc([]prompt.Suggest{
						{
							Text:        "character_name",
							Description: localization.LanguagePack.GetTranslation(localization.GetLocalLanguage(), "new_character_help"),
						},
						{
							Text:        "Keqing",
							Description: localization.LanguagePack.GetTranslation(localization.GetLocalLanguage(), "new_character_example"),
						},
					}...),
				),
				advisor.WithAdvisorFunctionChain(
					NewCharacterHandler(data.GetContext()),
				),
			)
		}
	}

	{
		// list entities command
		initializingAdvisor.PluginChildWithOpts(
			advisor.WithAdvisorEntrance("list"),
			advisor.WithAdvisorBuiltinSuggestions(AllEntityType...),
			advisor.WithAdvisorFunctionChain(func(argument string) {
				fmt.Println(argument)
			}),
		)
	}

	{
		// init mod command
		initializingAdvisor.PluginChildWithOpts(
			advisor.WithAdvisorEntrance("init"),
			advisor.WithAdvisorSuggesterFunctions(
				PackagePathLegalSuggesterFunc(),
				StaticSuggesterFunc([]prompt.Suggest{
					{
						Text:        "package_name",
						Description: localization.LanguagePack.GetTranslation(localization.GetLocalLanguage(), "init_mod_help"),
					},
					{
						Text:        "github.com/sunist-c/gisb-base-mod",
						Description: localization.LanguagePack.GetTranslation(localization.GetLocalLanguage(), "init_mod_public_example"),
					},
					{
						Text:        "gisb_base_mod",
						Description: localization.LanguagePack.GetTranslation(localization.GetLocalLanguage(), "init_mod_private_example"),
					},
				}...),
			),
			advisor.WithAdvisorFunctionChain(
				InitModHandler(data.GetContext()),
			),
		)
	}

	rootAdvisor = initializingAdvisor
}

func initTemplates() {
	// attackSkillTemplate
	if attackSkillParsingTemplate, err := template.ParseFiles(filepath.Join("./template", "attack_skill.tpl")); err != nil {
		logger.Panicf("failed to parse attack skill template: %v", err)
	} else {
		attackSkillTemplate = attackSkillParsingTemplate
	}

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

	logger.Logf("try to find metadata.yml in [%v]", workingDirectory)
	if _, err := os.Stat(filepath.Join(workingDirectory, "metadata.yml")); err != nil && !os.IsNotExist(err) {
		logger.Panicf("find metadata.yml in [%v] failed: %v", workingDirectory, err)
	} else if os.IsNotExist(err) {
		logger.Logf("no metadata.yml detected in [%v]", workingDirectory)
		data.GetContext().Initialized = false
	} else {
		logger.Logf("parsing metadata file [%v]", filepath.Join(workingDirectory, "metadata.yml"))
		logger.Logf("current mod id: %v", implement.ModID())
	}

	logger.Logf("initializing selectors")
	initSelectors()

	logger.Logf("initializing advisors")
	initAdvisors()

	logger.Logf("initializing templates")
	initTemplates()
}
