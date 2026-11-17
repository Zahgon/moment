package jsre

import (
	"fmt"
	"strings"
	"unicode/utf16"
)

type nodeKind int

const (
	nAlt nodeKind = iota
	nSeq
	nChar
	nClass
	nAny
	nStart
	nEnd
	nWordB
	nNotWordB
	nGroup
	nLook
	nBackref
	nRepeat
)

// node is one element of a parsed pattern.
type node struct {
	kind     nodeKind
	subs     []*node
	ch       uint16
	cls      *class
	capture  int // nGroup: capture index, 0 for a non-capturing group
	ahead    bool
	negate   bool
	ref      int
	refName  string
	min, max int // nRepeat: max < 0 means unbounded
	greedy   bool
	capFirst int // nRepeat: captures inside the repeated atom
	capLast  int
}

type classRange struct{ lo, hi uint16 }

type class struct {
	negate bool
	ranges []classRange
	// escapes are \d \D \w \W \s \S members of the class.
	escapes []byte
}

type parser struct {
	src      []uint16
	pos      int
	ncap     int
	names    map[string]int
	hasNames bool
	total    int
}

type syntaxError struct{ msg string }

func (e *syntaxError) Error() string { return e.msg }

func (p *parser) fail(msg string) {
	panic(&syntaxError{msg})
}

func (p *parser) more() bool         { return p.pos < len(p.src) }
func (p *parser) peek() uint16       { return p.src[p.pos] }
func (p *parser) peekIs(c byte) bool { return p.more() && p.src[p.pos] == uint16(c) }

// countGroups pre-scans the pattern for capturing groups and group names (backreferences may point forward).
func (p *parser) countGroups() {
	inClass := false
	for i := 0; i < len(p.src); i++ {
		c := p.src[i]
		switch {
		case c == '\\':
			i++
		case inClass:
			if c == ']' {
				inClass = false
			}
		case c == '[':
			inClass = true
		case c == '(':
			if i+1 < len(p.src) && p.src[i+1] == '?' {
				if i+2 < len(p.src) && p.src[i+2] == '<' && i+3 < len(p.src) && p.src[i+3] != '=' && p.src[i+3] != '!' {
					p.total++
					p.hasNames = true
				}
			} else {
				p.total++
			}
		}
	}
}

func parse(source string) (root *node, ncap int, names map[string]int, err error) {
	p := &parser{src: utf16.Encode([]rune(source)), names: map[string]int{}}
	p.countGroups()
	defer func() {
		if r := recover(); r != nil {
			se, ok := r.(*syntaxError)
			if !ok {
				panic(r)
			}
			err = se
		}
	}()
	root = p.disjunction()
	if p.more() {
		if p.peek() == ')' {
			p.fail("Unmatched ')'")
		}
		p.fail("Unexpected character")
	}
	return root, p.ncap, p.names, nil
}

func (p *parser) disjunction() *node {
	alts := []*node{p.alternative()}
	for p.peekIs('|') {
		p.pos++
		alts = append(alts, p.alternative())
	}
	if len(alts) == 1 {
		return alts[0]
	}
	return &node{kind: nAlt, subs: alts}
}

func (p *parser) alternative() *node {
	seq := &node{kind: nSeq}
	for p.more() && !p.peekIs('|') && !p.peekIs(')') {
		capBefore := p.ncap
		atom, quantifiable := p.term()
		if q := p.quantifier(); q != nil {
			if !quantifiable {
				p.fail("Nothing to repeat")
			}
			q.subs = []*node{atom}
			q.capFirst, q.capLast = capBefore+1, p.ncap
			atom = q
		}
		seq.subs = append(seq.subs, atom)
	}
	return seq
}

// readInt reads decimal digits; ok is false when there are none.
func (p *parser) readInt() (int, bool) {
	start := p.pos
	n := 0
	for p.more() && p.peek() >= '0' && p.peek() <= '9' {
		if n < 1<<30 {
			n = n*10 + int(p.peek()-'0')
		}
		p.pos++
	}
	return n, p.pos > start
}

// braceQuantifier parses {n}, {n,} or {n,m} at p.pos; it restores the position when the text is not one.
func (p *parser) braceQuantifier() (int, int, bool) {
	start := p.pos
	p.pos++
	lo, ok := p.readInt()
	if !ok {
		p.pos = start
		return 0, 0, false
	}
	hi := lo
	if p.peekIs(',') {
		p.pos++
		if v, ok := p.readInt(); ok {
			hi = v
		} else {
			hi = -1
		}
	}
	if !p.peekIs('}') {
		p.pos = start
		return 0, 0, false
	}
	p.pos++
	return lo, hi, true
}

