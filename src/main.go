package main

import "fmt"

func main() {
	func(tes string) {
		fmt.Println("测试",tes)
	}("1")
}
