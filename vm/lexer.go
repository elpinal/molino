package vm

import (
  _ "fmt"
  "regexp"
  "strconv"
)

const (
  EOF     = -1
  UNKNOWN = 0
)

var keywords = map[string]int{
  "def"  : DEF,
  "if"   : IF,
  "true" : TRUE,
  "false": FALSE,
  "nil"  : NIL,
  "fn"   : FN,
  "quote": QUOTE,
}

var intPat *regexp.Regexp = regexp.MustCompile("^([-+]?)(?:(0)|([1-9][0-9]*)|0[xX]([0-9A-Fa-f]+)|0([0-7]+)|([1-9][0-9]?)[rR]([0-9A-Za-z]+)|0[0-9]+)$")


type Position struct {
  Line   int
  Column int
}

type Scanner struct {
  src      []rune
  offset   int
  lineHead int
  line     int
}

func (s *Scanner) Init(src string) {
  s.src = []rune(src)
}

func (s *Scanner) Scan() (tok int, lit string, pos Position) {
  for isWhiteSpace(s.peek()) || s.peek() == ';' {
    s.skipWhiteSpace()
    s.skipComment()
  }
  pos = s.position()
  switch ch := s.peek(); {
  case isLetter(ch), ch == '*', ch == '=':
    lit = s.scanIdentifier()
    if keyword, ok := keywords[lit]; ok {
      tok = keyword
    } else {
      tok = IDENT
    }
  case isDigit(ch):
    tok, lit = NUMBER, s.scanNumber()
  default:
    switch ch {
    case -1:
      tok = EOF
    case '(', ')', '[', ']', '{', '}', '&':
      tok = int(ch)
      lit = string(ch)
    case '+', '-':
      s.next()
      if isDigit(s.peek()) {
        s.back()
        tok = NUMBER
        lit = s.scanNumber()
        s.back()
      } else {
        s.back()
        tok = IDENT
        lit = s.scanIdentifier()
      }
    case '"':
      tok = STRING
      lit = s.scanString()
    case ':':
      s.next()
      tok = KEYWORD
      lit = s.scanIdentifier()
      s.back()
    }
    s.next()
  }
  return
}

// ========================================

func isLetter(ch rune) bool {
  return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch rune) bool {
  return '0' <= ch && ch <= '9'
}

func isSign(ch rune) bool {
  return /*ch == '&' ||*/ ch == '?' || ch == '-' || ch == '*' || ch == '.' || ch == '+' || ch == '='
}

func isWhiteSpace(ch rune) bool {
  return ch == ' ' || ch == '\t' || ch == '\n'
}

func (s *Scanner) peek() rune {
  if !s.reachEOF() {
    return s.src[s.offset]
  } else {
    return -1
  }
}

func (s *Scanner) next() {
  if !s.reachEOF() {
    if s.peek() == '\n' {
      s.lineHead = s.offset + 1
      s.line++
    }
    s.offset++
  }
}

func (s *Scanner) back() {
  s.offset--
}

func (s *Scanner) reachEOF() bool {
  return len(s.src) <= s.offset
}

func (s *Scanner) position() Position {
  return Position{Line: s.line + 1, Column: s.offset - s.lineHead + 1}
}

func (s *Scanner) skipWhiteSpace() {
  for isWhiteSpace(s.peek()) {
    s.next()
  }
}

func (s *Scanner) scanIdentifier() string {
  var ret []rune
  for isLetter(s.peek()) || isDigit(s.peek()) || isSign(s.peek()) {
    ret = append(ret, s.peek())
    s.next()
  }
  return string(ret)
}

func (s *Scanner) scanNumber() string {
  var ret []rune
//  for isDigit(s.peek()) {
  loop:
  for {
    if ch := s.peek(); ch == -1 || isWhiteSpace(ch) || ch == '(' || ch == ')' {
      break loop
    }
    ret = append(ret, s.peek())
//    fmt.Println("009 -- ", string(ret))
    s.next()
  }
//  return string(ret)
  n, notmatch := matchNumber(string(ret))
  if !notmatch {
    panic("Invalid number: " + string(ret))
  }
  return strconv.FormatInt(n, 10)
}

func matchNumber(x string) (int64, bool) {
  m := intPat.FindStringSubmatch(x)
  if m != nil {
    if m[2] != "" {
      return 0, true
    }
    var negate bool = m[1] == "-"
    radix := 10
    var n string
    if n = m[3]; n != "" {
      radix = 10
    } else if n = m[4]; n != "" {
      radix = 16
    }
    if n == "" {
      return -1, false
    }
    ret, err := strconv.ParseInt(n, radix, 64)
    if err != nil {
      panic("error!")
    }
    if negate {
      ret = -ret
    }
    return ret, true
  }
  return -1, false
}

func (s *Scanner) scanString() string {
  var ret []rune
  s.next()
  for s.peek() != '"' {
    if s.reachEOF() {
      panic("unterminated string meets end of file\n       syntax error, unexpected end-of-input, expecting keyword_end")
    }
    if s.peek() == '\\' {
      s.next()
      if s.peek() == '"' {
        ret = append(ret, s.peek())
        s.next()
      } else {
        ret = append(ret, '\\', s.peek())
      }
    } else {
      ret = append(ret, s.peek())
      s.next()
    }
  }
  return string(ret)
}

func (s *Scanner) skipComment() {
  if s.peek() == ';' {
    for s.peek() != '\n' {
      s.next()
    }
    s.next()
  }
}
