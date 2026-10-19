// 统一负责打印当前运行进程、就绪队列和所有 PCB 信息。

package main

import "fmt"

func displayStep(algorithm string, step int, current *Process, readyQueue, allProcesses []*Process) {
	fmt.Printf("\n[%s] 第 %d 个时间片，运行进程：%s\n", algorithm, step, current.Name)

	fmt.Print("就绪队列：")
	if len(readyQueue) == 0 {
		fmt.Println("（空）")
	} else {
		for i, p := range readyQueue {
			if i > 0 {
				fmt.Print("、")
			}
			fmt.Print(p.Name)
		}
		fmt.Println()
	}

	fmt.Println("所有进程的 PCB：")
	fmt.Printf("%-8s %-8s %-12s %-12s %-8s\n", "进程名", "优先数", "需要运行时间", "已运行时间", "状态")
	for _, p := range allProcesses {
		fmt.Printf("%-8s %-8d %-12d %-12d %-8s\n",
			p.Name, p.Priority, p.RequiredTime, p.CPUTime, p.State)
	}
}
