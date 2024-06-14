package logi

import (
	"reflect"
	"strconv"
)

type optionFormat struct {
	DiscardNil bool
}

func (o *optionFormat) apply(opts ...OptionFormat) {
	for _, opt := range opts {
		opt(o)
	}
}

type OptionFormat func(*optionFormat)

func WithDiscardNil(v bool) OptionFormat {
	return func(o *optionFormat) {
		o.DiscardNil = v
	}
}

// Format discard log:"false" tag in struct values and return new struct.
func Format(v any, opts ...OptionFormat) any {
	o := optionFormat{}
	o.apply(opts...)

	// if v is struct
	if reflect.TypeOf(v) != nil {
		return formatStruct(v, o)
	}

	return v
}

func structStructure(v any, o optionFormat) []reflect.StructField {
	// filter struct's nil fields
	rv := reflect.ValueOf(v)
	rt := reflect.TypeOf(v)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
		rt = rt.Elem()
	}

	if rv.Kind() != reflect.Struct {
		return nil
	}

	structFields := make([]reflect.StructField, 0, rv.NumField())
	for i := 0; i < rv.NumField(); i++ {
		fv := rv.Field(i)

		// field has log false tag
		if tag := rt.Field(i).Tag.Get("log"); tag != "" {
			if v, _ := strconv.ParseBool(tag); !v {
				continue
			}
		}

		if o.DiscardNil && fv.Kind() == reflect.Ptr && fv.IsNil() {
			continue
		}

		// if field is struct, recursively format it
		if fv.Kind() == reflect.Struct {
			structFields = append(structFields, reflect.StructField{
				Name: rt.Field(i).Name,
				Type: reflect.StructOf(structStructure(fv.Interface(), o)),
			})

			continue
		}

		structFields = append(structFields, rt.Field(i))
	}

	return structFields
}

func formatStruct(v any, o optionFormat) any {
	// filter struct's nil fields
	rv := reflect.ValueOf(v)
	rt := reflect.TypeOf(v)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
		rt = rt.Elem()
	}

	if rv.Kind() != reflect.Struct {
		return v
	}

	structFields := structStructure(v, o)

	// create a new struct with non-nil fields
	nv := reflect.New(reflect.StructOf(structFields)).Elem()
	for i := 0; i < rv.NumField(); i++ {
		fv := rv.Field(i)

		// field has log false tag
		if tag := rt.Field(i).Tag.Get("log"); tag != "" {
			if v, _ := strconv.ParseBool(tag); !v {
				continue
			}
		}

		if o.DiscardNil && fv.Kind() == reflect.Ptr && fv.IsNil() {
			continue
		}

		// if field is struct, recursively format it
		if fv.Kind() == reflect.Struct {
			fv = reflect.ValueOf(formatStruct(fv.Interface(), o))
		}

		nv.FieldByName(rt.Field(i).Name).Set(fv)
	}

	return nv.Interface()
}
