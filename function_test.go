package validate

import (
	"fmt"
	"reflect"
	"testing"
)

func TestUrl(t *testing.T) {
	str := "http://www.google.cn"
	refStr := reflect.ValueOf(str)

	f := NewField(refStr, "attach", refStr, refStr.Kind(), "aaa")

	if false == url(f) {
		t.Error("匹配错误")
	} else {
		fmt.Println("test url ok")
	}
}

func TestSafeStr(t *testing.T) {
	//str := "P234234_2323"
	str := "mpay_6056002"
	refStr := reflect.ValueOf(str)

	f := NewField(refStr, "attach", refStr, refStr.Kind(), "aaa")

	if false == safe_str(f) {
		t.Error("匹配错误")
	} else {
		fmt.Println("test safe_str ok")
	}
}

func TestEmail(t *testing.T) {
	str := "xxx@qq.com"
	refStr := reflect.ValueOf(str)

	f := NewField(refStr, "attach", refStr, refStr.Kind(), "aaa")

	if false == email(f) {
		t.Error("匹配错误")
	} else {
		fmt.Println("test email ok")
	}
}
