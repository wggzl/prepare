package main

import (
	"os/exec"
	"fmt"
)

func main() {
	var (
		cmd *exec.Cmd
		output []byte
		err error
	)

	//生成cmd
	cmd = exec.Command("C:\\cygwin64\\bin\\bash.exe", "-c", "sleep 5;ls -l")
	//执行了命令，捕获子进程的输出(pipe)
	if output, err = cmd.CombinedOutput(); err != nil {
		fmt.Println(err)
		return
	}

	//打印子进程的输出
	fmt.Println(string(output))
}
