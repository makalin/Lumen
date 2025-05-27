package compiler

import (
	"fmt"
	"strings"
)

type NodeType int

const (
	NodeProgram NodeType = iota
	NodeLetStatement
	NodeFunctionLiteral
	NodeCallExpression
	NodeIdentifier
	NodeIntegerLiteral
	NodeFloatLiteral
	NodeStringLiteral
	NodeBooleanLiteral
	NodePrefixExpression
	NodeInfixExpression
	NodeIfExpression
	NodeBlockStatement
	NodeReturnStatement
	NodeExpressionStatement
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

func (p *Program) String() string {
	var out string
	for _, s := range p.Statements {
		out += s.String()
	}
	return out
}

type LetStatement struct {
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return "let" }
func (ls *LetStatement) String() string {
	var out string
	out += "let " + ls.Name.String() + " = "
	if ls.Value != nil {
		out += ls.Value.String()
	}
	out += ";"
	return out
}

type Identifier struct {
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Value }
func (i *Identifier) String() string       { return i.Value }

type IntegerLiteral struct {
	Value int64
}

func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return fmt.Sprintf("%d", il.Value) }
func (il *IntegerLiteral) String() string       { return il.TokenLiteral() }

type FloatLiteral struct {
	Value float64
}

func (fl *FloatLiteral) expressionNode()      {}
func (fl *FloatLiteral) TokenLiteral() string { return fmt.Sprintf("%f", fl.Value) }
func (fl *FloatLiteral) String() string       { return fl.TokenLiteral() }

type StringLiteral struct {
	Value string
}

func (sl *StringLiteral) expressionNode()      {}
func (sl *StringLiteral) TokenLiteral() string { return sl.Value }
func (sl *StringLiteral) String() string       { return "\"" + sl.Value + "\"" }

type BooleanLiteral struct {
	Value bool
}

func (bl *BooleanLiteral) expressionNode()      {}
func (bl *BooleanLiteral) TokenLiteral() string { return fmt.Sprintf("%t", bl.Value) }
func (bl *BooleanLiteral) String() string       { return bl.TokenLiteral() }

type PrefixExpression struct {
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode()      {}
func (pe *PrefixExpression) TokenLiteral() string { return pe.Operator }
func (pe *PrefixExpression) String() string {
	var out string
	out += "("
	out += pe.Operator
	out += pe.Right.String()
	out += ")"
	return out
}

type InfixExpression struct {
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *InfixExpression) expressionNode()      {}
func (ie *InfixExpression) TokenLiteral() string { return ie.Operator }
func (ie *InfixExpression) String() string {
	var out string
	out += "("
	out += ie.Left.String()
	out += " " + ie.Operator + " "
	out += ie.Right.String()
	out += ")"
	return out
}

type IfExpression struct {
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ie *IfExpression) expressionNode()      {}
func (ie *IfExpression) TokenLiteral() string { return "if" }
func (ie *IfExpression) String() string {
	var out string
	out += "if"
	out += ie.Condition.String()
	out += " "
	out += ie.Consequence.String()
	if ie.Alternative != nil {
		out += "else "
		out += ie.Alternative.String()
	}
	return out
}

type BlockStatement struct {
	Statements []Statement
}

func (bs *BlockStatement) statementNode()       {}
func (bs *BlockStatement) TokenLiteral() string { return "{" }
func (bs *BlockStatement) String() string {
	var out string
	for _, s := range bs.Statements {
		out += s.String()
	}
	return out
}

type FunctionLiteral struct {
	Parameters []*Identifier
	Body       *BlockStatement
}

func (fl *FunctionLiteral) expressionNode()      {}
func (fl *FunctionLiteral) TokenLiteral() string { return "fn" }
func (fl *FunctionLiteral) String() string {
	var out string
	params := []string{}
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}
	out += "fn"
	out += "("
	out += strings.Join(params, ", ")
	out += ") "
	out += fl.Body.String()
	return out
}

type CallExpression struct {
	Function  Expression
	Arguments []Expression
}

func (ce *CallExpression) expressionNode()      {}
func (ce *CallExpression) TokenLiteral() string { return "(" }
func (ce *CallExpression) String() string {
	var out string
	args := []string{}
	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}
	out += ce.Function.String()
	out += "("
	out += strings.Join(args, ", ")
	out += ")"
	return out
}

type ReturnStatement struct {
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return "return" }
func (rs *ReturnStatement) String() string {
	var out string
	out += "return "
	if rs.ReturnValue != nil {
		out += rs.ReturnValue.String()
	}
	out += ";"
	return out
}

type ExpressionStatement struct {
	Expression Expression
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Expression.TokenLiteral() }
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}
