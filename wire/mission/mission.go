package mission

import (
	"fmt"
	"github.com/yang201396/GoExamples/monster"
	"github.com/yang201396/GoExamples/player"
)

type Mission struct {
	Player  player.Player
	Monster monster.Monster
}

func NewMission(p player.Player, m monster.Monster) Mission {
	return Mission{p, m}
}

func (m Mission) Start() {
	fmt.Printf("%s defeats %s, world peace!\n", m.Player.Name, m.Monster.Name)
}
