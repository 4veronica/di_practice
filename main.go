package main

import (
	"awesomeProject/people"
	"fmt"
)

func main() {
	// 메인은 아래와 같음.
	weight := people.ReturnHeight()
	fmt.Println(weight)

	name := people.ReturnName()
	fmt.Println(name)
}
