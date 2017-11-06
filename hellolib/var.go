package hellolib

import (
	"fmt"
	"math"
	"strconv"
	"unsafe"
)

var tt = 1000

func TestConst() {

	const (
		x uint16 = 120
		y
		s = "abc"
		z
	)

	fmt.Println(x, y, s, z) //120,120,abc,abc

	const (
		x1 = iota
		x2
		x3
		x4
	)

	fmt.Println(x1, x2, x3, x4) //0,1,2,3

	const (
		_  = iota //0
		KB = 1 << (10 * iota)
		MB
		GB
	)

	fmt.Println(KB, MB, GB) //1024 1048576 1073741824

	const (
		a = iota
		b
		c = 100
		d        //100
		e = iota // 4，按行索引器
		f        //5
	)

	fmt.Println(a, b, c, d, e, f) //0 1 100 100 4 5

	type color byte
	const (
		black color = iota
		red
		blue
	)

	fmt.Println(black, red, blue) //0,1,2

	const (
		t1, t2 int  = 99, -999
		t3     byte = byte(t1) //t1 指定为int型，需显式转换为byte类型
		//t4 = uint8(t2)  // 常量-999,超出了unit8的范围

	)

	fmt.Println(t1, t2, t3)

	const (
		ptrSize = unsafe.Sizeof(uintptr(0)) //8
		strSize = len("hello world!")       //12
	)

	fmt.Println(ptrSize, strSize)

	const (
		_, _   = iota, iota
		aa, bb
		cc, dd
	)
	fmt.Println(aa, bb, cc, dd) // 1 1 2 2

	const (
		xx         = iota
		yy float32 = iota
		zz         = iota
	)
	fmt.Println(xx, yy, zz)

	const (
		y2 float32 = iota
		z2         = iota
	)
	fmt.Println(y2, z2) //0 1 2

	fmt.Printf("%T", y2)
}

func TestInt() {

	a, b, c := 100, 0144, 0x64
	fmt.Println(a, b, c)                    //100 100 100
	fmt.Printf("%b, %o,%x\n", a, a, a)      //0b1100100, 0144,0x64
	fmt.Println(math.MinInt8, math.MaxInt8) //-128,127

	x, _ := strconv.ParseInt("1100100", 2, 32) //
	y, _ := strconv.ParseInt("0144", 8, 32)
	z, _ := strconv.ParseInt("64", 16, 32)

	fmt.Println(x, y, z) //100 100 100

	fmt.Println(strconv.FormatInt(x, 2))
	fmt.Println(strconv.FormatInt(x, 8))
	fmt.Println(strconv.FormatInt(x, 16))

}

func TestFloat() {
	var a float32 = 1.1234567899
	var b float32 = 1.12345678
	var c float32 = 1.123456781

	fmt.Println(a, b, c)              // 1.1234568 1.1234568 1.1234568
	fmt.Println(a == b, b == c)       //true true
	fmt.Printf("%v %v %v\n", a, b, c) //1.1234568 1.1234568 1.1234568
}

func MkMap() map[string]int {
	m := make(map[string]int)
	m["a"] = 1

	fmt.Println(m) //map[a:1]

	//p := new(map[string]int)
	//t := *p
	//t["a"] = 100
	//fmt.Println(t) //panic: assignment to entry in nil map

	return m
}

func MkSlice() []int {
	s := make([]int, 0, 10)
	s = append(s, 100)

	fmt.Println(s) //[100]

	return s
}

func TypeConvert() {
	x := 100
	p := (*int)(&x)

	fmt.Println(p)
}

type node struct {
	_    int
	id   int
	next *node
}

func TestStruct() {
	var a struct {
		x int    `x`
		s string `s`
	}

	var b struct {
		x int
		s string
	}

	//b=a  //cannot use a (type struct { x int "x"; s string "s" })
	// as type struct { x int; s string } in assignment

	fmt.Println(a.x, b.x)

	n1 := node{
		id: 1,
	}

	n2 := node{
		id:   2,
		next: &n1,
	}

	fmt.Println(n1, n2)
}


func TestSlice() {
	var arr = [10]byte{81, 82, 83, 84, 85, 86} //array

	var a, b, c []byte //slice

	a = arr[1:3]
	b = arr[2:4]

	fmt.Println(a, b) //[98 99] [99 100]

	c = a[0:2]

	fmt.Println(c) //[82,83]

	a = arr[1:5]

	fmt.Println(a, c) //[82 83 84 85] [82 83]

	//append(c,34)

	fmt.Println(a)
}

func TestMap() {
	rating := map[string]float32{"C": 5, "Go": 1.83}

	fmt.Println(rating) //map[C:5 Go:1.83]

	m := make(map[string]string)
	m["Hello"] = "hello"
	m1 := m
	m1["Hello"] = "world"

	fmt.Println(m, m1) //map[Hello:world] map[Hello:world]
}

func TestString() {
	s := "hello"
	s = "c" + s[1:]

	fmt.Println(s)
}

func TestVar() {
	var t = 011i
	fmt.Println(t) //0+11i

	str := "012"
	fmt.Println(str[1]) //字符串可以利用索引求得

	var d1 [4]int //数组
	var d2 [4]int //数组

	d3 := [4]int{2, 3, 4}  //数组
	d4 := [4]int{5, 3: 10} //数组

	fmt.Println(d1 == d2, d3 == d4) //true false

	fmt.Println(&d4, &d4[0], &d4[3]) //&[5 0 0 10] 0xc4200680e0 0xc4200680f8

	d5 := [...]int{34, 8, 9} //数组，其长度为3
	d6 := [...]int{5, 3: 10} //数组，其长度为4，内容为[5 0 0 10]，可以在特定位置初始化
	d7 := [3][4]int{{2}, {3}}

	d11, d12 := 20, 10
	d13 := []*int{&d11, &d12} //指针数组
	d14 := &d13               //数组指针
	fmt.Println(d13, d14)     //[0xc420072250 0xc420072258] &[0xc420072250 0xc420072258]

	fmt.Println(d1, d2, d3, d4, d5, d6) //[0 0 0] [0 0 0 0] [2 3 4 0 0]
	fmt.Println(len(d7), cap(d7))       //3 3 求得第一维的长度
}

func TestComplex() {

	x := 0 + 9i //复数的默认类型是complex128
	fmt.Printf("%T %v", x, x)

	y := 1E6 + 0.i //(1e+06+0i)猜测复数的实部和虚部都是float64

	fmt.Println(y)
}

func TestOwnType() {
	type X int

	var m X = 23

	if m < 100 { // < not defined on X
		fmt.Println("Test")
	}
}

func TestPtr() {
	x := 100
	p := &x

	//p++  //不能加
	//p+1

	var p2 = p
	fmt.Println(p2 == p)
}

func TestChanVar() {
	var a, b chan int = make(chan int, 3), make(chan int)
	var c chan bool

	println(a == b)
	println(c == nil)

	fmt.Printf("%p, %d\n", a, unsafe.Sizeof(a))
}
