package cli

import (
	"fmt"
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

func readLine() (line string) {
	_, err := fmt.Scanf("%s", &line)
	if err != nil {
		printInputError("buff reader readline failed: %v", err)
		return readLine()
	} else {
		return line
	}
}

func printTranslation(key string, args ...interface{}) {
	fmt.Println(fmt.Sprintf(languagePack.GetTranslation(LocalLanguage(), key), args...))
}

func printInputError(formant string, args ...interface{}) {
	logger.Errorf(formant, args...)
	printTranslation("incorrect_input_format")
}

func yesNoParser() bool {
	input := strings.ToUpper(readLine())
	return strings.HasPrefix(input, "Y")
}

func uint16Parser() uint16 {
	input := readLine()
	uintVal, err := strconv.ParseUint(input, 10, 64)
	if err != nil {
		printInputError("parse input [%v] to uint16 failed: %v", input, err)
		return uint16Parser()
	} else {
		return uint16(uintVal & uint64(1<<16-1))
	}
}

func intParser() int {
	input := readLine()
	intVal, err := strconv.Atoi(input)
	if err != nil {
		printInputError("parse input [%v] to int failed: %v", input, err)
		return intParser()
	} else {
		return intVal
	}
}

func indexParser(lowerBound, upperBound int) int {
	intVal := intParser()
	if intVal > upperBound || intVal < lowerBound {
		printInputError("incorrect index [%v] to index range [%v, %v]", intVal, lowerBound, upperBound)
		return indexParser(lowerBound, upperBound)
	} else {
		return intVal
	}
}
