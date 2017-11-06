package hellolib

import (
	"fmt"
	"sync"
)

func LocalPtr() *int {
	x := 0x100
	return &x
}

func A() {} // 空函数
func B() {} // 空函数

func C(a, b int) {
	fmt.Println("Test")
}

func Add(x, y int) (z int) {
	{
		z := x + y
		return z //不可以直接使用隐式return
	}

	return
}

func TestDefer() (z int) {
	defer func() {
		println("defer", z) //100
		z += 100
	}()
	return 100 //200
}

func TestFunc() {
	fmt.Println(LocalPtr()) //变成了堆内存,0xc42007a078
	fmt.Println(A == nil)   //false

}

var m sync.Mutex

func TestReturn(x int) int {
	for {
		break
	}
	return 0
}

func TestDeferRegister() {
	x, y := 1, 2
	defer func(a int) {
		println("defer x,y =", a, y)
	}(x)

	x += 100
	y += 200

	println(x, y)
}

func TestMultiPara(s string, a ...int) {
	fmt.Println(s, a)
}
