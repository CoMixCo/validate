package main

import (
	"fmt"
	"reflect"
	"time"

	"github.com/coco-look/validate"
	"github.com/coco-look/validate/element"
)

func main() {
	v := validate.New()

	//自定义格式化方法

	v.AddFormatMethod("date", func(f *element.Field) bool {
		switch f.Kind {
		case reflect.String:
			if _, err := time.Parse("2006-01-02", f.Val.String()); err == nil {
				return true
			}
		}
		return false
	})

	data := struct {
		Account   string `validate:"format=email > 邮箱格式错误"`
		Name      string `validate:"empty=true | format=trim_space & gt=4 > 字符必须大于4个"`
		Age       int    `validate:"gte=10 & lte=100 > 年龄需要大于10小于100"`
		Mobile    string `validate:"format=cn_mobile > 手机格式错误"`
		Status    int    `validate:"in=0,1 >状态值错误"`
		DateStart string `validate:"format=date>日期格式错误"`
	}{
		Account:   "even@qq.com",
		Name:      "eventt ",
		Age:       6,
		Mobile:    "1361173787",
		Status:    -1,
		DateStart: "2022-05",
	}
	if !v.Struct(&data).Check() {
		for _, val := range v.GetErrors() {
			fmt.Println(val.Msg)
		}
	}

	fmt.Printf("被验证数据为：%+v \n", data)
}
