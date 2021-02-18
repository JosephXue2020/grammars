// Package main shows the practice on grammars.
package main

import (
	"fmt"
	"time"
)

// 主程序
func main() {
	// p1 := T{name: "John", age: 30}
	// fmt.Printf("The type of %+v is %T.\n", p1, p1)
	// var p2 T = T{name: "Kate", age: 20}
	// fmt.Printf("The type of %+v is %T.\n", p2, p2)
	// var p3 T
	// p3 = T{name: "Lily", age: 22}
	// fmt.Printf("The type of %+v is %T.\n", p3, p3)
	// p4 := new(T)
	// p4.name = "Rachel"
	// p4.age = 25
	// fmt.Printf("The type of %+v is %T.\n", p4, p4)

	// 引用包
	// fmt.Println("This is main package.")
	// a.PrintPackageName()
	// var input1, input2 string
	// fmt.Scanf("%s%s", &input1, &input2)
	// fmt.Println("The first input is:", input1)
	// fmt.Println("The second input is:", input2)

	// // 测试testing包
	// var b *testing.B
	// fmt.Printf("The type of *testing.B is %T.\n", b)
	// // fmt.Printf("%+v.\n", b.N)

	// // 测试通道
	// var t chan string
	// t = make(chan string)
	// t <- "xxxxxxxxxxxxxx"
	// go slowFunc(t)
	// msg := <-t
	// fmt.Printf(msg)

	// // 测试直接取函数调用的地址
	// test := &(aFunc())
	// fmt.Println(test)

	// 测试make函数
	// var arr [2]int
	// fmt.Printf("%+v\n", arr)
	// arr[1] = 99
	// fmt.Printf("%+v\n", arr)
	arr := make([]int, 2, 4)
	fmt.Printf("%+v\n", arr)
}

// T is some type.
type T struct {
	name string
	age  int
}

func slowFunc(c chan string) {
	time.Sleep(time.Second * 2)
	c <- "slowFunc finished."
}

func aFunc() int {
	a := 99
	return a + 3
}
