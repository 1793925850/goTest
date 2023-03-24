package setting

// 用于声明配置属性的结构体，并编写读取区段配置的配置方法

import "time"

type ServerSettingS struct {
	RunMode       string
	HttpPort      string
	ReadTimeout   time.Duration
	WriterTimeout time.Duration
}

type AppSettingS struct {
	DefaultPageSize int
	MaxPageSize     int
	LogSavePath     string
	LogFileName     string
	LogFileExt      string
}

type DatabaseSettingS struct {
	DBType       string
	Username     string
	Password     string
	Host         string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

// ReadSection 读取区段配置
func (s *Setting) ReadSection(k string, v interface{}) error {
	// UnmarshalKey 获取单个键并将其解组为 Struct
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}

	return nil
}

func (s *Setting) ReloadAllSection() error {}
