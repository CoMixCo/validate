package validate

import "strings"

type And map[string]string
type Or []And

type Tag struct {
	str    string
	expStr string
	msg    string
	exp    Or
}

func NewTag(str string) *Tag {
	return &Tag{
		str: str,
	}
}

func (t *Tag) Parse() *Tag {
	if b, a, f := strings.Cut(t.str, ">"); f {
		t.expStr = strings.TrimSpace(b)
		t.msg = strings.TrimSpace(a)
	} else {
		t.expStr = strings.TrimSpace(t.str)
	}
	t.exp = OrExp(t.expStr)
	return t
}

func (t *Tag) GetMsg() string {
	return t.msg
}

func (t *Tag) GetExp() Or {
	return t.exp
}

//empty=true | empty=false&len>0
func OrExp(str string) Or {
	or := Or{}
	slice := strings.Split(str, "|")
	for _, part := range slice {
		or = append(or, AndExp(part))
	}
	return or
}

// empty=false&len>0
func AndExp(str string) And {
	and := And{}
	str = strings.TrimSpace(str)
	slice := strings.Split(str, "&")
	for _, part := range slice {
		if b, a, f := strings.Cut(part, "="); f {
			and[strings.TrimSpace(b)] = strings.TrimSpace(a)
		}
	}
	return and

}
