package util

import (
	"crypto/md5"   // md5 包实现了MD5哈希算法
	"encoding/hex" // hex 包实现了16进制字符表示的编解码
)

/**
MD5 加密
*/

// EncodeMD5 用 md5 加密
func EncodeMD5(value string) string {
	m := md5.New() // 返回一个新的使用MD5校验的hash.Hash接口
	m.Write([]byte(value))

	// Sum 返回数据data的MD5校验和
	// EncodeToString 将数据 src 编码为字符串s
	return hex.EncodeToString(m.Sum(nil))
}
