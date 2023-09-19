package main

import "fmt"

type Character struct {
	Name string
	Height int
	Weight int
	Haircolor string
	Level int
	Hp int
	Power int
	Speed int
	Money int
}

func Init(name string, height int, weight int, haircolor string, level int, hp int,power int, speed int, money int) Character {
	return Character{
		Name: name,
		Height: height,
		Weight: weight,
		Haircolor: haircolor,
		Level: level,
		Hp: hp,
		Power: power,
		Speed: speed,
		Money: money,
	}
}

func main() {
	character := Init("Gasan", 180, 70, "brun", 1, 50, 1, 10, 10)
	fmt.Println(character)
}
