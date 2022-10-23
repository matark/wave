package parser

import "fmt"
import "strconv"
import "neweos.de/sube/ast"
import "neweos.de/sube/lexer"
import "neweos.de/sube/token"

type Parser struct {
  tokens  []lexer.Token
  current int
}

// func (p *Parser) setInput(tokens []token.Token) {
//   p.tokens = tokens
// }
// func (p *Parser) program() []ast.Statement {
//   statements := make([]ast.Statement, 0)

//   for {
//     if p.isEnd() { break }
//     statements = append(statements, p.declaration())
//   }

//   return statements
// }
func (p *Parser) expression() ast.Expression {
  return p.assignment()
}

func (p *Parser) assignment() ast.Expression {
  expression := p.logicOr()

  if p.match(token.Equal) {
  }

//     equals := p.previous()
//     value := p.assignment()

//     if result, ok := expression.(ast.VariableExpression); ok {
//       return ast.AssignExpression{
//         Name: result.Name,
//         Value: value,
//       }
//     }

//     if result, ok := expression.(ast.GetExpression); ok {
//       return ast.SetExpression{
//         Object: result.Object,
//         Name: result.Name,
//         Value: value,
//       }
//     }

//     panic(ParseError{token: equals, message: "Invalid assignment target"})
//   }

//   return expression
// }
}

func (p *Parser) logicOr() ast.Expression {

}
// func (p *Parser) or() ast.Expression {
//   expression := p.and()

//   for p.match(token.Or) {
//     operator := p.previous()
//     right := p.and()
//     expression = ast.LogicalExpression{
//       Operator: operator,
//       Left: expression,
//       Right: right,
//     }
//   }

//   return expression
// }

func (p *Parser) logicAnd() ast.Expression {

}
// func (p *Parser) and() ast.Expression {
//   expression := p.equality()

//   for p.match(token.And) {
//     operator := p.previous()
//     right := p.equality()
//     expression = ast.LogicalExpression{
//       Operator: operator,
//       Left: expression,
//       Right: right,
//     }
//   }

//   return expression
// }

func (p *Parser) equality() ast.Expression {
//   expression := p.comparison()

//   for p.match(token.BangEqual, token.EqualEqual) {
//     operator := p.previous()
//     right := p.comparison()
//     expression = ast.BinaryExpression{
//       Operator: operator,
//       Left: expression,
//       Right: right,
//     }
//   }

//   return expression
// }
}

func (p *Parser) comparison() ast.Expression {
//   expression := p.term()

//   for p.match(token.Greater, token.GreaterEqual, token.Less, token.LessEqual) {
//     operator := p.previous()
//     right := p.term()
//     expression = ast.BinaryExpression{
//       Operator: operator,
//       Left: expression,
//       Right: right,
//     }
//   }

//   return expression
// }
}

func (p *Parser) term() ast.Expression {

//   expression := p.factor()

//   for p.match(token.Minus, token.Plus) {
//     operator := p.previous()
//     right := p.factor()
//     expression = ast.BinaryExpression{
//       Operator: operator,
//       Left: expression,
//       Right: right,
//     }
//   }

//   return expression
// }
}

func (p *Parser) factor() ast.Expression {
  exp := p.unary()
  //   for p.match(token.Slash, token.Star) {
  //     operator := p.previous()
  //     right := p.unary()
  //     expression = ast.BinaryExpression{
  //       Operator: operator,
  //       Left: expression,
  //       Right: right,
  //     }
  //   }

  //   return expression
  // }
}

func (p *Parser) unary() ast.Expression {
  if p.match(token.Bang, token.Minus) {
    tok := p.previous()
    exp := p.unary()
    return ast.Unary{
      Operator: tok.Type,
      Right: exp,
    }
  }
  return p.call()
}

func (p *Parser) call() ast.Expression {
  exp := p.primary()

  for {
    if p.match(token.LeftParen) {
      exp = p.finishCall(exp)
      continue
    }

    if p.match(token.Dot) {
      tok := p.consume(token.Identifier, "expect property name after '.'")
      exp = ast.Get{Property: tok.Value, Object: exp}
      continue
    }
    break
  }
  return exp
}

func (p *Parser) finishCall(callee ast.Expression) ast.Expression {
  args := make([]ast.Expression, 0)

  if !p.check(token.RightParen) {
    for {
      if len(args) >= 255 {
        // panic(ParseError{token: p.peek(), message: "Can't have more than 255 arguments"})
        panic("error")
      }

      // arguments = append(arguments, p.expression())
      args = append(args, p.primary())
      if !p.match(token.Comma) {
        break
      }
    }
  }

  p.consume(token.RightParen, "expect ')' after args")
  return ast.Call{Arguments: args, Callee: callee}
}

