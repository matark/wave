
// package parser

// import "fmt"
// import "github.com/minisode/teaser/token"

// type ParseError struct {
//   token   token.Token
//   message string
// }

// func (p *ParseError) Error() string {
//   if p.token.Type == token.EOF {
//     return fmt.Sprintf("[%d]> at end, %s", p.token.Line, p.message)
//   }

//   return fmt.Sprintf("[%d]> at '%s', %s", p.token.Line, p.token.Lexeme, p.message)
// }
