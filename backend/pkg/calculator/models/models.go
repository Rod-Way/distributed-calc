package models

type Node interface {
	Eval() float64
}

type Operand struct {
	Value float64
}

func (o Operand) Eval() float64 {
	return o.Value
}

type Operator struct {
	Symbol     string
	Precedence int
	Operation  func(left, right float64) float64
}

type BinOp struct {
	Left  Node
	Right Node
	Op    Operator
}

func (b BinOp) Eval() float64 {
	return b.Op.Operation(b.Left.Eval(), b.Right.Eval())
}
