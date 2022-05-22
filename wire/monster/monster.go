package monster

type Monster struct {
	Name MonsterParam
}

type MonsterParam string

func NewMonster(name MonsterParam) Monster {
	return Monster{Name: name}
}
