/*
 * @author: onno
 * @date: 2020/4/28
 */
package main

type T interface{}

type Function func(input []T) T

type Filter func(input T) bool

type ToMap func(input []T) map[T]T

func Map(slice []T) {

}

