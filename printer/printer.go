package printer
// import "fmt"

// type Printer struct {
//   ast.Visitor
// }

// var printer = &Printer{}

// func Print(statements []ast.Statement) {
//   for _, statement := range statements {
//     printer.execute(statement)
//   }
// }

// func (p *Printer) VisitLiteralExpression(node ast.LiteralExpression) interface{} {
//   return node.Value
// }

// func (p *Printer) VisitBinaryExpression(node ast.BinaryExpression) interface{} {
//   p.evaluate(node.Left)
//   p.execute(node.Right)

//   return nil
// }

// func (p *Printer) VisitUnaryExpression(node ast.UnaryExpression) interface{} {
//   p.evaluate(node.Right)
//   return nil
// }

// func (p *Printer) VisitGroupingExpression(node ast.GroupingExpression) interface{} {
//   p.evaluate(node.Expression)
//   return nil
// }

// func (p *Printer) VisitLogicalExpression(node ast.LogicalExpression) interface{} {
//   p.evaluate(node.Left)
//   p.evaluate(node.Right)

//   return nil
// }

// func (p *Printer) VisitVariableExpression(node ast.VariableExpression) interface{} {
//   p.evaluate(node.Expression)
//   return nil
// }

// func (p *Printer) VisitAssignExpression(node ast.AssignExpression) interface{} {
//   p.evaluate(node.Value)
//   return nil
// }

// func (p *Printer) VisitCallExpression(node ast.CallExpression) interface{} {
//   p.evaluate(node.Callee)
//   // p.evaluate(node.Arguments)
//   return nil
// }

// func (p *Printer) VisitSuperExpression(node ast.SuperExpression) interface{} {
//   // p.execute(node.Method)
//   return nil
// }

// func (p *Printer) VisitSelfExpression(node ast.SelfExpression) interface{} {
//   return nil
// }

// func (p *Printer) VisitGetExpression(node ast.GetExpression) interface{} {
//   p.evaluate(node.Object)
//   return nil
// }

// func (p *Printer) VisitSetExpression(node ast.SetExpression) interface{} {
//   p.evaluate(node.Object)
//   p.evaluate(node.Value)

//   return nil
// }

// func (p *Printer) VisitIfStatement(node ast.IfStatement) interface{} {
//   if node.Else == nil {
//     p.execute(node.Condition)
//     p.execute(node.Then)
//     return nil
//   }

//   p.execute(node.Condition)
//   p.execute(node.Then)

//   return nil
// }

// func (p *Printer) VisitWhileStatement(node ast.WhileStatement) interface{} {
//   p.execute(node.Condition)
//   p.execute(node.Body)

//   return nil
// }

// func (p *Printer) VisitBlockStatement(node ast.BlockStatement) interface{} {
//   for _, statement := range node.Body {
//     p.execute(statement)
//   }

//   return nil
// }

// func (p *Printer) VisitExpressionStatement(node ast.ExpressionStatement) interface{} {
//   p.execute(node.Expression)
//   return nil
// }

// func (p *Printer) VisitVariableDeclaration(node ast.VariableDeclaration) interface{} {
//   if node.Initializer == nil {
//     return node.Name
//   }

//   p.execute(node.Initializer)
//   return node.Name
// }

// func (p *Printer) VisitFunctionDeclaration(node ast.FunctionDeclaration) interface{} {
//   for _, param := range node.Params {
//     fmt.Println(param.Type)
//   }

//   for _, body := range node.Body {
//     p.execute(body)
//   }

//   return nil
// }

// func (p *Printer) VisitReturnStatement(node ast.ReturnStatement) interface{} {
//   if node.Value == nil { return nil }
//   return node.Value
// }

// func (p *Printer) VisitEchoStatement(node ast.EchoStatement) interface{} {
//   p.evaluate(node.Expression)
//   return nil
// }

// func (p *Printer) VisitClassDeclaration(node ast.ClassDeclaration) interface{} {
//   return nil
// }

// func (p *Printer) evaluate(expression ast.Expression) interface{} {
//   return expression.Accept(p)
// }

// func (p *Printer) execute(statement ast.Statement) interface{} {
//   return statement.Accept(p)
// }
