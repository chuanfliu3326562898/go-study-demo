package main

import "fmt"


func funcA() {
	fmt.Println("func A")
}

func funcB() {

	//recover()必须搭配defer使用。
	//defer一定要在可能引发panic的语句之前定义。

	//painc 想到于抛出了一个异常
	//如果我们要让程序没有影响，就要在之前加上一个recover
	defer func() {
		//如果程序出出现了panic错误,可以通过recover恢复过来,并且在recover里面还能收集到刚刚出错的信息
		err := recover()
		if err != nil{
			fmt.Println("recover in B")
		}
	}()
	panic("panic in B")
}

func funcC() {
	fmt.Println("func C")
}

func main()  {

	funcA()
	funcB()
	funcC()

}
