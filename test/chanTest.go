package main

import "fmt"



type Sayer interface {
	say()
}


type dog struct {
	
}


type cat struct {

}
// cat实现了Sayer接口
func (c cat) say() {
	fmt.Println("喵喵喵")
}

func (d dog) say() {
	fmt.Println("汪汪汪")
}

var mymap = make(map[string]int)

func main()  {
	var x Sayer // 声明一个Sayer类型的变量x
	a := cat{}  // 实例化一个cat
	b := dog{}  // 实例化一个dog
	x = a       // 可以把cat实例直接赋值给x
	x.say()     // 喵喵喵
	x = b       // 可以把dog实例直接赋值给x
	x.say()     // 汪汪汪

	mymap["mytest1"] = 1
	mymap["mytest"] = 2

	fmt.Print(mymap)

	var a1 []string
	var b2 []int
	var c = []bool{false,true}
	fmt.Println(c,a1,b2)

}