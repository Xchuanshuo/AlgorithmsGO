package main

import (
	"fmt"
	"strings"
)

const (
	LeftBracket     = "<"
	LeftBracketSigh = "</"
	LeftBracketEnd  = ">"
	CDataStart      = "<![CDATA["
	CDataEnd        = "]]>"
)

func isValid(code string) bool {
	var validator = New(code)
	return validator.Validate()
}

type XmlValidator struct {
	idx   int
	stack []string
	str   string
}

func New(str string) *XmlValidator {
	parser := XmlValidator{
		idx:   0,
		str:   str,
		stack: make([]string, 0),
	}
	return &parser
}

func (x *XmlValidator) Validate() bool {
	var code = x.str
	for x.idx < len(x.str) {
		if code[x.idx] != '<' {
			if len(x.stack) == 0 {
				return false
			}
			x.idx += 1
			continue
		}
		if strings.HasPrefix(code[x.idx:], CDataStart) {
			if !x.parseCDATA() {
				return false
			}
		} else if strings.HasPrefix(code[x.idx:], LeftBracketSigh) {
			if !x.parseLeftBracketSigh() {
				return false
			}
		} else if strings.HasPrefix(code[x.idx:], LeftBracket) {
			if !x.parseLeftBracket() {
				return false
			}
		}
	}
	return len(x.stack) == 0
}

func (x *XmlValidator) parseCDATA() bool {
	if len(x.stack) == 0 {
		return false
	}
	x.idx += len(CDataStart)
	for x.idx < len(x.str) && !strings.HasPrefix(x.str[x.idx:], CDataEnd) {
		x.idx += 1
	}
	x.idx += len(CDataEnd)
	return x.idx < len(x.str)
}

func (x *XmlValidator) parseLeftBracket() bool {
	if x.idx > 0 && len(x.stack) == 0 {
		return false
	}
	x.idx += len(LeftBracket)
	tagName, ok := x.readTagName()
	if !ok {
		return false
	}
	x.stack = append(x.stack, tagName)
	x.idx += len(LeftBracketEnd)
	return true
}

func (x *XmlValidator) parseLeftBracketSigh() bool {
	if len(x.stack) == 0 {
		return false
	}
	x.idx += len(LeftBracketSigh)
	tagName, ok := x.readTagName()
	if !ok {
		return false
	}
	var top = x.stack[len(x.stack)-1]
	x.stack = x.stack[0 : len(x.stack)-1]
	if tagName != top {
		return false
	}
	x.idx += len(LeftBracketEnd)
	return true
}

func (x *XmlValidator) readTagName() (string, bool) {
	var bytes = make([]byte, 0)
	for x.idx < len(x.str) && x.str[x.idx] != '>' {
		var c = x.str[x.idx]
		x.idx += 1
		if c < 'A' || c > 'Z' {
			return "", false
		}
		bytes = append(bytes, c)
	}
	var str = string(bytes)
	return string(bytes), len(str) >= 1 && len(str) <= 9
}

func main() {
	var res = isValid("<DIV>This is the first line <![CDATA[<div>]]></DIV>")
	fmt.Println(res)
}
