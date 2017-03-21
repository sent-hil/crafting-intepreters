package scanner

import "github.com/sent-hil/crafting-interpreters/token"

type Scanner struct {
	Source string
	tokens []*token.Token
}

func NewScanner(source string) *Scanner {
	return &Scanner{
		Source: source,
		tokens: make([]*token.Token, 0),
	}
}
