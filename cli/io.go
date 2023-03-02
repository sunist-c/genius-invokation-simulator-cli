package cli

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/sunist-c/genius-invokation-simulator-cli/advisor"
	"github.com/sunist-c/genius-invokation-simulator-cli/localization"
	"strconv"
	"strings"
)

type indexOption[entity any] struct {
	key       string
	reference entity
}

type indexSelector[reference any] struct {
	indexes    map[int]string
	references map[int]reference
}

func (is *indexSelector[reference]) loadOptions(options ...indexOption[reference]) {
	for index, option := range options {
		is.indexes[index+1] = localization.LanguagePack.GetTranslation(localization.GetLocalLanguage(), option.key)
		is.references[index+1] = option.reference
	}
}

func (is *indexSelector[reference]) formatString(sep string) string {
	ts := make([]string, len(is.indexes)+1)
	for i := 1; i <= len(is.indexes); i++ {
		ts[i-1] = fmt.Sprintf("[%v] %v", i, is.indexes[i])
	}

	return strings.Join(ts, sep)
}

func (is *indexSelector[reference]) length() int {
	return len(is.indexes)
}

func (is *indexSelector[reference]) getContent(index int) string {
	return is.indexes[index]
}

func (is *indexSelector[reference]) getReference(index int) reference {
	return is.references[index]
}

func newIndexSelector[reference any]() *indexSelector[reference] {
	return &indexSelector[reference]{
		indexes:    map[int]string{},
		references: map[int]reference{},
	}
}

func readLine(handlers ...advisor.SuggesterFunc) (line string) {
	return prompt.Input("input: ", func(document prompt.Document) []prompt.Suggest {
		ctx := advisor.NewContext(&document, handlers...)
		ctx.Next()
		return ctx.Result()
	})
}

func printTranslation(key string, args ...interface{}) {
	fmt.Println(fmt.Sprintf(localization.LanguagePack.GetTranslation(localization.GetLocalLanguage(), key), args...))
}

func printInputError(formant string, args ...interface{}) {
	logger.Errorf(formant, args...)
	printTranslation("incorrect_input_format")
}

func stringParser() string {
	input := readLine(
		LegalJudgeSuggesterFunc(func(argument string) bool {
			return true
		}),
	)

	return input
}

func yesNoParser() bool {
	input := readLine(
		LegalJudgeSuggesterFunc(func(s string) bool {
			upperCase := strings.ToUpper(s)
			return strings.HasPrefix(upperCase, "Y") || strings.HasPrefix(upperCase, "N")
		}),
	)

	return strings.HasPrefix(strings.ToUpper(input), "Y")
}

func uint16Parser() uint16 {
	input := readLine(
		LegalJudgeSuggesterFunc(func(s string) bool {
			uintVal, err := strconv.ParseUint(s, 10, 64)
			return err == nil && uintVal < 1<<16
		}),
	)

	uintVal, err := strconv.ParseUint(input, 10, 64)
	if err != nil {
		printInputError("parse input [%v] to uint16 failed: %v", input, err)
		return uint16Parser()
	} else {
		return uint16(uintVal & uint64(1<<16-1))
	}
}

func intParser() int {
	input := readLine(
		LegalJudgeSuggesterFunc(func(s string) bool {
			_, err := strconv.Atoi(s)
			return err == nil
		}),
	)

	intVal, err := strconv.Atoi(input)
	if err != nil {
		printInputError("parse input [%v] to int failed: %v", input, err)
		return intParser()
	} else {
		return intVal
	}
}

func positiveIntParser() uint {
	input := readLine(
		LegalJudgeSuggesterFunc(func(s string) bool {
			intVal, err := strconv.Atoi(s)
			return err == nil && intVal >= 0
		}),
	)

	intVal, err := strconv.Atoi(input)
	if err != nil {
		printInputError("parse input [%v] to positive-int failed: %v", input, err)
		return positiveIntParser()
	} else if intVal < 0 {
		printInputError("input [%v] is not a positive number", intVal)
		return positiveIntParser()
	} else {
		return uint(intVal)
	}
}

func indexParser(lowerBound, upperBound int) int {
	input := readLine(
		LegalJudgeSuggesterFunc(func(s string) bool {
			intVal, err := strconv.Atoi(s)
			return err == nil && intVal >= lowerBound && intVal <= upperBound
		}),
	)

	intVal, err := strconv.Atoi(input)
	if err != nil {
		printInputError("parse input [%v] to int-index failed: %v", input, err)
		return indexParser(lowerBound, upperBound)
	}

	if intVal > upperBound || intVal < lowerBound {
		printInputError("incorrect index [%v] to index range [%v, %v]", intVal, lowerBound, upperBound)
		return indexParser(lowerBound, upperBound)
	} else {
		return intVal
	}
}