func (p *parser) quantifier() *node {
	if !p.more() {
		return nil
	}
	var lo, hi int
	switch p.peek() {
	case '*':
		lo, hi = 0, -1
		p.pos++
	case '+':
		lo, hi = 1, -1
		p.pos++
	case '?':
		lo, hi = 0, 1
		p.pos++
	case '{':
		var ok bool
		if lo, hi, ok = p.braceQuantifier(); !ok {
			return nil
		}
		if hi >= 0 && hi < lo {
			p.fail("numbers out of order in {} quantifier")
		}
	default:
		return nil
	}
	q := &node{kind: nRepeat, min: lo, max: hi, greedy: true}
	if p.peekIs('?') {
		p.pos++
		q.greedy = false
	}
	return q
}

func (p *parser) term() (*node, bool) {
	c := p.peek()
	switch c {
	case '^':
		p.pos++
		return &node{kind: nStart}, false
	case '$':
		p.pos++
		return &node{kind: nEnd}, false
	case '.':
		p.pos++
		return &node{kind: nAny}, true
	case '(':
		return p.group()
	case '[':
		return &node{kind: nClass, cls: p.class()}, true
	case '\\':
		return p.atomEscape()
	case '*', '+', '?':
		p.fail("Nothing to repeat")
	case '{':
		if _, _, ok := p.braceQuantifier(); ok {
			p.fail("Nothing to repeat")
		}
	}
	p.pos++
	return &node{kind: nChar, ch: c}, true
}

func (p *parser) group() (*node, bool) {
	p.pos++
	if p.peekIs('?') {
		p.pos++
		if !p.more() {
			p.fail("Invalid group")
		}
		switch p.peek() {
		case ':':
			p.pos++
			n := &node{kind: nGroup, subs: []*node{p.disjunction()}}
			p.closeGroup()
			return n, true
		case '=', '!':
			neg := p.peek() == '!'
			p.pos++
			n := &node{kind: nLook, ahead: true, negate: neg, subs: []*node{p.disjunction()}}
			p.closeGroup()
			return n, true
		case '<':
			p.pos++
			if p.peekIs('=') || p.peekIs('!') {
				neg := p.peek() == '!'
				p.pos++
				n := &node{kind: nLook, negate: neg, subs: []*node{p.disjunction()}}
				p.closeGroup()
				return n, false
			}
			name := p.groupName()
			p.ncap++
			idx := p.ncap
			if _, dup := p.names[name]; dup {
				p.fail("Duplicate capture group name")
			}
			p.names[name] = idx
			n := &node{kind: nGroup, capture: idx, subs: []*node{p.disjunction()}}
			p.closeGroup()
			return n, true
		}
		p.fail("Invalid group")
	}
	p.ncap++
	idx := p.ncap
	n := &node{kind: nGroup, capture: idx, subs: []*node{p.disjunction()}}
	p.closeGroup()
	return n, true
}

func (p *parser) groupName() string {
	var b strings.Builder
	for p.more() && p.peek() != '>' {
		b.WriteRune(rune(p.peek()))
		p.pos++
	}
	if !p.more() || b.Len() == 0 {
		p.fail("Invalid capture group name")
	}
	p.pos++
	return b.String()
}

func (p *parser) closeGroup() {
	if !p.peekIs(')') {
		p.fail("Unterminated group")
	}
	p.pos++
}

