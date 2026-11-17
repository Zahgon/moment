// Package jsre is an ECMAScript regular expression engine (non-unicode mode) that matches on UTF-16 code units,
// so patterns written for JavaScript, and patterns built at run time from locale strings, behave as in Node.js.
package jsre

import (
	"errors"
	"strings"
	"unicode/utf16"
)

// Regexp is a compiled ECMAScript regular expression. LastIndex is the lastIndex property used by Exec and Test
// on global or sticky expressions.
type Regexp struct {
	source    string
	flags     string
	root      *node
	ncap      int
	names     map[string]int
	global    bool
	icase     bool
	multiline bool
	dotAll    bool
	sticky    bool
	LastIndex int
}

// Match is the result of a successful match: Index is the UTF-16 index of the match, Groups[0] the matched text
// and Groups[i] capture i; Defined[i] is false for a capture that did not participate (undefined in JavaScript).
type Match struct {
	Index   int
	Input   string
	Groups  []string
	Defined []bool
}

// Compile parses source with flags as `new RegExp(source, flags)` does. The error carries the SyntaxError
// message of V8, such as "Invalid regular expression: /(/: Unterminated group".
func Compile(source, flags string) (*Regexp, error) {
	re := &Regexp{source: source}
	seen := map[rune]bool{}
	for _, f := range flags {
		if seen[f] || !strings.ContainsRune("dgimsuvy", f) || f == 'u' || f == 'v' {
			return nil, errors.New("Invalid flags supplied to RegExp constructor '" + flags + "'")
		}
		seen[f] = true
	}
	for _, f := range "dgimsy" {
		if seen[f] {
			re.flags += string(f)
		}
	}
	re.global, re.icase, re.multiline = seen['g'], seen['i'], seen['m']
	re.dotAll, re.sticky = seen['s'], seen['y']
	root, ncap, names, err := parse(source)
	if err != nil {
		return nil, errors.New(errorMessage(source, flags, err.Error()))
	}
	re.root, re.ncap, re.names = root, ncap, names
	return re, nil
}

// MustCompile is Compile for patterns known to be valid; it panics on a syntax error.
func MustCompile(source, flags string) *Regexp {
	re, err := Compile(source, flags)
	if err != nil {
		panic(err)
	}
	return re
}

// Source is RegExp.prototype.source: "/" outside classes and line terminators are escaped.
func (re *Regexp) Source() string {
	if re.source == "" {
		return "(?:)"
	}
	var b strings.Builder
	inClass := false
	units := []rune(re.source)
	for i := 0; i < len(units); i++ {
		c := units[i]
		switch {
		case c == '\\':
			b.WriteRune(c)
			if i+1 < len(units) {
				i++
				writeEscapedTerminator(&b, units[i], true)
			}
			continue
		case c == '[':
			inClass = true
		case c == ']':
			inClass = false
		case c == '/' && !inClass:
			b.WriteString(`\/`)
			continue
		}
		writeEscapedTerminator(&b, c, false)
	}
	return b.String()
}

