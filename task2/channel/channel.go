// 示例：使用通道在两个协程间通信（生产者/消费者模型）。
// - 生产者：生成 1..10 的整数并发送到通道
// - 消费者：从通道接收整数并打印；当通道关闭时结束
package main

import (
	"fmt"
	"sync"
)

// main 启动生产者与消费者两个协程，并等待其完成。
func main() {
	var wg sync.WaitGroup
	//编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。

	// 创建一个无缓冲通道，用于在生产者与消费者之间传递整数
	var ch = make(chan int)

	// 计数器置为 2，表示有两个协程会在完成时调用 Done()
	wg.Add(2)
	//生成
	go add_element(ch, &wg)
	//接收
	go reveive_element(ch, &wg)
	wg.Wait()

	//实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。

	//缓冲10的通道
	//无缓冲（make(chan T)）: 发送与接收必须同时发生，双方同步“握手”。发送方在没有接收者时会阻塞；接收方在没有发送者时也阻塞。
	// 有缓冲（make(chan T, n)）: 允许暂存至多 n 个元素。发送在缓冲未满时不阻塞；接收在缓冲不空时不阻塞。缓冲满/空时与无缓冲一致会阻塞。
	ch1 := make(chan int, 10)
	wg.Add(2)

	go producer(ch1, &wg)
	go consumer(ch1, &wg)
	wg.Wait()
}

func consumer(ch chan int, s *sync.WaitGroup) {
	defer s.Done()
	for value := range ch {
		fmt.Println("value:", value)
	}
}

func producer(ch chan int, s *sync.WaitGroup) {
	defer s.Done()
	for i := 1; i <= 100; i++ {
		ch <- i
	}
	close(ch)
}

// reveive_element 从只读通道 ch 持续接收数据并打印，直到通道被关闭。
// 参数说明：
// - ch：只读通道（<-chan int），消费者只能接收不能发送
// - wg：WaitGroup 指针，用于在函数结束时通知外部协程已完成
func reveive_element(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for value := range ch {
		fmt.Println("接收的整数：", value)
	}
}

// add_element 向只写通道 ch 依次发送 1..10，并在发送完成后关闭通道。
// 参数说明：
// - ch：只写通道（chan<- int），生产者只能发送不能接收
// - wg：WaitGroup 指针，用于在函数结束时通知外部协程已完成
func add_element(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		// 发送当前整数到通道
		ch <- i
	}
	// 关闭通道，告知消费者“不会再有新数据”
	close(ch)
}
