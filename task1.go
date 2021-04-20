package main

import (
	"fmt"
	"math/rand"
)

/*
	Есть слайс из 1000 элементов типа int,
	нужно его забить и отсортировать в обратном порядке
	без использования встроенных функций
*/

func main() {

	mySlice := createSlice(30)
	fmt.Println(sortSlice(mySlice))
}

func createSlice(randomNumber int) []int {
	var newSlice []int

	for i := 0; i < randomNumber; i++ {
		randomNumber := rand.Intn(1000)
		newSlice = append(newSlice, randomNumber)
	}
	return newSlice
}

func sortSlice(sl []int)  []int{
	for i := 0; i < len(sl); i++ {
		for j := i; j < len(sl); j++ {
			if sl[i] > sl[j] {
				tmp := sl[i]
				sl[i] = sl[j]
				sl[j] = tmp
			}
		}
	}
	return sl
}

