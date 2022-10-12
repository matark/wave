package lexer

import "neweos.de/sube/token"

type Tag [2]interface{}

type Lexer struct {
  tokens  []token.Token
  input   []rune
  tags    []Tag
  line    int
  start   int
  current int
}

func (l *Lexer) NextToken() token.Token {
  if l.scanToken() {
    return l.tokens[len(l.tokens)-1]
  }

  return token.EOF
}

func (l *Lexer) scanToken() bool {
  c := l.advance()
  return false
  // return l.identifierToken(c) ||
  //        l.commentToken(c)    ||
  //        l.whitespaceToken(c) ||
  //        l.stringToken(c)     ||
  //        l.numberToken(c)     ||
  //        l.regexToken(c)      ||
  //        l.literalToken(c)
}

func (l *Lexer) produce(tok token.Token) {

//   l.tokens = append(l.tokens, l.makeToken(tokenType, nil))
// }
}

func (l *Lexer) match(c rune) bool {
  if l.isEnd() { return false }
  if l.input[l.current] == c {
    l.current += 1
    return true
  }

  return false
}

func (l *Lexer) peek() rune {
  if l.isEnd() { return 0 }
  return l.input[l.current]
}

func (l *Lexer) peekNext() rune {
  if l.current + 1 < len(l.input) {
    return l.input[l.current+1]
  }

  return 0
}

func (l *Lexer) advance() rune {
  c := l.input[l.current]
  l.current += 1
  return c
}

func (l *Lexer) isEnd() bool {
  return l.current >= len(l.input)
}

func Tokenize() []string {
}
