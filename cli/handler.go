package cli

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
	"github.com/sunist-c/genius-invokation-simulator-backend/mod/implement"
	"github.com/sunist-c/genius-invokation-simulator-cli/data"
	"github.com/sunist-c/genius-invokation-simulator-cli/generator"
	"github.com/sunist-c/genius-invokation-simulator-cli/localization"
	"os"
	"path/filepath"
	"strings"
)

func generateLongID(shortID uint16) uint64 {
	return implement.NewEntityWithOpts(implement.WithEntityID(shortID)).TypeID()
}

type HandlerFunc func(argument string)

func ExitCliHandler() HandlerFunc {
	return func(argument string) {
		if strings.ToLower(argument) == "exit" {
			logger.Logf("exit gis-cli tools")
			defer func() {
				Exit(0)
			}()
		}
	}
}

func NewCharacterHandler(ctx *data.Context) HandlerFunc {
	return func(argument string) {
		characterName := argument
		logger.Logf("trying to create character %v", characterName)
		printTranslation("new_entity_warning")

		// ID分配
		var characterID uint16
		printTranslation("new_character_id_question")
		if yesNoParser() {
			printTranslation("new_character_id_description")

		inputID:
			want := uint16Parser()
			if success, result := implement.UseID(want); success {
				characterID = want
				printTranslation("new_character_id_success_result", want, generateLongID(want))
			} else {
				printTranslation("new_character_id_failed_result")
				if yesNoParser() {
					goto inputID
				} else {
					characterID = result
					printTranslation("new_character_id_result", result, generateLongID(result))
				}
			}
		} else {
			characterID = implement.NextID()
			printTranslation("new_character_id_result", characterID, generateLongID(characterID))
		}

		// 势力所属
		printTranslation("new_character_affiliation_description", affiliationSelector.formatString(" "))
		affiliationIndex := indexParser(1, affiliationSelector.length())
		printTranslation("new_character_affiliation_result", affiliationSelector.getContent(affiliationIndex))

		// 元素类型
		printTranslation("new_character_vision_description", elementTypeSelector.formatString(" "))
		elementTypeIndex := indexParser(1, elementTypeSelector.length())
		printTranslation("new_character_vision_result", elementTypeSelector.getContent(elementTypeIndex))

		// 武器类型
		printTranslation("new_character_weapon_description", weaponTypeSelector.formatString(" "))
		weaponTypeIndex := indexParser(1, weaponTypeSelector.length())
		printTranslation("new_character_weapon_result", weaponTypeSelector.getContent(weaponTypeIndex))

		// HP
		printTranslation("new_character_hp_description")
		characterHP := positiveIntParser()
		printTranslation("new_character_hp_result", characterHP)

		// MP
		printTranslation("new_character_mp_description")
		characterMP := positiveIntParser()
		printTranslation("new_character_mp_result", characterMP)

		// 技能生成
		printTranslation("new_character_with_skill_question")
		generateSkill := yesNoParser()
		var skillOutDirectory string
		if generateSkill {
			skillOutDirectory = filepath.Join(workingDirectory, "./skill", fmt.Sprintf("%v", characterName))
			if err := os.MkdirAll(skillOutDirectory, 0666); err != nil && !os.IsExist(err) {
				logger.Errorf("creating character skill directory [%v] failed: %v", skillOutDirectory, err)
				printTranslation("io_failed_description", skillOutDirectory)
				Exit(403)
			}
		}
		for generateSkill {
			printTranslation("new_character_with_skill_description")
			skillName := readLine(
				FirstLetterUpperCaseLegalSuggesterFunc(),
				StaticSuggesterFunc([]prompt.Suggest{
					{
						Text:        "skill_name",
						Description: localization.LanguagePack.GetTranslation(localization.GetLocalLanguage(), "new_character_with_skill_help"),
					},
					{
						Text:        "KeqingNormalAttack",
						Description: localization.LanguagePack.GetTranslation(localization.GetLocalLanguage(), "new_character_with_skill_example"),
					},
				}...))

			logger.Logf("try to create skill %v", skillName)

			// 技能ID
			var skillID uint16
			var skillLongID uint64
			printTranslation("new_character_with_skill_id_question")
			if yesNoParser() {
				printTranslation("new_character_with_skill_id_description")

			inputSKillID:
				want := uint16Parser()
				if success, result := implement.UseID(want); success {
					skillID = want
					skillLongID = generateLongID(want)
					printTranslation("new_character_with_skill_id_success_result", skillID, skillLongID)
				} else {
					printTranslation("new_character_with_skill_id_failed_result")
					if yesNoParser() {
						goto inputSKillID
					} else {
						skillID = result
						skillLongID = generateLongID(result)
						printTranslation("new_character_with_skill_id_result", skillID, skillLongID)
					}
				}
			} else {
				skillID = implement.NextID()
				skillLongID = generateLongID(skillID)
				printTranslation("new_character_with_skill_id_result", skillID, skillLongID)
			}

			// 技能类型
			printTranslation("new_character_with_skill_type_description", skillTypeSelector.formatString(" "))
			skillTypeIndex := indexParser(1, skillTypeSelector.length())
			printTranslation("new_character_with_skill_type_result", skillTypeSelector.getContent(skillTypeIndex))

			printTranslation("new_character_with_skill_continue_question")
			generateSkill = yesNoParser()

			skillFilePath := filepath.Join(skillOutDirectory, fmt.Sprintf("%v.go", skillName))
			skillContext := NewSkillContext{
				CharacterName: characterName,
				SkillName:     skillName,
				SkillID:       skillID,
				SkillType:     skillTypeSelector.getReference(skillTypeIndex),
			}

			var skillGenerator *generator.Generator[context]
			switch skillContext.SkillType {
			case enum.SkillNormalAttack, enum.SkillElementalSkill, enum.SkillElementalBurst:
				skillGenerator = generator.NewGeneratorWithOpts[context](
					generator.WithGeneratorWorkingDirectory[context](skillOutDirectory),
					generator.WithGeneratorTemplate[context](attackSkillTemplate),
				)
			case enum.SkillCooperative:
				skillGenerator = generator.NewGeneratorWithOpts[context](
					generator.WithGeneratorWorkingDirectory[context](skillOutDirectory),
					generator.WithGeneratorTemplate[context](cooperativeSkillTemplate),
				)
			case enum.SkillPassive:
				skillGenerator = generator.NewGeneratorWithOpts[context](
					generator.WithGeneratorWorkingDirectory[context](skillOutDirectory),
					generator.WithGeneratorTemplate[context](passiveSkillTemplate),
				)
			default:
				logger.Errorf("unsupported skill type in skill %v: %v", skillName, skillTypeSelector.getReference(skillTypeIndex))
				skillGenerator = generator.NewGeneratorWithOpts[context](
					generator.WithGeneratorWorkingDirectory[context](skillOutDirectory),
					generator.WithGeneratorTemplate[context](currencySkillTemplate),
				)
			}

			logger.Logf("generating skill %v in [%v]", skillName, skillFilePath)
			if err := skillGenerator.GenerateFile(fmt.Sprintf("%v.go", skillName), context{
				"Skill": skillContext,
				"Paths": ctx.Paths,
			}); err != nil {
				logger.Errorf("error generating skill %v in [%v]: %v", skillName, skillFilePath, err)
				printTranslation("no_files_created")
				Exit(500)
			} else {
				logger.Logf("register skill %v(%v) to context", skillName, skillLongID)
				ctx.RegisterSkill(
					data.EntityIndex{
						ShortID: skillID,
						LongID:  skillLongID,
					},
					data.CreateSkillArguments{
						OwnCharacterName: characterName,
						SkillName:        skillName,
						SkillType:        skillTypeSelector.getReference(skillTypeIndex),
					},
				)

				printTranslation("success_generated_file", skillFilePath)
			}
		}

		characterFilePath := filepath.Join(workingDirectory, "./character", fmt.Sprintf("./%v.go", characterName))
		logger.Logf("generating character %v in [%v]", characterName, characterFilePath)
	}
}

func InitModHandler(ctx *data.Context) HandlerFunc {
	return func(argument string) {
		packageName := argument
		logger.Logf("trying to init mod %v", packageName)
		printTranslation("new_entity_warning")

		implement.InitMetaData()
		modID := implement.ModID()
		ctx.Initialized = true
		ctx.Mod.InitArguments = data.InitModArguments{
			ID:          modID,
			PackagePath: packageName,
		}

		logger.Logf("init mod %v(%v) success", packageName, modID)
		printTranslation("init_mod_id_result", packageName, modID)
	}
}
