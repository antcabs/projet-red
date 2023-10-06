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
	Items []int
	Class	string
	Health 	int
	MaxHealth int
	Damage int
	IsDead	bool
	Resurrected bool
	Racket	string
	Balls 	string
	Water	string
	
}

type Personnage struct {
	nom string
	class string
	level int
	Hp int
}

type Teme struct {
	Name string
	Price int
}




type Equipement struct {
	Head string
	Torso string
	Feet string
}

type Inventory struct {
    Objects []string
	Racket string
	Balls string
	Water string
	Items []string
	Powerade string
	
}

type Ressources struct {
	CarbonRacket int
	TitaniumRacket int
}

type Weapon struct {
	name string
	power int
}

type Forgeron struct {
	availableWeapons map [string]int
}

type Player struct {
	inventory []string
}