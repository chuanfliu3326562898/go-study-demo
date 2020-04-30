package _interface

import "fmt"


type Dog struct {

}

type Dat struct {

}

func (d Dog)Say()  {
	fmt.Print("ddddddd 汪汪汪")
}

func (c Dat)Say()  {
	fmt.Print("ccccccc 喵喵喵")
}

type daren interface {
	do()
}

func(d *Dog) do() {
	fmt.Println("wangwangwang")
}
// Sayer 接口,只要实现了所有interface里面的方法，就可以当作是一种interface的类型，比如这里只要实现say方法，就理解这个类型就是一个Sayer类型
type Sayer interface {
	Say()
}

func do(arg Sayer)  {
	arg.Say()
}



//func main()  {
//	var x Sayer // 声明一个Sayer类型的变量x
//	a := cat{}  // 实例化一个cat
//	b := dog{}  // 实例化一个dog
//	x = a       // 可以把cat实例直接赋值给x
//	x.say()     // 喵喵喵
//	x = b       // 可以把dog实例直接赋值给x
//	x.say()     // 汪汪汪
//
//	var aa chan int
//	aa = make(chan int,10)
//	fmt.Println(aa)
//}
