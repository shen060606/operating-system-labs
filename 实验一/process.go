// 定义 PCB 结构体、进程状态常量、进程初始化等公共功能。

package main

type ProcessState string

const (
	Wait   ProcessState = "W"
	Run    ProcessState = "R"
	Finish ProcessState = "F"
)

type Process struct {
	Name         string       // 进程名
	Priority     int          // 优先数
	RequiredTime int          // 需要运行时间
	CPUTime      int          // 已运行 CPU 时间
	State        ProcessState // W、R 或 F
}

func newProcess(name string, priority, requiredTime int) *Process {
	return &Process{
		Name:         name,
		Priority:     priority,
		RequiredTime: requiredTime,
		CPUTime:      0,
		State:        Wait,
	}
}
