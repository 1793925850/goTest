package setting

/**
用于声明配置属性的结构体，并编写读取区段配置的配置方法
*/

import "time"

// ServerSettingS 服务器端配置
type ServerSettingS struct {
	RunMode       string        // 运行模式
	HttpPort      string        // 服务端监听端口号
	ReadTimeout   time.Duration // 读超时
	WriterTimeout time.Duration // 写超时
}

// AppSettingS 应用程序配置
type AppSettingS struct {
	DefaultPageSize      int      // 默认每页数量
	MaxPageSize          int      // 最大每页数量
	LogSavePath          string   // 日志保存路径
	LogFileName          string   // 日志文件名称
	LogFileExt           string   // 日志文件类型
	UploadSavePath       string   // 上传文件的最终保存目录
	UploadServerUrl      string   // 上传文件后用于展示的文件服务地址
	UploadImageMaxSize   int      // 上传文件所允许的最大空间大小（MB）
	UploadImageAllowExts []string // 上传文件所允许的文件后缀
}

type DatabaseSettingS struct {
	DBType       string // 数据库类型
	Username     string // 数据库用户名
	Password     string // 数据库密码
	Host         string // 数据库连接端口
	DBName       string // 数据库名
	TablePrefix  string // 表前缀
	Charset      string // 字符集
	ParseTime    bool   // 表示是否将 MySql 的时间自动转换为 time.Time；如果为 false，则转换为 []byte 或 string
	MaxIdleConns int    // 连接池中保持的最大空闲连接数
	MaxOpenConns int    // 打开数据库的最大连接数
}

// sections 记录区段键，同时也用来保存区段配置的值
var sections = make(map[string]interface{})

// ReadSection 读取区段配置
func (s *Setting) ReadSection(k string, v interface{}) error {
	// UnmarshalKey 获取单个键并将其解组为 Struct
	// UnmarshalKey 根据 k，来找到数据并解析加载到变量 v 中
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}

	if _, ok := sections[k]; !ok {
		sections[k] = v
	}

	return nil
}

// ReloadAllSection 读取所有区段配置
func (s *Setting) ReloadAllSection() error {
	for k, v := range sections {
		err := s.ReadSection(k, v)
		if err != nil {
			return err
		}
	}

	return nil
}
