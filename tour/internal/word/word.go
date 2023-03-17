package word

import (
	"strings"
	"unicode"

	"golang.org/x/text/cases"    // cases 包提供了通用的和特定语言的案例映射器，也就是一种映射方法
	"golang.org/x/text/language" // language 包实现了 BCP47 语言标签和相关功能
)

// ToUpper 把单词全部转为大写
func ToUpper(s string) string {
	// 将 s 中的所有元素全部大写
	return strings.ToUpper(s)
}

// ToLower 把单词全部转为小写
func ToLower(s string) string {
	// 将 s 中的所有元素全部小写
	return strings.ToLower(s)
}

// UnderscoreToUpperCamelCase 下划线单词转大写驼峰单词
func UnderscoreToUpperCamelCase(s string) string {
	// 如果第四个参数小于0，则对替换前后的数量没有限制
	s = strings.Replace(s, "_", " ", -1)
	// Caser 是一种映射样式
	s = cases.Title(language.English).String(s)

	return strings.Replace(s, " ", "", -1)
}

// UnderscoreToLowerCamelCase 下划线单词转小写驼峰单词
func UnderscoreToLowerCamelCase(s string) string {
	s = UnderscoreToUpperCamelCase(s)

	// rune 就是 int32
	// ToLower 返回对应的小写字母
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

// CamelCaseToUnderscore 驼峰单词转下划线单词
func CamelCaseToUnderscore(s string) string {
	var output []rune

	// 每个字符都进行判断然后添加
	for i, r := range s {
		if i == 0 {
			output = append(output, unicode.ToLower(r))
			continue
		}
		if unicode.IsUpper(r) { // 返回字符是否是大写字母
			output = append(output, '_')
		}
		output = append(output, unicode.ToLower(r))
	}

	return string(output)
}
