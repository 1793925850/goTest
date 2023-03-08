package logic

import (
	"strings" // strings 包用来对 string 进行操作

	"chatroom/global"
)

// 过滤敏感词
func FilterSensitive(content string) string {
	for _, word := range global.SensitiveWords { // 遍历敏感词表
		// ReplaceAll 返回将 content 中所有不重叠 word 子串都替换为 new 的新字符串
		// 即将 content 中的所有 word 替换成 new
		content = strings.ReplaceAll(content, word, "**") // 如果 content 中有 word，则将 content 的词改为 **
	}

	return content
}
