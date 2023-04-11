package validate

import (
	"fmt"
	"reflect"
	"strings"
)

type Field struct {
	RefStruct reflect.Value
	Name      string
	Val       reflect.Value
	Kind      reflect.Kind
	Tag       string
	State     bool
	Msg       string
}

func NewField(struct_value reflect.Value, field_name string, field_val reflect.Value, field_kind reflect.Kind, field_tag string) *Field {
	f := &Field{
		RefStruct: struct_value,
		Name:      field_name,
		Val:       field_val,
		Kind:      field_kind,
		Tag:       field_tag,
		State:     false,
	}
	f.parse()
	return f
}

// exp:[map[empty:true] map[format:email gt:3]]
func (f *Field) parse() *Field {
	t := NewTag(f.Tag)
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
			if !f.State {
				break
			}
		}
		// or条件有true就满足
		if f.State {
			break
		}
	}
	if !f.State {
		if DebugModel {
			f.Msg = fmt.Sprintf("field:%s value:%v verify:%s", SnakeString(f.Name), f.Val, t.GetMsg())
		} else {
			f.Msg = fmt.Sprintf("field:%s verify:%s", SnakeString(f.Name), t.GetMsg())
		}

	}

	return f
}

// 转为蛇形文
func SnakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		// or通过ASCII码进行大小写的转化
		// 65-90（A-Z），97-122（a-z）
		//判断如果字母为大写的A-Z就在前面拼接一个_
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	//ToLower把大写字母统一转小写
	return strings.ToLower(string(data[:]))
}
