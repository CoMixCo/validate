示例：

```
package main

import (
	"fmt"
	"utils/validate"
)

func main() {

	//validate.DebugModel = true
	v := validate.New()
	data := struct {
		Account string `validate:"empty=false & format=email > 邮箱格式错误"`
		Name    string `validate:"empty=true | gt=4 > 字符必须大于4个"`
		Age     int    `validate:"section=10,100 > 年龄需要大于10小于100"`
		Mobile  string `validate:"format=cn_mobile > 手机格式错误"`
		Status  int    `validate:"in=0,1 >状态值错误"`
	}{
		Account: "even@qq.com",
		Name:    "even",
		Age:     6,
		Mobile:  "1361173787",
		Status:  -1,
	}
	rs := v.Struct(&data).Check()
	if !rs {
		for _, val := range v.GetErrors() {
			fmt.Println(val.Msg)
		}
		//fmt.Println(v.Error())
	}
}
```
