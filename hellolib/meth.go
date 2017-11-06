package hellolib

import (
	"fmt"
	"sync"
	"reflect"
)

type N int

func (n N) toString() string {
	return fmt.Sprintf("%#x", n)
}

func TestMethod() {
	var a N = 25
	p := &a

	a.value()
	a.pointer()
	p.value()
	p.pointer()

	//var b * N
	//b.value() //b 相当于nil,可以调用，但是是错误的，因为没有指向具体实例
	//N{}.pointer()

	//d :=data{}
	//d.Lock()
	//defer d.Unlock()
}

func (n N) value() {
	n++
	fmt.Printf("v: %p,%v\n", &n, n)
}

func (n *N) pointer() {
	(*n)++;
	fmt.Printf("p: %p, %v\n", n, *n)
}

type data struct {
	sync.Mutex
}

type S struct{}

type T struct {
	S
}

func (S) sVal()  {}
func (*S) sPtr() {}
func (T) tVale() {}
func (*T) tPtr() {}

func MethodSet(a interface{}) {
	t := reflect.TypeOf(a)

	for i, n := 0, t.NumMethod(); i < n; i++ {
		m := t.Method(i)
		fmt.Println(m.Name, m.Type)
	}
}

func TestMethodSet() {
	var t T
	MethodSet(t)
	MethodSet(&t)
}
