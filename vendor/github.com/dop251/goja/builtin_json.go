package goja

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"
	"unicode/utf16"

	"github.com/dop251/goja/unistring"
)

const hex = "0123456789abcdef"

func (r *Runtime) builtinJSON_parse(call FunctionCall) Value ***REMOVED***
	d := json.NewDecoder(bytes.NewBufferString(call.Argument(0).toString().String()))

	value, err := r.builtinJSON_decodeValue(d)
	if err != nil ***REMOVED***
		panic(r.newError(r.global.SyntaxError, err.Error()))
	***REMOVED***

	if tok, err := d.Token(); err != io.EOF ***REMOVED***
		panic(r.newError(r.global.SyntaxError, "Unexpected token at the end: %v", tok))
	***REMOVED***

	var reviver func(FunctionCall) Value

	if arg1 := call.Argument(1); arg1 != _undefined ***REMOVED***
		reviver, _ = arg1.ToObject(r).self.assertCallable()
	***REMOVED***

	if reviver != nil ***REMOVED***
		root := r.NewObject()
		createDataPropertyOrThrow(root, stringEmpty, value)
		return r.builtinJSON_reviveWalk(reviver, root, stringEmpty)
	***REMOVED***

	return value
***REMOVED***

func (r *Runtime) builtinJSON_decodeToken(d *json.Decoder, tok json.Token) (Value, error) ***REMOVED***
	switch tok := tok.(type) ***REMOVED***
	case json.Delim:
		switch tok ***REMOVED***
		case '***REMOVED***':
			return r.builtinJSON_decodeObject(d)
		case '[':
			return r.builtinJSON_decodeArray(d)
		***REMOVED***
	case nil:
		return _null, nil
	case string:
		return newStringValue(tok), nil
	case float64:
		return floatToValue(tok), nil
	case bool:
		if tok ***REMOVED***
			return valueTrue, nil
		***REMOVED***
		return valueFalse, nil
	***REMOVED***
	return nil, fmt.Errorf("Unexpected token (%T): %v", tok, tok)
***REMOVED***

func (r *Runtime) builtinJSON_decodeValue(d *json.Decoder) (Value, error) ***REMOVED***
	tok, err := d.Token()
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return r.builtinJSON_decodeToken(d, tok)
***REMOVED***

func (r *Runtime) builtinJSON_decodeObject(d *json.Decoder) (*Object, error) ***REMOVED***
	object := r.NewObject()
	for ***REMOVED***
		key, end, err := r.builtinJSON_decodeObjectKey(d)
		if err != nil ***REMOVED***
			return nil, err
		***REMOVED***
		if end ***REMOVED***
			break
		***REMOVED***
		value, err := r.builtinJSON_decodeValue(d)
		if err != nil ***REMOVED***
			return nil, err
		***REMOVED***

		object.self._putProp(unistring.NewFromString(key), value, true, true, true)
	***REMOVED***
	return object, nil
***REMOVED***

func (r *Runtime) builtinJSON_decodeObjectKey(d *json.Decoder) (string, bool, error) ***REMOVED***
	tok, err := d.Token()
	if err != nil ***REMOVED***
		return "", false, err
	***REMOVED***
	switch tok := tok.(type) ***REMOVED***
	case json.Delim:
		if tok == '***REMOVED***' ***REMOVED***
			return "", true, nil
		***REMOVED***
	case string:
		return tok, false, nil
	***REMOVED***

	return "", false, fmt.Errorf("Unexpected token (%T): %v", tok, tok)
***REMOVED***

func (r *Runtime) builtinJSON_decodeArray(d *json.Decoder) (*Object, error) ***REMOVED***
	var arrayValue []Value
	for ***REMOVED***
		tok, err := d.Token()
		if err != nil ***REMOVED***
			return nil, err
		***REMOVED***
		if delim, ok := tok.(json.Delim); ok ***REMOVED***
			if delim == ']' ***REMOVED***
				break
			***REMOVED***
		***REMOVED***
		value, err := r.builtinJSON_decodeToken(d, tok)
		if err != nil ***REMOVED***
			return nil, err
		***REMOVED***
		arrayValue = append(arrayValue, value)
	***REMOVED***
	return r.newArrayValues(arrayValue), nil
***REMOVED***

