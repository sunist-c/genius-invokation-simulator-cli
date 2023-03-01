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

type advisorImpl struct {
	suggesters []SuggesterFunc
	entrance   string
	depth      int
	nodes      map[string]Advisor
	handler    func(argument string)
	document   *prompt.Document
}

func (impl *advisorImpl) Entrance() string {
	return impl.entrance
}

func (impl *advisorImpl) NodeDepth() int {
	return impl.depth
}

func (impl *advisorImpl) SyncDocument(document *prompt.Document) {
	impl.document = document
}

func (impl *advisorImpl) Suggestions() []prompt.Suggest {
	if impl.suggesters == nil {
		impl.suggesters = []SuggesterFunc{}
	}

	ctx := &SuggesterContext{
		index:    0,
		handlers: impl.suggesters,
		Document: impl.document,
		Result:   []prompt.Suggest{},
	}

	ctx.Next()

	return ctx.Result
}

func (impl *advisorImpl) ParseEntrance(nextEntrance string) Advisor {
	if impl.nodes == nil {
		impl.nodes = map[string]Advisor{}
	}

	return impl.nodes[nextEntrance]
}

func (impl *advisorImpl) PluginChildWithOpts(childOptions ...Options) (child Advisor) {
	if impl.nodes == nil {
		impl.nodes = map[string]Advisor{}
	}

	childEntity := &advisorImpl{
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

func (impl *advisorImpl) Handler() func(argument string) {
	if impl.handler == nil {
		impl.handler = func(argument string) {}
	}

	return impl.handler
}

type Options func(option *advisorImpl)

func WithAdvisorDepth(depth int) Options {
	return func(option *advisorImpl) {
		option.depth = depth
	}
}

func WithAdvisorEntrance(entrance string) Options {
	return func(option *advisorImpl) {
		option.entrance = entrance
	}
}

func WithAdvisorBuiltinSuggestions(keywords ...string) Options {
	return func(option *advisorImpl) {
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

			ctx.Result = append(appended, prompt.FilterHasPrefix(appended, ctx.Document.GetWordBeforeCursor(), true)...)
		})
	}
}

func WithAdvisorSuggesterFunctions(functions ...SuggesterFunc) Options {
	return func(option *advisorImpl) {
		option.suggesters = functions
	}
}

func WithAdvisorSuggesterAppend(functions ...SuggesterFunc) Options {
	return func(option *advisorImpl) {
		option.suggesters = append(option.suggesters, functions...)
	}
}

func WithAdvisorFunctionChain(functions ...func(argument string)) Options {
	return func(option *advisorImpl) {
		option.handler = func(argument string) {
			for _, function := range functions {
				function(argument)
			}
		}
	}
}

func WithAdvisorChildren(nodes ...Advisor) Options {
	return func(option *advisorImpl) {
		for _, node := range nodes {
			option.nodes[node.Entrance()] = node
		}
	}
}

func NewAdvisorWithOpts(options ...Options) Advisor {
	impl := &advisorImpl{
		nodes: map[string]Advisor{},
	}

	for _, option := range options {
		option(impl)
	}

	return impl
}
