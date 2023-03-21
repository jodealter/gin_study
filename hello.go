package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建一个通道，用于接收通知
	ch := make(chan int)

	// 启动5个goroutine等待通知
	for i := 0; i < 5; i++ {
		go func(id int, ch chan int) {
			fmt.Printf("goroutine %d 等待通知...\n", id)
			<-ch // 阻塞等待通知
			fmt.Printf("goroutine %d 收到通知\n", id)
		}(i, ch)
	}

	// 等待1秒后通知指定的goroutine
	time.Sleep(1 * time.Second)
	fmt.Println("开始通知...")
	ch <- 3 // 通知ID为3的goroutine
	fmt.Println("通知结束")
	time.Sleep(1 * time.Second) // 等待所有goroutine输出
}
