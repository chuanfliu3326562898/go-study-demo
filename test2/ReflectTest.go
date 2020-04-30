package main

import (
	"math/rand"
	"reflect"
	"strings"
	"sync"
	"time"
)
import "fmt"

var wg sync.WaitGroup

type student struct {
	Name string `json:"name"`
	Score int `json:"score"`
}

type person struct {
	Name string
	Age int
	Score float32
	Next *person
}

type Person1 struct {
	Name string
	Age int
}

func Newperson(name string,age int) *Person1 {
	return &Person1{
		Name: name,
		Age:  age,
	}
}



func transvers(head *person) {
	fmt.Printf("%d",&head)
	for head.Next != nil{
		fmt.Println(head)
		head = head.Next
	}
	(*head).Score = 80
	(*head).Age = 80
	(*head).Name = "罗志祥"

}
func insertHead(head *person)(*person){
	for i:=0;i<10;i++{
		////声明Student对象
		var stu person
		stu.Name = fmt.Sprintf("stu%d",i)
		stu.Age= rand.Intn(100)
		stu.Score= rand.Float32()*100
		stu.Next = head
		head = &stu
	}
	fmt.Println("insertHead函数内部遍历")
	transvers(head)
	//fmt.Println(headresult)
	fmt.Println(head)
	return head
}


func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type:%v\n", v)
}

func main()  {

	var head *person = new(person)
	(*head).Name = "hua"
	(*head).Age = 18
	(*head).Score = 100
	head = insertHead(head)
	fmt.Println("main函数内部遍历")
	transvers(head)


	var a = 12
	reflectType(a)
	var b float32 = 3.14
	reflectType(b)
	for i := 0; i < 10; i++ {
		wg.Add(1) // 启动一个goroutine就登记+1
		go hello(i)
	}
	fmt.Println("Hello Goroutine! start wait" )


	wg.Wait() // 等待所有登记的goroutine都结束
	fmt.Println("Hello Goroutine! end wait")

	//定义个变量作为func a的返回值，funca就是一个闭包
	funca := afunc()
	//相当于执行了afunc里面的匿名函数，我这里的funca就能有一个饮用指向刚刚的name
	funca()
}

func hello(i int) {
	fmt.Println("Hello Goroutine!", i)
	//if i == 2 {
	//	time.Sleep(time.Second*10)
	//}
	//匿名函数
	niming := func(){
		if i == 2 {
			fmt.Println("匿名函数")
			time.Sleep(time.Second*10)
		}
	}//()
	niming()
	defer wg.Done() // goroutine结束就登记-1

	fmt.Println("Hello Goroutine done!", i)
}
//定义一个函数，直接返回一个匿名函数，把函数作为一个返回值
func afunc() func() {
	//闭包 = 函数 + 外部变量
	name := "罗志祥"
	makefun := makeSuffixFunc(".txt")
	result := makefun("罗志祥")
	fmt.Println(result)
	return func(){
		//如果我要打印一个变量的name ,会向上找
		fmt.Println("hello 这是函数里面的匿名函数",name)
	}

}

func makeSuffixFunc(suffix string) func(name string) string  {
	return func(name string) string {
		if strings.HasSuffix(name,suffix) {
			return name
		}
		return name + suffix
	}
}




