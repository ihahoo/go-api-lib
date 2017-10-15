package validate

import (
	"reflect"
	"regexp"
	"strconv"

	valid "github.com/asaskevich/govalidator"
)

const (
	mobile string = "^(\\+?0?86\\-?)?1[345789]\\d{9}$"
)

var (
	rxMobile = regexp.MustCompile(mobile)
)

// E ...
type E map[string]interface{}

// IsEmpty 检查值是否是空
func IsEmpty(s interface{}) bool {
	v := reflect.ValueOf(s)
	switch v.Kind() {
	case reflect.String, reflect.Array:
		return v.Len() == 0
	case reflect.Map, reflect.Slice:
		return v.Len() == 0 || v.IsNil()
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	default:
		return reflect.DeepEqual(s, reflect.Zero(v.Type()).Interface())
	}
}

// Required 检查是否为空
func Required(v interface{}, name string) (E, bool) {
	if IsEmpty(v) {
		return E{"status": 422, "code": 400001, "msg": name + "不能为空"}, false
	}

	return nil, true
}

// Email 检查是否是Email
func Email(v string, isRequired bool, name string) (E, bool) {
	if isRequired {
		if info, ok := Required(v, name); !ok {
			return info, false
		}
	}
	if !valid.IsEmail(v) {
		return E{"status": 422, "code": 400004, "msg": name + "格式不正确"}, false
	}

	return nil, true
}

// MinLength 字符的最小长度
func MinLength(v string, length int, name string) (E, bool) {
	if len(v) < length {
		return E{"status": 422, "code": 400005, "msg": name + "不能少于" + strconv.Itoa(length) + "个字符"}, false
	}

	return nil, true
}

// Mobile 手机号码
func Mobile(v string, isRequired bool, name string) (E, bool) {
	if isRequired {
		if info, ok := Required(v, name); !ok {
			return info, false
		}
	}
	rxMobile := regexp.MustCompile("^(\\+?0?86\\-?)?1[345789]\\d{9}$")
	if !rxMobile.MatchString(v) {
		return E{"status": 422, "code": 400021, "msg": name + "错误"}, false
	}

	return nil, true
}

// UUID uuid
func UUID(v string, isRequired bool, name string) (E, bool) {
	if isRequired {
		if info, ok := Required(v, name); !ok {
			return info, false
		}
	}
	if !valid.IsUUID(v) {
		return E{"status": 422, "code": 400002, "msg": name + "格式错误"}, false
	}

	return nil, true
}

// RealName 汉字姓名
func RealName(v string, isRequired bool, name string) (E, bool) {
	if isRequired {
		if info, ok := Required(v, name); !ok {
			return info, false
		}
	}
	rxRealName := regexp.MustCompile("^[\u2E80-\uFE4F]{2,10}$")
	if !rxRealName.MatchString(v) {
		return E{"status": 422, "code": 400002, "msg": name + "格式错误"}, false
	}

	return nil, true
}

// IDCard 身份证
func IDCard(v string, isRequired bool, name string) (E, bool) {
	if isRequired {
		if info, ok := Required(v, name); !ok {
			return info, false
		}
	}
	rxIDCard := regexp.MustCompile("(^\\d{18}$)|(^\\d{17}(\\d|X|x)$)")
	if !rxIDCard.MatchString(v) {
		return E{"status": 422, "code": 400002, "msg": name + "格式错误"}, false
	}

	return nil, true
}
