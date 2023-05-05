package method

import (
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/coco-look/validate/element"
)

type FormatFunc func(f *element.Field) bool

/**
 * 格式化函数
 */
var FormatFuncMap = map[string]FormatFunc{
	"email":      email,
	"cn_mobile":  cn_mobile,
	"url":        url,
	"safe_str":   safe_str,
	"trim_space": trim_space,
	"date":       date,
	"date_time":  date_time,
}

/**
 * 格式化：邮箱
 * 适用类型：字符串
 * eg: format=email
 */
func email(f *element.Field) bool {
	switch f.Kind {
	case reflect.String:
		reg := regexp.MustCompile(`^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`)
		return reg.MatchString(f.Val.String())
	}
	return false
}

/**
 * 格式化：中国手机
 * 适用类型：字符串
 * eg: format=cn_mobile
 */
func cn_mobile(f *element.Field) bool {
	switch f.Kind {
	case reflect.String:
		reg := regexp.MustCompile(`^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\d{8}$`)
		return reg.MatchString(f.Val.String())
	}
	return false
}

/**
 * 格式化：网址
 * 使用类型：字符串
 * eg: format=cn_mobile
 */
func url(f *element.Field) bool {
	switch f.Kind {
	case reflect.String:
		reg := regexp.MustCompile(`^(http|https):\/\/[\w\-_]+(\.[\w\-_]+)+([\w\-\.,@?^=%&:/~\+#]*[\w\-\@?^=%&/~\+#])?$`)
		return reg.MatchString(f.Val.String())
	}
	return false
}

/**
 * 安全的字符串
 * eg: format=safe_str
 */
func safe_str(f *element.Field) bool {
	switch f.Kind {
	case reflect.String:
		reg := regexp.MustCompile(`^[A-Za-z0-9_\.]+$`)
		return reg.MatchString(f.Val.String())
	}
	return false
}

/**
 * 过滤首尾空格
 * eg: format=trim_space
 */
func trim_space(f *element.Field) bool {
	switch f.Kind {
	case reflect.String:
		trim_str := strings.TrimSpace(f.Val.String())
		f.RefStruct.FieldByName(f.Name).SetString(trim_str)
	}
	return true
}

/**
 * 判断日期
 * eg: format=date
 */
func date(f *element.Field) bool {
	switch f.Kind {
	case reflect.String:
		if _, err := time.Parse("2006-01-02", f.Val.String()); err == nil {
			return true
		}
	}
	return false
}

/**
 * 判断日期时间
 * eg: format=date_time
 */
func date_time(f *element.Field) bool {
	switch f.Kind {
	case reflect.String:
		if _, err := time.Parse("2006-01-02 15:04:05", f.Val.String()); err == nil {
			return true
		}
	}
	return false
}
