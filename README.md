仅针对GO结构体进行数据验证，使用反射实现。功能描述：

支持逻辑运算 且[&] 和 或[|]
```
逻辑：eq=0 | gte=10 & lte=100
```

支持比较运算, 字符串比较长度，数字比较大小:
```
等于： eq=6   
大于：gt=6 
大于等于：gte=6 
小于：lt=6
小于等于：lte=6
```

支持包含比较
```
包含：in=1,0
```

支持字段比较
```
比较字段 eq_field=Password
```

支持字符串格式化校验：
```
format=email
format=cn_mobile
format=url
format=safe_str
format=trim_space
format=date
format=date_time
```

支持自定义比较方法：v.AddCompareMethod(tagName string, func(f *element.Field, args ...string) bool)

支持自定义格式化方法：v.AddFormatMethod(tagName string, func(f *element.Field, args ...string) bool)
示例：
```
v.AddCompareMethod("lt_field", func(f *element.Field, args ...string) bool {
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
```

示例一：

```
package main

import (
	"fmt"
	"validate"
	"validate/element"
)

func main() {

	v := validate.New()
	data := struct {
		Account string `validate:"format=email > 邮箱格式错误"`
		Name    string `validate:"gt=4 > 字符必须大于4个"`
		Age     int    `validate:"gt=10 & lt=100 > 年龄需要大于10小于100"`
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
	"time"
	"validate"
	"validate/element"
)

func main() {

	validate.DebugModel = true
	v := validate.New()
	data := struct {
		Account        string `validate:"format=email >邮箱格式错误"`
		Name           string `validate:"gte=4 >字符必须大于等于4个"`
		FirstName      string `validate:"lt_field=Name >姓氏长度需要小于名字长度"`
		Age            int    `validate:"eq=0 | gt=10 & lt=100 >年龄需要大于等于10小于等于100"`
		Password       string `validate:"gte=6>密码长度需要大于6"`
		PasswordRepeat string `validate:"eq_field=Password>两次密码不相同"`
		DateStart      string `validate:"format=date>日期格式错误"`
	}{
		Account:        "even@qq.com",
		Name:           "even cc",
		FirstName:      "ccsdsdsd",
		Age:            0,
		Password:       "123456",
		PasswordRepeat: "1234567",
		DateStart:      "2022-05",
	}
	v.AddCompareMethod("lt_field", func(f *element.Field, arg string) bool {
		compare_val := f.RefStruct.FieldByName(arg)
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
	v.AddFormatMethod("date", func(f *element.Field) bool {
		switch f.Kind {
		case reflect.String:
			if _, err := time.Parse("2006-01-02", f.Val.String()); err == nil {
				return true
			}
		}
		return false
	})
	if !v.Struct(&data).Check() {
		fmt.Println(v.Error())
	}
}

```
