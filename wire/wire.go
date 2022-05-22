//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/yang201396/GoExamples/mission"
	"github.com/yang201396/GoExamples/monster"
	"github.com/yang201396/GoExamples/player"
)

func InitMission(MonsterName monster.MonsterParam, PlayerNamp player.PlayerParam) mission.Mission {
	wire.Build(monster.NewMonster, player.NewPlayer, mission.NewMission)
	return mission.Mission{}
}
