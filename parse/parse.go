package parse

import (
	"errors"
	. "github.com/caziz/calc/ast"
	. "github.com/caziz/calc/lex"
)

func Parse(tokens []Token) (ENode, error) {
	tokensE, e, errE := parseE(tokens, nil)
	if errE != nil {
		return ENode{}, errE
	}
	if tokensE[0].Type != End {
		return ENode{}, errors.New("failed to parse all tokens")
	}
	return e, nil
}

func parseE(tokens []Token, err error) ([]Token, ENode, error) {
	if err != nil {
		return nil, ENode{}, err
	}
	switch tokens[0].Type {
	case Minus, LeftParen, Nat, Id:
		tokensT, t, errT := parseT(tokens, nil)
		tokensA, a, errA := parseA(tokensT, errT)
		return tokensA, ENode{&t, &a}, errA
	}
	return nil, ENode{}, errors.New("failed to parse E node")
}

func parseA(tokens []Token, err error) ([]Token, ANode, error) {
	if err != nil {
		return nil, ANode{}, err
	}
	switch tokens[0].Type {
	case End, RightParen:
		return tokens, ANode{AEpsilon, nil, nil}, nil
	case Plus, Minus:
		aType := AAddition
		if tokens[0].Type == Minus {
			aType = ASubtraction
		}
		tokensT, t, errT := parseT(tokens[1:], nil)
		tokensA, a, errA := parseA(tokensT, errT)
		return tokensA, ANode{aType, &t, &a}, errA
	}
	return nil, ANode{}, errors.New("failed to parse A node")
}

func parseT(tokens []Token, err error) ([]Token, TNode, error) {
	if err != nil {
		return nil, TNode{}, err
	}
	switch tokens[0].Type {
	case Minus, LeftParen, Nat, Id:
		tokensF, f, errF := parseF(tokens, nil)
		tokensB, b, errB := parseB(tokensF, errF)
		return tokensB, TNode{&f, &b}, errB
	}
	return nil, TNode{}, errors.New("failed to parse T node")
}

func parseB(tokens []Token, err error) ([]Token, BNode, error) {
	if err != nil {
		return nil, BNode{}, err
	}
	switch tokens[0].Type {
	case Plus, Minus, RightParen, End:
		return tokens, BNode{BEpsilon, nil, nil}, nil
	case Multiply, Divide:
		bType := BMultiplication
		if tokens[0].Type == Divide {
			bType = BDivision
		}
		tokensF, f, errF := parseF(tokens[1:], nil)
		tokensB, b, errB := parseB(tokensF, errF)
		return tokensB, BNode{bType, &f, &b}, errB
	}
	return nil, BNode{}, errors.New("failed to parse B node")
}

func parseF(tokens []Token, err error) ([]Token, FNode, error) {
	if err != nil {
		return nil, FNode{}, nil
	}
	switch tokens[0].Type {
	case LeftParen:
		tokensE, e, errE := parseE(tokens[1:], nil)
		if errE != nil {
			return nil, FNode{}, errE
		}
		if len(tokensE) == 0 || tokensE[0].Type != RightParen {
			return nil, FNode{}, errors.New("failed to parse F node")
		}
		return tokensE[1:], FNode{FGroup, &e, nil, nil}, nil
	case Id:
		return tokens[1:], FNode{FId, nil, &(tokens[0].Id), nil}, nil
	case Nat:
		return tokens[1:], FNode{FNat, nil, nil, &(tokens[0].Nat)}, nil
	case Minus:
		if len(tokens) == 1 {
			return nil, FNode{}, errors.New("failed to parse F node")
		}
		if tokens[1].Type == Id {
			return tokens[2:], FNode{FNegId, nil, &(tokens[1].Id), nil}, nil
		}
		if tokens[1].Type == Nat {
			return tokens[2:], FNode{FNegNat, nil, nil, &(tokens[1].Nat)}, nil
		}
		return nil, FNode{}, errors.New("failed to parse F node")
	}
	return nil, FNode{}, errors.New("failed to parse F node")
}