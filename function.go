package validate

import (
	"regexp"
	"strconv"
)

type CallFunc func(f *Field, args ...string) bool

var formatFunc = map[string]CallFunc{
	"email": email,
}

var expFunc = map[string]CallFunc{
	"gt":    gt,
	"eq":    eq,
	"lt":    lt,
	"empty": empty,
}

func gt(f *Field, args ...string) bool {
	switch fdata := f.Val.(type) {
	case uint, uint8, uint16, uint32, uint64:
		compare, _ := strconv.ParseUint(args[0], 10, 64)
		return uint64(fdata.(uint)) > compare
	case int, int8, int16, int32, int64:
		compare, _ := strconv.ParseInt(args[0], 10, 64)
		return int64(fdata.(int)) > compare
	case string:
		compare, _ := strconv.Atoi(args[0])
		return len(fdata) > compare
	}
	return false
}

func eq(f *Field, args ...string) bool {
	switch fdata := f.Val.(type) {
	case uint, uint8, uint16, uint32, uint64:
		compare, _ := strconv.ParseUint(args[0], 10, 64)
		return uint64(fdata.(uint)) == compare
	case int, int8, int16, int32, int64:
		compare, _ := strconv.ParseInt(args[0], 10, 64)
		return int64(fdata.(int)) == compare
	case string:
		compare, _ := strconv.Atoi(args[0])
		return len(fdata) == compare
	}
	return false
}

func lt(f *Field, args ...string) bool {

	switch fdata := f.Val.(type) {
	case uint, uint8, uint16, uint32, uint64:
		compare, _ := strconv.ParseUint(args[0], 10, 64)
		return uint64(fdata.(uint)) < compare
	case int, int8, int16, int32, int64:
		compare, _ := strconv.ParseInt(args[0], 10, 64)
		return int64(fdata.(int)) < compare
	case string:
		compare, _ := strconv.Atoi(args[0])
		return len(fdata) < compare
	}
	return false
}

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

func email(f *Field, args ...string) bool {
	switch fdata := f.Val.(type) {
	case string:
		reg := regexp.MustCompile(`\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`)
		return reg.MatchString(fdata)
	}
	return false
}
