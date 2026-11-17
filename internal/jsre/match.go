package jsre

import "sort"

var canonicalTable = func() map[uint16]uint16 {
	t := map[uint16]uint16{}
	for _, r := range canonicalRuns {
		for c := int(r.lo); c <= int(r.hi); c += int(r.step) {
			t[uint16(c)] = uint16(int32(c) + r.delta)
		}
	}
	return t
}()

// sameCanonical lists, for each canonical value, the code units that canonicalize to it.
var sameCanonical = func() map[uint16][]uint16 {
	m := map[uint16][]uint16{}
	for c, k := range canonicalTable {
		m[k] = append(m[k], c)
	}
	for k, list := range m {
		if _, mapped := canonicalTable[k]; !mapped {
			list = append(list, k)
		}
		sort.Slice(list, func(i, j int) bool { return list[i] < list[j] })
		m[k] = list
	}
	return m
}()

func canonicalize(c uint16) uint16 {
	if k, ok := canonicalTable[c]; ok {
		return k
	}
	return c
}

func isWord(c uint16) bool {
	return c == '_' || (c >= '0' && c <= '9') || (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func isSpace(c uint16) bool {
	switch c {
	case '\t', '\n', '\v', '\f', '\r', ' ', 0xA0, 0x1680, 0x2028, 0x2029, 0x202F, 0x205F, 0x3000, 0xFEFF:
		return true
	}
	return c >= 0x2000 && c <= 0x200A
}

func isLineTerminator(c uint16) bool {
	return c == '\n' || c == '\r' || c == 0x2028 || c == 0x2029
}

func escapeMatches(esc byte, c uint16) bool {
	switch esc {
	case 'd':
		return c >= '0' && c <= '9'
	case 'D':
		return !(c >= '0' && c <= '9')
	case 'w':
		return isWord(c)
	case 'W':
		return !isWord(c)
	case 's':
		return isSpace(c)
	case 'S':
		return !isSpace(c)
	}
	return false
}

func (cls *class) contains(c uint16) bool {
	for _, r := range cls.ranges {
		if c >= r.lo && c <= r.hi {
			return true
		}
	}
	for _, e := range cls.escapes {
		if escapeMatches(e, c) {
			return true
		}
	}
	return false
}

func (cls *class) matches(c uint16, icase bool) bool {
	found := cls.contains(c)
	if !found && icase {
		k := canonicalize(c)
		candidates := sameCanonical[k]
		if candidates == nil {
			candidates = []uint16{k}
		}
		for _, other := range candidates {
			if other != c && canonicalize(other) == k && cls.contains(other) {
				found = true
				break
			}
		}
	}
	return found != cls.negate
}

// matcher runs a compiled pattern over UTF-16 input with backtracking.
type matcher struct {
	re    *Regexp
	input []uint16
	caps  []int
}

func (m *matcher) unitEqual(a, b uint16) bool {
	if a == b {
		return true
	}
	return m.re.icase && canonicalize(a) == canonicalize(b)
}

func (m *matcher) wordAt(i int) bool {
	return i >= 0 && i < len(m.input) && isWord(m.input[i])
}

// match tries n at pos and then the continuation k; captures are restored when it fails.
func (m *matcher) match(n *node, pos int, k func(int) bool) bool {
	switch n.kind {
	case nSeq:
		return m.seq(n.subs, 0, pos, k)
	case nAlt:
		for _, alt := range n.subs {
			if m.match(alt, pos, k) {
				return true
			}
		}
		return false
	case nChar:
		if pos < len(m.input) && m.unitEqual(m.input[pos], n.ch) {
			return k(pos + 1)
		}
		return false
	case nAny:
		if pos < len(m.input) && (m.re.dotAll || !isLineTerminator(m.input[pos])) {
			return k(pos + 1)
		}
		return false
	case nClass:
		if pos < len(m.input) && n.cls.matches(m.input[pos], m.re.icase) {
			return k(pos + 1)
		}
		return false
	case nStart:
		if pos == 0 || (m.re.multiline && isLineTerminator(m.input[pos-1])) {
			return k(pos)
		}
		return false
	case nEnd:
		if pos == len(m.input) || (m.re.multiline && isLineTerminator(m.input[pos])) {
			return k(pos)
		}
		return false
	case nWordB, nNotWordB:
		boundary := m.wordAt(pos-1) != m.wordAt(pos)
		if boundary == (n.kind == nWordB) {
			return k(pos)
		}
		return false
	case nGroup:
		if n.capture == 0 {
			return m.match(n.subs[0], pos, k)
		}
		i := n.capture * 2
		oldStart, oldEnd := m.caps[i], m.caps[i+1]
		if m.match(n.subs[0], pos, func(end int) bool {
			s, e := m.caps[i], m.caps[i+1]
			m.caps[i], m.caps[i+1] = pos, end
			if k(end) {
				return true
			}
			m.caps[i], m.caps[i+1] = s, e
			return false
		}) {
			return true
		}
		m.caps[i], m.caps[i+1] = oldStart, oldEnd
		return false
	case nLook:
		return m.look(n, pos, k)
	case nBackref:
		return m.backref(n, pos, k)
	case nRepeat:
		return m.repeat(n, pos, 0, k)
	}
	return false
}

func (m *matcher) seq(subs []*node, i, pos int, k func(int) bool) bool {
	if i == len(subs) {
		return k(pos)
	}
	return m.match(subs[i], pos, func(next int) bool { return m.seq(subs, i+1, next, k) })
}

func (m *matcher) look(n *node, pos int, k func(int) bool) bool {
	saved := append([]int(nil), m.caps...)
	var ok bool
	if n.ahead {
		ok = m.match(n.subs[0], pos, func(int) bool { return true })
	} else {
		for start := pos; start >= 0 && !ok; start-- {
			ok = m.match(n.subs[0], start, func(end int) bool { return end == pos })
		}
	}
	if n.negate {
		copy(m.caps, saved)
		if ok {
			return false
		}
		return k(pos)
	}
	if !ok {
		copy(m.caps, saved)
		return false
	}
	if k(pos) {
		return true
	}
	copy(m.caps, saved)
	return false
}

func (m *matcher) backref(n *node, pos int, k func(int) bool) bool {
	idx := n.ref
	if n.refName != "" {
		idx = m.re.names[n.refName]
	}
	s, e := m.caps[idx*2], m.caps[idx*2+1]
	if s < 0 || e < 0 {
		return k(pos)
	}
	length := e - s
	if pos+length > len(m.input) {
		return false
	}
	for i := 0; i < length; i++ {
		if !m.unitEqual(m.input[s+i], m.input[pos+i]) {
			return false
		}
	}
	return k(pos + length)
}

func (m *matcher) clearCaptures(n *node) {
	for c := n.capFirst; c <= n.capLast; c++ {
		m.caps[c*2], m.caps[c*2+1] = -1, -1
	}
}

// repeat matches the remaining iterations of a quantified atom; count is the number done. Once the minimum is
// reached, an iteration that consumes nothing fails, which ends the loop.
func (m *matcher) repeat(n *node, pos, count int, k func(int) bool) bool {
	if n.max >= 0 && count >= n.max {
		return k(pos)
	}
	try := func() bool {
		saved := make([]int, 0, 2*(n.capLast-n.capFirst+1))
		for c := n.capFirst; c <= n.capLast; c++ {
			saved = append(saved, m.caps[c*2], m.caps[c*2+1])
		}
		m.clearCaptures(n)
		if m.match(n.subs[0], pos, func(next int) bool {
			if next == pos && count >= n.min {
				return false
			}
			return m.repeat(n, next, count+1, k)
		}) {
			return true
		}
		for i, c := 0, n.capFirst; c <= n.capLast; i, c = i+2, c+1 {
			m.caps[c*2], m.caps[c*2+1] = saved[i], saved[i+1]
		}
		return false
	}
	if count < n.min {
		return try()
	}
	if n.greedy {
		return try() || k(pos)
	}
	return k(pos) || try()
}

// execAt runs the pattern anchored at start and returns the capture positions or nil.
func (re *Regexp) execAt(input []uint16, start int) []int {
	m := &matcher{re: re, input: input, caps: make([]int, 2*(re.ncap+1))}
	for i := range m.caps {
		m.caps[i] = -1
	}
	if m.match(re.root, start, func(end int) bool {
		m.caps[0], m.caps[1] = start, end
		return true
	}) {
		return m.caps
	}
	return nil
}
