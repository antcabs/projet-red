package main

func (c *Character) Init(name string, height int, weight int, haircolor string, level int, hp int, power int, speed int, money int, class string, health int, maxHealth int, damage int, isDead bool, resurrected bool) {
	c.Name = name
	c.Height = height
	c.Weight = weight
	c.Haircolor = haircolor
	c.Level = level
	c.Hp = hp
	c.Power = power
	c.Speed = speed
	c.Money = money
	c.Class = class
	c.Health = health
	c.MaxHealth = maxHealth
	c.Damage = damage
	c.IsDead = isDead 
	c.Resurrected = resurrected

}

func (c *Item) Init(name string, price int) {
	c.Name = name
	c.Price = price
}

func (c *Equipement) Init(head string, torso string, feet string) {
	c.Head = head
	c.Torso = torso
	c.Feet = feet
}


