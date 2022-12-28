package cmd

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/geekr-dev/go-cli-app/internal/timer"
	"github.com/spf13/cobra"
)

const FORMAT = "2006-01-02 15:04:05"

var calculateTime string
var duration string

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "时间格式处理",
	Long:  "时间格式处理",
}

// time now 子命令
var nowTimeCmd = &cobra.Command{
	Use:   "now",
	Short: "获取当前时间",
	Long:  "获取当前时间",
	Run: func(cmd *cobra.Command, args []string) {
		nowTime := timer.Now()
		log.Printf("输出结果：%s, %d", nowTime.Format(FORMAT), nowTime.Unix())
	},
}

// time calc 子命令
var calcTimeCmd = &cobra.Command{
	Use:   "calc",
	Short: "计算所需时间",
	Long:  "计算所需时间",
	Run: func(cmd *cobra.Command, args []string) {
		var current time.Time
		var err error
		var layout = FORMAT
		if calculateTime == "" {
			current = timer.Now()
		} else {
			if !strings.Contains(calculateTime, " ") {
				layout = "2006-01-02"
			}
			current, err = time.Parse(layout, calculateTime)
			if err != nil {
				t, _ := strconv.Atoi(calculateTime)
				current = time.Unix(int64(t), 0)
			}
		}
		calculateTime, err := timer.CalculateTime(current, duration)
		if err != nil {
			log.Fatalf("timer.GetCalculateTime err: %v", err)
		}

		log.Printf("输出结果：%s, %d", calculateTime.Format(layout), calculateTime.Unix())
	},
}

func init() {
	timeCmd.AddCommand(nowTimeCmd)
	timeCmd.AddCommand(calcTimeCmd)

	calcTimeCmd.Flags().StringVarP(&calculateTime, "calculate", "c", "", "需要计算的时间，可以是时间戳或者格式化后的时间")
	calcTimeCmd.Flags().StringVarP(&duration, "duration", "d", "", "持续时间, 有效单位为ns/us/ms/s/m/h")
}
