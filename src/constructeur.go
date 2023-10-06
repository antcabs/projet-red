package main

func (c *Character) Init(name string, height int, weight int, haircolor string, level int, hp int, power int, speed int, money int, skills [] string, class string, health int, maxHealth int, damage int, isDead bool, resurrected bool, racket string, water string, balls string) {
	c.Name = name
	c.Height = height
	c.Weight = weight
	c.Haircolor = haircolor
	c.Level = level
	c.Hp = hp
	c.Power = power
	c.Speed = speed
	c.Money = money
	c.Skills = skills
	c.Class = class
	c.Health = health
	c.MaxHealth = maxHealth
	c.Damage = damage
	c.IsDead = isDead 
	c.Resurrected = resurrected
	c.Racket = racket 
	c.Balls = balls 
	c.Water = water

}

func (c *Teme) Init(name string, price int) {
	c.Name = name
	c.Price = price
}

func (c *Equipement) Init(head string, torso string, feet string) {
	c.Head = head
	c.Torso = torso
	c.Feet = feet
}

func (c *Inventory) Init(objects []string,racket string, water string, balls string, items [] string, powerade string) {
	c.Objects = objects
	c.Racket = racket
	c.Water = water
	c.Balls = balls
	c.Items = items
	c.Powerade = powerade
}

func (c *Personnage) Init(nom string, class string, level int, Hp int) {
	c.nom = nom
	c.class = class
	c.level = level
	c.Hp = Hp
}

func (c *Ressources) Init(CarbonRacket int, TitaniumRacket int ) {
	c.CarbonRacket = CarbonRacket
	c.TitaniumRacket = TitaniumRacket
}

