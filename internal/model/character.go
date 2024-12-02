package model

type Character struct {
	name string

	stats     Stats
	inventory Inventory
}

type Stats struct {
	attack, mana, defense, health, speed, luck int
}

type Inventory struct {
	weapon   Weapon
	spell    Spell
	armor    Armor
	backpack []Equipment
}

type Equipment interface {
	Equip(Character) bool
	Unequip() bool
	Store()
	Discard()
}

type equipmentBase struct{}

type Weapon struct {
	equipmentBase
	damage     int
	durability int
}

type Armor struct {
	equipmentBase
	defense    int
	durability int
}

type Spell struct {
	damage   int
	manaCost int
}

func (e equipmentBase) Equip(c Character) bool {
	return true // stub
}

func (e equipmentBase) Unequip() bool {
	return true // stub
}

func (e equipmentBase) Store() {
	// stub
}

func (e equipmentBase) Discard() {
	// stub
}
