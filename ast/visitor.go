package ast

type Acceptable interface {
  Accept(v Visitor) any
}

type Visitor interface {

}

// type Visitor interface {
//   VisitLiteralExpression(node LiteralExpression) interface{}
//   VisitBinaryExpression(node BinaryExpression) interface{}
//   VisitUnaryExpression(node UnaryExpression) interface{}
//   VisitGroupingExpression(node GroupingExpression) interface{}
//   VisitLogicalExpression(node LogicalExpression) interface{}
//   VisitVariableExpression(node VariableExpression) interface{}
//   VisitAssignExpression(node AssignExpression) interface{}
//   VisitCallExpression(node CallExpression) interface{}
//   VisitSuperExpression(node SuperExpression) interface{}
//   VisitSelfExpression(node SelfExpression) interface{}
//   VisitGetExpression(node GetExpression) interface{}
//   VisitSetExpression(node SetExpression) interface{}
//   VisitIfStatement(node IfStatement) interface{}
//   VisitWhileStatement(node WhileStatement) interface{}
//   VisitBlockStatement(node BlockStatement) interface{}
//   VisitExpressionStatement(node ExpressionStatement) interface{}
//   VisitVariableDeclaration(node VariableDeclaration) interface{}
//   VisitFunctionDeclaration(node FunctionDeclaration) interface{}
//   VisitReturnStatement(node ReturnStatement) interface{}
//   VisitEchoStatement(node EchoStatement) interface{}
//   VisitClassDeclaration(node ClassDeclaration) interface{}
// }


// func (node LiteralExpression) Accept(v Visitor) interface{} {
//   return v.VisitLiteralExpression(node)
// }

// func (node BinaryExpression) Accept(v Visitor) interface{} {
//   return v.VisitBinaryExpression(node)
// }

// func (node UnaryExpression) Accept(v Visitor) interface{} {
//   return v.VisitUnaryExpression(node)
// }

// func (node GroupingExpression) Accept(v Visitor) interface{} {
//   return v.VisitGroupingExpression(node)
// }

// func (node LogicalExpression) Accept(v Visitor) interface{} {
//   return v.VisitLogicalExpression(node)
// }

// func (node VariableExpression) Accept(v Visitor) interface{} {
//   return v.VisitVariableExpression(node)
// }

// func (node AssignExpression) Accept(v Visitor) interface{} {
//   return v.VisitAssignExpression(node)
// }

// func (node CallExpression) Accept(v Visitor) interface{} {
//   return v.VisitCallExpression(node)
// }

// func (node SuperExpression) Accept(v Visitor) interface{} {
//   return v.VisitSuperExpression(node)
// }

// func (node SelfExpression) Accept(v Visitor) interface{} {
//   return v.VisitSelfExpression(node)
// }

// func (node GetExpression) Accept(v Visitor) interface{} {
//   return v.VisitGetExpression(node)
// }

// func (node SetExpression) Accept(v Visitor) interface{} {
//   return v.VisitSetExpression(node)
// }


// func (node IfStatement) Accept(v Visitor) interface{} {
//   return v.VisitIfStatement(node)
// }

// func (node WhileStatement) Accept(v Visitor) interface{} {
//   return v.VisitWhileStatement(node)
// }

// func (node BlockStatement) Accept(v Visitor) interface{} {
//   return v.VisitBlockStatement(node)
// }

// func (node ExpressionStatement) Accept(v Visitor) interface{} {
//   return v.VisitExpressionStatement(node)
// }

// func (node VariableDeclaration) Accept(v Visitor) interface{} {
//   return v.VisitVariableDeclaration(node)
// }

// func (node FunctionDeclaration) Accept(v Visitor) interface{} {
//   return v.VisitFunctionDeclaration(node)
// }

// func (node ReturnStatement) Accept(v Visitor) interface{} {
//   return v.VisitReturnStatement(node)
// }

// func (node EchoStatement) Accept(v Visitor) interface{} {
//   return v.VisitEchoStatement(node)
// }

// func (node ClassDeclaration) Accept(v Visitor) interface{} {
//   return v.VisitClassDeclaration(node)
// }
