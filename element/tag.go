package element

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
	t := &Tag{
		str: str,
	}
	t.logicOperation()
	return t
}

/**
 * tag 的逻辑运算处理
 */
func (t *Tag) logicOperation() *Tag {
	//干掉所有空格
	if b, a, f := strings.Cut(t.str, ">"); f {
		t.msg = strings.TrimSpace(a)
		t.expStr = strings.Replace(b, " ", "", -1)
	} else {
		t.expStr = strings.Replace(t.str, " ", "", -1)
	}
	t.exp = OrExp(t.expStr)
	return t
}

/**
 * 获取错误提示信息, 表达式>后面的字符串为错误提示信息
 */
func (t *Tag) GetMsg() string {
	return t.msg
}

/**
 * 获取表达式的map条件
 */
func (t *Tag) GetExp() Or {
	return t.exp
}

/**
 * or表达式
 * eg: eq=0 | eq=1
 */
func OrExp(exp_str string) Or {
	or := Or{}
	slice := strings.Split(exp_str, "|")
	for _, part := range slice {
		or = append(or, AndExp(part))
	}
	return or
}

/**
 * and 表达式
 * eg: format=url & lte=255
 */
func AndExp(exp_str string) And {
	and := And{}
	slice := strings.Split(exp_str, "&")
	for _, part := range slice {
		if b, a, f := strings.Cut(part, "="); f {
			and[b] = a
		}
	}
	return and
}
