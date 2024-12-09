package model

type Dungeon struct {
	startRoom DungeonNode
}

type DungeonNode interface {
	Generate() DungeonNode
	Enter(Character)
	Leave(Character)
}

type DungeonNodeBase struct {
	// enter and leave need to be implemented
	grid [][]*int
}

type Room struct {
	DungeonNodeBase
	corridors []*Connector
}

type Connector struct {
	DungeonNodeBase
	connectedRooms []*Room // 2 rooms connected by this connector
}

type Grid struct {
	grid [][]*cell
}

type CellType int

const (
	Wall CellType = iota
	Floor
	Door
)

type cell struct {
	cellType  CellType
	character *Character
}

func (node DungeonNodeBase) Enter(c Character) {
	// stub
}

func (node DungeonNodeBase) Leave(c Character) {
	// stub
}

// custom grid representation?
