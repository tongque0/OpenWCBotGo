package task

import (
	"fmt"
	"time"

	"github.com/eatmoreapple/openwechat"
)

type Task struct {
	Name   string
	Run    func(*Task) // 修改 Run 方法，接收 Task 自身
	RunAt  time.Time
	Update func(*Task) // 新增 Update 方法来更新 RunAt
}

var tasks []Task

func addTask(task Task) {
	tasks = append(tasks, task)
}

// taskInit 初始化任务
func taskInit(self *openwechat.Self) {
	addTask(Task{
		Name: "DailyTask",
		Run: func(t *Task) {
			fmt.Println("执行每日任务")
			t.Update(t) // 执行后更新下一次执行时间
		},
		RunAt: dayRunTime(17, 51, 0),
		Update: func(t *Task) { // 每天固定时间任务的更新逻辑
			t.RunAt = dayRunTime(8, 0+randnumber(), 0+randnumber())
		},
	})
	addTask(Task{
		Name: "NextTask",
		Run: func(t *Task) {
			fmt.Println("执行每日任务1")
			t.Update(t)
		},
		RunAt: nextRunTime(0, 0, 1),
		Update: func(t *Task) {
			t.RunAt = nextRunTime(0, 0, 1)
		},
	})
}