func (p *Parser) primary() ast.Expression {
  if p.match(token.Nil) {
    return ast.Literal{Type: token.Nil, Value: nil}
  }

  if p.match(token.Bool) {
    tok := p.previous()
    value, _ := strconv.ParseBool(tok.Value)
    return ast.Literal{Type: token.Bool, Value: value}
  }

  if p.match(token.Int) {
    tok := p.previous()
    value, _ := strconv.Atoi(tok.Value)
    return ast.Literal{Type: token.Int, Value: value}
  }

  if p.match(token.String) {
    tok := p.previous()
    return ast.Literal{Type: token.String, Value: tok.Value}
  }

  if p.match(token.Super) {
    p.consume(token.Dot, "expect '.' after 'super'")
    tok := p.consume(token.Identifier, "expect superclass method name")
    return ast.Super{Method: tok.Value}
  }

  if p.match(token.Self) {
    return ast.Self{}
  }

  if p.match(token.Identifier) {
    tok := p.previous()
    return ast.Variable{Name: tok.Value}
  }

  if p.match(token.LeftParen) {
    exp := p.primary()
    p.consume(token.RightParen, "expect ')' after expression")
    return ast.Grouping{Expr: exp}
  }

  //   panic(ParseError{token: p.peek(), message: "Expect expression"})
  panic("error")
}

func (p *Parser) consume(tok token.Token, message string) lexer.Token {
  if p.check(tok) { return p.advance() }
  //   panic(ParseError{token: p.peek(), message: message})
  panic(message)
}

func (p *Parser) match(types ...token.Token) bool {
  for _, tokenType := range types {
    if p.check(tokenType) {
      p.advance()
      return true
    }
  }
  return false
}

func (p *Parser) check(tok token.Token) bool {
  if p.isEnd() { return false }
  return p.peek().Type == tok
}

func (p *Parser) advance() lexer.Token {
  if !p.isEnd() { p.current += 1 }
  return p.previous()
}

func (p *Parser) previous() lexer.Token {
  return p.tokens[p.current-1]
}

func (p *Parser) peek() lexer.Token {
  return p.tokens[p.current]
}

func (p *Parser) isEnd() bool {
  return p.peek().Type == token.EOF
}

// func Parse(tokens []token.Token) []ast.Statement {
//   p := Parser{current: 0}
//   p.setInput(tokens)

//   return p.program()
// }

// func (p *Parser) declaration() ast.Statement {
//   defer func() {
//     result, ok := recover().(ParseError)

//     if ok {
//       _ = result
//       fmt.Println(result.Error())
//       p.synchronize()
//     }
//   }()

//   if p.match(token.Class) { return p.classDeclaration() }
//   if p.match(token.Func) { return p.functionDeclaration("function") }
//   if p.match(token.Var) { return p.variableDeclaration() }

//   return p.statement()
// }

// func (p *Parser) ifStatement() ast.Statement {
//   p.consume(token.LeftParen, "Expect '(' after 'if'")
//   condition := p.expression()

//   p.consume(token.RightParen, "Expect ')' after if condition")
//   thenBranch := p.statement()

//   var elseBranch ast.Statement // nil eq null

//   if p.match(token.Else) {
//     elseBranch = p.statement()
//   }

//   return ast.IfStatement{
//     Condition: condition,
//     Then: thenBranch,
//     Else: elseBranch,
//   }
// }

// func (p *Parser) forStatement() ast.Statement {
//   p.consume(token.LeftParen, "Expect '(' after 'for'")
//   var initializer ast.Statement

//   if p.match(token.Semicolon) {
//     initializer = nil
//   } else if p.match(token.Var) {
//     initializer = p.variableDeclaration()
//   } else {
//     initializer = p.expressionStatement()
//   }

//   var condition ast.Expression = nil

//   if !p.check(token.Semicolon) {
//     condition = p.expression()
//   }

//   p.consume(token.Semicolon, "Expect ';' after loop condition")
//   var increment ast.Expression = nil

//   if !p.check(token.RightParen) {
//     increment = p.expression()
//   }

//   p.consume(token.RightParen, "Expect ')' after for clauses")
//   body := p.statement()

//   if increment != nil {
//     body = ast.BlockStatement{
//       Body: []ast.Statement{
//         body,
//         ast.ExpressionStatement{
//           Expression: increment,
//         },
//       },
//     }
//   }

//   if condition == nil {
//     condition = ast.LiteralExpression{Value: true}
//   }

//   body = ast.WhileStatement{Condition: condition, Body: body}

//   if initializer != nil {
//     body = ast.BlockStatement{
//       Body: []ast.Statement{
//         initializer,
//         body,
//       },
//     }
//   }

//   return body
// }

