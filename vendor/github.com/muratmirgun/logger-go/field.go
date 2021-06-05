package log

import (
	"reflect"
	"strconv"
)

// Field stores a name/value pair to be formatted by the emitter.
type Field struct {
	Name        string
	Type        reflect.Kind
	BoolValue   bool
	IntValue    int64
	StringValue string
}

// Bool returns a Field that contains a boolean value.
func Bool(name string, value bool) Field {
	return Field{Name: name, Type: reflect.Bool, BoolValue: value}
}

// Int returns a Field that contains a default integer.
func Int(name string, value int) Field {
	return Field{Name: name, Type: reflect.Int, IntValue: int64(value)}
}

// Int8 returns a Field that contains an 8-bit integer
func Int8(name string, value int8) Field {
	return Field{Name: name, Type: reflect.Int8, IntValue: int64(value)}
}

// Int16 returns a Field that contains an 16-bit integer
func Int16(name string, value int16) Field {
	return Field{Name: name, Type: reflect.Int16, IntValue: int64(value)}
}

// Int32 returns a Field that contains an 32-bit integer
func Int32(name string, value int32) Field {
	return Field{Name: name, Type: reflect.Int32, IntValue: int64(value)}
}

// Int64 returns a Field that contains an 64-bit integer
func Int64(name string, value int64) Field {
	return Field{Name: name, Type: reflect.Int64, IntValue: value}
}

// String returns a Field that contains a string.
func String(name string, value string) Field {
	return Field{Name: name, Type: reflect.String, StringValue: value}
}

// Err returns a Field that contains the message from an error.
func Err(value error) Field {
	return Field{Name: "error", Type: reflect.String, StringValue: value.Error()}
}

// Json returns the contents of the Field as `"key":value`.
func (field Field) Json() string {
	return "\"" + field.Name + "\":" + stringValue(field, true)
}

// String returns the contents of the Field as `key=value`.
func (field Field) String() string {
	return field.Name + "=" + stringValue(field, false)
}

// stringValue converts the field's value into a string.
func stringValue(field Field, quoted bool) string {
	switch field.Type {
	case reflect.Bool:
		if field.BoolValue {
			return "true"
		}
		return "false"
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(field.IntValue, 10)
	case reflect.String:
		if quoted {
			return "\"" + field.StringValue + "\""
		}
		return field.StringValue
	default:
		panic("Invalid type")
	}
}
