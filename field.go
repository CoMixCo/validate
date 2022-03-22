package validate

type Field struct {
	Name  string
	Val   any
	Kind  string
	Tag   string
	State bool
	Msg   string
}

func NewField(n string, v any, k string, t string) *Field {
	return &Field{
		Name:  n,
		Val:   v,
		Kind:  k,
		Tag:   t,
		State: true,
	}
}

//exp:[map[empty:true] map[format:email gt:3]]
func (f *Field) Parse() *Field {
	t := NewTag(f.Tag)
	t.Parse()
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

			if f.State == false {
				break
			}
		}
		if f.State == true {
			break
		}
	}
	if f.State == false {
		f.Msg = t.GetMsg()
	}

	return f
}
