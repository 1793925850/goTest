package cmd

import (
	"log"
	"strconv"
	"strings"
	"time"

	"tour/internal/timer"

	"github.com/spf13/cobra"
)

// 因为 time 中有两个子命令，所以专门设置一个 timeCmd 作为另外两个子命令的 Parent
var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "时间格式处理",
	Long:  "时间格式处理",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var nowTimeCmd = &cobra.Command{
	Use:   "now",
	Short: "获取当前时间",
	Long:  "根据指定时区，获取当前时间",
	Run: func(cmd *cobra.Command, args []string) {
		nowTime := timer.GetNowTime()
		log.Printf("输出结果：%s，%d", nowTime.Format("2006-01-01 15:04:05"), nowTime.Unix())
	},
}

var calculateTimeCmd = &cobra.Command{
	Use:   "calc",
	Short: "计算所需时间",
	Long:  "计算所需时间",
	Run: func(cmd *cobra.Command, args []string) {
		var currentTimer time.Time
		var layout = "2006-01-01 15:04:05"

		if calculateTime == "" {
			currentTimer = timer.GetNowTime()
		} else {
			var err error
			// Count 计数子串在主串中的非重叠出现次数；如果子串为空字符串，则返回主串的元素数+1
			space := strings.Count(calculateTime, " ")

			if space == 0 {
				layout = "2006-01-01"
			}
			if space == 1 {
				layout = "2006-01-01 15:04"
			}

			// Parse 解析一个格式化的时间字符串并返回它代表的时间
			currentTimer, err = time.Parse(layout, calculateTime) // 如果这里的 calculateTime 内容是时间戳的话，那这一句将失败
			if err != nil {                                       // 如果 calculateTime 里的时间解析失败，使用下列方法
				t, _ := strconv.Atoi(calculateTime)
				currentTimer = time.Unix(int64(t), 0) // Unix 直接将时间戳转化为当前时间格式
			}
		}

		t, err := timer.GetCalculateTime(currentTimer, duration)
		if err != nil { // 这里也就意味着 ParseDuration 解析失败
			log.Fatalf("timer.GetCalculateTime err: %v", err)
		}

		log.Printf("输出结果：%s，%d", t.Format(layout), t.Unix())
	},
}

var (
	calculateTime string // 需要被计算的时间
	duration      string // 持续时间
)

func init() {
	timeCmd.AddCommand(nowTimeCmd)
	timeCmd.AddCommand(calculateTimeCmd)

	calculateTimeCmd.Flags().StringVarP(&calculateTime, "calculate", "c", "", "需要计算的时间，有效单位为时间戳或已格式化后的时间")
	calculateTimeCmd.Flags().StringVarP(&duration, "duration", "d", "", `持续时间，有效时间单位为"ns", "us" (or "µs"), "ms", "s", "m", "h"`)
}
