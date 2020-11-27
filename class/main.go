package main

import (
	"fmt"
)

// Go에선 클래스 키워드 없이 struct를 사용
// 다만 struct 내부에 메소드를 선언할 수 없음
type Character struct {
	// Uppercase -> public
	Health int
	Name   string
	// lowercase -> private
	// 같은 패키지 내에서만 사용가능
	id int
}

type Magician struct {
	// Golang의 상속 -> Composition 이라고 함
	// 그냥 구조체 안에 부모 구조체를 넣는다
	// 이를 부모 구조체를 embedding 한다고 표현
	Character
	Class string
}

func New(
	health int,
	name string,
	id int,
) *Character {
	return &Character{Health: health, Name: name, id: id}
}

// "리시버"를 이용해 구조체 바깥에서 메소드 선언
// 아래 함수에서 (r *Character)에 해당
func (r *Character) ShowInfo() {
	fmt.Printf("Character %s\n", r.Name)
	fmt.Printf("Health %d\n", r.Health)
	fmt.Printf("id %d\n", r.id)
}

func (r *Character) ShowHealth() {
	fmt.Println("Health", r.Health)
}

func attacks(target *Character, i ...CharacterAction) {
	for _, a := range i {
		a.Attack(target)
		target.ShowHealth()
	}
}

func AttackTarget(c CharacterAction, target *Character) {
	c.Attack(target)
}

func main() {
	// 객체 생성 방법1 - {}안에 매개변수를 전달 (구조체 리터럴이라고 함)
	// 구조체이름 앞에 &를 붙이면 구조체 포인터로 초기화됨
	c1 := &Character{100, "c1", 1}
	// 객체 생성 방법2 - New 메소드를 외부에 선언(builder 패턴)
	// 이때 메소드 New의 반환타입이 포인터이므로 c2역시 구조체 포인터가 됨
	c2 := New(100, "c2", 2)
	// 객체 생성 방법3 - new 키워드를 사용 (좀 번거롭다)
	// new 키워드는 기본적으로 포인터를 반환
	c3 := new(Magician)
	c3.Health = 80
	c3.Name = "c3"
	c3.id = 3
	c3.Class = "Magician"
	// 1번 방법에 멤버변수를 순서대로 쓰기 싫으면 아래와 같이 하면 됨
	c4 := &Character{Health: 100, id: 4, Name: "c4"}

	c1.ShowInfo()
	c2.ShowInfo()
	c3.ShowInfo()
	c4.ShowInfo()

	fmt.Println("\ninitial health points")
	fmt.Println(c1.Health, c2.Health, c3.Health, c4.Health)
	// Character 타입 c1의 Attack메소드 호출시 Health -= 10
	c1.Attack(c2)
	fmt.Println(c1.Health, c2.Health, c3.Health, c4.Health)
	// Character 타입을 상속한 Magician 타입 c3의 Attack메소드 호출시 Health -= 30
	c3.Attack(c1)
	fmt.Println(c1.Health, c2.Health, c3.Health, c4.Health)

	// 인터페이스를 이용한 다형성 구현
	// 인터페이스 안에 선언된 메소드를 모두 구현한 객체만 할당될 수 있음
	var characterArr [4]CharacterAction
	characterArr[0] = c1
	characterArr[1] = c2
	characterArr[2] = c3
	characterArr[3] = c4

	for _, c := range characterArr {
		// Magician 타입 c3의 메소드 실행시 Character 타입 객체들의 실행결과와 다름
		AttackTarget(c, c4)
		c4.ShowHealth()
	}
}
