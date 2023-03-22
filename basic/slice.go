// code resourse: https://hackmd.io/@BKLiang/golang_ds
package main

import "fmt"

func main() {
	// ====== 宣告
	s1 := []int{1, 2, 3, 4, 5} // #1 看起來是陣列，但事實上是切片
	// #2、#3是利用make函式初始化
	s2 := make([]int, 5)     // #2 表示動態配置一切片，長度5，
	s3 := make([]int, 5, 10) // #3 差別在當前長度為5，並且預留10個int空間

	// #4到#8利用已有的陣列進行初始化
	arr := [5]int{1, 2, 3, 4, 5}

	s4 := arr[:]   // #4 表示取整個arr為初始
	s5 := arr[1:3] // #5 取陣列[start_index: end_index - 1]
	s6 := arr[1:]  // #6
	s7 := arr[:3]  // #7
	s8 := s4[:]    // #8

	s9 := []int{} // #9 初始化長度為0沒有值的切片(會預留4個空間)

	// 加入資料
	s1 = append(s1, 4, 5, 6, 7, 8)
	for i := 0; i < len(s1); i++ {
		fmt.Printf("%d ", s1[i])
	}

	// output: 1 2 3 4 5 4 5 6 7 8

	// 刪除，沒有delete可用，可用append()進行刪除
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	slice = append(slice[:3], slice[4:]...) // 刪除第三個元素

}
