package options

import (
	"fmt"
	"github.com/spf13/cast"
	"reflect"
)

type StringOption func(interface{})

func ApplyStringOptions(s interface{}, opts ...StringOption) {
	for _, opt := range opts {
		opt(s) // 逐个应用 Option
	}
}

func WithStringFieldStr(fieldName string, fieldValue string) StringOption {
	return func(s interface{}) {
		v := reflect.ValueOf(s).Elem() // 获取结构体指针的值

		field := v.FieldByName(fieldName) // 根据字段名查找字段
		if field.IsValid() && field.CanSet() && field.Kind() == reflect.Int64 {
			field.SetInt(cast.ToInt64(fieldValue)) // 将字符串转换为 int64 并设置
		} else {
			fmt.Printf("Field %s is invalid or cannot be set\n", fieldName)
		}
	}
}
