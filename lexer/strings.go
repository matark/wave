package lexer

func isAlphaNumeric(c rune) bool {
  return isAlpha(c) || isDigit(c)
}

func isAlpha(c rune) bool {
  return ('A' <= c && c <= 'Z') ||
         ('a' <= c && c <= 'z') ||
         ('_' == c)
}

func isDigit(c rune) bool {
  return '0' <= c && c <= '9'
}
