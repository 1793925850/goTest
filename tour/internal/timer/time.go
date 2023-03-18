package timer

import "time"

// GetNowTime 获得指定时区的当前时间
func GetNowTime() time.Time {
	// LoadLocation 根据名称获取特定时区的 Location 实例
	location, _ := time.LoadLocation("Asia/Shanghai")

	return time.Now().In(location)
}

// GetCalculateTime 根据指定持续时间来推算时间
func GetCalculateTime(currentTimer time.Time, d string) (time.Time, error) {
	// ParseDuration 解析持续时间(可以解析后缀来确定时间的单位)
	duration, err := time.ParseDuration(d)
	if err != nil {
		return time.Time{}, err // 解析出错，则返回一个空的 Time 的实例和错误
	}

	return currentTimer.Add(duration), nil
}