func (r *Runtime) builtinJSON_reviveWalk(reviver func(FunctionCall) Value, holder *Object, name Value) Value ***REMOVED***
	value := nilSafe(holder.get(name, nil))

	if object, ok := value.(*Object); ok ***REMOVED***
		if isArray(object) ***REMOVED***
			length := toLength(object.self.getStr("length", nil))
			for index := int64(0); index < length; index++ ***REMOVED***
				name := asciiString(strconv.FormatInt(index, 10))
				value := r.builtinJSON_reviveWalk(reviver, object, name)
				if value == _undefined ***REMOVED***
					object.delete(name, false)
				***REMOVED*** else ***REMOVED***
					createDataProperty(object, name, value)
				***REMOVED***
			***REMOVED***
		***REMOVED*** else ***REMOVED***
			for _, name := range object.self.stringKeys(false, nil) ***REMOVED***
				value := r.builtinJSON_reviveWalk(reviver, object, name)
				if value == _undefined ***REMOVED***
					object.self.deleteStr(name.string(), false)
				***REMOVED*** else ***REMOVED***
					createDataProperty(object, name, value)
				***REMOVED***
			***REMOVED***
		***REMOVED***
	***REMOVED***
	return reviver(FunctionCall***REMOVED***
		This:      holder,
		Arguments: []Value***REMOVED***name, value***REMOVED***,
	***REMOVED***)
***REMOVED***

type _builtinJSON_stringifyContext struct ***REMOVED***
	r                *Runtime
	stack            []*Object
	propertyList     []Value
	replacerFunction func(FunctionCall) Value
	gap, indent      string
	buf              bytes.Buffer
***REMOVED***

func (r *Runtime) builtinJSON_stringify(call FunctionCall) Value ***REMOVED***
	ctx := _builtinJSON_stringifyContext***REMOVED***
		r: r,
	***REMOVED***

	replacer, _ := call.Argument(1).(*Object)
	if replacer != nil ***REMOVED***
		if isArray(replacer) ***REMOVED***
			length := toLength(replacer.self.getStr("length", nil))
			seen := map[string]bool***REMOVED******REMOVED***
			propertyList := make([]Value, length)
			length = 0
			for index := range propertyList ***REMOVED***
				var name string
				value := replacer.self.getIdx(valueInt(int64(index)), nil)
				switch v := value.(type) ***REMOVED***
				case valueFloat, valueInt, valueString:
					name = value.String()
				case *Object:
					switch v.self.className() ***REMOVED***
					case classNumber, classString:
						name = value.String()
					default:
						continue
					***REMOVED***
				default:
					continue
				***REMOVED***
				if seen[name] ***REMOVED***
					continue
				***REMOVED***
				seen[name] = true
				propertyList[length] = newStringValue(name)
				length += 1
			***REMOVED***
			ctx.propertyList = propertyList[0:length]
		***REMOVED*** else if c, ok := replacer.self.assertCallable(); ok ***REMOVED***
			ctx.replacerFunction = c
		***REMOVED***
	***REMOVED***
	if spaceValue := call.Argument(2); spaceValue != _undefined ***REMOVED***
		if o, ok := spaceValue.(*Object); ok ***REMOVED***
			switch oImpl := o.self.(type) ***REMOVED***
			case *primitiveValueObject:
				switch oImpl.pValue.(type) ***REMOVED***
				case valueInt, valueFloat:
					spaceValue = o.ToNumber()
				***REMOVED***
			case *stringObject:
				spaceValue = o.ToString()
			***REMOVED***
		***REMOVED***
		isNum := false
		var num int64
		if i, ok := spaceValue.(valueInt); ok ***REMOVED***
			num = int64(i)
			isNum = true
		***REMOVED*** else if f, ok := spaceValue.(valueFloat); ok ***REMOVED***
			num = int64(f)
			isNum = true
		***REMOVED***
		if isNum ***REMOVED***
			if num > 0 ***REMOVED***
				if num > 10 ***REMOVED***
					num = 10
				***REMOVED***
				ctx.gap = strings.Repeat(" ", int(num))
			***REMOVED***
		***REMOVED*** else ***REMOVED***
			if s, ok := spaceValue.(valueString); ok ***REMOVED***
				str := s.String()
				if len(str) > 10 ***REMOVED***
					ctx.gap = str[:10]
				***REMOVED*** else ***REMOVED***
					ctx.gap = str
				***REMOVED***
			***REMOVED***
		***REMOVED***
	***REMOVED***

	if ctx.do(call.Argument(0)) ***REMOVED***
		return newStringValue(ctx.buf.String())
	***REMOVED***
	return _undefined
***REMOVED***

func (ctx *_builtinJSON_stringifyContext) do(v Value) bool ***REMOVED***
	holder := ctx.r.NewObject()
	createDataPropertyOrThrow(holder, stringEmpty, v)
	return ctx.str(stringEmpty, holder)
***REMOVED***

