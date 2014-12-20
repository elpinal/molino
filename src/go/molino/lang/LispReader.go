package vm

import (
  "regexp"
  "strconv"
)

var QUOTE Symbol = intern("quote")

var macros = map[rune]int {
  '"' : StringReader(),
  ';' : CommentReader(),
  '\'' : WrappingReader(QUOTE),
//  '@' : WrappingReader(DEREF),
//  '^' : MetaReader(),
//  '`' : SyntaxQuoteReader(),
//  '~' : UnquoteReader(),
  '(' : ListReader(),
  ')' : UnmatchedDelimiterReader(),
  '[' : VectorReader(),
  ']' : UnmatchedDelimiterReader(),
  '{' : MapReader(),
  '}' : UnmatchedDelimiterReader(),
  '\\' : CharacterReader(),
//  '%' : ArgReader(),
//  '#' : DispatchReader(),
}

var symbolPat *regexp.Regexp = regexp.MustCompile("^[:]?([^/0-9].*/)?(/|[^/0-9][^/]*)$")
var intPat    *regexp.Regexp = regexp.MustCompile("^([-+]?)(?:(0)|([1-9][0-9]*)|0[xX]([0-9A-Fa-f]+)|0([0-7]+)|([1-9][0-9]?)[rR]([0-9A-Za-z]+)|0[0-9]+)$")


type Position struct {
  Line   int
  Column int
}

type Scanner struct {
  src      []rune // source
  offset   int    // 
  lineHead int    // 
  line     int    // 
}

func (s *Scanner) Init(src string) {
  s.src = []rune(src)
}

func (s *Scanner) Scan() interface{} { //(tok int, lit string, pos Position)
  for {
    pos = s.position()
    ch := s.read()
    for isWhitespace(ch) {
      ch = s.read()
    }
    if ch == -1 {
      //
    }
    if isDigit(ch) {
      var n int64 = s.readNumber(ch)
      return n
    }
    macroFn, ismacro := getMacro(ch)
    if ismacro {
      ret := macroFn()
      //
      if ret == s {
        continue
      }
      return ret
    }
    if ch == '+' || ch == '-' {
      ch2 := s.read()
      if isDigit(ch2) {
        s.unread()
        var n int64 = s.readNumber(ch)
        return n
      }
      s.unread()
    }
    var token string = s.readToken(ch)
    return interpretToken(token)
  }
}


func isDigit(ch rune) bool {
  return '0' <= ch && ch <= '9'
}

func isWhitespace(ch rune) bool {
  return ch == ' ' || ch == '\t' || ch == '\n' || ch == ','
}

func getMacro(ch rune) (int, bool) {
  m, n := macros[ch]
  return m, n
}

func isMacro(ch rune) bool {
  _, n := macros[ch]
  return n
}

func isTerminatingMacro(ch rune) bool {
  return ch != '#' && ch != '\'' && ch != '%' && isMacro(ch)
}

func (s *Scanner) read() rune {
  if !s.reachEOF() {
    s.offset++
    ch := s.src[s.offset]
    if ch == '\n' {
      s.lineHead = s.offset
      s.line++
    }
    return ch
  }
  return -1
}

func (s *Scanner) unread() {
  s.offset--
}

func (s *Scanner) reachEOF() bool {
  return len(s.src) <= s.offset
}

func (s *Scanner) position() Position {
  return Position{Line: s.line + 1, Column: s.offset - s.lineHead + 1}
}

func (s *Scanner) readToken(initch rune) string {
  var ret []rune = []rune{initch}
  for {
    if ch := s.read(); ch == -1 || isWhitespace(ch) || isTerminatingMacro(ch) {
      s.unread()
      return string(ret)
    } else {
      ret = append(ret, ch)
    }
  }
}

func (s *Scanner) readNumber(initch rune) int64 {
  var ret []rune = []rune{initch}
  loop:
  for {
    if ch := s.read(); ch == -1 || isWhitespace(ch) || isMacro(ch) {
      s.unread()
      break loop
    } else {
      ret = append(ret, ch)
    }
  }
  n, notmatch := matchNumber(string(ret))
  if !notmatch {
    panic("Invalid number: " + string(ret))
  }
  return n //strconv.FormatInt(n, 10)
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

func interpretToken(s string) interface{} {
  if s == "nil" {
    return nil
  } else if s == "true" {
    return true
  } else if s == "false" {
    return false
  }
  ret, isValid := matchSymbol(s)
  if isValid {
    return ret
  }

  panic("Invalid token: " + s)
}

func matchSymbol(s string) (interface{}, bool) {
  m := symbolPat.FindStringSubmatch(s)
  if m != nil {
    isKeyword := s[0] == ':'
    if isKeyword {
      sym := intern(s[1:])
      return /*Keyword.intern(s)*/ sym, true
    } else {
      sym := intern(s)
      return sym, true
    }
  }
  return nil, false
}
