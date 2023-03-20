package main

import (
	"fmt"
)

/*
//===== Pass by value
func add(x int) {
	x = x + 10
	fmt.Println("Add function", x)
}
func main() {
	var a int = 10
	add(a) //Pass by value ，20
	fmt.Println("Main function",a) //10
}
*/
// ====== Pass by pointer
func add(xPtr *int) {
	*xPtr = *xPtr + 10 //反解指標=資料
	fmt.Println("Add function", *xPtr)

}
func main() {
	var b int = 10
	var bPtr *int = &b
	add(bPtr) // or add(&b) is ok
	fmt.Println("Main function", b)

	//和使用者要求輸入，運用到指標參數 pass by pointer
	var msg string
	fmt.Scanln(&msg) //傳遞字串變數的指標 (記憶體位置)
	fmt.Println(msg)
}
