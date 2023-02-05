package utils

import (
	"encoding"
	"reflect"
	"strconv"
)

// Scan T must be a struct
func Scan[T any](mp map[string]string) *T {
	var t T
	scan(mp, reflect.ValueOf(&t).Elem())
	return &t
}

func scan(mp map[string]string, st reflect.Value) {
	if st.Type().Kind() == reflect.Pointer {
		st = st.Elem()
	}
	ty := st.Type()
	for i := 0; i < ty.NumField(); i++ {
		fieldType := ty.Field(i)
		fieldVal := st.Field(i)
		k := fieldType.Tag.Get("redis")
		if k == "" {
			k = fieldType.Name
		}

		literal, ok := mp[k]
		if !ok && fieldType.Type.Kind() == reflect.Struct {
			scan(mp, st.Field(i))
			continue
		}

		// try encoding.TextUnmarshaler
		var addr reflect.Value
		if fieldType.Type.Kind() == reflect.Pointer {
			addr = fieldVal
		} else if fieldVal.CanAddr() {
			addr = fieldVal.Addr()
		}
		if addr.Type().NumMethod() > 0 && addr.CanInterface() {
			switch scan := addr.Interface().(type) {
			case encoding.TextUnmarshaler:
				_ = scan.UnmarshalText([]byte(literal))
				continue
			}
		}
		fieldVal = addr.Elem()

		switch fieldVal.Kind() {
		case reflect.String:
			fieldVal.SetString(literal)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if v, err := strconv.ParseInt(literal, 10, 64); err == nil && !fieldVal.OverflowInt(v) {
				fieldVal.SetInt(v)
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			if v, err := strconv.ParseUint(literal, 10, 64); err == nil && !fieldVal.OverflowUint(v) {
				fieldVal.SetUint(v)
			}
		case reflect.Float32, reflect.Float64:
			if v, err := strconv.ParseFloat(literal, 64); err == nil && !fieldVal.OverflowFloat(v) {
				fieldVal.SetFloat(v)
			}
		case reflect.Bool:
			if v, err := strconv.ParseBool(literal); err == nil {
				fieldVal.SetBool(v)
			}
		case reflect.Struct:
			scan(mp, fieldVal)
		}

	}

}