// func (p *Parser) whileStatement() ast.Statement {
//   p.consume(token.LeftParen, "Expect '(' after 'while'")
//   condition := p.expression()

//   p.consume(token.RightParen, "Expect ')' after condition")
//   body := p.statement()

//   return ast.WhileStatement{
//     Condition: condition,
//     Body: body,
//   }
// }

// func (p *Parser) blockStatement() []ast.Statement {
//   statements := make([]ast.Statement, 0)

//   for {
//     if p.check(token.RightBrace) { break }
//     if p.isEnd() { break }

//     statements = append(statements, p.declaration())
//   }

//   p.consume(token.RightBrace, "Expect '}' after block")
//   return statements
// }

// func (p *Parser) expressionStatement() ast.Statement {
//   expression := p.expression()
//   p.consume(token.Semicolon, "Expect ';' after expression")
//   return ast.ExpressionStatement{Expression: expression}
// }

// func (p *Parser) variableDeclaration() ast.Statement {
//   name := p.consume(token.Identifier, "Expect variable name")
//   var initializer ast.Expression = nil

//   if p.match(token.Equal) {
//     initializer = p.expression()
//   }

//   p.consume(token.Semicolon, "Expect ';' after variable declaration")

//   return ast.VariableDeclaration{
//     Initializer: initializer,
//     Name: name,
//   }
// }

// func (p *Parser) functionDeclaration(kind string) ast.FunctionDeclaration {
//   name := p.consume(token.Identifier, "Expect " + kind + " name")
//   p.consume(token.LeftParen, "Expect '(' after " + kind + " name")
//   parameters := make([]token.Token, 0)

//   if !p.check(token.RightParen) {
//     for {
//       if len(parameters) >= 255 {
//         panic(ParseError{token: p.peek(), message: "Can't have more than 255 parameters"})
//       }

//       parameters = append(parameters, p.consume(token.Identifier, "Expect parameter name"))
//       if !p.match(token.Comma) { break }
//     }
//   }

//   p.consume(token.RightParen, "Expect ')' after parameters")
//   p.consume(token.LeftBrace, "Expect '{' before " + kind + " body")

//   body := p.blockStatement()
//   return ast.FunctionDeclaration{
//     Params: parameters,
//     Name: name,
//     Body: body,
//   }
// }

// func (p *Parser) returnStatement() ast.Statement {
//   var value ast.Expression = nil
//   keyword := p.previous()

//   if !p.check(token.Semicolon) {
//     value = p.expression()
//   }

//   p.consume(token.Semicolon, "Expect ';' after return value")
//   return ast.ReturnStatement{
//     Keyword: keyword,
//     Value: value,
//   }
// }

// func (p *Parser) echoStatement() ast.Statement {
//   expression := p.expression()
//   p.consume(token.Semicolon, "Expect ';' after value")
//   return ast.EchoStatement{Expression: expression}
// }

// func (p *Parser) classDeclaration() ast.Statement {
//   name := p.consume(token.Identifier, "Expect class name")
//   var superclass ast.VariableExpression // nil

//   if p.match(token.Less) {
//     p.consume(token.Identifier, "Expect superclass name")
//     superclass = ast.VariableExpression{
//       Name: p.previous(),
//     }
//   }

//   p.consume(token.LeftBrace, "Expect '{' before class body")
//   methods := make([]ast.ClassMethod, 0)

//   for {
//     if p.check(token.RightBrace) { break }
//     if p.isEnd() { break }

//     declaration := p.functionDeclaration("method")
//     methods = append(methods, ast.ClassMethod{
//       Declaration: declaration,
//       Kind: "method",
//     })
//   }

//   p.consume(token.RightBrace, "Expect '}' after class body")

//   return ast.ClassDeclaration{
//     SuperClass: &superclass,
//     Methods: methods,
//     Name: name,
//   }
// }

// func (p *Parser) statement() ast.Statement {
//   if p.match(token.For)       { return p.forStatement()    }
//   if p.match(token.If)        { return p.ifStatement()     }
//   if p.match(token.Echo)      { return p.echoStatement()   }
//   if p.match(token.Return)    { return p.returnStatement() }
//   if p.match(token.While)     { return p.whileStatement()  }

//   if p.match(token.LeftBrace) {
//     return ast.BlockStatement{Body: p.blockStatement()}
//   }

//   return p.expressionStatement()
// }

// func (p *Parser) synchronize() {
//   p.advance()

//   for {
//     if p.isEnd() { break }
//     if p.previous().Type == token.Semicolon {
//       return
//     }

//     switch p.peek().Type {
//       case token.Class:  return
//       case token.Func:   return
//       case token.Var:    return
//       case token.For:    return
//       case token.If:     return
//       case token.While:  return
//       case token.Echo:   return
//       case token.Return: return
//     }

//     p.advance()
//   }
// }
