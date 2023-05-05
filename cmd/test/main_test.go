package main

import (
	"fmt"
	"testing"

	"github.com/coco-look/validate"
)

func TestUrl(t *testing.T) {
	v := validate.New()

	val := "http://www.google.cn"

	data := struct {
		Url string `validate:"format=url"`
	}{
		Url: val,
	}
	if v.Struct(&data).Check() {
		fmt.Println("test url ok")
	} else {
		t.Error("URL匹配错误", v.Error())
	}
}

func TestSafeStr(t *testing.T) {

	v := validate.New()

	val := "mpa+6056002"

	data := struct {
		OrderNo string `validate:"format=safe_str"`
	}{
		OrderNo: val,
	}
	if !v.Struct(&data).Check() {
		fmt.Println("test safe_str ok")
	} else {
		t.Error("test safe_str fail")
	}
}

func TestEmail(t *testing.T) {

	v := validate.New()

	val := "xxx@qq.com"

	data := struct {
		Email string `validate:"format=email"`
	}{
		Email: val,
	}
	if v.Struct(&data).Check() {
		fmt.Println("test email ok")
	} else {
		t.Error("匹配错误", v.Error())
	}
}

func TestTrimSpace(t *testing.T) {

	v := validate.New()

	val := "A232423 "

	data := struct {
		CardNo string `validate:"format=trim_space"`
	}{
		CardNo: val,
	}
	v.Struct(&data).Check()

	if len(val) != len(data.CardNo) {
		fmt.Println("test trim_space ok")
	} else {
		t.Error("test trim_space fail")
	}
}
