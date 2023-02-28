package cli

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/mod/implement"
	"os"
	"strings"
	"time"
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
				time.Sleep(time.Millisecond * 100)
				os.Exit(0)
			}()
		}
	}
}

func NewCharacterHandler() HandlerFunc {
	return func(argument string) {
		characterName := argument
		logger.Logf("trying to create character %v", characterName)

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

		logger.Logf("generating character %v.go", characterName)

	}
}
