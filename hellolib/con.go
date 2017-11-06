package hellolib

import (
	"fmt"
	"runtime"
)

func Say(s string) {
	for i := 0; i < 2; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

func TestSay() {
	go Say("xiao") //两个goroutine 运行在同一个进程里面
	Say("world!")
}
