package main

import (
"os/exec"
"fmt"
)

func main()  {
	var (
		cmd *exec.Cmd
		err error
	)
	cmd = exec.Command("C:\\cygwin64\\bin\\bash.exe", "-c", "echo 1")
	err = cmd.Run()

	fmt.Println(err)
}
