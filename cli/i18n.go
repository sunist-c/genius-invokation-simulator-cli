package cli

import (
	"fmt"
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
	"github.com/sunist-c/genius-invokation-simulator-cli/log"
	"github.com/sunist-c/genius-invokation-simulator-cli/parser"
	"gopkg.in/yaml.v2"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var (
	languagePack SuggestionLanguageSupport

	localLanguage               = enum.Unknown
	fullStringToLanguageAdapter = parser.NewFullStringToLanguageAdapter()
	languageToFullStringAdapter = parser.NewLanguageToFullStringAdapter()
)

func init() {
	logger := log.Default()
	pathOfLocalization := path.Join(filepath.Dir(os.Args[0]), "./i18n/localization.yml")

	logger.Logf("try to read localization file [%v]", pathOfLocalization)
	bytesOfLocalizationFile, readErr := os.ReadFile(pathOfLocalization)
	if readErr != nil {
		logger.Panicf("failed to read localization file [%v]", pathOfLocalization)
	}

	languagePack = SuggestionLanguageSupport{dictionary: map[string]map[string]string{}}
	unmarshalErr := yaml.Unmarshal(bytesOfLocalizationFile, &languagePack.dictionary)
	if unmarshalErr != nil {
		logger.Panicf("failed to unmarshal localization file [%v]", pathOfLocalization)
	}

	logger.Logf("read localization file success")
}

type SuggestionLanguageSupport struct {
	dictionary map[string]map[string]string
}

func (support SuggestionLanguageSupport) GetTranslation(language enum.Language, key string) (translation string) {
	_, lang := languageToFullStringAdapter.Convert(language)
	return support.dictionary[lang][key]
}

func LocalLanguage() enum.Language {
	if localLanguage == enum.Unknown {
		localLanguage, _ = getLocalLanguage()
	}

	return localLanguage
}

func getLocalLanguage() (language enum.Language, err error) {
	logger := log.Default()
	logger.Logf("trying to get local language")
	defer func() {
		if err != nil {
			logger.Errorf("get local language failed: %v", err)
		} else {
			_, stringLanguage := languageToFullStringAdapter.Convert(language)
			logger.Logf("get local language %v success", stringLanguage)
		}
	}()

	envLang := os.Getenv("LANG")
	if envLang == "" {
		return enum.English, fmt.Errorf("no local language detected")
	} else if args := strings.Split(envLang, "."); args == nil || len(args) != 2 {
		return enum.English, fmt.Errorf("cannot dial local language")
	} else {
		languageCode := strings.ToLower(args[0])
		if strings.HasPrefix(languageCode, "zh_cn") ||
			strings.HasPrefix(languageCode, "zh_sg") ||
			strings.HasPrefix(languageCode, "zh_hans") {
			return enum.ChineseSimplified, nil
		} else if strings.HasPrefix(languageCode, "zh_tw") ||
			strings.HasPrefix(languageCode, "zh_hk") ||
			strings.HasPrefix(languageCode, "zh_mo") ||
			strings.HasPrefix(languageCode, "zh_hant") {
			return enum.ChineseTraditional, nil
		} else if strings.HasPrefix(languageCode, "en") {
			return enum.English, nil
		} else if strings.HasPrefix(languageCode, "fr") {
			return enum.French, nil
		} else if strings.HasPrefix(languageCode, "de") {
			return enum.German, nil
		} else if strings.HasPrefix(languageCode, "ja") {
			return enum.Japanese, nil
		} else if strings.HasPrefix(languageCode, "ko") {
			return enum.Korean, nil
		} else if strings.HasPrefix(languageCode, "ru") {
			return enum.Russian, nil
		} else {
			return enum.English, fmt.Errorf("unsupported language: %s", languageCode)
		}
	}
}
