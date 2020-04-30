/*
 * @author: onno
 * @date: 2020/4/30
 */
package collection

import (
	"fmt"
	"git.baijiahulian.com/go-learn/user"
)

func ToSlice(data ...int) []int {
	return data
}

func arrayTest1(array1 [3]int, array2 *[3]int) {
	println(&array1)
	println(array2)
	fmt.Println(*array2)
	println(&array2)
	array1[0] = 1
	array2[1] = -1
	array1 = [3]int{100, 100, 100}
	array2 = &[3]int{100, 100, 100}

}

func slice() {
	slice := []int{1, 2, 3}
	fmt.Printf("slice, address: %p\n", &slice)
	sliceTest2(slice, &slice)
	fmt.Println(slice)
}

func sliceTest2(slice []int, slicePtr *[]int) {
	println(slice, &slice)
	println(slicePtr, &slicePtr)
	slice[0] = -1
	(*slicePtr)[1] = -1
	slice = []int{100, 100, 100}
	slicePtr = &[]int{100, 100, 100}
}

func testMap() {
	userScoreMap := map[string]int{
		"Allen": 10,
		"Bob":   8,
	}
	fmt.Println(userScoreMap)
	for k, v := range userScoreMap {
		fmt.Printf("%s: %d\n", k, v)
	}
	for k := range userScoreMap {
		fmt.Printf("key: %s\n", k)
	}
	for _, v := range userScoreMap {
		fmt.Printf("value: %d\n", v)
	}
	value, ok := userScoreMap["Carle"]
	if ok {
		println(value)
	} else {
		println("Init Carle")
		userScoreMap["Carle"] = 8
		delete(userScoreMap, "Allen")
	}

	clazz2StudentsMap := make(map[int][]string)
	clazz2StudentsMap[0] = []string{"Allen", "Bob"}
	_, ok2 := clazz2StudentsMap[1]
	if !ok2 {
		clazz2StudentsMap[1] = []string{"Cindy"}
	}
	value3, ok3 := clazz2StudentsMap[1]
	if ok3 {
		clazz2StudentsMap[1] = append(value3, "Danny")
	}
	fmt.Println(clazz2StudentsMap)
}
