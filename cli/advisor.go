package cli

import (
	"fmt"
	"github.com/c-bata/go-prompt"
)

type Advisor interface {
	Entrance() string
	NodeDepth() int
	SyncDocument(document *prompt.Document)
	Suggestions() []prompt.Suggest
	ParseEntrance(nextEntrance string) Advisor
	PluginChildWithOpts(childOptions ...AdvisorOptions) (child Advisor)
	Handler() func(argument string)
}

type AdvisorImpl struct {
	suggestionFunction func(d *prompt.Document) []prompt.Suggest
	entrance           string
	depth              int
	nodes              map[string]Advisor
	handler            func(argument string)
	document           *prompt.Document
}

func (impl *AdvisorImpl) Entrance() string {
	return impl.entrance
}

func (impl *AdvisorImpl) NodeDepth() int {
	return impl.depth
}

func (impl *AdvisorImpl) SyncDocument(document *prompt.Document) {
	impl.document = document
}

func (impl *AdvisorImpl) Suggestions() []prompt.Suggest {
	if impl.suggestionFunction == nil {
		impl.suggestionFunction = func(d *prompt.Document) []prompt.Suggest {
			return []prompt.Suggest{}
		}
	}

	return impl.suggestionFunction(impl.document)
}

func (impl *AdvisorImpl) ParseEntrance(nextEntrance string) Advisor {
	if impl.nodes == nil {
		impl.nodes = map[string]Advisor{}
	}

	return impl.nodes[nextEntrance]
}

func (impl *AdvisorImpl) PluginChildWithOpts(childOptions ...AdvisorOptions) (child Advisor) {
	if impl.nodes == nil {
		impl.nodes = map[string]Advisor{}
	}

	childEntity := &AdvisorImpl{
		entrance: "",
		depth:    impl.depth + 1,
		nodes:    map[string]Advisor{},
	}
	for _, option := range childOptions {
		option(childEntity)
	}

	impl.nodes[childEntity.Entrance()] = childEntity
	return childEntity
}

func (impl *AdvisorImpl) Handler() func(argument string) {
	if impl.handler == nil {
		impl.handler = func(argument string) {}
	}

	return impl.handler
}

type AdvisorOptions func(option *AdvisorImpl)

func WithAdvisorDepth(depth int) AdvisorOptions {
	return func(option *AdvisorImpl) {
		option.depth = depth
	}
}

func WithAdvisorEntrance(entrance string) AdvisorOptions {
	return func(option *AdvisorImpl) {
		option.entrance = entrance
	}
}

func WithAdvisorBuiltinSuggestions(keywords ...string) AdvisorOptions {
	return func(option *AdvisorImpl) {
		option.suggestionFunction = func(d *prompt.Document) []prompt.Suggest {
			result := make([]prompt.Suggest, len(keywords))
			for i, keyword := range keywords {
				result[i] = prompt.Suggest{
					Text:        keyword,
					Description: languagePack.GetTranslation(LocalLanguage(), fmt.Sprintf("%v_desc", keyword)),
				}
			}

			return prompt.FilterHasPrefix(result, d.GetWordBeforeCursor(), true)
		}
	}
}

func WithAdvisorSuggestionFunction(function func(d *prompt.Document) []prompt.Suggest) AdvisorOptions {
	return func(option *AdvisorImpl) {
		option.suggestionFunction = function
	}
}

func WithAdvisorFunctionChain(functions ...func(argument string)) AdvisorOptions {
	return func(option *AdvisorImpl) {
		option.handler = func(argument string) {
			for _, function := range functions {
				function(argument)
			}
		}
	}
}

func WithAdvisorChildren(nodes ...Advisor) AdvisorOptions {
	return func(option *AdvisorImpl) {
		for _, node := range nodes {
			option.nodes[node.Entrance()] = node
		}
	}
}

func NewAdvisorWithOpts(options ...AdvisorOptions) Advisor {
	impl := &AdvisorImpl{
		nodes: map[string]Advisor{},
	}

	for _, option := range options {
		option(impl)
	}

	return impl
}
