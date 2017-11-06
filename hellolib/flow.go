package hellolib

import "fmt"

func TestIF() {

	x := 3

	if x > 3 {
		fmt.Println("A")
	} else if x < 3 && x > 0 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

	if a, b := x+1, x+10; a < b {
		fmt.Println(a, b)
	} else {
		fmt.Println("Test")
	}

	if InitIF(); x < 10 {
		fmt.Println("initif")
	}

	for x < 5 { // 只有条件表达式
		println(x)
		x++
	}

	for { //没有条件表达式
		println(x)
		x--

		if x == 0 {
			break
		}
	}
}

func InitIF() {
	fmt.Println("InitFI is called!")
}

func TestSwitch() {
	a, b, c, x := 1, 2, 3, 2

	switch x {
	case a, b:
		fmt.Println("a|b")
	case c:
		fmt.Println("c")
	case 4:
		fmt.Println("d")
	default:
		fmt.Println("z")
	}

	switch y := 5; y {
	default:
		y += 100
		fmt.Println(y)
	case 5:
		y += 50
		fmt.Println(y)
	}

	switch z := 20; z {
	case 20:
		z += 10
		fmt.Println(z)
		fallthrough
	case 21:
		z += 88
		fmt.Println(z)
	default:
		fmt.Println("default") //即使使用了fallthrough，而且defualt在case21下，
		// 也不会被执行
	}

	switch z1 := 5; { //没有条件表达式可以将swtich当成if-else使用
	case z1 > 5:
		fmt.Println("z1>5")
	case z1 > 0 && z1 <= 5:
		fmt.Println("z1>0 && z1<5")
	case z1 <= 0:
		fmt.Println("z1>=2 && z1<5")
	default:
		fmt.Println("Test")
	}
}

func TestFor() {
	for i, c := 0, count(); i < c; i++ { //c的作用域只在for循环中
		println("a", i)
	}

	c := 0 //新的变量c
	for c < count() {
		println("b", c)
		c++
	}

}

func count() int {
	print("count.")
	return 3
}

func TestRange() {
	data := [3]string{"a", "b", "c"}

	for i, s := range data {
		println(i, s)

	}
	for i := range data {
		println(i, data[i])
	}
	for _, s := range data {
		println(s)
	}
	for range data {
	}
	fmt.Println(data)
}
