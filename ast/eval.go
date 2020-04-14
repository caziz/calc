package ast

func (eNode ENode) Eval(env map[string]int) int {
	return eNode.A.evalRhs(eNode.T.eval(env), env)
}

func (aNode ANode) evalRhs(lhs int, env map[string]int) int {
	switch aNode.Type {
	case AAddition:
		return lhs + aNode.A.evalRhs(aNode.T.eval(env), env)
	case ASubtraction:
		return lhs - aNode.A.evalRhs(aNode.T.eval(env), env)
	case AEpsilon:
		return lhs
	}
	panic("unrecognized A Node type")
}

func (tNode TNode) eval(env map[string]int) int {
	return tNode.B.evalRhs(tNode.F.eval(env), env)
}

func (bNode BNode) evalRhs(lhs int, env map[string]int) int {
	switch bNode.Type {
	case BMultiplication:
		return lhs * bNode.B.evalRhs(bNode.F.eval(env), env)
	case BDivision:
		return lhs / bNode.B.evalRhs(bNode.F.eval(env), env)
	case BEpsilon:
		return lhs
	}
	panic("unrecognized B Node type")
}

func (fNode FNode) eval(env map[string]int) int {
	switch fNode.Type {
	case FGroup:
		return fNode.E.Eval(env)
	case FNegNat:
		return -(*fNode.Nat)
	case FNegId:
		return -0
	case FNat:
		return *fNode.Nat
	case FId:
		return env[*fNode.Id]
	}
	panic("unrecognized F Node type")
}
