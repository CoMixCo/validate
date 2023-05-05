package validate

import (
	"fmt"
	"reflect"

	"github.com/coco-look/validate/method"

	"github.com/coco-look/validate/element"
)

var DebugModel bool

type Validate struct {
	errors map[string]*element.Field
}

func New() *Validate {
	return &Validate{
		errors: map[string]*element.Field{},
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
			f := element.NewField(struct_value, field_type.Name, struct_value.Field(i), field_type.Type.Kind(), validate_tag)
			f = v.Parse(f)
			if !f.State {
				v.errors[f.Name] = f
			}
		}
	}
	return v
}

/**
 * 解析表达式逻辑
 * eg: exp=[map[eq:0], map[format:url, lte:255]]
 */
func (v *Validate) Parse(f *element.Field) *element.Field {
	t := element.NewTag(f.Tag)
	exp := t.GetExp()

	for _, part := range exp {
		for k, v := range part {
			if k == "format" {
				if call, ok := method.FormatFuncMap[v]; ok {
					f.State = call(f)
				}
			} else {
				if call, ok := method.CompareFuncMap[k]; ok {
					f.State = call(f, v)
				}
			}
			// and 条件有返回false就中断后续逻辑判断
			if !f.State {
				break
			}
		}
		// or条件有返回true就中断后续逻辑判断
		if f.State {
			break
		}
	}
	if !f.State {
		if DebugModel {
			f.Msg = fmt.Sprintf("field:%s value:%v error:%s", element.SnakeString(f.Name), f.Val, t.GetMsg())
		} else {
			f.Msg = fmt.Sprintf("field:%s error:%s", element.SnakeString(f.Name), t.GetMsg())
		}

	}

	return f
}

// 检测验证错误是否存在
func (v *Validate) Check() bool {
	return len(v.errors) == 0
}

// 如果有多个错误，随机获取一个错误
func (v *Validate) Error() string {
	if len(v.errors) > 0 {
		for _, v := range v.errors {
			return v.Msg
		}
	}
	return ""
}

// 获取所有错误
func (v *Validate) GetErrors() map[string]*element.Field {
	return v.errors
}

// 添加自定义比较方法
func (v *Validate) AddCompareMethod(name string, f method.CompareFunc) {
	method.CompareFuncMap[name] = f
}

// 添加自定义格式化方法
func (v *Validate) AddFormatMethod(name string, f method.FormatFunc) {
	method.FormatFuncMap[name] = f
}
