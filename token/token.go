package token

import "fmt"

type Token struct {
	ID         TokenID
	Lexme      string
	Literal    interface{}
	LineNumber int
}

func NewToken(id TokenID, lexme string, literal interface{}, lineNumber int) *Token {
	return &Token{
		ID:         id,
		Lexme:      lexme,
		Literal:    literal,
		LineNumber: lineNumber,
	}
}

func (t *Token) String() string {
	return fmt.Sprintf("%s %s %s", t.ID, t.Lexme, t.Literal)
}
