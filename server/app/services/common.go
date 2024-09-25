package services

import (
	"fmt"
	"reflect"
)

func CheckReq(req interface{}) error {

	// 遍历结构体中的字段并自动检查
	v := reflect.ValueOf(req).Elem() // 获取结构体的值
	t := v.Type()                    // 获取结构体的类型

	for i := 0; i < v.NumField(); i++ {
		fieldValue := v.Field(i)
		fieldName := t.Field(i).Name

		// 检查字段是否有 `check` 标签
		if t.Field(i).Tag.Get("check") == "true" {
			// 使用之前的 checkField 函数来验证字段值
			if err := checkField(fieldValue.Interface(), fieldName); err != nil {
				return err
			}
		}
	}

	return nil
}

// 通用字段校验函数
func checkField(value interface{}, fieldName string) error {
	v := reflect.ValueOf(value)

	// 检查不同类型的零值
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if v.Int() == 0 {
			return fmt.Errorf("%s参数错误", fieldName)
		}
	case reflect.String:
		if v.String() == "" {
			return fmt.Errorf("%s参数错误", fieldName)
		}
	case reflect.Slice, reflect.Array:
		if v.Len() == 0 {
			return fmt.Errorf("%s参数错误", fieldName)
		}
	default:
		return fmt.Errorf("不支持的类型: %s", fieldName)
	}

	return nil
}