func isHex(c uint16) bool {
	return (c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')
}

func hexVal(c uint16) uint16 {
	switch {
	case c >= '0' && c <= '9':
		return c - '0'
	case c >= 'a' && c <= 'f':
		return c - 'a' + 10
	}
	return c - 'A' + 10
}

// hexEscape reads n hex digits after p.pos; ok is false (position unchanged) when they are not there.
func (p *parser) hexEscape(n int) (uint16, bool) {
	if p.pos+n > len(p.src) {
		return 0, false
	}
	var v uint16
	for i := 0; i < n; i++ {
		c := p.src[p.pos+i]
		if !isHex(c) {
			return 0, false
		}
		v = v*16 + hexVal(c)
	}
	p.pos += n
	return v, true
}

// legacyOctal reads up to three octal digits forming a value below 256.
func (p *parser) legacyOctal() uint16 {
	var v uint16
	for i := 0; i < 3 && p.more() && p.peek() >= '0' && p.peek() <= '7'; i++ {
		next := v*8 + (p.peek() - '0')
		if next > 255 {
			break
		}
		v = next
		p.pos++
	}
	return v
}

// characterEscape parses the escape after a backslash that denotes one code unit (shared by atoms and classes).
func (p *parser) characterEscape(inClass bool) uint16 {
	c := p.peek()
	p.pos++
	switch c {
	case 't':
		return '\t'
	case 'n':
		return '\n'
	case 'v':
		return '\v'
	case 'f':
		return '\f'
	case 'r':
		return '\r'
	case 'c':
		if p.more() {
			l := p.peek()
			if (l >= 'a' && l <= 'z') || (l >= 'A' && l <= 'Z') || (inClass && ((l >= '0' && l <= '9') || l == '_')) {
				p.pos++
				return l % 32
			}
		}
		p.pos--
		return '\\'
	case 'x':
		if v, ok := p.hexEscape(2); ok {
			return v
		}
		return 'x'
	case 'u':
		if v, ok := p.hexEscape(4); ok {
			return v
		}
		return 'u'
	case '0':
		if !p.more() || p.peek() < '0' || p.peek() > '9' {
			return 0
		}
		p.pos--
		return p.legacyOctal()
	}
	if c >= '1' && c <= '7' {
		p.pos--
		return p.legacyOctal()
	}
	return c
}

func (p *parser) atomEscape() (*node, bool) {
	p.pos++
	if !p.more() {
		p.fail("\\ at end of pattern")
	}
	c := p.peek()
	switch c {
	case 'b':
		p.pos++
		return &node{kind: nWordB}, false
	case 'B':
		p.pos++
		return &node{kind: nNotWordB}, false
	case 'd', 'D', 'w', 'W', 's', 'S':
		p.pos++
		return &node{kind: nClass, cls: &class{escapes: []byte{byte(c)}}}, true
	case 'k':
		if p.hasNames {
			p.pos++
			if !p.peekIs('<') {
				p.fail("Invalid named reference")
			}
			p.pos++
			return &node{kind: nBackref, refName: p.groupName()}, true
		}
	}
	if c >= '1' && c <= '9' {
		start := p.pos
		n, _ := p.readInt()
		if n <= p.total {
			return &node{kind: nBackref, ref: n}, true
		}
		p.pos = start
		if c >= '8' {
			p.pos++
			return &node{kind: nChar, ch: c}, true
		}
	}
	return &node{kind: nChar, ch: p.characterEscape(false)}, true
}

// classAtom parses one class member: a code unit, or a class escape reported in esc.
func (p *parser) classAtom() (ch uint16, esc byte) {
	c := p.peek()
	if c != '\\' {
		p.pos++
		return c, 0
	}
	p.pos++
	if !p.more() {
		p.fail("\\ at end of pattern")
	}
	switch e := p.peek(); e {
	case 'd', 'D', 'w', 'W', 's', 'S':
		p.pos++
		return 0, byte(e)
	case 'b':
		p.pos++
		return '\b', 0
	case '-':
		p.pos++
		return '-', 0
	}
	if e := p.peek(); e == '8' || e == '9' {
		p.pos++
		return e, 0
	}
	return p.characterEscape(true), 0
}

func (p *parser) class() *class {
	p.pos++
	cls := &class{}
	if p.peekIs('^') {
		p.pos++
		cls.negate = true
	}
	for {
		if !p.more() {
			p.fail("Unterminated character class")
		}
		if p.peek() == ']' {
			p.pos++
			return cls
		}
		lo, loEsc := p.classAtom()
		if p.peekIs('-') && p.pos+1 < len(p.src) && p.src[p.pos+1] != ']' {
			p.pos++
			hi, hiEsc := p.classAtom()
			if loEsc != 0 || hiEsc != 0 {
				cls.add(lo, loEsc)
				cls.add('-', 0)
				cls.add(hi, hiEsc)
				continue
			}
			if hi < lo {
				p.fail("Range out of order in character class")
			}
			cls.ranges = append(cls.ranges, classRange{lo, hi})
			continue
		}
		cls.add(lo, loEsc)
	}
}

func (c *class) add(ch uint16, esc byte) {
	if esc != 0 {
		c.escapes = append(c.escapes, esc)
		return
	}
	c.ranges = append(c.ranges, classRange{ch, ch})
}

func errorMessage(source, flags, msg string) string {
	return fmt.Sprintf("Invalid regular expression: /%s/%s: %s", source, flags, msg)
}
