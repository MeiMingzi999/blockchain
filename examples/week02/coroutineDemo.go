package main

import (
	"fmt"
	"sync"
	"time"
)

func printOddNum(limit int, wait *sync.WaitGroup) {
	defer wait.Done()
	for i := 0; i <= limit; i++ {
		if i%2 == 1 {
			fmt.Printf("奇数为: %d\n", i)
		}
	}

}

func printEvenNum(limit int, wait *sync.WaitGroup) {
	defer wait.Done()
	for i := 0; i <= limit; i++ {
		if i%2 == 0 {
			fmt.Printf("偶数为: %d\n", i)
		}
	}
}

type SchedulerType struct {
	task []Task
}

type Task struct {
	name string
	fn   func()
}

func scheduler(schedulertype *SchedulerType, wait *sync.WaitGroup) {
	wait.Add(len(schedulertype.task))
	for index, curTask := range schedulertype.task {
		go func(task Task, index int) {
			defer wait.Done()
			startTime := time.Now()
			task.fn()
			cost := time.Since(startTime)
			fmt.Printf("正在执行第 %d 个任务, 耗时: %v\n", index, cost)
		}(curTask, index)
	}
}

func main() {
	var wait sync.WaitGroup
	// wait.Add(2)
	// //奇偶数打印
	// go printEvenNum(100, &wait)
	// go printOddNum(100, &wait)
	// wait.Wait()
	// fmt.Println("程序执行完成!")

	//任务调度器
	st := SchedulerType{
		task: []Task{
			{
				name: "任务A",
				fn: func() {
					time.Sleep(200 * time.Millisecond)
				},
			},
			{
				name: "任务B",
				fn: func() {
					time.Sleep(400 * time.Millisecond)
				},
			},
			{
				name: "任务C",
				fn: func() {
					time.Sleep(300 * time.Millisecond)
				},
			},
		},
	}

	scheduler(&st, &wait)
	wait.Wait()
	fmt.Printf("全部任务执行完成!\n")

}
