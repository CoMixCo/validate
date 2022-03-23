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
	//干掉所有空格
	if b, a, f := strings.Cut(t.str, ">"); f {
		t.msg = a
		t.expStr = strings.Replace(b, " ", "", -1)
	} else {
		t.expStr = strings.Replace(t.str, " ", "", -1)
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

//eg: empty=true | empty=false&len>0
func OrExp(str string) Or {
	or := Or{}
	slice := strings.Split(str, "|")
	for _, part := range slice {
		or = append(or, AndExp(part))
	}
	return or
}

// eg: empty=false&len>0
func AndExp(str string) And {
	and := And{}
	slice := strings.Split(str, "&")
	for _, part := range slice {
		if b, a, f := strings.Cut(part, "="); f {
			and[b] = a
		}
	}
	return and
}
