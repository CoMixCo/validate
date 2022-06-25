package validate

import (
	"reflect"
)

var DebugModel bool

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

func (v *Validate) Struct(s interface{}) *Validate {
	struct_value := reflect.ValueOf(s)
	if struct_value.Kind() == reflect.Ptr {
		struct_value = struct_value.Elem()
	}
	if struct_value.Kind() != reflect.Struct {
		panic("validate data expect struct or struct point")
	}
	struct_type := struct_value.Type()
	for i := 0; i < struct_type.NumField(); i++ {
		field_type := struct_type.Field(i)
		if validate_tag, ok := field_type.Tag.Lookup("validate"); ok {
			field_value := struct_value.Field(i)
			f := NewField(struct_value, field_type.Name, field_value, field_type.Type.Kind(), validate_tag).Parse()
			if f.State == false {
				v.errors[f.Name] = f
			}
		}
	}
	return v
}

func (v *Validate) Check() bool {
	return len(v.errors) == 0
}

func (v *Validate) Error() string {
	if len(v.errors) > 0 {
		//随机获取一个错误
		for _, v := range v.errors {
			return v.Msg
		}
	}
	return ""
}

func (v *Validate) GetErrors() map[string]*Field {
	return v.errors
}
