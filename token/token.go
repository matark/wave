package token

type Token int

const (
  Int Token = iota
  Float
  Complex
  Identifier
  String
  Rune
  Bool
  Nil
  Plus
  Minus
  Star
  Slash
  Modulo
  Question
  QuestionDot
  HashBang
  AndAnd
  OrOr
  And
  Or
  Dot
  Semi
  Colon
  LeftParen
  LeftSquare
  LeftBrace
  RightParen
  RightSquare
  RightBrace
  BackQuote
  Pipeline
  Comma
  Bang
  BangEq
  Equal
  EqualEq
  Greater
  GreaterEq
  Less
  LessEq
  AndAndEq
  OrOrEq
  Tilde
  Fun
  If
  Else
  For
  In
  Loop
  Break
  Return
  Continue
  Switch
  Case
  Default
  Struct
  Class
  Super
  Self
  Let
  Mut
  Throw
  Import
  Println
  Illegal
  EOF
)

func (tok Token) String() string {
  return tokens[tok]
}

func (tok Token) Precedence() int {
  switch tok {
    case OrOr:   return 1
    case AndAnd: return 2
    case BangEq,
         EqualEq,
         Greater,
         GreaterEq,
         Less,
         LessEq: return 3
    case Or,
         Plus,
         Minus:  return 4
    case And,
         Star,
         Slash,
         Modulo: return 5
    default:     return 0
  }
}

func (tok Token) IsOperator() bool {
  return Nil < tok && tok <= Tilde
}

func (tok Token) IsKeyword() bool {
  return Tilde < tok && tok < Illegal
}

func (tok Token) IsLiteral() bool {
  return Int <= tok && tok < Nil
}

func Lookup(name string) Token {
  tok, ok := keywords[name]
  if ok { return tok }
  return Identifier
}
