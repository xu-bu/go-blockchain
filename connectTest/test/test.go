package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background() // 创建一个空的背景上下文

	// 创建可取消的上下文
	ctx, cancel := context.WithCancel(ctx)

	// 将上下文作为参数传递给其他函数
	go performTask(ctx)

	// 模拟一些操作，然后在一定时间后取消操作
	time.Sleep(3 * time.Second)
	cancel() // 取消操作

	// 等待一段时间，以观察被取消的操作是否终止
	time.Sleep(2 * time.Second)
}

func performTask(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("任务已被取消")
			return
		default:
			fmt.Println("正在执行任务...")
			time.Sleep(1 * time.Second)
		}
	}
}
