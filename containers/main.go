package main

import (
	"container/list"
	"container/ring"
	"fmt"
)

func main() {
	// arrays
	var arr [5][5]int
	arr[0][1] = 1
	arr2 := [...]string{"1", "2", "3"}
	fmt.Println(len(arr))

	// slices
	var sl []int // nil можно добавлять
	sl2 := []int{1, 2, 3}
	sl3 := make([]int, 5, 10) // size, capacity

	sl2 = append(sl2, 1)
	subSl := sl2[0:5]
	copy(sl3, sl2)
	fmt.Println(subSl)

	fmt.Println(len(sl))

	// maps
	var m map[int]float64 // nul-map не писать не читать
	fmt.Println(len(m))

	m = map[int]float64{1: 1.0, 2: 2.0}
	for key, value := range m {
		fmt.Printf("%d: %f\n", key, value)
	}
	key := 2
	delete(m, key) // удалить по ключу

	m2 := make(map[rune]string, 2) // capacity
	fmt.Println(len(m2))

	// list
	l := list.New()
	l.PushBack(2)
	l.PushFront(1)
	el := l.PushBack(3)
	l.Remove(el)

	var x int = l.Front().Value.(int)
	x++
	for cur := l.Front(); cur != nil; cur = cur.Next() {
		fmt.Println(cur.Value)
	}

	// ring
	rng := ring.New(5)
	for i := 0; i < 5; i++ {
		rng.Value = i * 10000
		rng = rng.Next()
	}
	rng.Do(func(el interface{}) { fmt.Println(el.(int)) })

	// string
	// rune -> int32 unicode point
	var str string = "привет"
	for byteIdx, rune := range str {
		fmt.Println(byteIdx, string(rune))
	}
	r := []rune(str)
	fmt.Println(len(r))

	// loops
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	var i int = 10
	for i <= 20 {
		fmt.Println(i)
		i++
	}

	for {
		break
	}

	for idx, val := range arr2 {
		fmt.Printf("arr[%d] = %s\n", idx, val)
	}

}
