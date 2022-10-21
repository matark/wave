package lexer

import "neweos.de/sube/token"

type Token struct {
  Type  token.Token
  Value string
  token.Position
}

func (tok Token) String() string {
  str := tok.Type.String()
  if tok.Type.IsLiteral() {
    return str + "(" + tok.Lexeme() + ")"
  }
  return str
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
