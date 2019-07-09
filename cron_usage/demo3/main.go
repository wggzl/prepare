package main

import (
	"context"
	"os/exec"
	"time"
	"fmt"
)

type result struct {
	err error
	output []byte
}

func main() {
	//执行一个cmd 让它在一个协程里去执行，让它执行2秒 sleep 2;echo hello;

	//1秒的时候 我们杀死cmd

	var (
		ctx context.Context
		cancelFunc context.CancelFunc
		cmd *exec.Cmd
		resultChan chan *result
		res *result
	)

	//创建一个结果队列
	resultChan = make(chan *result, 1000)

	//context chan byte
	//cancelFunc close(chan byte)
	ctx, cancelFunc = context.WithCancel(context.TODO())

	go func() {
		var (
			output []byte
			err error
		)
		cmd = exec.CommandContext(ctx, "C:\\cygwin64\\bin\\bash.exe", "-c", "sleep 2;echo hello;")

		//select {case <- ctx.Done}
		//kill pid 进程id 杀死子进程
		//捕获输出
		output, err = cmd.CombinedOutput()

		//把任务输出传给main协程
		resultChan <- &result{
			err:err,
			output:output,
		}

	}()

	//继续执行
	time.Sleep(1 * time.Second)

	//取消上下文
	cancelFunc()

	//在main协程 等待子协程退出 并打印结果
	res = <- resultChan

	fmt.Println(res.err, string(res.output))
}
