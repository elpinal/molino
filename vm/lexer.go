package vm

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
  case isLetter(ch):
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
    case '(', ')', '[', ']', '{', '}', '+', '-', '*', '/', '%', '=', '&':
      tok = int(ch)
      lit = string(ch)
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
  return ch == '&' || ch == '?' || ch == '-' || ch == '*'
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
  for isDigit(s.peek()) {
    ret = append(ret, s.peek())
    s.next()
  }
  return string(ret)
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
