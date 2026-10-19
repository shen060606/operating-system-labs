// 实现最高优先数优先算法，包括优先数动态调整。

package main

import (
	"sort"
)

func sortPcb(pcb []*Process) *Process {
	sort.SliceStable(pcb, func(i, j int) bool { return pcb[i].Priority > pcb[j].Priority })
	return pcb[0]
}

func priority(queue, allProcesses []*Process) {
	step := 0
	for len(queue) > 0 {
		current := sortPcb(queue)
		queue = queue[1:]
		step++

		current.State = Run
		current.CPUTime++

		if current.CPUTime >= current.RequiredTime {
			current.State = Finish
		} else {
			current.State = Wait
			current.Priority--
			queue = append(queue, current)
		}

		displayStep("最高优先数优先", step, current, queue, allProcesses)
	}
}
