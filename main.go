package main

import (
	"fmt"
	"math"
	"test1/interface"
)


const MAX int = 3

func main()  {

	car := _interface.Dat{}

	var say _interface.Sayer
	say = car
	say.Say()


	a := []int{10,100,200}
	var i int
	var ptr [MAX]*int;
	for  i = 0; i < MAX; i++ {
		ptr[i] = &a[i] /* 整数地址赋值给指针数组 */
	}
	getNumber(ptr)
	sqrtDomo();
}
func getNumber(a [3]*int) {
	var i int
	for  i = 0;i<3;i++ {
		fmt.Printf("a[%d] = %d\n",i,*a[i]);
	}

}
func sqrtDomo(){

	var a,b = 3,4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Printf("%d\n",c)
}


