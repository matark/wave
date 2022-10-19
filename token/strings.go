package token

var tokens = [...]string{
  Int:         "<int>",
  Float:       "<float>",
  Complex:     "<complex>",
  Identifier:  "Identifier",
  String:      "<string>",
  Rune:        "<rune>",
  Bool:        "<bool>",
  Nil:         "nil",
  Plus:        "+",
  Minus:       "-",
  Star:        "*",
  Slash:       "/",
  Modulo:      "%",
  Question:    "?",
  QuestionDot: "?.",
  HashBang:    "#!",
  AndAnd:      "&&",
  OrOr:        "||",
  And:         "&",
  Or:          "|",
  Dot:         ".",
  Semi:        ";",
  Colon:       ":",
  LeftParen:   "(",
  LeftSquare:  "[",
  LeftBrace:   "{",
  RightParen:  ")",
  RightSquare: "]",
  RightBrace:  "}",
  BackQuote:   "`",
  Pipeline:    "|>",
  Comma:       ",",
  Bang:        "!",
  BangEq:      "!=",
  Equal:       "=",
  EqualEq:     "==",
  Greater:     ">",
  GreaterEq:   ">=",
  Less:        "<",
  LessEq:      "<=",
  AndAndEq:    "&&=",
  OrOrEq:      "||=",
  Tilde:       "~",
  Fun:         "fun",
  If:          "if",
  Else:        "else",
  For:         "for",
  In:          "in",
  Loop:        "loop",
  Break:       "break",
  Return:      "return",
  Continue:    "continue",
  Switch:      "switch",
  Case:        "case",
  Default:     "default",
  Struct:      "struct",
  Class:       "class",
  Super:       "super",
  Self:        "self",
  Let:         "let",
  Mut:         "mut",
  Throw:       "throw",
  Import:      "import",
  Println:     "println",
  Illegal:     "Illegal",
  EOF:         "EOF",
}

var keywords map[string]Token

func setKeywords() {
  keywords = make(map[string]Token, 21)

  for i := Fun; i < Illegal; i++ {
    keywords[tokens[i]] = i
  }
}

func init() {
  if len(tokens) != 65 {}
  setKeywords()
}
