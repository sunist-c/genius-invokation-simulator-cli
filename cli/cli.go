package cli

import (
	"github.com/c-bata/go-prompt"
	"strings"
)

func Completer(d prompt.Document) []prompt.Suggest {
	return Complete(d, advisor, parseArgs(d.TextBeforeCursor()))
}

func parseArgs(arg string) []string {
	splits := strings.Split(arg, " ")
	args := make([]string, 0, len(splits))

	for i := range splits {
		if i != len(splits)-1 && splits[i] == "" {
			continue
		}
		args = append(args, splits[i])
	}

	return args
}

func Complete(d prompt.Document, advisor Advisor, args []string) (suggestions []prompt.Suggest) {
	advisor.SyncDocument(&d)

	if len(args) <= 1 {
		return advisor.Suggestions()
	}

	for len(args) > 1 {
		if nextAdvisor := advisor.ParseEntrance(args[0]); nextAdvisor != nil {
			return Complete(d, nextAdvisor, args[1:])
		} else {
			return []prompt.Suggest{}
		}
	}

	return []prompt.Suggest{}
}

func FindHandler(advisor Advisor, args []string) (handler func(argument string)) {
	if len(args) <= 1 {
		return advisor.Handler()
	}

	for len(args) > 1 {
		if nextAdvisor := advisor.ParseEntrance(args[0]); nextAdvisor != nil {
			return FindHandler(nextAdvisor, args[1:])
		} else {
			return advisor.Handler()
		}
	}

	return advisor.Handler()
}
