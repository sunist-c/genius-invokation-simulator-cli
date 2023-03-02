package advisor

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/sunist-c/genius-invokation-simulator-cli/localization"
)

type Advisor interface {
	Entrance() string
	NodeDepth() int
	SyncDocument(document *prompt.Document)
	Suggestions() []prompt.Suggest
	ParseEntrance(nextEntrance string) Advisor
	PluginChildWithOpts(childOptions ...Options) (child Advisor)
	Handler() func(argument string)
}

type implement struct {
	suggesters []SuggesterFunc
	entrance   string
	depth      int
	nodes      map[string]Advisor
	handler    func(argument string)
	document   *prompt.Document
}

func (impl *implement) Entrance() string {
	return impl.entrance
}

func (impl *implement) NodeDepth() int {
	return impl.depth
}

func (impl *implement) SyncDocument(document *prompt.Document) {
	impl.document = document
}

func (impl *implement) Suggestions() []prompt.Suggest {
	if impl.suggesters == nil {
		impl.suggesters = []SuggesterFunc{}
	}

	ctx := NewContext(impl.document, impl.suggesters...)

	ctx.Next()

	return ctx.Result()
}

func (impl *implement) ParseEntrance(nextEntrance string) Advisor {
	if impl.nodes == nil {
		impl.nodes = map[string]Advisor{}
	}

	return impl.nodes[nextEntrance]
}

func (impl *implement) PluginChildWithOpts(childOptions ...Options) (child Advisor) {
	if impl.nodes == nil {
		impl.nodes = map[string]Advisor{}
	}

	childEntity := &implement{
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

func (impl *implement) Handler() func(argument string) {
	if impl.handler == nil {
		impl.handler = func(argument string) {}
	}

	return impl.handler
}

type Options func(option *implement)

func WithAdvisorDepth(depth int) Options {
	return func(option *implement) {
		option.depth = depth
	}
}

func WithAdvisorEntrance(entrance string) Options {
	return func(option *implement) {
		option.entrance = entrance
	}
}

func WithAdvisorBuiltinSuggestions(keywords ...string) Options {
	return func(option *implement) {
		if option.suggesters == nil {
			option.suggesters = []SuggesterFunc{}
		}

		option.suggesters = append(option.suggesters, func(ctx *SuggesterContext) {
			ctx.Next()
			appended := make([]prompt.Suggest, len(keywords))
			for i := range keywords {
				appended[i] = prompt.Suggest{
					Text:        keywords[i],
					Description: localization.LanguagePack.GetTranslation(localization.GetLocalLanguage(), fmt.Sprintf("%v_desc", keywords[i])),
				}
			}

			ctx.AppendSuggestAfter(prompt.FilterHasPrefix(appended, ctx.Document.GetWordBeforeCursor(), true)...)
		})
	}
}

func WithAdvisorSuggesterFunctions(functions ...SuggesterFunc) Options {
	return func(option *implement) {
		option.suggesters = functions
	}
}

func WithAdvisorSuggesterAppend(functions ...SuggesterFunc) Options {
	return func(option *implement) {
		option.suggesters = append(option.suggesters, functions...)
	}
}

func WithAdvisorFunctionChain(functions ...func(argument string)) Options {
	return func(option *implement) {
		option.handler = func(argument string) {
			for _, function := range functions {
				function(argument)
			}
		}
	}
}

func WithAdvisorChildren(nodes ...Advisor) Options {
	return func(option *implement) {
		for _, node := range nodes {
			option.nodes[node.Entrance()] = node
		}
	}
}

func NewAdvisorWithOpts(options ...Options) Advisor {
	impl := &implement{
		nodes: map[string]Advisor{},
	}

	for _, option := range options {
		option(impl)
	}

	return impl
}
