package amoeba

import (
	"amoeba/gjson"
	"bytes"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"unicode/utf8"
)

// A Number represents a JSON number literal.
type Number string

var numberType = reflect.TypeOf(Number(""))

func Marshal(v *Schema, rg string) ([]byte, error) {
	e := newEncodeState()

	err := e.marshal(v, rg)
	if err != nil {
		return nil, err
	}
	buf := append([]byte(nil), e.Bytes()...)

	encodeStatePool.Put(e)

	return buf, nil
}

type UnsupportedTypeError struct {
	Type Type
}

func (e *UnsupportedTypeError) Error() string {
	return fmt.Sprintf("json: unsupported type: %+v", e.Type)
}

type UnsupportedValueError struct {
	Value *Schema
	Str   string
}

func (e *UnsupportedValueError) Error() string {
	return "json: unsupported value: " + e.Str
}

var hex = "0123456789abcdef"

// An encodeState encodes JSON into a bytes.Buffer.
type encodeState struct {
	bytes.Buffer // accumulated output
	scratch      [64]byte
}

var encodeStatePool sync.Pool

func newEncodeState() *encodeState {
	if v := encodeStatePool.Get(); v != nil {
		e := v.(*encodeState)
		e.Reset()
		return e
	}
	return &encodeState{}
}

// jsonError is an error wrapper type for internal use only.
// Panics with errors are wrapped in jsonError so that the top-level recover
// can distinguish intentional panics from this package.
type jsonError struct{ error }

func (e *encodeState) marshal(v *Schema, rg string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			if je, ok := r.(jsonError); ok {
				err = je.error
			} else {
				panic(r)
			}
		}
	}()
	e.reflectValue(v, rg)
	return nil
}

// error aborts the encoding by panicking with err wrapped in jsonError.
func (e *encodeState) error(err error) {
	panic(jsonError{err})
}

func (e *encodeState) reflectValue(v *Schema, rg string) {
	valueEncoder(v)(e, v, rg)
}

type encoderFunc func(e *encodeState, v *Schema, rg string)

var encoderCache sync.Map // map[reflect.Type]encoderFunc

func valueEncoder(v *Schema) encoderFunc {
	return typeEncoder(v)
}

func typeEncoder(t *Schema) encoderFunc {
	if fi, ok := encoderCache.Load(t); ok {
		return fi.(encoderFunc)
	}

	// To deal with recursive types, populate the map with an
	// indirect func before we build it. This type waits on the
	// real func (f) to be ready and then calls it. This indirect
	// func is only used for recursive types.
	var (
		wg sync.WaitGroup
		f  encoderFunc
	)
	wg.Add(1)
	fi, loaded := encoderCache.LoadOrStore(t, encoderFunc(func(e *encodeState, v *Schema, rg string) {
		wg.Wait()
		f(e, v, rg)
	}))
	if loaded {
		return fi.(encoderFunc)
	}

	// Compute the real encoder and replace the indirect func with it.
	f = newTypeEncoder(t)
	wg.Done()
	encoderCache.Store(t, f)
	return f
}

// newTypeEncoder constructs an encoderFunc for a type.
func newTypeEncoder(t *Schema) encoderFunc {
	switch t.kind() {
	case TypeBoolean:
		return boolEncoder
	case TypeInt:
		return intEncoder
	case TypeFloat:
		return float64Encoder
	case TypeString:
		return stringEncoder
	case TypeStruct:
		return newStructEncoder(t)
	case TypeMap:
		return newMapEncoder(t)
	case TypeArray:
		return newArrayEncoder(t)
	default:
		return unsupportedTypeEncoder
	}
}

func boolEncoder(e *encodeState, v *Schema, rg string) {
	value := gjson.Get(rg, v.Value).Value()
	b, ok := value.(bool)
	if !ok {
		// TODO Convert Error
	}

	if b {
		e.WriteString("true")
	} else {
		e.WriteString("false")
	}
}

func intEncoder(e *encodeState, v *Schema, rg string) {
	value := gjson.Get(rg, v.Value)
	b := strconv.AppendInt(e.scratch[:0], value.Int(), 10)
	e.Write(b)
}

type floatEncoder int // number of bits

