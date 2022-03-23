package validate

import (
	"fmt"
	"reflect"
)

type Field struct {
	RefValue reflect.Value
	Name     string
	Val      any
	Kind     string
	Tag      string
	State    bool
	Msg      string
}

func NewField(r reflect.Value, n string, v any, k string, t string) *Field {
	return &Field{
		RefValue: r,
		Name:     n,
		Val:      v,
		Kind:     k,
		Tag:      t,
		State:    false,
	}
}

//exp:[map[empty:true] map[format:email gt:3]]
func (f *Field) Parse() *Field {
	t := NewTag(f.Tag).Parse()
	exp := t.GetExp()

	for _, part := range exp {
		for k, v := range part {
			if k == "format" {
				if call, ok := formatFunc[v]; ok {
					f.State = call(f)
				}
			} else {
				if call, ok := expFunc[k]; ok {
					f.State = call(f, v)
				}
			}
			// and 条件有false就不满足
			if f.State == false {
				break
			}
		}
		// or条件有true就满足
		if f.State == true {
			break
		}
	}
	if f.State == false {
		if DebugModel {
			f.Msg = fmt.Sprintf("字段:%s 传值:%v 校验:%s", f.Name, f.Val, t.GetMsg())
		} else {
			f.Msg = fmt.Sprintf("字段:%s 校验:%s", f.Name, t.GetMsg())
		}

	}

	return f
}
