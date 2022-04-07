package main

import (
	"fmt"
	"reflect"
	"time"
)

func foo() {
	defer fmt.Println("A")
	fmt.Println("C")
	defer func() {
		defer func() {
			fmt.Println("D")
		}()
		panic(2)
	}()
	defer func() {
		pv := recover()
		fmt.Println("recover err :", pv)
		fmt.Println("11:", reflect.TypeOf(pv))
	}()
	panic(1)

}

func Panic() {
	defer func() {
		pv := recover()
		fmt.Println("recover err :", pv)
		fmt.Println("11:", reflect.TypeOf(pv))
	}()

	foo()
	fmt.Println("end")
}

func test() {
	fmt.Println("A")
	defer func() {
		defer func() {
			fmt.Println("B")
		}()
		fmt.Println("C")
	}()
}

func test2(a []int) []int {
	s := a[:1]
	s[0] = 10
	s = append(s, 1)
	s = append(s, 1)
	s = append(s, 1)
	s = append(s, 1)
	s = append(s, 1)

	return s
}

type A interface {
	hello(a int, b int) int
}

type Student struct {
	Name string
	Age  int
}

func (s *Student) SetAge() {
	s.Age = 18
}
func (s Student) SetName() {
	s.Name = "yy"
}

func (s *Student) hello(a int, b int) int {

	return a + b
}

type IGreeting interface {
	sayHello()
}

func sayHello(i IGreeting) {
	i.sayHello()
}

type Go struct{}

func (g *Go) sayHello() {
	fmt.Println("我实现了这个方法")
}
func main() {
	// var _ IGreeting = (*Go)(nil)
	// g := new(Go)
	// sayHello(g)

	// f := map[string]int{"a": 13, "b": 24}
	// fmt.Println(f["a"])
	// if v, ok := f["a"]; ok {
	// 	fmt.Println(v)
	// }

	// a := []int{1, 2, 3, 4, 5, 6}
	// b := a[:4]
	// b = append(b, 1, 2, 3, 4, 5)
	// b[0] = 3
	// fmt.Println(b)
	// fmt.Println(a)

	// n := 0
	// f := func() int {

	// 	n++
	// 	return n
	// }
	// fmt.Println(f())
	// fmt.Println(n)
	// var wg sync.WaitGroup
	// for i := 0; i < 5; i++ {
	// 	wg.Add(1)
	// 	go func(i int) {
	// 		fmt.Println(i)
	// 		wg.Done()
	// 	}(i)
	// }
	// wg.Wait()

	// var s1 []int
	// s1 = append(s1, 1)
	// s := make([]int, 0)
	// fmt.Println(s)

	// a := []int{1, 2, 3}
	// a := make([]int, 0)
	// a = append(a, 1)
	// a = append(a, 1)
	// a = append(a, 1)
	// fmt.Printf("a原来的值%v \n", a)
	// b := test2(a)
	// fmt.Printf("函数操作之后的值%v \n", b)
	// fmt.Printf("a现在的值%v \n", a)
	// var a int
	// f := func(x int) {
	// 	a++
	// 	x = x + a
	// 	fmt.Println(a)
	// }
	// f(10)

	// a := 3
	// if a > 3 {
	// 	fmt.Println(a)
	// } else if a == 3 {
	// 	fmt.Println()
	// } else {

	// }

	// switch a {
	// case 1:
	// 	fmt.Println("1")
	// case 3:
	// 	fmt.Println("2")
	// 	fallthrough
	// case 2:
	// 	fmt.Println("3")
	// default:
	// 	fmt.Println()

	// }

	// a和b是值类型
	// a, b := 0, 1
	// // c和d是引用类型，指向a和b的指针变量
	// c := &a
	// d := &b

	// // 指针变量的赋值操作
	// *c = 4
	// *d = 6
	// // c1和d1是引用变量
	// c1 := c
	// d1 := d
	// // 通过引用变量修改数据
	// *c1 = 8
	// *d1 = 10

	// student := &Student{
	// 	Name: "aa",
	// 	Age:  0,
	// }
	// // 引用类型的方法
	// student.SetAge()
	// // 值方法
	// student.SetName()

	// fmt.Println(student)

	// var a interface{}

	// a = 3
	// fmt.Println(a)
	// if v, ok := a.(int); ok {
	// 	fmt.Println("a是整型", v)
	// }
	// su := new(Student)
	// // su.hello(1, 2)
	// fmt.Println(su.hello(1, 2))

	// done := make(chan struct{})

	// go func() {
	// 	fmt.Println("hello")
	// 	done <- struct{}{}
	// }()

	// <-done

	// 使用三个chan传来你起来

	// natural := make(chan int)
	// squerts := make(chan int)
	// // 产生自然数
	// go func() {
	// 	for x := 0; x < 6; x++ {
	// 		natural <- x
	// 	}
	// 	close(natural)
	// }()

	// // 接收自然数计算平方
	// go func() {
	// 	for x := range natural {
	// 		squerts <- x * x
	// 	}
	// 	close(squerts)
	// }()

	// for x := range squerts {
	// 	fmt.Println(x)
	// }
	// ticker := time.Tick(1 * time.Second)
	// time2 := time.After(3 * time.Second)
	// for {
	// 	select {
	// 	case <-ticker:
	// 		fmt.Println("hello")
	// 	case <-time2:
	// 		return
	// 	}
	// }

	// 协程
	// taskCH := make(chan int, 100)

	// for i := 0; i < 10; i++ {
	// 	taskCH <- i
	// }
	// go worker(taskCH)
	// select {
	// case <-time.After(time.Hour):
	// 	return
	// }

	//
	// fmt.Println(JieCheng(4))

	// t := &dataStruct.TreeNode{Data: "A"}
	// t.Left = &dataStruct.TreeNode{Data: "B"}
	// t.Right = &dataStruct.TreeNode{Data: "C"}
	// t.Left.Left = &dataStruct.TreeNode{Data: "D"}
	// t.Left.Right = &dataStruct.TreeNode{Data: "E"}
	// t.Right.Left = &dataStruct.TreeNode{Data: "F"}
	// fmt.Println("\n层次排序")
	// res := dataStruct.LayerOrder(t)
	// fmt.Println(res)
	list := []int{5, 9, 1, 6, 8}
	fmt.Println(reverse(list))

}
func reverse(a []int) []int {

	n := len(a)
	for i := 0; i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
	return a

}

// 递归
func JieCheng(n int) int {
	if n == 1 {
		return 1
	}
	return n * JieCheng(n-1)
}

// 尾递归
func JieCheng2(n int, a int) int {
	if n == 1 {
		return a
	}
	return JieCheng2(n-1, a*n)
}

func worker(taskCh <-chan int) {
	const N = 5
	for i := 0; i < N; i++ {
		go func(id int) {
			for {
				task := <-taskCh
				fmt.Printf("finish task: %d by worker %d\n", task, id)
				time.Sleep(time.Second)
			}
		}(i)
	}
}
