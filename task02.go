//go:build task02
// +build task02

// Goroutine
// 题目1：使用 go 关键字启动两个协程，一个打印 1-10 的奇数，另一个打印 2-10 的偶数。
// 题目2：设计一个任务调度器，接收一组任务（函数），并发执行并统计每个任务的耗时。

package main

import (
	"fmt"
	"sync"
	"time"
)

// printOdds 并发打印 [start, end] 范围内的奇数
func printOdds(start, end int, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	for i := start; i <= end; i++ {
		if i%2 == 1 {
			fmt.Printf("odd: %d\n", i)
			// 轻微休眠，便于观察并发交错输出（非必须）
			time.Sleep(5 * time.Millisecond)
		}
	}
}

// printEvens 并发打印 [start, end] 范围内的偶数
func printEvens(start, end int, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	for i := start; i <= end; i++ {
		if i%2 == 0 {
			fmt.Printf("even: %d\n", i)
			time.Sleep(5 * time.Millisecond)
		}
	}
}

// Task 表示可执行任务
type Task func()

// TaskResult 保存任务执行结果（名称与耗时）
type TaskResult struct {
	Name     string
	Duration time.Duration
}

// RunTasksConcurrently 并发执行任务集合，返回每个任务的耗时
func RunTasksConcurrently(tasks map[string]Task) []TaskResult {
	var waitGroup sync.WaitGroup
	resultsChan := make(chan TaskResult, len(tasks))

	for name, task := range tasks {
		waitGroup.Add(1)
		go func(taskName string, t Task) {
			defer waitGroup.Done()
			startedAt := time.Now()
			t()
			resultsChan <- TaskResult{
				Name:     taskName,
				Duration: time.Since(startedAt),
			}
		}(name, task)
	}

	waitGroup.Wait()
	close(resultsChan)

	results := make([]TaskResult, 0, len(tasks))
	for r := range resultsChan {
		results = append(results, r)
	}
	return results
}

func main() {
	// 题目1演示：两个协程并发打印奇偶数
	var wg sync.WaitGroup
	wg.Add(2)
	go printOdds(1, 10, &wg)
	go printEvens(2, 10, &wg)
	wg.Wait()

	fmt.Println("---- 并发任务调度器演示 ----")

	// 题目2演示：任务调度器并发执行任务并统计耗时
	demoTasks := map[string]Task{
		"task-sleep-50ms": func() {
			time.Sleep(50 * time.Millisecond)
		},
		"task-calc": func() {
			sum := 0
			for i := 0; i < 500000; i++ {
				sum += i
			}
			// 使用结果以避免编译器优化掉循环
			if sum == -1 {
				fmt.Println("impossible")
			}
		},
		"task-sleep-120ms": func() {
			time.Sleep(120 * time.Millisecond)
		},
	}

	results := RunTasksConcurrently(demoTasks)
	for _, r := range results {
		fmt.Printf("%s took %v\n", r.Name, r.Duration)
	}
}

// Goroutine
// 题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
// 考察点 ： go 关键字的使用、协程的并发执行。
// 题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
// 考察点 ：协程原理、并发任务调度。
