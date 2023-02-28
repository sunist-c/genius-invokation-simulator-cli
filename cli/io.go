package cli

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"strconv"
	"strings"
)

type indexSelector struct {
	indexes map[int]string
}

func (is *indexSelector) loadTranslation(keys ...string) {
	for index, key := range keys {
		is.indexes[index+1] = languagePack.GetTranslation(LocalLanguage(), key)
	}
}

func (is *indexSelector) formatString(sep string) string {
	ts := make([]string, len(is.indexes)+1)
	for i := 1; i <= len(is.indexes); i++ {
		ts[i-1] = fmt.Sprintf("[%v] %v", i, is.indexes[i])
	}

	return strings.Join(ts, sep)
}

func readLine(legalJudge func(string) bool) (line string) {
	return prompt.Input(
		"input: ",
		ConvertSuggesterFuncToCompleter(PrintLegalSuggesterFunc(legalJudge)),
	)
}

func printTranslation(key string, args ...interface{}) {
	fmt.Println(fmt.Sprintf(languagePack.GetTranslation(LocalLanguage(), key), args...))
}

func printInputError(formant string, args ...interface{}) {
	logger.Errorf(formant, args...)
	printTranslation("incorrect_input_format")
}

func yesNoParser(prefix ...string) bool {
	input := readLine(func(s string) bool { return true })
	return strings.HasPrefix(strings.ToUpper(input), "Y")
}

func uint16Parser() uint16 {
	input := readLine(func(s string) bool {
		uintVal, err := strconv.ParseUint(s, 10, 64)
		if err != nil {
			return false
		} else {
			return uintVal < 1<<16
		}
	})

	uintVal, err := strconv.ParseUint(input, 10, 64)
	if err != nil {
		printInputError("parse input [%v] to uint16 failed: %v", input, err)
		return uint16Parser()
	} else {
		return uint16(uintVal & uint64(1<<16-1))
	}
}

func intParser() int {
	input := readLine(func(s string) bool {
		_, err := strconv.Atoi(s)
		if err != nil {
			return false
		} else {
			return true
		}
	})

	intVal, err := strconv.Atoi(input)
	if err != nil {
		printInputError("parse input [%v] to int failed: %v", input, err)
		return intParser()
	} else {
		return intVal
	}
}

func indexParser(lowerBound, upperBound int) int {
	input := readLine(func(s string) bool {
		intVal, err := strconv.Atoi(s)
		if err != nil {
			return false
		} else {
			return intVal <= upperBound && intVal >= lowerBound
		}
	})

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
