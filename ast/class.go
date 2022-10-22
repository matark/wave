package ast

import "neweos.de/sube/token"

type ClassDeclaration struct {
  //   Name        token.Token
  //   Methods     []ClassMethod
  //   SuperClass  *VariableExpression
  //   Declaration
  // }
}

type Super struct {
  Method token.Token
  Expression
}

type Self struct {
  Expression
}
//   Keyword    token.Token
// type ClassMethod struct {
//   Kind        string // "constructor" | "method"
//   Declaration FunctionDeclaration
// }
