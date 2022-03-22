package validate

import (
	"reflect"
)

type Validate struct {
	errors map[string]*Field
}

func New() *Validate {
	return &Validate{
		errors: map[string]*Field{},
	}
}

func (v *Validate) Use(name string, f CallFunc) {
	if name == "format" {
		formatFunc[name] = f
	} else {
		expFunc[name] = f
	}
}

func (v *Validate) Check() bool {
	return len(v.errors) == 0
}

func (v *Validate) Error() string {
	if len(v.errors) > 0 {
		for _, v := range v.errors {
			return v.Msg
		}
	}
	return ""
}

func (v *Validate) GetErrors() map[string]*Field {
	return v.errors
}

func (v *Validate) Struct(s any) *Validate {
	svalue := reflect.ValueOf(s)
	if svalue.Kind() == reflect.Ptr {
		svalue = svalue.Elem()
	}
	stype := svalue.Type()
	for i := 0; i < stype.NumField(); i++ {
		ftype := stype.Field(i)
		fvalue := svalue.Field(i)
		f := NewField(ftype.Name, fvalue.Interface(), ftype.Type.Kind().String(), ftype.Tag.Get("validate"))
		f.Parse()
		if f.State == false {
			v.errors[f.Name] = f
		}

	}
	return v
}
