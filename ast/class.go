package ast

type ClassDeclaration struct {
  //   Name        token.Token
  //   Methods     []ClassMethod
  //   SuperClass  *VariableExpression
  //   Declaration
  // }
}

type Super struct {
  Method string
  Expression
}

type Self struct {
  Expression
}

type Get struct {
  Property string
  Object   Expression
  Expression
}

type Set struct {
  Property string
  Object   Expression
  Value    Expression
  Expression
}
//   Keyword    token.Token
// type ClassMethod struct {
//   Kind        string // "constructor" | "method"
//   Declaration FunctionDeclaration
// }
