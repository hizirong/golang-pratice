package main

import "fmt"

func main() {
	// 宣告
	users := make(map[string]string) // [key: string, value: string]

	// 加入資料
	users["7458"] = "Charlie"
	users["6422"] = "POKA"
	users["9738"] = "Wendy"
	users["7458"] = "York"

	if users["7458"] == "Charlie" {
		fmt.Println("KO NO Charlie DA!")
	} else if users["7458"] == "York" {
		fmt.Println("KO NO York DA!") //run this
	}

	// 檢查資料是否存在(利用key)
	val, ok := users["6233"]
	if !ok {
		fmt.Println("Value does not exit.") //run this
	} else {
		fmt.Println("Value is ", val)
	}

	// 假如你只是要檢查值在不在，不需要值的結果
	_, ok = users["6233"] // 注意不是":="，是"="

	// Map的大小，與陣列的用法一樣
	users_sz := len(users)
	fmt.Println("Count of users: ", users_sz)

	// 刪除資料
	delete(users, "6422")
	_, ok = users["6422"] //檢查是否存在
	if !ok {
		fmt.Println("POKA has been killed")
	}

}
