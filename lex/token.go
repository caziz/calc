package lex

type TokenType int
const (
	LeftParen TokenType = iota
	RightParen
	Equals
	Plus
	Minus
	Multiply
	Divide
	Nat
	Id
	End
)

type Token struct {
	Type TokenType
	Nat  int
	Id   string
}