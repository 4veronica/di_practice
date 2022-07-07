package dowon

import "awesomeProject/people"

func GetHeightNumber(person people.Person) int {
	//2.위의 함수는 people 패키지의 Person 이라는 구조체를 반드시 사용해야 하는 상황.
	return person.Weight
}

func GetName(person people.Person) string {
	return person.Name
}
