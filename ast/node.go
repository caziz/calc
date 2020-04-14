package ast

// E := T A

type ENode struct {
	T *TNode
	A *ANode
}

// A := + T A | - T A | epsilon

type AType int
const (
	AAddition AType = iota
	ASubtraction
	AEpsilon
)

type ANode struct {
	Type AType
	T *TNode
	A *ANode
}

// T := F B

type TNode struct {
	F *FNode
	B *BNode
}

// B := * F B | / F B | epsilon

type BType int
const (
	BMultiplication BType = iota
	BDivision
	BEpsilon
)

type BNode struct {
	Type BType
	F *FNode
	B *BNode
}

// F := ( E ) | Id | Nat | -Id | -Nat

type FType int
const (
	FGroup FType = iota
	FId
	FNat
	FNegId
	FNegNat
)

type FNode struct {
	Type FType
	E    *ENode
	Id   *string
	Nat  *int
}