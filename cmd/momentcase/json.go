package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"

	"github.com/moment/moment/internal/js"
)

// readJSON decodes one JSON document into JavaScript values: objects become *js.Object with their key order,
// arrays []any, numbers float64.
func readJSON(r io.Reader) (any, error) {
	dec := json.NewDecoder(r)
	dec.UseNumber()
	v, err := readValue(dec)
	if err != nil {
		return nil, err
	}
	switch _, err := dec.Token(); {
	case err == io.EOF:
		return v, nil
	case err != nil:
		return nil, fmt.Errorf("reading after the case: %w", err)
	}
	return nil, fmt.Errorf("trailing data after the case")
}

func readValue(dec *json.Decoder) (any, error) {
	tok, err := dec.Token()
	if err != nil {
		return nil, err
	}
	switch t := tok.(type) {
	case json.Delim:
		switch t {
		case '{':
			o := js.NewObject()
			for dec.More() {
				key, err := dec.Token()
				if err != nil {
					return nil, err
				}
				v, err := readValue(dec)
				if err != nil {
					return nil, err
				}
				o.Set(key.(string), v)
			}
			_, err := dec.Token()
			return o, err
		case '[':
			list := []any{}
			for dec.More() {
				v, err := readValue(dec)
				if err != nil {
					return nil, err
				}
				list = append(list, v)
			}
			_, err := dec.Token()
			return list, err
		}
	case json.Number:
		// A range error only reports overflow to ±Inf or underflow to ±0, which is what JSON.parse returns.
		f, _ := strconv.ParseFloat(string(t), 64)
		return f, nil
	case string, bool, nil:
		return t, nil
	}
	return nil, fmt.Errorf("unexpected JSON token %v", tok)
}

// writeJSON appends the JSON text of an encoded value: nil, bool, float64 (finite), string, []any and
// *js.Object.
func writeJSON(b *bytes.Buffer, v any) {
	switch x := v.(type) {
	case nil:
		b.WriteString("null")
	case bool:
		b.WriteString(strconv.FormatBool(x))
	case float64:
		if math.IsNaN(x) || math.IsInf(x, 0) {
			b.WriteString("null")
			return
		}
		b.WriteString(js.NumberToString(x))
	case string:
		writeString(b, x)
	case []any:
		b.WriteByte('[')
		for i, e := range x {
			if i > 0 {
				b.WriteByte(',')
			}
			writeJSON(b, e)
		}
		b.WriteByte(']')
	case *js.Object:
		b.WriteByte('{')
		for i, k := range x.Keys() {
			if i > 0 {
				b.WriteByte(',')
			}
			writeString(b, k)
			b.WriteByte(':')
			writeJSON(b, x.Get(k))
		}
		b.WriteByte('}')
	}
}

func writeString(b *bytes.Buffer, s string) {
	var sb strings.Builder
	enc := json.NewEncoder(&sb)
	enc.SetEscapeHTML(false)
	// Encoding a string into a strings.Builder cannot fail.
	_ = enc.Encode(s)
	b.WriteString(strings.TrimSuffix(sb.String(), "\n"))
}
