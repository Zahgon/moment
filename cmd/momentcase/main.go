// Command momentcase runs one moment case: it reads a case from standard input, evaluates its steps against the
// Go moment package and writes one JSON line per output, deprecation, warning and exception, exactly as the
// reference command of the original JavaScript library does. Local time follows the TZ environment variable.
//
//	momentcase < case.json
package main

import (
	"bufio"
	"bytes"
	"errors"
	"log"
	"math"
	"os"
	"strings"

	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
	_ "github.com/moment/moment/locale"
)

// inputError reports a malformed case (an unknown name, variable or step); it is not a JavaScript exception.
type inputError struct{ msg string }

func (e *inputError) Error() string { return e.msg }

type runner struct {
	vars     map[string]any
	out      *bufio.Writer
	writeErr error
}

// emit writes one output line; the first write error is kept and reported when the case ends.
func (r *runner) emit(key string, v any) {
	var b bytes.Buffer
	b.WriteString("{")
	writeString(&b, key)
	b.WriteString(":")
	writeJSON(&b, v)
	b.WriteString("}\n")
	if _, err := r.out.Write(b.Bytes()); err != nil && r.writeErr == nil {
		r.writeErr = err
	}
}

var functions = map[string]func(float64) float64{
	"Math.floor": math.Floor, "Math.ceil": math.Ceil, "Math.round": js.Round, "Math.trunc": math.Trunc,
}

func tagged(o *js.Object) (string, any, bool) {
	keys := o.Keys()
	if len(keys) != 1 || !strings.HasPrefix(keys[0], "$") {
		return "", nil, false
	}
	return keys[0], o.Get(keys[0]), true
}

// decode turns a case argument into a JavaScript value.
func (r *runner) decode(a any) (any, error) {
	switch x := a.(type) {
	case []any:
		out := make([]any, len(x))
		for i, e := range x {
			v, err := r.decode(e)
			if err != nil {
				return nil, err
			}
			out[i] = v
		}
		return out, nil
	case *js.Object:
		if tag, t, ok := tagged(x); ok {
			return r.decodeTag(tag, t)
		}
		o := js.NewObject()
		for _, k := range x.Keys() {
			v, err := r.decode(x.Get(k))
			if err != nil {
				return nil, err
			}
			o.Set(k, v)
		}
		return o, nil
	}
	return a, nil
}

func (r *runner) decodeTag(tag string, t any) (any, error) {
	switch tag {
	case "$undefined":
		return js.Undef, nil
	case "$num":
		return js.StringToNumber(js.ToString(t)), nil
	case "$date":
		if t == nil {
			return js.NewDate(js.NaN), nil
		}
		return js.NewDate(js.ToNumber(t)), nil
	case "$ref":
		v, ok := r.vars[js.ToString(t)]
		if !ok {
			return nil, &inputError{"unknown variable " + js.ToString(t)}
		}
		return v, nil
	case "$x":
		return r.evaluate(t)
	case "$const":
		switch t {
		case "moment.ISO_8601":
			return moment.ISO8601, nil
		case "moment.RFC_2822":
			return moment.RFC2822, nil
		}
		return nil, &inputError{"unknown constant " + js.ToString(t)}
	case "$fn":
		name := js.ToString(t)
		f, ok := functions[name]
		if !ok {
			return nil, &inputError{"unknown function " + name}
		}
		return &js.Func{Name: name, Call: func(_ any, args []any) any {
			if len(args) == 0 {
				return js.NaN
			}
			return f(js.ToNumber(args[0]))
		}}, nil
	}
	o := js.NewObject()
	v, err := r.decode(t)
	o.Set(tag, v)
	return o, err
}

func (r *runner) evaluate(x any) (any, error) {
	o, ok := x.(*js.Object)
	if !ok || (!o.Has("call") && !o.Has("get") && !o.Has("set")) {
		return r.decode(x)
	}
	if o.Has("get") {
		v, err := moment.GetProperty(js.ToString(o.Get("get")))
		return v, nameError(err)
	}
	if o.Has("set") {
		v, err := r.decode(o.Get("value"))
		if err != nil {
			return nil, err
		}
		return js.Undef, nameError(moment.SetProperty(js.ToString(o.Get("set")), v))
	}
	var args []any
	if list, ok := o.Get("args").([]any); ok {
		for _, a := range list {
			v, err := r.decode(a)
			if err != nil {
				return nil, err
			}
			args = append(args, v)
		}
	}
	name := js.ToString(o.Get("call"))
	if !o.Has("on") {
		v, err := moment.CallStatic(name, args)
		return v, nameError(err)
	}
	recv, err := r.decode(o.Get("on"))
	if err != nil {
		return nil, err
	}
	return moment.CallMethod(recv, name, args)
}

