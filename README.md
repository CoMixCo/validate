v1版本：
自定义验证标签：v.Use(tagName string, func(f *validate.Field, args ...string) bool)
示例：
```
v.Use("lt_field", func(f *validate.Field, args ...string) bool {
	diff_field := f.RefValue.FieldByName(args[0])
	if len(f.Val.(string)) < len(string(diff_field.Interface().(string))) {
		return true
	}
	return false
})
```

支持逻辑运算 且[&] 和 或[|]
示例：
```
Account  string `validate:"empty=false & format=email >邮箱格式错误"`
Age            int    `validate:"eq=0 | section=10,100 >年龄需要大于10小于100"`
```

支持比较运算:
等于： eq=6   
大于：gt=6 
小于：lt=6

支持包含验证
包含：in=1,0

支持区间验证
区间：section=0,100  大于0小于100

支持字段比较
比较字段 eq_field
示例：
```
Password       string `validate:"gt=6>密码长度需要大于6"`
PasswordRepeat string `validate:"eq_field=Password>两次密码不相同"`
```
格式化校验：
format: email
format:cn_mobile
```
Account string `validate:"format=email > 邮箱格式错误"`
Mobile  string `validate:"format=cn_mobile > 手机格式错误"`
```

示例一：

```
package main

import (
	"fmt"
	"utils/validate"
)

func main() {

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
	if v.Struct(&data).Check() {
		for _, val := range v.GetErrors() {
			fmt.Println(val.Msg)
		}
	}
}
```

示例二
```
package main

import (
	"fmt"
	"utils/validate"
)

func main() {

	validate.DebugModel = true
	v := validate.New()
	data := struct {
		Account        string `validate:"empty=false & format=email >邮箱格式错误"`
		Name           string `validate:"empty=true | gt=4 >字符必须大于4个"`
		FirstName      string `validate:"lt_field=Name > 姓名必须小于全名"`
		Age            int    `validate:"eq=0 | section=10,100 >年龄需要大于10小于100"`
		Password       string `validate:"gt=6>密码长度需要大于6"`
		PasswordRepeat string `validate:"eq_field=Password>两次密码不相同"`
	}{
		Account:        "even@qq.com",
		Name:           "even cc",
		FirstName:      "ccsdsdsd",
		Age:            0,
		Password:       "1qaz@2wsx",
		PasswordRepeat: "1qaz@2wsx1",
	}
	v.Use("lt_field", func(f *validate.Field, args ...string) bool {
		compare_val := f.RefStruct.FieldByName(args[0])
		switch f.Kind {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return f.Val.Int() < compare_val.Int()
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return f.Val.Uint() < compare_val.Uint()
		case reflect.Float32, reflect.Float64:
			return f.Val.Float() < compare_val.Float()
		case reflect.String:
			return f.Val.Len() < compare_val.Len()
		}
		return false
	})
	if !v.Struct(&data).Check() {
		fmt.Println(v.Error())
	}
}

```
