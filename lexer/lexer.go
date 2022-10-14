package lexer

import "neweos.de/sube/token"

type Token struct {
  Value string
  Kind  token.Token
  Pos   token.Position
}

type Lexer struct {
  tokens  []Token
  input   []rune
  line    int
  start   int
  current int
}

func (l *Lexer) ScanTokens() []Token {
  for {
    if l.isEnd() { break }
    l.start = l.current
    l.scanToken()
  }

  l.produce(token.EOF)
  return l.tokens
}

func (l *Lexer) scanToken() bool {
  c := l.advance()

  if l.skipWhitespace(c) {
    return false
  }

  if l.skipComment(c) {
    return false
  }

  return l.singleToken(c) ||
         l.doubleToken(c) ||
         l.tripleToken(c) ||
         l.stringToken(c) ||
         l.numberToken(c) ||
         l.identifierToken(c)
}

func (l *Lexer) skipWhitespace(c rune) bool {
  if c == '\n' {
    l.line += 1
    return true
  }

  return c == '\r' ||
         c == '\t' ||
         c == ' '
}

func (l *Lexer) skipComment(c rune) bool {
  if c == '/' && l.match('/') {
    for {
      if l.peek() == '\n' || l.isEnd() {
        break
      }
      l.advance()
    }
    return true
  }
  return false
}

func (l *Lexer) identifierToken(c rune) bool {
  if isAlpha(c) {
    for isAlphaNumeric(l.peek()) {
      l.advance()
    }

    value := string(l.input[l.start:l.current])
    l.produce(token.Lookup(value))
  }
  return false
}

func (l *Lexer) stringToken(c rune) bool {
  if c == '"' {
    for {
      if l.peek() == '"' || l.isEnd() { break }
      if l.peek() == '\n' { l.line += 1 }
      l.advance()
    }

    if l.isEnd() {
      // report(l.line, "unterminated string")
      return false
    }

    value := string(l.input[l.start+1:l.current-1])
    l.addToken(token.String, value)
    l.advance()
    return true
  }
  return false
}

func (l *Lexer) numberToken(c rune) bool {
  if isDigit(c) {
    for isDigit(l.peek()) {
      l.advance()
    }

    if l.peek() == '.' && isDigit(l.peekNext()) {
      l.advance()

      for isDigit(l.peek()) {
        l.advance()
      }
    }

    lexeme := string(l.input[l.start:l.current])
    l.addToken(token.Int, lexeme)

    return true
  }
  return false
}

func (l *Lexer) singleToken(c rune) bool {
  switch c {
    case '+': l.produce(token.Plus)
    case '-': l.produce(token.Minus)
    case '*': l.produce(token.Star)
    case '/': l.produce(token.Slash)
    case '%': l.produce(token.Modulo)
    case '.': l.produce(token.Dot)
    case ';': l.produce(token.Semi)
    case ':': l.produce(token.Colon)
    case '(': l.produce(token.LeftParen)
    case '[': l.produce(token.LeftSquare)
    case '{': l.produce(token.LeftBrace)
    case ')': l.produce(token.RightParen)
    case ']': l.produce(token.RightSquare)
    case '}': l.produce(token.RightBrace)
    case ',': l.produce(token.Comma)
    case '~': l.produce(token.Tilde)
    default:
      return false
  }
  return true
}

func (l *Lexer) doubleToken(c rune) bool {
  switch c {
    case '!':
      if l.match('=') {
        l.produce(token.BangEq)
      } else {
        l.produce(token.Bang)
      }
    case '=':
      if l.match('=') {
        l.produce(token.EqualEq)
      } else {
        l.produce(token.Equal)
      }
    case '>':
      if l.match('=') {
        l.produce(token.GreaterEq)
      } else {
        l.produce(token.Greater)
      }
    case '<':
      if l.match('=') {
        l.produce(token.LessEq)
      } else {
        l.produce(token.Less)
      }
    case '?':
      if l.match('.') {
        l.produce(token.QuestionDot)
      } else {
        l.produce(token.Question)
      }
    default:
      return false
  }
  return true
}

func (l *Lexer) tripleToken(c rune) bool {
  switch c {
    case '&':
      if l.match('&') {
        if l.match('=') {
          l.produce(token.AndAndEq)
        } else {
          l.produce(token.AndAnd)
        }
      } else {
        l.produce(token.And)
      }
    case '|':
      if l.match('|') {
        if l.match('=') {
          l.produce(token.OrOrEq)
        } else {
          l.produce(token.OrOr)
        }
      } else {
        if l.match('>') {
          l.produce(token.Pipeline)
        } else {
          l.produce(token.Or)
        }
      }
    default:
      return false
  }
  return true
}

func (l *Lexer) makeToken(tok token.Token, literal any) Token {
  pos := token.Position{Line: l.line}
  value, _ := literal.(string)

  return Token{
    Value: value,
    Kind: tok,
    Pos: pos,
  }
}

func (l *Lexer) addToken(tok token.Token, value string) {
  l.tokens = append(l.tokens, l.makeToken(tok, value))
}

func (l *Lexer) produce(tok token.Token) {
  l.tokens = append(l.tokens, l.makeToken(tok, nil))
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

func Tokenize(code string) []string {
  tokens := New(code).ScanTokens()
  values := make([]string, 0)

  for _, tok := range tokens {
    values = append(values, tok.Value)
  }

  return values
}

func New(code string) *Lexer {
  return &Lexer{
    tokens: make([]Token, 0),
    input: []rune(code),
    line: 1,
  }
}
