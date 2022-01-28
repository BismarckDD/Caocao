package c2char

type Aligment int

const (
	AligmentWe = iota
	AligmentFriend
	AligmentEnemy
)

// S-figure: Scenario Figure.
// R-figure: Battle Field Figure.
// char 侧重描述形象
// charId: decide
//
type Char struct {
	CharId              int16
	Name                string
	Power               uint8
	Wisdom              uint8
	Leadership          uint8
	Agility             uint8
	Luck                uint8
	ExtraMP             uint8
	ExtraHP             uint8
	BattleFieldFigureId uint16
}

// X:100
// S: 90
// A: 

const (
	AbilityIncX uint8 = 5
	AbilityIncS uint8 = 4
	AbilityIncA uint8 = 3
	AbilityIncB uint8 = 2
	AbilityIncC uint8 = 1
	AbilityIncD uint8 = 0
	AbilityStdX uint8 = 100
	AbilityStdS uint8 = 90
	AbilityStdA uint8 = 75
	AbilityStdB uint8 = 60
	AbilityStdC uint8 = 30
)

func CalculateRealInc(expectedInc, ablility uint8) uint8 {
	switch expectedInc {
	case AbilityIncX:
		if ability >=  AbilityStdX {
			return AbilityIncX
		} else {
			return AbilityIncA
		}
	case AbilityIncS:
		if ability >=  AbilityStdS {
			return AbilityIncS
		} else {
			return AbilityIncA
		}
	case AbilityIncA:
		if ability >=  AbilityStdX {
			return AbilityIncS
		} else if ability >= AbilityStdA {
			return AblilityIncA
		} else {
			return AbilityIncB
		}
	case AbilityIncB:
		if ability >=  AbilityStdS {
			return AbilityIncA
		} else if ability >= AbilityStdB {
			return abilityIncB
		} else {
			return AbilityIncC
		}
	case AbilityIncB:
		return AbilityIncC
	}
}
var CharList = []Char{
	{
		CharId:              1,
		Name:                "步兵", // 能力全部合格即可
		Power:               70,
		Wisdom:              90,
		Leadership:          94,
		Agility:             70,
		Luck:                78,
		ExtraHP:             2,
		ExtraMP:             0,
		BattleFieldFigureId: 100,
	}
	{
		CharId:              1,
		Name:                "曹操",
		Power:               82,
		Wisdom:              90,
		Leadership:          94,
		Agility:             70,
		Luck:                78,
		ExtraHP:             2,
		ExtraMP:             0,
		BattleFieldFigureId: 100,
	},
	{
		CharId:              2,
		Name:                "夏侯惇",
		Power:               98,
		Wisdom:              52,
		Leadership:          94,
		Agility:             90,
		Luck:                78,
		ExtraHP:             6,
		ExtraMP:             0,
		BattleFieldFigureId: 101,
	},
	{
		CharId:              1,
		Name:                "曹操",
		Power:               82,
		Wisdom:              90,
		Leadership:          94,
		Agility:             70,
		Luck:                78,
		ExtraHP:             2,
		ExtraMP:             0,
		BattleFieldFigureId: 100,
	},
	{
		CharId:              1,
		Name:                "曹操",
		Power:               82,
		Wisdom:              90,
		Leadership:          94,
		Agility:             70,
		Luck:                78,
		ExtraHP:             2,
		ExtraMP:             0,
		BattleFieldFigureId: 100,
	},
	{
		CharId:              1,
		Name:                "曹操",
		Power:               82,
		Wisdom:              90,
		Leadership:          94,
		Agility:             70,
		Luck:                78,
		ExtraHP:             2,
		ExtraMP:             0,
		BattleFieldFigureId: 100,
	},
	{
		CharId:              1,
		Name:                "曹操",
		Power:               82,
		Wisdom:              90,
		Leadership:          94,
		Agility:             70,
		Luck:                78,
		ExtraHP:             2,
		ExtraMP:             0,
		BattleFieldFigureId: 100,
	},
}
