package advisor

import (
	"github.com/c-bata/go-prompt"
)

var abortIndex int = 1145141919810

type SuggesterContext struct {
	index    int
	handlers []SuggesterFunc
	Document *prompt.Document
	result   []prompt.Suggest
}

func (s *SuggesterContext) Next() {
	s.index++
	for s.index < len(s.handlers) {
		s.handlers[s.index](s)
		s.index++
	}
}

func (s *SuggesterContext) Abort() {
	s.index = abortIndex
}

func (s *SuggesterContext) IsAborted() bool {
	return s.index >= abortIndex
}

func (s *SuggesterContext) Result() []prompt.Suggest {
	if s.result == nil {
		s.result = []prompt.Suggest{}
	}

	return s.result
}

func (s *SuggesterContext) AppendSuggestBefore(appends ...prompt.Suggest) {
	//fmt.Println("before:", s.result)
	s.result = append(appends, s.result...)
	//fmt.Println("after:", s.result)
}

func (s *SuggesterContext) AppendSuggestAfter(appends ...prompt.Suggest) {
	//fmt.Println("before:", s.result)
	s.result = append(s.result, appends...)
	//fmt.Println("after:", s.result)
}

func (s *SuggesterContext) ToCompleter() func(d prompt.Document) []prompt.Suggest {
	s.Next()
	return func(d prompt.Document) []prompt.Suggest {
		return s.Result()
	}
}

func NewContext(d *prompt.Document, handlers ...SuggesterFunc) *SuggesterContext {
	return &SuggesterContext{
		index:    -1,
		handlers: handlers,
		Document: d,
		result:   []prompt.Suggest{},
	}
}

type SuggesterFunc func(ctx *SuggesterContext)
