package main

import "os/exec"
import "fmt"

func main() {
	cmd := exec.Command("ls", "-l", "-h")
	cmdOut, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(cmdOut))
}
