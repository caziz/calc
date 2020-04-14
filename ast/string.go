package ast

import (
	"strconv"
)

func (eNode ENode) String() string {
	return eNode.T.String() + eNode.A.String()
}

func (aNode ANode) String() string {
	switch aNode.Type {
	case AAddition:
		return " + " + aNode.T.String() + aNode.A.String()
	case ASubtraction:
		return " - " + aNode.T.String() + aNode.A.String()
	case AEpsilon:
		return ""
	}
	panic("unrecognized A Node type")
}

func (bNode BNode) String() string {
	switch bNode.Type {
	case BMultiplication:
		return " * " + bNode.F.String() + bNode.B.String()
	case BDivision:
		return " / " + bNode.F.String() + bNode.B.String()
	case BEpsilon:
		return ""
	}
	panic("unrecognized A Node type")
}

func (tNode TNode) String() string {
	return tNode.F.String() + tNode.B.String()
}

func (fNode FNode) String() string {
	switch fNode.Type {
	case FGroup:
		return "(" + fNode.E.String() + ")"
	case FId:
		return *fNode.Id
	case FNegId:
		return "-" + *fNode.Id
	case FNat:
		return strconv.Itoa(*fNode.Nat)
	case FNegNat:
		return strconv.Itoa(-(*fNode.Nat))
	}
	panic ("unrecognized F node type")
}