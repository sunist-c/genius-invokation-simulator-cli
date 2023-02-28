package cli

import (
	"github.com/c-bata/go-prompt"
)

func Exec() {
	logger.Logf("running gis-cli tools")
	for {
		input := prompt.Input("gis-cli> ", Completer)
		logger.Logf("executing command [%v]", input)
		splits := parseArgs(input)
		handler := FindHandler(advisor, splits)
		handler(splits[len(splits)-1])
	}
}
