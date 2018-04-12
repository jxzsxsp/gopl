package main

import (
	"fmt"
	"os"
)

/**
 * 初始化函数，会在主函数执行之前执行
 */
func init()  {
	fmt.Print("Hello ")
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: helloworld ${name}")
		os.Exit(1)
	}

	fmt.Println(os.Args[1])
}
