package main

import "fmt"

func main() {
	sl := make([]int, 1024)
	tmp := sl[1000:1025]
	fmt.Println(tmp)
}
