package moment

import (
	"reflect"

	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

// Locale is a moment locale. Like the JavaScript object it holds the config's data under "_" + key and the
// config's functions under their own names, next to the parse caches it builds lazily.
type Locale struct {
	own *js.Object
}

var (
	locales        = js.NewObject()
	localeFamilies = map[string][]familyMember{}
	globalLocale   *Locale
	localeModules  = map[string]func(){}
	loadedModules  = map[string]bool{}
)

type familyMember struct {
	name   string
	config *js.Object
}

// RegisterLocale makes a locale module available to the lazy locale loader: define runs, once, the first time
// the locale is needed, and defines it with DefineLocale (the Go form of a file under locale/).
func RegisterLocale(name string, define func()) {
	localeModules[name] = define
}

func isFunc(v Value) bool {
	if _, ok := v.(*js.Func); ok {
		return true
	}
	return v != nil && reflect.ValueOf(v).Kind() == reflect.Func
}

func isPlainObject(v Value) bool {
	_, ok := v.(*js.Object)
	return ok
}

// Prop returns the locale's data property key (the JavaScript this["_" + key]), js.Undef when absent.
func (l *Locale) Prop(key string) Value { return l.own.Get("_" + key) }

// Abbr is the locale's name.
func (l *Locale) Abbr() Value { return l.own.Get("_abbr") }

func (l *Locale) weekSetting(key string) Value {
	if w, ok := l.own.Get("_week").(*js.Object); ok {
		return w.Get(key)
	}
	return js.Undef
}

func (l *Locale) set(config *js.Object) {
	for _, prop := range monthsParseProperties {
		if !config.Has(prop) {
			l.own.Delete("_" + prop)
		}
	}
	for _, prop := range weekdaysParseProperties {
		if !config.Has(prop) {
			l.own.Delete("_" + prop)
		}
	}
	for _, k := range config.Keys() {
		v := config.Get(k)
		if isFunc(v) {
			l.own.Set(k, v)
		} else {
			l.own.Set("_"+k, v)
		}
	}
	l.own.Set("_config", config)
	source := ""
	if re, ok := l.own.Get("_dayOfMonthOrdinalParse").(*jsre.Regexp); ok {
		source = re.Source()
	} else if re, ok := l.own.Get("_ordinalParse").(*jsre.Regexp); ok {
		source = re.Source()
	}
	if re, err := jsre.Compile(source+`|\d{1,2}`, ""); err == nil {
		l.own.Set("_dayOfMonthOrdinalParseLenient", re)
	}
}

func newLocale(config *js.Object) *Locale {
	l := &Locale{own: js.NewObject()}
	if config != nil {
		l.set(config)
	}
	return l
}

func (l *Locale) config() *js.Object {
	c, _ := l.own.Get("_config").(*js.Object)
	return c
}

func mergeConfigs(parent, child *js.Object) *js.Object {
	res := parent.Clone()
	for _, prop := range child.Keys() {
		pv, cv := parent.Get(prop), child.Get(prop)
		switch {
		case isPlainObject(pv) && isPlainObject(cv):
			merged := js.NewObject()
			merged.Extend(pv.(*js.Object))
			merged.Extend(cv.(*js.Object))
			res.Set(prop, merged)
		case !js.IsNullish(cv):
			res.Set(prop, cv)
		default:
			res.Delete(prop)
		}
	}
	for _, prop := range parent.Keys() {
		if !child.Has(prop) && isPlainObject(parent.Get(prop)) {
			copied := js.NewObject()
			if o, ok := res.Get(prop).(*js.Object); ok {
				copied.Extend(o)
			}
			res.Set(prop, copied)
		}
	}
	return res
}

func normalizeLocale(key Value) Value {
	s, ok := key.(string)
	if !ok || s == "" {
		return key
	}
	return js.ReplaceFirst(js.ToLower(s), "_", "-")
}

var saneLocaleName = jsre.MustCompile(`^[a-z0-9]+(?:-[a-z0-9]+)*$`, "")

func localeEntry(name string) (*Locale, bool) {
	if !locales.Has(name) {
		return nil, false
	}
	l, _ := locales.Get(name).(*Locale)
	return l, true
}

// loadLocale returns a defined locale, running its module the first time; unknown sane names are remembered
// as missing.
func loadLocale(name Value) *Locale {
	if l, ok := localeEntry(js.ToString(name)); ok {
		return l
	}
	normalized := js.ToString(normalizeLocale(name))
	if l, ok := localeEntry(normalized); ok {
		return l
	}
	if _, isString := normalizeLocale(name).(string); isString && saneLocaleName.Test(normalized) {
		define, exists := localeModules[normalized]
		if !exists {
			locales.Set(normalized, nil)
		} else if !loadedModules[normalized] {
			loadedModules[normalized] = true
			old := globalLocale.Abbr()
			define()
			GetSetGlobalLocale(old, js.Undef)
		}
	}
	l, _ := localeEntry(normalized)
	return l
}

func commonPrefix(a, b []string) int {
	n := min(len(a), len(b))
	for i := 0; i < n; i++ {
		if a[i] != b[i] {
			return i
		}
	}
	return n
}

func chooseLocale(names []any) *Locale {
	for i := 0; i < len(names); i++ {
		split := js.Split(js.ToString(normalizeLocale(names[i])), "-")
		var next []string
		if i+1 < len(names) {
			if n := normalizeLocale(names[i+1]); js.Truthy(n) {
				next = js.Split(js.ToString(n), "-")
			}
		}
		for j := len(split); j > 0; j-- {
			if l := loadLocale(joinParts(split[:j])); l != nil {
				return l
			}
			if next != nil && len(next) >= j && commonPrefix(split, next) >= j-1 {
				break
			}
		}
	}
	return globalLocale
}

func joinParts(parts []string) string {
	out := ""
	for i, p := range parts {
		if i > 0 {
			out += "-"
		}
		out += p
	}
	return out
}

// GetSetGlobalLocale is moment.locale(key, values): it defines (with values) or selects the global locale and
// returns the global locale's name; an unknown key keeps the current locale and writes a warning.
func GetSetGlobalLocale(key, values Value) Value {
	if js.Truthy(key) {
		var data *Locale
		if js.IsUndefined(values) {
			data = getLocale(key)
		} else if l, ok := DefineLocaleValue(js.ToString(key), values).(*Locale); ok {
			data = l
		}
		if data != nil {
			globalLocale = data
		} else {
			Warn("Locale " + js.ToString(key) + " not found. Did you forget to load it?")
		}
	}
	return globalLocale.Abbr()
}

// DefineLocale is moment.defineLocale(name, config): it creates the locale from config merged over its parent
// (the base config by default) and makes it the global locale.
func DefineLocale(name string, config *js.Object) *Locale {
	l, _ := DefineLocaleValue(name, config).(*Locale)
	return l
}

// DefineLocaleValue is moment.defineLocale with any config value: null removes the locale and returns null.
func DefineLocaleValue(name string, config Value) Value {
	cfg, ok := config.(*js.Object)
	if !ok {
		locales.Delete(name)
		return nil
	}
	parentConfig := baseConfig
	cfg.Set("abbr", name)
	if existing, ok := localeEntry(name); ok && existing != nil {
		deprecateSimple("defineLocaleOverride", "use moment.updateLocale(localeName, config) to change an "+
			"existing locale. moment.defineLocale(localeName, config) should only be used for creating a new "+
			"locale See http://momentjs.com/guides/#/warnings/define-locale/ for more info.")
		parentConfig = existing.config()
	} else if parent := cfg.Get("parentLocale"); !js.IsNullish(parent) {
		parentName := js.ToString(parent)
		if p, ok := localeEntry(parentName); ok && p != nil {
			parentConfig = p.config()
		} else if p := loadLocale(parent); p != nil {
			parentConfig = p.config()
		} else {
			localeFamilies[parentName] = append(localeFamilies[parentName], familyMember{name, cfg})
			return nil
		}
	}
	locales.Set(name, newLocale(mergeConfigs(parentConfig, cfg)))
	if family, ok := localeFamilies[name]; ok {
		for _, x := range family {
			DefineLocaleValue(x.name, x.config)
		}
	}
	GetSetGlobalLocale(name, js.Undef)
	return locales.Get(name)
}

// UpdateLocale is moment.updateLocale(name, config); a null config undoes the last update.
func UpdateLocale(name string, config Value) Value {
	tmp := loadLocale(name)
	if tmp != nil {
		name = js.ToString(tmp.Abbr())
	}
	if cfg, ok := config.(*js.Object); ok {
		current, _ := localeEntry(name)
		if current != nil && !js.IsNullish(current.own.Get("parentLocale")) {
			current.set(mergeConfigs(current.config(), cfg))
		} else {
			parentConfig := baseConfig
			if tmp != nil {
				parentConfig = tmp.config()
			}
			merged := mergeConfigs(parentConfig, cfg)
			if tmp == nil {
				merged.Set("abbr", name)
			}
			l := newLocale(merged)
			if current != nil {
				l.own.Set("parentLocale", current)
			} else {
				l.own.Set("parentLocale", js.Undef)
			}
			locales.Set(name, l)
		}
		GetSetGlobalLocale(name, js.Undef)
	} else if current, _ := localeEntry(name); current != nil {
		if parent, ok := current.own.Get("parentLocale").(*Locale); ok {
			locales.Set(name, parent)
			if js.StrictEquals(name, GetSetGlobalLocale(js.Undef, js.Undef)) {
				GetSetGlobalLocale(name, js.Undef)
			}
		} else {
			locales.Delete(name)
		}
	}
	if l, ok := localeEntry(name); ok {
		if l == nil {
			return nil
		}
		return l
	}
	return js.Undef
}

// getLocale resolves a locale key (a name, an array of names, or a moment whose locale is used); nil means not
// found.
func getLocale(key Value) *Locale {
	switch x := key.(type) {
	case *Moment:
		if x.locale != nil && js.Truthy(x.locale.Abbr()) {
			key = x.locale.Abbr()
		}
	case *Duration:
		if x.locale != nil && js.Truthy(x.locale.Abbr()) {
			key = x.locale.Abbr()
		}
	}
	if !js.Truthy(key) {
		return globalLocale
	}
	if !isArray(key) {
		if l := loadLocale(key); l != nil {
			return l
		}
		return chooseLocale([]any{key})
	}
	return chooseLocale(asArray(key))
}

// LocaleData is moment.localeData(key).
func LocaleData(key Value) Value {
	if l := getLocale(key); l != nil {
		return l
	}
	return js.Undef
}

// Locales is moment.locales(): every name the registry knows, in JavaScript property order, including names
// that failed to load.
func Locales() []string { return locales.Keys() }

func listGet(format Value, index Value, field string, setter string) Value {
	utc := utcMoment()
	utc.Set(setter, index)
	return globalLocale.callList(field, utc, format)
}

func (l *Locale) callList(field string, m *Moment, format Value) Value {
	switch field {
	case "months":
		return l.Months(m, format)
	case "monthsShort":
		return l.MonthsShort(m, format)
	case "weekdays":
		return l.Weekdays(m, format)
	case "weekdaysShort":
		return l.WeekdaysShort(m)
	}
	return l.WeekdaysMin(m)
}

// ListMonths implements moment.months(format, index) and moment.monthsShort (field "monthsShort").
func ListMonths(field string, format, index Value) Value {
	if _, ok := format.(float64); ok {
		index = format
		format = js.Undef
	}
	if !js.Truthy(format) {
		format = ""
	}
	if !js.IsNullish(index) {
		return listGet(format, index, field, "month")
	}
	out := make([]any, 12)
	for i := range out {
		out[i] = listGet(format, float64(i), field, "month")
	}
	return out
}

// ListWeekdays implements moment.weekdays(localeSorted, format, index), weekdaysShort and weekdaysMin.
func ListWeekdays(field string, localeSorted, format, index Value) Value {
	sorted := false
	if b, ok := localeSorted.(bool); ok {
		sorted = b
		if _, ok := format.(float64); ok {
			index = format
			format = js.Undef
		}
	} else {
		format = localeSorted
		index = format
		if _, ok := format.(float64); ok {
			format = js.Undef
		}
	}
	if !js.Truthy(format) {
		format = ""
	}
	shift := 0.0
	if sorted {
		shift = js.ToNumber(globalLocale.weekSetting("dow"))
	}
	if !js.IsNullish(index) {
		return listGet(format, js.Mod(js.ToNumber(js.Add(index, shift)), 7), field, "day")
	}
	out := make([]any, 7)
	for i := range out {
		out[i] = listGet(format, js.Mod(float64(i)+shift, 7), field, "day")
	}
	return out
}
