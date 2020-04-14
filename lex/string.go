package lex

import "strconv"

func (token Token) String() string {
	switch token.Type {
	case LeftParen:
		return "("
	case RightParen:
		return ")"
	case Equals:
		return "="
	case Plus:
		return "+"
	case Minus:
		return "-"
	case Multiply:
		return "*"
	case Divide:
		return "/"
	case Nat:
		return strconv.Itoa(token.Nat)
	case Id:
		return token.Id
	case End:
		return "$"
	}
	panic("unrecognized token type")
}