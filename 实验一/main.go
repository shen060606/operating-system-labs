// 最终主程序：输入进程信息、选择算法、调用调度函数、展示结果。

package main

import "fmt"

func main() {
	fcfsAll := []*Process{
		newProcess("P1", 3, 2),
		newProcess("P2", 2, 1),
		newProcess("P3", 1, 3),
		newProcess("P4", 4, 2),
		newProcess("P5", 2, 4),
	}
	fcfsQueue := append([]*Process(nil), fcfsAll...)

	fmt.Println("========== FCFS 调度 ==========")
	fcfs(fcfsQueue, fcfsAll)

	priorityAll := []*Process{
		newProcess("P1", 3, 2),
		newProcess("P2", 2, 1),
		newProcess("P3", 1, 3),
		newProcess("P4", 4, 2),
		newProcess("P5", 2, 4),
	}
	priorityQueue := append([]*Process(nil), priorityAll...)

	fmt.Println("\n========== 最高优先数优先调度 ==========")
	priority(priorityQueue, priorityAll)
}
