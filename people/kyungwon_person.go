package people

import "awesomeProject/people/dowon"

type Person struct {
	Name   string
	Weight int
}

func ReturnHeight() int {
	//1. people 패키지에서는 dowon 패키지의 어떠한 함수를 사용해야 하는 상황 (아래는 GetHeightNumber 함수)
	a := Person{Name: "kyungwon", Weight: 60}
	height := dowon.GetHeightNumber(a)
	return height
}

func ReturnName() string {
	b := Person{Name: "dowon", Weight: 70}
	height := dowon.GetName(b)
	return height
}
