package main

import "fmt"

func main() {
	// 示例1。
	//value1 := [...]int8{0, 1, 2, 3, 4, 5, 6}
	//而下面case表达式都是int8类型，不能转换成无类型常量
	//switch 1 + 3 { // 这条语句无法编译通过。
	//case value1[0], value1[1]:
	//	fmt.Println("0 or 1")
	//case value1[2], value1[3]:
	//	fmt.Println("2 or 3")
	//case value1[4], value1[5], value1[6]:
	//	fmt.Println("4 or 5 or 6")
	//}


	//有多个case满足情况下，如果没有fallthrogh则只会执行一个 不会向下传递

	// 示例2。
	value2 := [...]int8{0, 1, 2, 3, 4, 5, 6}
	//switch表达式是int8类型，下面case表达式是无类型常量，无类型常量是可以转成int8类型的
	var _ int8 = 0 //这条语句能通过编译是因为无类型常量0可以被转换成int8类型
	switch value2[4] {
	case 0, 1:
		fmt.Println("0 or 1")
	case 2, 3:
		fmt.Println("2 or 3")
	case 4, 5, 6:
		fmt.Println("4 or 5 or 6")
	}
}
