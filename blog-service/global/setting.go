package global

// 将配置信息和应用程序关联起来

import "blog-service/pkg/setting"

// 配置的全局变量
var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
)
