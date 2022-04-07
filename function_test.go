package validate

import (
	"fmt"
	"reflect"
	"testing"
)

func TestUrl(t *testing.T) {
	str := "http://www.google.cn"
	refStr := reflect.ValueOf(str)

	f := NewField(refStr, "callback_url", refStr, refStr.Kind(), "aaa")

	if false == url(f) {
		t.Error("匹配错误")
	} else {
		fmt.Println("test url ok")
	}
}
