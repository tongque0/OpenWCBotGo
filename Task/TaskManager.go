package task

import (
	"math/rand"
	"time"

	"github.com/eatmoreapple/openwechat"
)
// TaskManager 任务管理器
func TaskManager(self *openwechat.Self) {
	taskInit(self)
	for {
		now := time.Now()
		for i := range tasks {
			task := &tasks[i]
			if now.After(task.RunAt) {
				go task.Run(task) // 执行任务
			}
		}
		time.Sleep(1 * time.Second)
	}
}

// 计算固定时间间隔的下一次执行时间
func nextRunTime(hours, minutes, seconds int) time.Time {
	interval := time.Duration(hours)*time.Hour + time.Duration(minutes)*time.Minute + time.Duration(seconds)*time.Second
	return time.Now().Add(interval)
}

// 计算每天特定时间的下一次执行时间
func dayRunTime(hour, minute, second int) time.Time {
	now := time.Now()
	next := time.Date(now.Year(), now.Month(), now.Day(), hour, minute, second, 0, now.Location())

	// 如果今天的执行时间已经过了，设置为明天
	if now.After(next) {
		next = next.Add(24 * time.Hour)
	}
	return next
}

func randnumber() int {
	rand.Seed(time.Now().UnixNano())
	randomInt := rand.Intn(10)
	return randomInt
}