func (bits floatEncoder) encode(e *encodeState, v *Schema, rs string) {
	f := gjson.Get(rs, v.Value).Float()

	if math.IsInf(f, 0) || math.IsNaN(f) {
		e.error(&UnsupportedValueError{v, strconv.FormatFloat(f, 'g', -1, int(bits))})
	}

	// Convert as if by ES6 number to string conversion.
	// This matches most other JSON generators.
	// See golang.org/issue/6384 and golang.org/issue/14135.
	// Like fmt %g, but the exponent cutoffs are different
	// and exponents themselves are not padded to two digits.
	b := e.scratch[:0]
	abs := math.Abs(f)
	fmt := byte('f')
	// Note: Must use float32 comparisons for underlying float32 value to get precise cutoffs right.
	if abs != 0 {
		if bits == 64 && (abs < 1e-6 || abs >= 1e21) || bits == 32 && (float32(abs) < 1e-6 || float32(abs) >= 1e21) {
			fmt = 'e'
		}
	}
	b = strconv.AppendFloat(b, f, fmt, -1, int(bits))
	if fmt == 'e' {
		// clean up e-09 to e-9
		n := len(b)
		if n >= 4 && b[n-4] == 'e' && b[n-3] == '-' && b[n-2] == '0' {
			b[n-2] = b[n-1]
			b = b[:n-1]
		}
	}

	e.Write(b)
}

var (
	float32Encoder = (floatEncoder(32)).encode
	float64Encoder = (floatEncoder(64)).encode
)

func formatArr(value string, rs string) string {
	strArr := strings.Split(value, ",")

	// 单个，不需要
	if len(strArr) == 1 {
		vTrim := strings.TrimSpace(strArr[0])
		value := gjson.Get(rs, vTrim).String()
		return value
	}

	// 多个，有格式的场景
	valArr := make([]interface{}, 0)
	for _, v := range strArr[1:] {
		vTrim := strings.TrimSpace(v)
		value := gjson.Get(rs, vTrim).String()
		valArr = append(valArr, value)
	}
	result := fmt.Sprintf(strArr[0], valArr...)
	return result
}

func stringEncoder(e *encodeState, v *Schema, rs string) {
	result := formatArr(v.Value, rs)
	e.string(result, true)
}

func unsupportedTypeEncoder(e *encodeState, v *Schema, _ string) {
	e.error(&UnsupportedTypeError{v.kind()})
}

type structEncoder struct {
	fields []*Schema
}

func (se structEncoder) encode(e *encodeState, v *Schema, rs string) {
	next := byte('{')
	for key, s := range v.Properties {
		//if isEmptyValue(*s) {
		//	continue
		//}

		e.WriteByte(next)
		next = ','

		e.WriteString(`"` + key + `":`)

		s.encoder(e, s, rs)
	}
	if next == '{' {
		e.WriteString("{}")
	} else {
		e.WriteByte('}')
	}
}

func newStructEncoder(t *Schema) encoderFunc {
	se := structEncoder{fields: cachedTypeFields(t)}
	return se.encode
}

type mapEncoder struct {
	elemEnc encoderFunc
}

func (mp mapEncoder) encode(e *encodeState, v *Schema, rs string) {
	e.WriteByte('{')
	i := 0
	mapp := gjson.Get(rs, v.Value).Map()
	if len(mapp) == 0 {
		// TODO DEAL ERROR
	}

	for _, vv := range mapp {
		if i > 0 {
			e.WriteByte(',')
		}
		i++

		newKey := formatArr(v.Key, vv.Raw)
		e.string(newKey, false)
		e.WriteByte(':')
		mp.elemEnc(e, v.Items, vv.Raw)
	}
	e.WriteByte('}')
}

func newMapEncoder(t *Schema) encoderFunc {
	switch t.kind() {
	case TypeInt,
		TypeBoolean, TypeFloat, TypeMap, TypeArray, TypeStruct, TypeString:
	default:
		return unsupportedTypeEncoder
	}
	me := mapEncoder{typeEncoder(t.Items)}
	return me.encode
}

type arrayEncoder struct {
	elemEnc encoderFunc
}

func (ae arrayEncoder) encode(e *encodeState, v *Schema, rs string) {
	e.WriteByte('[')
	arr := gjson.Get(rs, v.Value).Array()
	for i, vv := range arr {
		if i > 0 {
			e.WriteByte(',')
		}
		ae.elemEnc(e, v.Items, vv.Raw)
	}

	e.WriteByte(']')
}

func newArrayEncoder(t *Schema) encoderFunc {
	enc := arrayEncoder{typeEncoder(t.Items)}
	return enc.encode
}

type reflectWithString struct {
	v reflect.Value
	s string
}

func (w *reflectWithString) resolve() error {
	if w.v.Kind() == reflect.String {
		w.s = w.v.String()
		return nil
	}

	switch w.v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		w.s = strconv.FormatInt(w.v.Int(), 10)
		return nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		w.s = strconv.FormatUint(w.v.Uint(), 10)
		return nil
	}
	panic("unexpected map key type")
}

