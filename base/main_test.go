/*
 * @author: onno
 * @date: 2020/4/28
 */
package base

import (
	"fmt"
    "testing"
)

func TestBranch(t *testing.T) {
    result := Branchs(1)
    if result != "one" {
        t.Error("not equals")
    }
}

func TestLoop(t *testing.T) {
    Loop()
    t.Log("")
}

func TestSum(t *testing.T) {
    result, _, _ := Sum(1, 2)
    t.Log("result: ", result)
    if result != 3 {
        t.Fail()
    }
}

func genericType() {
    kv := map[interface{}]interface{}{}
    kv["name"] = "hello"
    kv["id"] = 1
    kv[1] = 1234
    fmt.Println(kv)
    println("hello")
}

func error() {
    defer func() {
        if err := recover(); err != nil {
            fmt.Println("error: ", err)
        }
    }()
    a, b := 0, 10
    if (a == 0) {
        b = 0
    }
    print(a / b)
    panic("wrong")
}