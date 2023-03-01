package advisor

import (
	"github.com/c-bata/go-prompt"
)

type SuggesterContext struct {
	index    int
	handlers []SuggesterFunc
	Document *prompt.Document
	Result   []prompt.Suggest
}

func (s *SuggesterContext) Next() {
	s.index += 1
	for s.index > 0 && s.index <= len(s.handlers) {
		s.handlers[s.index](s)
		s.index += 1
	}
}

func (s *SuggesterContext) Abort() {
	s.index = -1
}

func (s *SuggesterContext) IsAborted() bool {
	return s.index == -1
}

func (s *SuggesterContext) ToCompleter() func(d prompt.Document) []prompt.Suggest {
	s.Next()
	return func(d prompt.Document) []prompt.Suggest {
		return s.Result
	}
}

func NewContext(d *prompt.Document, handlers ...SuggesterFunc) *SuggesterContext {
	return &SuggesterContext{
		index:    0,
		handlers: handlers,
		Document: d,
		Result:   []prompt.Suggest{},
	}
}

type SuggesterFunc func(ctx *SuggesterContext)
