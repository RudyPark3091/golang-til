package main

import (
	"fmt"
)

// 인터페이스의 네이밍 컨벤션
// 상속받을 객체가 수행할 동작에 er을 붙임
type Attacker interface {
	// Attacker 를 구현한 객체를 대상으로 Attack() 수행
	Attack(Attacker)
	GetDamage(int) string
}

// 부모객체 Character
type Character struct {
	Health int
	Damage int
	Name   string
}

// Character를 상속받은 Magician
type Magician struct {
	mana int
	Character
}

// Attacker를 구현하는 Magician
// target 자리에 Attacker 인터페이스를 구현하는 객체 어느것이든 들어갈 수 있음
func (m *Magician) Attack(target Attacker) {
	fmt.Printf("%s used Magic to %s\n", m.Name, target.GetDamage(m.Damage))
	m.mana -= 10
}

func (m *Magician) GetDamage(damage int) string {
	m.Health -= damage
	return m.Name
}

// Character를 상속받은 Warrior
type Warrior struct {
	stamina int
	Character
}

// Attacker를 구현하는 Warrior
func (w *Warrior) Attack(target Attacker) {
	fmt.Printf("%s crashed %s\n", w.Name, target.GetDamage(w.Damage))
	w.stamina -= 30
}

func (w *Warrior) GetDamage(damage int) string {
	w.Health -= damage
	return w.Name
}

func main() {
	magician := &Magician{
		100,
		Character{100, 10, "magician"},
	}

	warrior := &Warrior{
		100,
		Character{100, 30, "warrior"},
	}

	// Warrior 객체 warrior가 Attack 동작을 수행하는 맥락
	warrior.Attack(magician)
	fmt.Printf("%s's Health: %d\n", magician.Name, magician.Health)
	fmt.Printf("%s's Health: %d\n", warrior.Name, warrior.Health)
}
