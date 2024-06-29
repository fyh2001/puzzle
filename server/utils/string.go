package utils

import (
	"bytes"
	"unicode"
)

// CamelToSnake 将小驼峰命名转换为蛇形命名
func CamelToSnake(s string) string {
	var buf bytes.Buffer
	buf.WriteByte(byte(unicode.ToLower(rune(s[0]))))
	for i := 1; i < len(s); i++ {
		if unicode.IsUpper(rune(s[i])) {
			buf.WriteRune('_')
			buf.WriteByte(byte(unicode.ToLower(rune(s[i]))))
		} else {
			buf.WriteByte(s[i])
		}
	}
	return buf.String()
}
