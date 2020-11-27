package main

import (
	"fmt"
)

// 인터페이스를 이용해 가상 메소드 구현
// OOP의 다형성에 해당
// Golang에서 인터페이스는 자신이 가지고 있는 메소드를 모두 구현하는 객체는 자신을 상속했음을 의미
type CharacterAction interface {
	Attack(*Character)
}

func (attacker Character) Attack(target *Character) {
	fmt.Printf("%s attacked %s\n", attacker.Name, target.Name)
	target.Health -= 10
}

func (attacker Magician) Attack(target *Character) {
	fmt.Printf("%s (%s) used fireball on %s\n", attacker.Name, attacker.Class, target.Name)
	target.Health -= 30
}
