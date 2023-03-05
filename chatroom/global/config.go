package global

import (
	"github.com/fsnotify/fsnotify" // fsnotify 库用于在系统上提供跨平台文件系统通知
	"github.com/spf13/viper"
)

var (
	SensitiveWords []string

	MessageQueueLen = 1024
)

func initConfig() {
	viper.SetConfigName("chatroom")          // SetConfigName 设置配置文件的名称。不包括扩展
	viper.AddConfigPath(RootDir + "/config") // 添加配置文件的路径，以便在该路径下搜索配置文件

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	// GetStringSlice 提供字符串切片
	// GetInt 提供整型数
	SensitiveWords = viper.GetStringSlice("sensitive") // 提供敏感词
	MessageQueueLen = viper.GetInt("message-queue")    // 提供消息队列长度

	viper.WatchConfig() // WatchConfig 监视配置文件的更改
	viper.OnConfigChange(func(e fsnotify.Event) { // 当一个配置文件被改动时，OnConfigChange 用来设置这个 event 的 handler
		viper.ReadInConfig() // 在指定路径下加载配置文件

		SensitiveWords = viper.GetStringSlice("sensitive")
	})
}
