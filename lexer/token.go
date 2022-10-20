package lexer

import "neweos.de/sube/token"

type Token struct {
  Type  token.Token
  Value string
  token.Position
}

func (tok Token) String() string {
  if tok.Type.IsLiteral() {
    return tok.Type.String() + "(" + tok.Lexeme() + ")"
  }
  return tok.Type.String()
}

func (tok Token) Lexeme() string {
  if tok.Value == "" {
    return tok.Type.String()
  }

  switch tok.Type {
    case token.String:
      return "\"" + tok.Value + "\""
    default:
      return tok.Value
  }
}
