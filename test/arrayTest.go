package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

var arr0 [5]int = [5]int{1,2,3}
var arr1 = [5]int{1, 2, 3, 4, 5}
var arr2 []int
var arr3 []string = []string{3:"helloworld",4:"linhua"}

func main()  {

	d:= []struct {
		name string
		age int
	}{
		{"user1", 10},
		{"linhua",25},
	}

	//对应切片来说，是左必右开，切片其实是对数组的封装可以随机取一个范围内的数据
	if arr2 == nil{
		fmt.Printf("切片为[%d]\n",0)
	}
	fmt.Println(arr0, arr1, arr2, arr3,d)

	s1 := make([]int,3)

	s2 := s1
	s2[0] = 100
	fmt.Println(s1) //[100 0 0]
	fmt.Println(s2) //[100 0 0]

	for index,value := range s1 {
		fmt.Println(index,value)
	}

	var num []int
	for i:=0;i<10;i++ {
		//append是追加
		num = append(num, i)
		fmt.Printf("%v  len:%d  cap:%d  ptr:%p\n", num, len(num), cap(num), num)
	}
	fmt.Println("eeeeeeeeeeee")
	serverIpi()
	forEachMap()
}

func serverIpi()  {
	var a = make([]string, 5,10)

	fmt.Println(len(a))
	for i := 0; i < 10; i++ {
		a = append(a, fmt.Sprintf("%v", i))
	}
	fmt.Println(len(a))
	fmt.Println(a)
}

func forEachMap()  {

	rand.Seed(time.Now().UnixNano()) //初始化随机数种子

	var scoreMap = make(map[string]int, 200)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串
		value := rand.Intn(100)          //生成0~99的随机整数
		scoreMap[key] = value
	}
	//取出map中的所有key存入切片keys
	var keys = make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}
	//对切片进行排序
	sort.Strings(keys)
	//按照排序后的key遍历map
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}
	var f = adder()
	fmt.Println(f(10)) //10
	fmt.Println(f(20))
	prt := 10
	var pr *int
	pr = &prt
	ptrTest(&prt)
	fmt.Println(prt,*pr)
}

func adder() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

func adder1() func(int) int  {
	var x int
	return func(i int) int {
		x=x+i
		return x
	}
}
func ptrTest(a *int){
	*a = 100
}
