//go:build task01
// +build task01

// 指针
// 题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
// 考察点 ：指针的使用、值传递与引用传递的区别。
// 题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
// 考察点 ：指针运算、切片操作。
package main

import "fmt"

// addTen 接收一个整数指针，将其值加 10
func addTen(n *int) {
	if n == nil {
		return
	}
	*n += 10
}

// doubleSlice 接收一个整数切片的指针，将切片中每个元素乘以 2
func doubleSlice(nums *[]int) {
	if nums == nil {
		return
	}
	s := *nums
	for i := range s {
		s[i] *= 2
	}
}

func main() {
	// 示例 1：指针参数修改数值
	value := 5
	addTen(&value)
	fmt.Printf("加10后：%d\n", value)

	// 示例 2：切片指针修改切片内容
	arr := []int{1, 2, 3, 4}
	doubleSlice(&arr)
	fmt.Printf("切片乘2后：%v\n", arr)
}
