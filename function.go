package validate

import (
	"regexp"
	"strconv"
	"strings"
)

type CallFunc func(f *Field, args ...string) bool

/**
 * 格式化函数
 */
var formatFunc = map[string]CallFunc{
	"email":     email,
	"cn_mobile": cn_mobile,
}

/**
 * 表达式比较函数
 */
var expFunc = map[string]CallFunc{
	"gt":       gt,
	"eq":       eq,
	"lt":       lt,
	"empty":    empty,
	"section":  section,
	"in":       in,
	"eq_field": eq_field,
}

/**
 * 适用数字和字符串
 */
func gt(f *Field, args ...string) bool {
	switch fdata := f.Val.(type) {
	case uint, uint8, uint16, uint32, uint64:
		compare, _ := strconv.ParseUint(args[0], 10, 64)
		return uint64(fdata.(uint)) > compare
	case int, int8, int16, int32, int64:
		compare, _ := strconv.ParseInt(args[0], 10, 64)
		return int64(fdata.(int)) > compare
	case float32, float64:
		compare, _ := strconv.ParseFloat(args[0], 64)
		return float64(fdata.(float64)) > compare
	case string:
		compare, _ := strconv.Atoi(args[0])
		return len(fdata) > compare
	}
	return false
}

/**
 * 适用数字和字符串
 */
func eq(f *Field, args ...string) bool {
	switch fdata := f.Val.(type) {
	case uint, uint8, uint16, uint32, uint64:
		compare, _ := strconv.ParseUint(args[0], 10, 64)
		return uint64(fdata.(uint)) == compare
	case int, int8, int16, int32, int64:
		compare, _ := strconv.ParseInt(args[0], 10, 64)
		return int64(fdata.(int)) == compare
	case float32, float64:
		compare, _ := strconv.ParseFloat(args[0], 64)
		return float64(fdata.(float64)) == compare
	case string:
		compare, _ := strconv.Atoi(args[0])
		return len(fdata) == compare
	}
	return false
}

/**
 * 适用数字和字符串
 */
func lt(f *Field, args ...string) bool {

	switch fdata := f.Val.(type) {
	case uint, uint8, uint16, uint32, uint64:
		compare, _ := strconv.ParseUint(args[0], 10, 64)
		return uint64(fdata.(uint)) < compare
	case int, int8, int16, int32, int64:
		compare, _ := strconv.ParseInt(args[0], 10, 64)
		return int64(fdata.(int)) < compare
	case float32, float64:
		compare, _ := strconv.ParseFloat(args[0], 64)
		return float64(fdata.(float64)) < compare
	case string:
		compare, _ := strconv.Atoi(args[0])
		return len(fdata) < compare
	}
	return false
}

/**
 * 字符串
 */
func empty(f *Field, args ...string) bool {
	switch fdata := f.Val.(type) {
	case string:
		len := len(fdata)
		if args[0] == "true" {
			return len == 0
		} else {
			return len != 0
		}
	}
	return false
}

/**
 * 适用数字和字符串 枚举 in=1,0  in=active,frozen
 */
func in(f *Field, args ...string) bool {
	switch fdata := f.Val.(type) {
	case uint, uint8, uint16, uint32, uint64:
		val := uint64(fdata.(uint))
		inSlice := strings.Split(args[0], ",")
		for _, v := range inSlice {
			compare, _ := strconv.ParseUint(v, 10, 64)
			if val == compare {
				return true
			}
		}
	case int, int8, int16, int32, int64:
		val := int64(fdata.(int))
		inSlice := strings.Split(args[0], ",")
		for _, v := range inSlice {
			compare, _ := strconv.ParseInt(v, 10, 64)
			if val == compare {
				return true
			}
		}
	case float32, float64:
		val := float64(fdata.(float64))
		inSlice := strings.Split(args[0], ",")
		for _, v := range inSlice {
			compare, _ := strconv.ParseFloat(v, 64)
			if val == compare {
				return true
			}
		}
	case string:
		inSlice := strings.Split(args[0], ",")
		for _, v := range inSlice {
			if fdata == v {
				return true
			}
		}
	}
	return false
}

/**
 * 数字：区间 section=min,max  min<val<max
 */
func section(f *Field, args ...string) bool {
	switch fdata := f.Val.(type) {
	case uint, uint8, uint16, uint32, uint64:
		val := uint64(fdata.(uint))
		if before, after, found := strings.Cut(args[0], ","); found {
			b, _ := strconv.ParseUint(before, 10, 64)
			a, _ := strconv.ParseUint(after, 10, 64)
			return val > b && val < a
		}
	case int, int8, int16, int32, int64:
		val := int64(fdata.(int))
		if before, after, found := strings.Cut(args[0], ","); found {
			b, _ := strconv.ParseInt(before, 10, 64)
			a, _ := strconv.ParseInt(after, 10, 64)
			return val > b && val < a
		}
	case float32, float64:
		val := float64(fdata.(float64))
		if before, after, found := strings.Cut(args[0], ","); found {
			b, _ := strconv.ParseFloat(before, 64)
			a, _ := strconv.ParseFloat(after, 64)
			return val > b && val < a
		}
	}
	return false
}

/**
 * 字符串
 */
func email(f *Field, args ...string) bool {
	switch fdata := f.Val.(type) {
	case string:
		reg := regexp.MustCompile(`\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`)
		return reg.MatchString(fdata)
	}
	return false
}

/**
 * 字符串
 */
func cn_mobile(f *Field, args ...string) bool {
	switch fdata := f.Val.(type) {
	case string:
		reg := regexp.MustCompile(`^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\d{8}$`)
		return reg.MatchString(fdata)
	}
	return false
}

/**
 * 字段比较
 */
func eq_field(f *Field, args ...string) bool {
	diff_field := f.RefValue.FieldByName(args[0])
	if f.Val == diff_field.Interface() {
		return true
	}
	return false
}
