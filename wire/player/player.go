package player

type Player struct {
	Name PlayerParam
}

type PlayerParam string

func NewPlayer(name PlayerParam) Player {
	return Player{Name: name}
}
