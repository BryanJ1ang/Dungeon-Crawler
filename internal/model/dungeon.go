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
}

type Room struct {
	DungeonNodeBase
	corridors []*Connector
	grid      [][]*int // 0 = wall, 1 = room

}

type Connector struct {
	DungeonNodeBase
	connectedRooms []*Room  // 2 rooms connected by this connector
	grid           [][]*int // 0 = wall, 1 = room
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
	character *Character
	cellType  CellType
}

// custom grid representation?
