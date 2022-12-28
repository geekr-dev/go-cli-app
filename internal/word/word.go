package word

import (
	"strings"
	"unicode"
)

// 单词转化为大写
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// 单词转化为小写
func ToLower(s string) string {
	return strings.ToLower(s)
}

// 下划线转大写驼峰
func SnakeToCamelCase(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	s = strings.Title(s)
	return strings.Replace(s, " ", "", -1)
}

// 下划线转小写驼峰
func SnakeToLowerCamelCase(s string) string {
	s = SnakeToCamelCase(s)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

// 驼峰转下划线
func CamelCaseToSnake(s string) string {
	var output []rune
	for i, r := range s {
		if i == 0 {
			output = append(output, unicode.ToLower(r))
			continue
		}
		if unicode.IsUpper(r) {
			output = append(output, '_')
		}
		output = append(output, unicode.ToLower(r))
	}
	return string(output)
}