// NOTE: keep in sync with stringBytes below.
func (e *encodeState) string(s string, escapeHTML bool) {
	e.WriteByte('"')
	start := 0
	for i := 0; i < len(s); {
		if b := s[i]; b < utf8.RuneSelf {
			if htmlSafeSet[b] || (!escapeHTML && safeSet[b]) {
				i++
				continue
			}
			if start < i {
				e.WriteString(s[start:i])
			}
			e.WriteByte('\\')
			switch b {
			case '\\', '"':
				e.WriteByte(b)
			case '\n':
				e.WriteByte('n')
			case '\r':
				e.WriteByte('r')
			case '\t':
				e.WriteByte('t')
			default:
				// This encodes bytes < 0x20 except for \t, \n and \r.
				// If escapeHTML is set, it also escapes <, >, and &
				// because they can lead to security holes when
				// user-controlled strings are rendered into JSON
				// and served to some browsers.
				e.WriteString(`u00`)
				e.WriteByte(hex[b>>4])
				e.WriteByte(hex[b&0xF])
			}
			i++
			start = i
			continue
		}
		c, size := utf8.DecodeRuneInString(s[i:])
		if c == utf8.RuneError && size == 1 {
			if start < i {
				e.WriteString(s[start:i])
			}
			e.WriteString(`\ufffd`)
			i += size
			start = i
			continue
		}
		// U+2028 is LINE SEPARATOR.
		// U+2029 is PARAGRAPH SEPARATOR.
		// They are both technically valid characters in JSON strings,
		// but don't work in JSONP, which has to be evaluated as JavaScript,
		// and can lead to security holes there. It is valid JSON to
		// escape them, so we do so unconditionally.
		// See http://timelessrepo.com/json-isnt-a-javascript-subset for discussion.
		if c == '\u2028' || c == '\u2029' {
			if start < i {
				e.WriteString(s[start:i])
			}
			e.WriteString(`\u202`)
			e.WriteByte(hex[c&0xF])
			i += size
			start = i
			continue
		}
		i += size
	}
	if start < len(s) {
		e.WriteString(s[start:])
	}
	e.WriteByte('"')
}

// NOTE: keep in sync with string above.
func (e *encodeState) stringBytes(s []byte, escapeHTML bool) {
	e.WriteByte('"')
	start := 0
	for i := 0; i < len(s); {
		if b := s[i]; b < utf8.RuneSelf {
			if htmlSafeSet[b] || (!escapeHTML && safeSet[b]) {
				i++
				continue
			}
			if start < i {
				e.Write(s[start:i])
			}
			e.WriteByte('\\')
			switch b {
			case '\\', '"':
				e.WriteByte(b)
			case '\n':
				e.WriteByte('n')
			case '\r':
				e.WriteByte('r')
			case '\t':
				e.WriteByte('t')
			default:
				// This encodes bytes < 0x20 except for \t, \n and \r.
				// If escapeHTML is set, it also escapes <, >, and &
				// because they can lead to security holes when
				// user-controlled strings are rendered into JSON
				// and served to some browsers.
				e.WriteString(`u00`)
				e.WriteByte(hex[b>>4])
				e.WriteByte(hex[b&0xF])
			}
			i++
			start = i
			continue
		}
		c, size := utf8.DecodeRune(s[i:])
		if c == utf8.RuneError && size == 1 {
			if start < i {
				e.Write(s[start:i])
			}
			e.WriteString(`\ufffd`)
			i += size
			start = i
			continue
		}
		// U+2028 is LINE SEPARATOR.
		// U+2029 is PARAGRAPH SEPARATOR.
		// They are both technically valid characters in JSON strings,
		// but don't work in JSONP, which has to be evaluated as JavaScript,
		// and can lead to security holes there. It is valid JSON to
		// escape them, so we do so unconditionally.
		// See http://timelessrepo.com/json-isnt-a-javascript-subset for discussion.
		if c == '\u2028' || c == '\u2029' {
			if start < i {
				e.Write(s[start:i])
			}
			e.WriteString(`\u202`)
			e.WriteByte(hex[c&0xF])
			i += size
			start = i
			continue
		}
		i += size
	}
	if start < len(s) {
		e.Write(s[start:])
	}
	e.WriteByte('"')
}

var fieldCache sync.Map // map[reflect.Type]structFields

// cachedTypeFields is like typeFields but uses a cache to avoid repeated work.
func cachedTypeFields(t *Schema) []*Schema {
	if f, ok := fieldCache.Load(t); ok {
		return f.([]*Schema)
	}
	f, _ := fieldCache.LoadOrStore(t, typeFields(t))
	return f.([]*Schema)
}

func typeFields(schema *Schema) []*Schema {
	if schema == nil {
		return nil
	}

	schemaList := make([]*Schema, 0, 0)
	for _, sc := range schema.Properties {
		sc.encoder = typeEncoder(sc)
		schemaList = append(schemaList, sc)
	}

	return schemaList
}
