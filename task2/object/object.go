//go:build shapes
// +build shapes

package main

import (
	"fmt"
	"math"
)

// Shape 接口：定义面积 Area 与周长 Perimeter 的计算行为。
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle 表示长方形，包含宽与高。
type Rectangle struct {
	width, height float64
}

// Circle 表示圆形，包含半径。
type Circle struct {
	radius float64
}

// Area 计算并返回长方形面积。
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

// Perimeter 计算并返回长方形周长。
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

// Area 计算并返回圆的面积。
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

// Perimeter 计算并返回圆的周长。
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID string
}

func (employee *Employee) PrintInfo() {
	fmt.Printf("员工姓名：%s,员工年龄：%d, 员工工号：%s ", employee.Person.Name, employee.Person.Age, employee.EmployeeID)

}

// main 演示：创建 Rectangle 与 Circle 并打印它们的面积与周长。
func main() {
	//定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。
	//在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
	r := Rectangle{width: 10, height: 5}
	r_area := r.Area()
	r_perimeter := r.Perimeter()
	fmt.Println("长方形的面积:", r_area)
	fmt.Println("长方形的周长:", r_perimeter)

	c := Circle{radius: 1.5}
	c_area := c.Area()
	c_perimeter := c.Perimeter()
	fmt.Println("圆的面积:", c_area)
	fmt.Println("圆的周长:", c_perimeter)

	//使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，
	//组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。

	emp := Employee{
		EmployeeID: "100",
		Person: Person{
			Name: "张三",
			Age:  20,
		},
	}
	emp.PrintInfo()
}
