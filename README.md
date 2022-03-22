```
package main

import (
	"fmt"
	"utils/validate"
)

func main() {

	v := validate.New()
	data := struct {
		Account string `validate:"empty=false & format=email >邮箱格式错误"`
		Name    string `validate:"empty=true | gt=4 >字符必须大于4个"`
		Age     int    `validate:"eq=0 | gt=10 >年龄需要大于10"`
	}{
		Account: "even@qq.com",
		Name:    "even",
		Age:     6,
	}
	rs := v.Struct(&data).Check()
	if !rs {
		for _, val := range v.GetErrors() {
			fmt.Printf("field is %v, val is %v, msg is %v", val.Name, val.Val, val.Msg)
			fmt.Println()
		}
	}
}
```