func (ctx *_builtinJSON_stringifyContext) str(key Value, holder *Object) bool ***REMOVED***
	value := nilSafe(holder.get(key, nil))

	if object, ok := value.(*Object); ok ***REMOVED***
		if toJSON, ok := object.self.getStr("toJSON", nil).(*Object); ok ***REMOVED***
			if c, ok := toJSON.self.assertCallable(); ok ***REMOVED***
				value = c(FunctionCall***REMOVED***
					This:      value,
					Arguments: []Value***REMOVED***key***REMOVED***,
				***REMOVED***)
			***REMOVED***
		***REMOVED***
	***REMOVED***

	if ctx.replacerFunction != nil ***REMOVED***
		value = ctx.replacerFunction(FunctionCall***REMOVED***
			This:      holder,
			Arguments: []Value***REMOVED***key, value***REMOVED***,
		***REMOVED***)
	***REMOVED***

	if o, ok := value.(*Object); ok ***REMOVED***
		switch o1 := o.self.(type) ***REMOVED***
		case *primitiveValueObject:
			switch pValue := o1.pValue.(type) ***REMOVED***
			case valueInt, valueFloat:
				value = o.ToNumber()
			default:
				value = pValue
			***REMOVED***
		case *stringObject:
			value = o.toString()
		case *objectGoReflect:
			if o1.toJson != nil ***REMOVED***
				value = ctx.r.ToValue(o1.toJson())
			***REMOVED*** else if v, ok := o1.origValue.Interface().(json.Marshaler); ok ***REMOVED***
				b, err := v.MarshalJSON()
				if err != nil ***REMOVED***
					panic(err)
				***REMOVED***
				ctx.buf.Write(b)
				return true
			***REMOVED*** else ***REMOVED***
				switch o1.className() ***REMOVED***
				case classNumber:
					value = o1.toPrimitiveNumber()
				case classString:
					value = o1.toPrimitiveString()
				case classBoolean:
					if o.ToInteger() != 0 ***REMOVED***
						value = valueTrue
					***REMOVED*** else ***REMOVED***
						value = valueFalse
					***REMOVED***
				***REMOVED***
			***REMOVED***
		***REMOVED***
	***REMOVED***

	switch value1 := value.(type) ***REMOVED***
	case valueBool:
		if value1 ***REMOVED***
			ctx.buf.WriteString("true")
		***REMOVED*** else ***REMOVED***
			ctx.buf.WriteString("false")
		***REMOVED***
	case valueString:
		ctx.quote(value1)
	case valueInt:
		ctx.buf.WriteString(value.String())
	case valueFloat:
		if !math.IsNaN(float64(value1)) && !math.IsInf(float64(value1), 0) ***REMOVED***
			ctx.buf.WriteString(value.String())
		***REMOVED*** else ***REMOVED***
			ctx.buf.WriteString("null")
		***REMOVED***
	case valueNull:
		ctx.buf.WriteString("null")
	case *Object:
		for _, object := range ctx.stack ***REMOVED***
			if value1 == object ***REMOVED***
				ctx.r.typeErrorResult(true, "Converting circular structure to JSON")
			***REMOVED***
		***REMOVED***
		ctx.stack = append(ctx.stack, value1)
		defer func() ***REMOVED*** ctx.stack = ctx.stack[:len(ctx.stack)-1] ***REMOVED***()
		if _, ok := value1.self.assertCallable(); !ok ***REMOVED***
			if isArray(value1) ***REMOVED***
				ctx.ja(value1)
			***REMOVED*** else ***REMOVED***
				ctx.jo(value1)
			***REMOVED***
		***REMOVED*** else ***REMOVED***
			return false
		***REMOVED***
	default:
		return false
	***REMOVED***
	return true
***REMOVED***

func (ctx *_builtinJSON_stringifyContext) ja(array *Object) ***REMOVED***
	var stepback string
	if ctx.gap != "" ***REMOVED***
		stepback = ctx.indent
		ctx.indent += ctx.gap
	***REMOVED***
	length := toLength(array.self.getStr("length", nil))
	if length == 0 ***REMOVED***
		ctx.buf.WriteString("[]")
		return
	***REMOVED***

	ctx.buf.WriteByte('[')
	var separator string
	if ctx.gap != "" ***REMOVED***
		ctx.buf.WriteByte('\n')
		ctx.buf.WriteString(ctx.indent)
		separator = ",\n" + ctx.indent
	***REMOVED*** else ***REMOVED***
		separator = ","
	***REMOVED***

	for i := int64(0); i < length; i++ ***REMOVED***
		if !ctx.str(asciiString(strconv.FormatInt(i, 10)), array) ***REMOVED***
			ctx.buf.WriteString("null")
		***REMOVED***
		if i < length-1 ***REMOVED***
			ctx.buf.WriteString(separator)
		***REMOVED***
	***REMOVED***
	if ctx.gap != "" ***REMOVED***
		ctx.buf.WriteByte('\n')
		ctx.buf.WriteString(stepback)
		ctx.indent = stepback
	***REMOVED***
	ctx.buf.WriteByte(']')
