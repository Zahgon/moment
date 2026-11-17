package moment

import (
	"math"
	"strings"

	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

// Eras is Locale#eras(): the locale's eras (English ones by default) with since and until resolved, in place,
// to the start of their local day in milliseconds.
func (l *Locale) Eras() []any {
	eras, ok := l.own.Get("_eras").([]any)
	if !ok {
		if en := getLocale("en"); en != nil {
			eras, _ = en.own.Get("_eras").([]any)
		}
	}
	for _, e := range eras {
		era, ok := e.(*js.Object)
		if !ok {
			continue
		}
		if s, ok := era.Get("since").(string); ok {
			era.Set("since", localMoment(s).StartOf("day").ValueOf())
		}
		switch u := era.Get("until").(type) {
		case js.Undefined:
			era.Set("until", math.Inf(1))
		case string:
			era.Set("until", localMoment(u).StartOf("day").ValueOf())
		}
	}
	return eras
}

func eraField(era *js.Object, key string) string { return js.ToUpper(js.ToString(era.Get(key))) }

// ErasParse is Locale#erasParse(eraName, format, strict): the matching era object, or undefined.
func (l *Locale) ErasParse(eraName string, format Value, strict bool) Value {
	upper := js.ToUpper(eraName)
	f := js.ToString(format)
	for _, e := range l.Eras() {
		era, ok := e.(*js.Object)
		if !ok {
			continue
		}
		name, abbr, narrow := eraField(era, "name"), eraField(era, "abbr"), eraField(era, "narrow")
		if strict {
			switch f {
			case "N", "NN", "NNN":
				if abbr == upper {
					return era
				}
			case "NNNN":
				if name == upper {
					return era
				}
			case "NNNNN":
				if narrow == upper {
					return era
				}
			}
		} else if upper == name || upper == abbr || upper == narrow {
			return era
		}
	}
	return js.Undef
}

// ErasConvertYear is Locale#erasConvertYear(era, year): the calendar year of the era's year.
func (l *Locale) ErasConvertYear(era, year Value) Value {
	o, ok := era.(*js.Object)
	if !ok {
		return js.NaN
	}
	dir := -1.0
	if js.ToNumber(o.Get("since")) <= js.ToNumber(o.Get("until")) {
		dir = 1
	}
	base := localMoment(o.Get("since")).Year()
	if js.IsUndefined(year) {
		return base
	}
	return base + float64((js.ToNumber(year)-js.ToNumber(o.Get("offset")))*dir)
}

func (l *Locale) computeErasParse() {
	var names, abbrs, narrows, mixed []string
	for _, e := range l.Eras() {
		era, ok := e.(*js.Object)
		if !ok {
			continue
		}
		name := regexEscape(js.ToString(era.Get("name")))
		abbr := regexEscape(js.ToString(era.Get("abbr")))
		narrow := regexEscape(js.ToString(era.Get("narrow")))
		names = append(names, name)
		abbrs = append(abbrs, abbr)
		narrows = append(narrows, narrow)
		mixed = append(mixed, name, abbr, narrow)
	}
	l.own.Set("_erasRegex", compileOrNil("^("+strings.Join(mixed, "|")+")", "i"))
	l.own.Set("_erasNameRegex", compileOrNil("^("+strings.Join(names, "|")+")", "i"))
	l.own.Set("_erasAbbrRegex", compileOrNil("^("+strings.Join(abbrs, "|")+")", "i"))
	l.own.Set("_erasNarrowRegex", compileOrNil("^("+strings.Join(narrows, "|")+")", "i"))
}

func (l *Locale) erasRegexFor(key string, strict bool) *jsre.Regexp {
	if !l.own.Has(key) {
		l.computeErasParse()
	}
	if strict {
		return regexProp(l.own.Get(key))
	}
	return regexProp(l.own.Get("_erasRegex"))
}

// ErasNameRegex is Locale#erasNameRegex(strict).
func (l *Locale) ErasNameRegex(strict bool) *jsre.Regexp {
	return l.erasRegexFor("_erasNameRegex", strict)
}

// ErasAbbrRegex is Locale#erasAbbrRegex(strict).
func (l *Locale) ErasAbbrRegex(strict bool) *jsre.Regexp {
	return l.erasRegexFor("_erasAbbrRegex", strict)
}

// ErasNarrowRegex is Locale#erasNarrowRegex(strict).
func (l *Locale) ErasNarrowRegex(strict bool) *jsre.Regexp {
	return l.erasRegexFor("_erasNarrowRegex", strict)
}
