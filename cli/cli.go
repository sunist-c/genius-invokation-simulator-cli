package cli

import (
	"github.com/c-bata/go-prompt"
	"github.com/sunist-c/genius-invokation-simulator-cli/advisor"
	"strings"
)

// Completer go-prompt的提示接口
func Completer(d prompt.Document) []prompt.Suggest {
	return Complete(d, rootAdvisor, parseArgs(d.TextBeforeCursor()))
}

// parseArgs 解析当前命令行参数，分割为若干个字符串
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

// Complete 递归解析Advisor的语法树结构，进行提示，O(n)
func Complete(d prompt.Document, adv advisor.Advisor, args []string) (suggestions []prompt.Suggest) {
	adv.SyncDocument(&d)

	if len(args) <= 1 {
		return adv.Suggestions()
	}

	for len(args) > 1 {
		if nextAdvisor := adv.ParseEntrance(args[0]); nextAdvisor != nil {
			return Complete(d, nextAdvisor, args[1:])
		} else {
			invalidAdvisor := advisor.InvalidAdvisor()
			invalidAdvisor.SyncDocument(&d)
			return invalidAdvisor.Suggestions()
		}
	}

	invalidAdvisor := advisor.InvalidAdvisor()
	invalidAdvisor.SyncDocument(&d)
	return invalidAdvisor.Suggestions()
}

// FindHandler 递归解析Advisor的语法树结构，查找命令的执行函数，O(n)
func FindHandler(adv advisor.Advisor, args []string) (handler func(argument string)) {
	defer func() {
		if handler == nil {
			logger.Errorf("no handler detected")
			handler = func(argument string) {}
		}
	}()
	if len(args) <= 1 {
		return adv.Handler()
	}

	for len(args) > 1 {
		if nextAdvisor := adv.ParseEntrance(args[0]); nextAdvisor != nil {
			return FindHandler(nextAdvisor, args[1:])
		} else {
			return adv.Handler()
		}
	}

	return adv.Handler()
}
