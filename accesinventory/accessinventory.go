package main

import "fmt"

type Item struct {
	name string
}
type Character struct {
	name string
    Height int
    Weight int
    Haircolor string
    Level int
    Hp int
    Power int
    Speed int
    Money int
	inventory []Item 
}

func (c*Character) accesInventory() {
	fmt.Println("Inventaire de", c.name+":")
	for _,item := range c.inventory {
		fmt.Println(item.name)
	}
}

func main() {
	character := Character{
		name: "gasan",
		inventory: []Item{
		{name: "Racket"},
		{name: "balls"},
		{name: "water"},
		},
	}
	character.accesInventory()
}
