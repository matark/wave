package lexer

import "neweos.de/sube/token"

type Token struct {
  Type  token.Token
  Value string
}

func (tok Token) String() string {
  return ""
}

func (tok Token) Lexeme() string {
  if tok.Type == token.String {
    return "" + tok.Value + ""
  }

  return tok.Value
}

// //   Pos   token.Position
