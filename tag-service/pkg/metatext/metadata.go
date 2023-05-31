package metatext

import (
	"google.golang.org/grpc/metadata"
	"strings"
)

/**
基于 TextMap 模式，对照实现了 metadata 的设置和读取方法
*/

type MetadataTextMap struct {
	metadata.MD
}

func (m MetadataTextMap) ForeachKey(handler func(key, val string) error) error {
	for k, vs := range m.MD {
		for _, v := range vs {
			if err := handler(k, v); err != nil {
				return err
			}
		}
	}
	return nil
}

func (m MetadataTextMap) Set(key, value string) {
	key = strings.ToLower(key)
	m.MD[key] = append(m.MD[key], value)
}
