package setting

/**
对读取配置的行为进行封装
*/

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Setting 是操作配置文件的工具
type Setting struct {
	// Viper 是一个优先配置注册表。它维护一组配置源，获取值来填充这些配置源，并根据源的优先级提供这些配置源
	// 简单地说，Viper 就是个帮忙找东西的
	vp *viper.Viper
}

// NewSetting 用于初始化本项目配置的基础属性
func NewSetting(configs ...string) (*Setting, error) {
	vp := viper.New()

	// SetConfigName 设置要查找的配置文件的名称。不包括扩展。
	vp.SetConfigName("config")
	// config 是配置文件所在的文件夹
	for _, config := range configs {
		if config != "" {
			// AddConfigPath 为 Viper 添加一个路径，以便在中搜索配置文件
			// 可以添加多条配置路径
			vp.AddConfigPath(config)
		}
	}
	// SetConfigType 设置要查找的配置文件的文件类型
	vp.SetConfigType("yaml")

	// ReadInConfig 加载配置文件
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}

	s := &Setting{vp}
	s.WatchSettingChange()

	return s, nil
}

// WatchSettingChange 监视设置更改
func (s *Setting) WatchSettingChange() {
	go func() {
		// WatchConfig 开始监视配置文件
		s.vp.WatchConfig()
		// OnConfigChange 设置配置文件更改时所调用的事件处理程序。
		s.vp.OnConfigChange(func(in fsnotify.Event) {
			_ = s.ReloadAllSection() // 配置文件发生更改时，重新加载配置文件中的所有数据
		})
	}()
}
