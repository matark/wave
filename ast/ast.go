package ast

import "neweos.de/sube/token"

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

type Grouping struct {
  Expr Expression
  Expression
}

type Binary struct {
  Operator token.Token
  Right    Expression
  Left     Expression
  Expression
}

type Unary struct {
  Operator token.Token
  Right    Expression
  Expression
}

// type Identifier struct {
//   Expression
// }

// type Function struct {
//   Params []token.Token
//   Body   BlockStatement
// }

// type BlockStatement struct {
//   Statement
//   Body []Statement
// }

// type FunctionDeclaration struct {
//   Declaration
//   *Function
//   Identifier token.Token
// }

// type FunctionExpression struct {
//   Expression
//   *Function
// }

// type ExpressionStatement struct {
//   Statement
//   Body Expression
// }

// type ReturnStatement struct {
//   Statement
//   Argument Expression
//   // Keyword   token.Token
//   // Value     Expression
// }

// type BreakStatement struct {
//   Statement
//   Label Expression
// }

// type ContinueStatement struct {
//   Statement
//   Label Expression
// }

// type IfStatement struct {
//   Statement
//   Condition Expression
//   Then Statement
//   Else Statement
// }

// type ForStatement struct {
// }

// type LoopStatement struct {
//   Statement
//   Condition Expression
//   Body Statement
// }

// type Program struct {
//   Node
// }

// type VariableDeclaration struct {
//   Name        token.Token
//   Initializer Expression
//   Declaration
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

// type Literal struct {
//   Type  token.Token
//   Value string
//   Expression
// }

// type CallExpression struct {
//   Paren      token.Token
//   Callee     Expression
//   Arguments  []Expression
//   Expression
// }
