package base

import (
	"fmt"
)

func branch(x int) string {
	if x > 0 {
		return "positive"
	} else if x < 0 {
		return "negative"
	} else {
		return "zero"
	}
}

/**
param int
return string
*/
func Branchs(x int) string {
	var result string
	switch x {
	case 0:
		println("branch1")
		result = "zero"
	case 1:
		println("branch2")
		result = "one"
	case 2:
		println("branch3")
		result = "two"
	case 3, 4:
		result = "three or four"
	default:
		result = "unknow"
	}
	return result
}

func branch2(x int) string {
	switch {
	case x == 1:
		return "one"
	case x == 2:
		return "two"
	default:
		return "unknow"
	}
}

func Loop() {
	for i := 0; i < 10; i++ {
		println(i)
	}
	i := 0
	for {
		i++
		println(i)
		if i > 10 {
			break
		}
	}
	i = 0
	for i < 10 {
		i++
		println(i)
	}
}

func Sum(a, b int) (int, int, int) {
	return a + b, a, b
}

type Predicate func(num int) bool

func filter(slice []int, predicate Predicate) []int {
	result := []int{}
	for _, v := range slice {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

func testFilter() {
	predicate := func(n int) bool {
		return n > 0
	}
	slice := []int{-1, 0, 1, 2, 3}
	result := filter(slice, predicate)
	fmt.Println(result)
}

func f() (r int) {
	defer func(r int) int {
		r += 5
		return 10
	}(r)
	return 1
}

func add(base int) func(int) int {
	fmt.Println("base: ", base)
	for i := 0; i < base; i++ {
		println("for")
		defer func() {
			fmt.Println("in defer", i)
		}()
	}
	return func(i int) int {
		base += i
		return base
	}
}

type Test struct {
	name string
}

func (t *Test) Close() {
	fmt.Printf("t address: %p, t point address: %p\n", &t, t)
	fmt.Println(t.name, " closed")
}

func (t Test) Close1() {
	fmt.Println(t.name, " closed")
}

func Close(t Test) {
	println(&t)
	t.Close()
}

func Close1(t *Test) {
	println(t)
	t.Close()
}

func main() {
	ts := []Test{{"a"}, {"b"}, {"c"}}

	for i := 0; i < len(ts); i++ {
		defer ts[i].Close1()
	}

	// for _, t := range ts {
	//     t2 := t
	//     fmt.Printf("t: %p, t2: %p\n", &t, &t2)
	//     // defer t.Close1()
	//     defer t.Close()
	//     // defer t2.Close()
	//     // defer Close(t)
	//     // defer Close1(&t)
	// }
}