func nameError(err error) error {
	var unknown *moment.ErrUnknownName
	if errors.As(err, &unknown) {
		return &inputError{unknown.Error()}
	}
	return err
}

func (r *runner) run(c *js.Object) error {
	now := js.ToNumber(c.Get("now"))
	moment.Now = func() float64 { return now }
	moment.SuppressDeprecationWarnings = true
	moment.DeprecationHandler = func(name any, msg string) { r.emit("deprecation", []any{name, msg}) }
	moment.Warn = func(msg string) { r.emit("warn", msg) }
	steps, _ := c.Get("steps").([]any)
	for _, s := range steps {
		step, ok := s.(*js.Object)
		if !ok {
			return &inputError{"bad step"}
		}
		var err error
		switch {
		case step.Has("let"):
			var v any
			v, err = r.evaluate(step.Get("x"))
			r.vars[js.ToString(step.Get("let"))] = v
		case step.Has("do"):
			_, err = r.evaluate(step.Get("do"))
		case step.Has("out"):
			var v any
			if v, err = r.evaluate(step.Get("out")); err == nil {
				r.emit("out", encode(v))
			}
		default:
			err = &inputError{"bad step"}
		}
		var thrown *moment.ThrowError
		if errors.As(err, &thrown) {
			r.emit("throw", []any{thrown.Name, thrown.Message})
			return r.writeErr
		}
		if err != nil {
			return err
		}
	}
	return r.writeErr
}

func main() {
	out := bufio.NewWriter(os.Stdout)
	c, err := readJSON(bufio.NewReader(os.Stdin))
	if err == nil {
		obj, ok := c.(*js.Object)
		if !ok {
			err = &inputError{"the case is not a JSON object"}
		} else {
			err = (&runner{vars: map[string]any{}, out: out}).run(obj)
		}
	}
	if flushErr := out.Flush(); err == nil {
		err = flushErr
	}
	if err != nil {
		log.SetFlags(0)
		log.SetPrefix("momentcase: ")
		log.Print(err)
		os.Exit(2)
	}
}

// encode converts a JavaScript value to its output form.
func encode(v any) any {
	switch x := v.(type) {
	case js.Undefined:
		return js.Obj("$undefined", true)
	case nil, bool, string:
		return x
	case float64:
		switch {
		case math.IsNaN(x):
			return js.Obj("$num", "NaN")
		case math.IsInf(x, 1):
			return js.Obj("$num", "Infinity")
		case math.IsInf(x, -1):
			return js.Obj("$num", "-Infinity")
		case x == 0 && math.Signbit(x):
			return js.Obj("$num", "-0")
		}
		return x
	case *js.Func:
		return js.Obj("$function", true)
	case *moment.Moment:
		return js.Obj("$moment", []any{encode(x.ValueOf()), encode(x.UTCOffset()), x.IsUtcOffset(), encode(x.Locale())})
	case *moment.Duration:
		values := []any{x.Years(), x.Months(), x.Days(), x.Hours(), x.Minutes(), x.Seconds(), x.Milliseconds(),
			x.AsMilliseconds(), x.IsValid(), x.Locale()}
		for i, e := range values {
			values[i] = encode(e)
		}
		return js.Obj("$duration", values)
	case *js.Date:
		return js.Obj("$date", encode(x.Time()))
	case *jsre.Regexp:
		return js.Obj("$regexp", []any{x.Source(), x.Flags()})
	case *moment.Locale:
		return js.Obj("$locale", encode(x.Abbr()))
	case []string:
		out := make([]any, len(x))
		for i, s := range x {
			out[i] = s
		}
		return out
	case []any:
		out := make([]any, len(x))
		for i, e := range x {
			out[i] = encode(e)
		}
		return out
	case *js.Object:
		o := js.NewObject()
		for _, k := range x.Keys() {
			o.Set(k, encode(x.Get(k)))
		}
		return o
	}
	return js.Obj("$function", true)
}
