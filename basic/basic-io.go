package main

import "fmt"

func main() {
	// ===== basic output : fmt.Scanln(變數名稱,...)
	// fmt.Scanln(3,"Hello",true)

	// ===== basic input : fmt.Scamln(&變數名稱,...)
	// var x int
	// fmt.Print("enter a number: ")
	// fmt.Scanln(&x)
	// fmt.Println(x)

	// ===== 整合練習 :輸入兩個數字 並輸出相加結果
	/* 方法1
	var x int
	var y int
	fmt.Print("first number: ")
	fmt.Scanln(&x)
	fmt.Print("second number: ")
	fmt.Scanln(&y)
	//var result int=x+y
	//fmt.Print(result)
	fmt.Print(x + y)
	*/
	// 2
	var x int
	var y int
	fmt.Print("輸入兩個數字用空格隔開: ")
	fmt.Scanln(&x, &y)
	var result int = x + y
	fmt.Print(result)

}
