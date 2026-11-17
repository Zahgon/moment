// Package js reproduces the parts of the JavaScript language and runtime that moment's results depend on: the
// value model (undefined, plain objects with their property order), number conversion and printing, UTF-16
// string operations and case mapping, and the Date object with local time as Node.js computes it.
package js

import (
	"math"
	"strconv"
)

// Undefined is the JavaScript undefined value; null is represented by Go nil.
type Undefined struct{}

// Undef is the single undefined value.
var Undef = Undefined{}

// IsUndefined reports whether v is undefined.
func IsUndefined(v any) bool {
	_, ok := v.(Undefined)
	return ok
}

// IsNullish reports whether v is null or undefined (the `v == null` test of JavaScript).
func IsNullish(v any) bool {
	return v == nil || IsUndefined(v)
}

// Func is a JavaScript function value: it receives `this` and the call's arguments.
type Func struct {
	Name string
	Call func(this any, args []any) any
}

// Object is a plain JavaScript object: string keys in property order and their values.
type Object struct {
	keys []string
	vals map[string]any
}

// NewObject returns an empty plain object.
func NewObject() *Object {
	return &Object{vals: map[string]any{}}
}

// Obj builds a plain object from key, value pairs given in insertion order.
func Obj(kv ...any) *Object {
	o := NewObject()
	for i := 0; i+1 < len(kv); i += 2 {
		o.Set(kv[i].(string), kv[i+1])
	}
	return o
}

// arrayIndex reports whether key is an array index ("0" to "4294967294" without leading zeros), the keys
// JavaScript enumerates first and in ascending numeric order.
func arrayIndex(key string) (uint64, bool) {
	if key == "" || len(key) > 10 || (len(key) > 1 && key[0] == '0') {
		return 0, false
	}
	n, err := strconv.ParseUint(key, 10, 64)
	if err != nil || n >= math.MaxUint32 {
		return 0, false
	}
	return n, true
}

// Set creates or replaces a property, keeping JavaScript property order.
func (o *Object) Set(key string, v any) {
	if _, ok := o.vals[key]; !ok {
		if n, isIndex := arrayIndex(key); isIndex {
			pos := 0
			for pos < len(o.keys) {
				m, other := arrayIndex(o.keys[pos])
				if !other || m > n {
					break
				}
				pos++
			}
			o.keys = append(o.keys, "")
			copy(o.keys[pos+1:], o.keys[pos:])
			o.keys[pos] = key
		} else {
			o.keys = append(o.keys, key)
		}
	}
	o.vals[key] = v
}

// Get returns a property value, or undefined when the object has no such own property.
func (o *Object) Get(key string) any {
	if v, ok := o.vals[key]; ok {
		return v
	}
	return Undef
}

// Has reports whether key is an own property.
func (o *Object) Has(key string) bool {
	_, ok := o.vals[key]
	return ok
}

// Delete removes an own property.
func (o *Object) Delete(key string) {
	if _, ok := o.vals[key]; !ok {
		return
	}
	delete(o.vals, key)
	for i, k := range o.keys {
		if k == key {
			o.keys = append(o.keys[:i], o.keys[i+1:]...)
			break
		}
	}
}

// Keys returns the own property keys in JavaScript order.
func (o *Object) Keys() []string {
	return append([]string(nil), o.keys...)
}

// Len is the number of own properties.
func (o *Object) Len() int {
	return len(o.keys)
}

// Clone returns a shallow copy (the result of `extend({}, o)`).
func (o *Object) Clone() *Object {
	c := NewObject()
	for _, k := range o.keys {
		c.Set(k, o.vals[k])
	}
	return c
}

// Extend copies every own property of b onto o, in b's order, and returns o.
func (o *Object) Extend(b *Object) *Object {
	if b == nil {
		return o
	}
	for _, k := range b.keys {
		o.Set(k, b.vals[k])
	}
	return o
}

// Truthy is JavaScript's ToBoolean.
func Truthy(v any) bool {
	switch x := v.(type) {
	case nil, Undefined:
		return false
	case bool:
		return x
	case float64:
		return x != 0 && !math.IsNaN(x)
	case string:
		return x != ""
	}
	return true
}

// StrictEquals is the === operator for the value kinds moment compares.
func StrictEquals(a, b any) bool {
	switch x := a.(type) {
	case nil:
		return b == nil
	case Undefined:
		return IsUndefined(b)
	case float64:
		y, ok := b.(float64)
		return ok && x == y
	case string:
		y, ok := b.(string)
		return ok && x == y
	case bool:
		y, ok := b.(bool)
		return ok && x == y
	case []any, []string:
		return false
	}
	switch b.(type) {
	case []any, []string:
		return false
	}
	return a == b
}