***REMOVED***

func (ctx *_builtinJSON_stringifyContext) jo(object *Object) ***REMOVED***
	var stepback string
	if ctx.gap != "" ***REMOVED***
		stepback = ctx.indent
		ctx.indent += ctx.gap
	***REMOVED***

	ctx.buf.WriteByte('***REMOVED***')
	mark := ctx.buf.Len()
	var separator string
	if ctx.gap != "" ***REMOVED***
		ctx.buf.WriteByte('\n')
		ctx.buf.WriteString(ctx.indent)
		separator = ",\n" + ctx.indent
	***REMOVED*** else ***REMOVED***
		separator = ","
	***REMOVED***

	var props []Value
	if ctx.propertyList == nil ***REMOVED***
		props = object.self.stringKeys(false, nil)
	***REMOVED*** else ***REMOVED***
		props = ctx.propertyList
	***REMOVED***

	empty := true
	for _, name := range props ***REMOVED***
		off := ctx.buf.Len()
		if !empty ***REMOVED***
			ctx.buf.WriteString(separator)
		***REMOVED***
		ctx.quote(name.toString())
		if ctx.gap != "" ***REMOVED***
			ctx.buf.WriteString(": ")
		***REMOVED*** else ***REMOVED***
			ctx.buf.WriteByte(':')
		***REMOVED***
		if ctx.str(name, object) ***REMOVED***
			if empty ***REMOVED***
				empty = false
			***REMOVED***
		***REMOVED*** else ***REMOVED***
			ctx.buf.Truncate(off)
		***REMOVED***
	***REMOVED***

	if empty ***REMOVED***
		ctx.buf.Truncate(mark)
	***REMOVED*** else ***REMOVED***
		if ctx.gap != "" ***REMOVED***
			ctx.buf.WriteByte('\n')
			ctx.buf.WriteString(stepback)
			ctx.indent = stepback
		***REMOVED***
	***REMOVED***
	ctx.buf.WriteByte('***REMOVED***')
***REMOVED***

func (ctx *_builtinJSON_stringifyContext) quote(str valueString) ***REMOVED***
	ctx.buf.WriteByte('"')
	reader := &lenientUtf16Decoder***REMOVED***utf16Reader: str.utf16Reader(0)***REMOVED***
	for ***REMOVED***
		r, _, err := reader.ReadRune()
		if err != nil ***REMOVED***
			break
		***REMOVED***
		switch r ***REMOVED***
		case '"', '\\':
			ctx.buf.WriteByte('\\')
			ctx.buf.WriteByte(byte(r))
		case 0x08:
			ctx.buf.WriteString(`\b`)
		case 0x09:
			ctx.buf.WriteString(`\t`)
		case 0x0A:
			ctx.buf.WriteString(`\n`)
		case 0x0C:
			ctx.buf.WriteString(`\f`)
		case 0x0D:
			ctx.buf.WriteString(`\r`)
		default:
			if r < 0x20 ***REMOVED***
				ctx.buf.WriteString(`\u00`)
				ctx.buf.WriteByte(hex[r>>4])
				ctx.buf.WriteByte(hex[r&0xF])
			***REMOVED*** else ***REMOVED***
				if utf16.IsSurrogate(r) ***REMOVED***
					ctx.buf.WriteString(`\u`)
					ctx.buf.WriteByte(hex[r>>12])
					ctx.buf.WriteByte(hex[(r>>8)&0xF])
					ctx.buf.WriteByte(hex[(r>>4)&0xF])
					ctx.buf.WriteByte(hex[r&0xF])
				***REMOVED*** else ***REMOVED***
					ctx.buf.WriteRune(r)
				***REMOVED***
			***REMOVED***
		***REMOVED***
	***REMOVED***
	ctx.buf.WriteByte('"')
***REMOVED***

func (r *Runtime) initJSON() ***REMOVED***
	JSON := r.newBaseObject(r.global.ObjectPrototype, "JSON")
	JSON._putProp("parse", r.newNativeFunc(r.builtinJSON_parse, nil, "parse", nil, 2), true, false, true)
	JSON._putProp("stringify", r.newNativeFunc(r.builtinJSON_stringify, nil, "stringify", nil, 3), true, false, true)
	JSON._putSym(SymToStringTag, valueProp(asciiString(classJSON), false, false, true))

	r.addToGlobal("JSON", JSON.val)
***REMOVED***
