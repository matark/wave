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
         l.numberToken(c)
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
      //     report(l.line, "unterminated string")
      //     return false
      return false
    }

    //   value := string(l.input[l.start+1:l.current-1])
    //   l.tokens = append(l.tokens, l.makeToken(token.String, value))
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

    //   value, _ := strconv.ParseFloat(string(l.input[l.start:l.current]), 64)
    //   l.tokens = append(l.tokens, l.makeToken(token.Number, value))
    l.produce(token.Int)

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

func (l *Lexer) produce(tok token.Token) {
//   l.tokens = append(l.tokens, l.makeToken(tokenType, nil))
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

// func [](code string) []token.[Token] {
//   return New(code).scanTokens()
// func New(code string) *Lexer {
//   return &Lexer{
//     tokens: make([]token.Token, 0),
//     input: []rune(code),
//     current: 0,
//     start: 0,
//     line: 1,
//   }
// func (l *Lexer) literalToken(c rune) bool {
// return l.identifierToken(c) ||
//        l.literalToken(c)
