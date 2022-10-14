package ast

type Node interface {}

type Expression interface {
  Acceptable
  Node
}

type Statement interface {
  Acceptable
  Node
}

type Declaration interface {
  Statement
}

// type Program struct {
//   Node
// }

// type IfStatement struct {
//   Condition Expression
//   Then      Statement
//   Else      Statement
//   Statement
// }

// type WhileStatement struct {
//   Condition Expression
//   Body      Statement
//   Statement
// }

// type BlockStatement struct {
//   Body      []Statement
//   Statement
// }

// type ExpressionStatement struct {
//   Expression Expression
//   Statement
// }

// type VariableDeclaration struct {
//   Name        token.Token
//   Initializer Expression
//   Declaration
// }

// type FunctionDeclaration struct {
//   Name   token.Token
//   Params []token.Token
//   Body   []Statement
//   Declaration
// }

// type ReturnStatement struct {
//   Keyword   token.Token
//   Value     Expression
//   Statement
// }

// type EchoStatement struct {
//   Expression Expression
//   Statement
// }

// type ClassMethod struct {
//   Kind        string // "constructor" | "method"
//   Declaration FunctionDeclaration
// }

// type ClassDeclaration struct {
//   Name        token.Token
//   Methods     []ClassMethod
//   SuperClass  *VariableExpression
//   Declaration
// }

// type LiteralExpression struct {
//   Value      interface{}
//   Expression
// }

// type BinaryExpression struct {
//   Left       Expression
//   Right      Expression
//   Operator   token.Token
//   Expression
// }

// type UnaryExpression struct {
//   Right      Expression
//   Operator   token.Token
//   Expression
// }

// type GroupingExpression struct {
//   Expression Expression
//   Acceptable
//   Node
// }

// type LogicalExpression struct {
//   Left       Expression
//   Right      Expression
//   Operator   token.Token
//   Expression
// }

// type VariableExpression struct {
//   Name       token.Token
//   Expression
// }

// type AssignExpression struct {
//   Name       token.Token
//   Value      Expression
//   Expression
// }

// type CallExpression struct {
//   Paren      token.Token
//   Callee     Expression
//   Arguments  []Expression
//   Expression
// }

// type SuperExpression struct {
//   Method     token.Token
//   Keyword    token.Token
//   Expression
// }

// type SelfExpression struct {
//   Keyword    token.Token
//   Expression
// }

// type GetExpression struct {
//   Name       token.Token
//   Object     Expression
//   Expression
// }

// type SetExpression struct {
//   Name       token.Token
//   Value      Expression
//   Object     Expression
//   Expression
// }
