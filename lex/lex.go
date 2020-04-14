package lex

import (
	"errors"
	"unicode"
	"unicode/utf8"
)

func Lex(lexeme string) ([]Token, error) {
	var tokens []Token
	for len(lexeme) > 0 {
		r, w := utf8.DecodeRuneInString(lexeme)
		if r == utf8.RuneError && w == 1 {
			return nil, errors.New("invalid encoding")
		}
		lexeme = lexeme[w:]
		switch r {
		case '(':
			tokens = append(tokens, Token{Type: LeftParen})
		case ')':
			tokens = append(tokens, Token{Type: RightParen})
		case '+':
			tokens = append(tokens, Token{Type: Plus})
		case '-':
			tokens = append(tokens, Token{Type: Minus})
		case '*':
			tokens = append(tokens, Token{Type: Multiply})
		case '/':
			tokens = append(tokens, Token{Type: Divide})
		case '=':
			tokens = append(tokens, Token{Type: Equals})
		case ' ', '\n', '\t':
			break
		default:
			if unicode.IsNumber(r) {
				num := 0
				for {
					num *= 10
					num += int(r) - int('0')
					nextR, nextW := utf8.DecodeRuneInString(lexeme)
					if unicode.IsLetter(nextR) {
						return nil, errors.New("identifier cannot start with a number")
					}
					if !unicode.IsNumber(nextR) {
						break
					}
					lexeme = lexeme[nextW:]
					r, w = nextR, nextW
				}
				tokens = append(tokens, Token{Type: Nat, Nat: num})
			} else if unicode.IsLetter(r) {
				id := ""
				for {
					id += string(r)
					nextR, nextW := utf8.DecodeRuneInString(lexeme)
					if !unicode.IsLetter(nextR) && !unicode.IsNumber(nextR) {
						break
					}
					lexeme = lexeme[nextW:]
					r, w = nextR, nextW
				}
				tokens = append(tokens, Token{Type: Id,  Id: id})
			} else {
				return nil, errors.New("unknown character: \"" + string(r) + "\"")
			}
		}
	}
	tokens = append(tokens, Token{Type:End})
	return tokens, nil
}

