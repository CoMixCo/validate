package main

import (
	"fmt"
	"validate"
)

func main() {
	v := validate.New()
	data := struct {
		Account string `validate:"format=email > 邮箱格式错误"`
		Name    string `validate:"empty=true | format=trim_space & gt=4 > 字符必须大于4个"`
		Age     int    `validate:"o_interval=10,100 > 年龄需要大于10小于100"`
		Mobile  string `validate:"format=cn_mobile > 手机格式错误"`
		Status  int    `validate:"in=0,1 >状态值错误"`
	}{
		Account: "even@qq.com",
		Name:    "eventt ",
		Age:     6,
		Mobile:  "1361173787",
		Status:  -1,
	}
	if !v.Struct(&data).Check() {
		for _, val := range v.GetErrors() {
			fmt.Println(val.Msg)
		}
	}

	fmt.Println(data.Name + "end")
}
