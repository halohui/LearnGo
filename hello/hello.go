package main

import (
	"fmt"
)

func main() {

	//hellolib.TestInt()
	//hellolib.TestFloat()
	//hellolib.MkMap()
	//hellolib.MkSlice()

	//hellolib.TestStruct()
	//hellolib.TestSlice()
	//hellolib.TestMap()

	//hellolib.TestIF()
	//hellolib.TestSwitch()
	//hellolib.TestFor()
	//hellolib.TestString()

	//hellolib.TestRange()

	//hellolib.TestVar()

	//hellolib.TestComplex()

	//hellolib.TestPtr()

	//hellolib.TestDiv()
	//hellolib.TestPanic()
	//hellolib.TestMethod()

	//var t hellolib.T
	//hellolib.MethodSet(t)
	//hellolib.MethodSet(&t)
	//hellolib.TestSay()

	//fmt.Println(runtime.NumCPU())
	//hellolib.TestChanVar()
	//hellolib.TestDeferRegister()
	//
	//hellolib.TestMultiPara("abc",1,2,3,4,5)
	//
	//hellolib.TestStruct()
	//

	//hellolib.TestSort()

	a := 1.0 << 3
	fmt.Println(a)
	var s uint = 3
	//b := 1.0 << s //shift count type unit, must be unsigned integer
	b := 1 << s
	fmt.Println(s, b)
	fmt.Printf("%T,%T\n", a, b) // int

	fmt.Println(Add88(12,34))

}

func (n N) test() {
	fmt.Printf("test.n:,%p,%d\n", &n, n)
}

type N int

func (n N) value() {}

func (n *N) pointer() {}

func call(m func()) {
	m()
}

func Add88(x,y int) (z int){
	{
		z:=x+y  //重新定义新的变量
		fmt.Println(z)
		return x  //不可以直接使用隐式return
	}
	return 100
}