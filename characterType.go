package main

type Character struct {
	Name      string
	Height    int
	Weight    int
	Haircolor string
	Level     int
	Hp        int
	Power     int
	Speed     int
	Money     int
	Skills []string
	Items []Item
	Class	string
	Health 	int
	MaxHealth int
	Damage int
	IsDead	bool
	Resurrected bool
}

type Item struct {
	Name string
	Price int
}

type Merchant struct {
	Items []Item
}

type Equipement struct {
	Head string
	Torso string
	Feet string
}