package main

import "fmt"

func main() {
	//===== continue
	// var x int
	// for x = 0; x < 5; x++ {
	// 	if x == 2 {
	// 		continue
	// 	}
	// 	fmt.Print(x)
	// }

	//無窮迴圈，有條件結束迴圈
	//計算總和值到輸入0為止
	var res int = 0
	for true {
		var n int
		fmt.Scanln(&n)
		if n == 0 {
			break
		}
		res += n
	}
	fmt.Println("sum ", res)
}
