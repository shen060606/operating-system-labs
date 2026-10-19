// 实现先来先服务（FCFS）调度算法。

package main

func fcfs(queue, allProcesses []*Process) {
	step := 0
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		step++

		current.State = Run
		current.CPUTime++

		if current.CPUTime >= current.RequiredTime {
			current.State = Finish
		} else {
			current.State = Wait
			queue = append(queue, current)
		}

		displayStep("FCFS", step, current, queue, allProcesses)
	}
}