func writeEscapedTerminator(b *strings.Builder, c rune, escaped bool) {
	prefix := `\`
	if escaped {
		prefix = ""
	}
	switch c {
	case '\n':
		b.WriteString(prefix + "n")
	case '\r':
		b.WriteString(prefix + "r")
	case 0x2028:
		b.WriteString(prefix + "u2028")
	case 0x2029:
		b.WriteString(prefix + "u2029")
	default:
		b.WriteRune(c)
	}
}

// Flags is RegExp.prototype.flags.
func (re *Regexp) Flags() string { return re.flags }

// Global reports the g flag.
func (re *Regexp) Global() bool { return re.global }

// ExecUTF16 searches s from index from and returns the capture positions [start0, end0, start1, end1, ...] in
// UTF-16 units, -1 for a capture that did not participate, or nil when there is no match.
func (re *Regexp) ExecUTF16(s []uint16, from int) []int {
	if from < 0 {
		from = 0
	}
	for start := from; start <= len(s); start++ {
		if caps := re.execAt(s, start); caps != nil {
			return caps
		}
		if re.sticky {
			break
		}
	}
	return nil
}

func toMatch(u []uint16, caps []int, input string) *Match {
	n := len(caps) / 2
	m := &Match{Index: caps[0], Input: input, Groups: make([]string, n), Defined: make([]bool, n)}
	for i := 0; i < n; i++ {
		if caps[2*i] >= 0 && caps[2*i+1] >= 0 {
			m.Groups[i] = string(utf16.Decode(u[caps[2*i]:caps[2*i+1]]))
			m.Defined[i] = true
		}
	}
	return m
}

func (re *Regexp) stateful() bool { return re.global || re.sticky }

// Exec is RegExp.prototype.exec, honouring and updating LastIndex for global and sticky expressions.
func (re *Regexp) Exec(s string) *Match {
	u := utf16.Encode([]rune(s))
	from := 0
	if re.stateful() {
		from = re.LastIndex
		if from > len(u) {
			re.LastIndex = 0
			return nil
		}
	}
	caps := re.ExecUTF16(u, from)
	if caps == nil {
		if re.stateful() {
			re.LastIndex = 0
		}
		return nil
	}
	if re.stateful() {
		re.LastIndex = caps[1]
	}
	return toMatch(u, caps, s)
}

// Test is RegExp.prototype.test, honouring and updating LastIndex for global and sticky expressions.
func (re *Regexp) Test(s string) bool {
	return re.Exec(s) != nil
}

// allMatches returns every match of a global search from index 0.
func (re *Regexp) allMatches(u []uint16) [][]int {
	var out [][]int
	for from := 0; from <= len(u); {
		caps := re.ExecUTF16(u, from)
		if caps == nil {
			break
		}
		out = append(out, caps)
		if caps[1] == caps[0] {
			from = caps[1] + 1
		} else {
			from = caps[1]
		}
		if re.sticky && from != caps[1] {
			break
		}
	}
	re.LastIndex = 0
	return out
}

// Match is String.prototype.match: for a global expression it returns every match text (nil when none) in
// Groups with Index -1; otherwise it is Exec.
func (re *Regexp) Match(s string) *Match {
	if !re.global {
		return re.Exec(s)
	}
	u := utf16.Encode([]rune(s))
	all := re.allMatches(u)
	if len(all) == 0 {
		return nil
	}
	m := &Match{Index: -1, Input: s}
	for _, caps := range all {
		m.Groups = append(m.Groups, string(utf16.Decode(u[caps[0]:caps[1]])))
		m.Defined = append(m.Defined, true)
	}
	return m
}

func (re *Regexp) replace(s string, fn func(u []uint16, caps []int) []uint16) string {
	u := utf16.Encode([]rune(s))
	var matches [][]int
	if re.global {
		matches = re.allMatches(u)
	} else {
		from := 0
		if re.sticky {
			from = re.LastIndex
		}
		if caps := re.ExecUTF16(u, from); caps != nil {
			matches = [][]int{caps}
			if re.sticky {
				re.LastIndex = caps[1]
			}
		} else if re.sticky {
			re.LastIndex = 0
		}
	}
	var out []uint16
	last := 0
	for _, caps := range matches {
		out = append(out, u[last:caps[0]]...)
		out = append(out, fn(u, caps)...)
		last = caps[1]
	}
	out = append(out, u[last:]...)
	return string(utf16.Decode(out))
}

// Replace is String.prototype.replace(re, replacement) with the $-patterns of the replacement string.
func (re *Regexp) Replace(s, replacement string) string {
	rep := utf16.Encode([]rune(replacement))
	return re.replace(s, func(u []uint16, caps []int) []uint16 {
		return expand(rep, u, caps, re.names)
	})
}

func expand(rep, u []uint16, caps []int, names map[string]int) []uint16 {
	m := len(caps)/2 - 1
	group := func(i int) []uint16 {
		if caps[2*i] < 0 {
			return nil
		}
		return u[caps[2*i]:caps[2*i+1]]
	}
	var out []uint16
	for i := 0; i < len(rep); i++ {
		c := rep[i]
		if c != '$' || i+1 >= len(rep) {
			out = append(out, c)
			continue
		}
		next := rep[i+1]
		switch {
		case next == '$':
			out = append(out, '$')
			i++
		case next == '&':
			out = append(out, group(0)...)
			i++
		case next == '`':
			out = append(out, u[:caps[0]]...)
			i++
		case next == '\'':
			out = append(out, u[caps[1]:]...)
			i++
		case next >= '0' && next <= '9':
			d1 := int(next - '0')
			if i+2 < len(rep) && rep[i+2] >= '0' && rep[i+2] <= '9' {
				if d2 := d1*10 + int(rep[i+2]-'0'); d2 >= 1 && d2 <= m {
					out = append(out, group(d2)...)
					i += 2
					continue
				}
			}
			if d1 >= 1 && d1 <= m {
				out = append(out, group(d1)...)
				i++
				continue
			}
			out = append(out, c)
		case next == '<' && len(names) > 0:
			end := i + 2
			for end < len(rep) && rep[end] != '>' {
				end++
			}
			if end >= len(rep) {
				out = append(out, c)
				continue
			}
			if idx, ok := names[string(utf16.Decode(rep[i+2:end]))]; ok {
				out = append(out, group(idx)...)
			}
			i = end
		default:
			out = append(out, c)
		}
	}
	return out
}

// ReplaceFunc is String.prototype.replace(re, fn): fn receives each match and returns its replacement.
func (re *Regexp) ReplaceFunc(s string, fn func(m *Match) string) string {
	return re.replace(s, func(u []uint16, caps []int) []uint16 {
		return utf16.Encode([]rune(fn(toMatch(u, caps, s))))
	})
}

// Split is String.prototype.split(re): captures are spliced into the result.
func (re *Regexp) Split(s string) []string {
	u := utf16.Encode([]rune(s))
	if len(u) == 0 {
		if re.execAt(u, 0) != nil {
			return []string{}
		}
		return []string{s}
	}
	var out []string
	p, q := 0, 0
	for q < len(u) {
		caps := re.execAt(u, q)
		if caps == nil || caps[1] == p || caps[1] > len(u) {
			q++
			continue
		}
		out = append(out, string(utf16.Decode(u[p:q])))
		for i := 1; i < len(caps)/2; i++ {
			if caps[2*i] >= 0 {
				out = append(out, string(utf16.Decode(u[caps[2*i]:caps[2*i+1]])))
			} else {
				out = append(out, "")
			}
		}
		p = caps[1]
		q = p
	}
	return append(out, string(utf16.Decode(u[p:])))
}
