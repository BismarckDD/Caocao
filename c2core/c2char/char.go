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
	CharExtraMP         uint8
	CharExtraHP         uint8
	BattleFieldFigureId uint16
}

const (
	AbilityIncX uint8 = 5
	AbilityIncS uint8 = 4
	AbilityIncA uint8 = 3
	AbilityIncB uint8 = 2
	AbilityIncC uint8 = 1
	AbilityIncD uint8 = 0
	AbilityStdX uint8 = 100 // 按 100-90-70-50-30 达标
	AbilityStdS uint8 = 90
	AbilityStdA uint8 = 70
	AbilityStdB uint8 = 50
	AbilityStdC uint8 = 30
)

// 计算能力的实际增长数值
func CalculateRealInc(expectedInc, ability uint8) uint8 {
	switch expectedInc {
	case AbilityIncX:
		if ability >= AbilityStdX {
			return AbilityIncX
		} else {
			return AbilityIncA
		}
	case AbilityIncS:
		if ability >= AbilityStdS {
			return AbilityIncS
		} else {
			return AbilityIncA
		}
	case AbilityIncA:
		if ability >= AbilityStdX {
			return AbilityIncS
		} else if ability >= AbilityStdA {
			return AbilityIncA
		} else {
			return AbilityIncB
		}
	case AbilityIncB:
		if ability >= AbilityStdS {
			return AbilityIncA
		} else if ability >= AbilityStdB {
			return AbilityIncB
		} else {
			return AbilityIncC
		}
	case AbilityIncC:
		return AbilityIncC
	default:
		return AbilityIncC
	}
}

var CharList = []Char{
	{
		CharId:              68,          // 照片形象
		Name:                "步兵",        // 能力全部合格即可
		Power:               AbilityStdB, // BASBB -2
		Wisdom:              AbilityStdA,
		Leadership:          AbilityStdS,
		Agility:             AbilityStdB,
		Luck:                AbilityStdB,
		CharExtraHP:         0,
		CharExtraMP:         0,
		BattleFieldFigureId: 0, // 小兵都没有特殊的战场形象，依靠Unit的战场形象
	},
	{
		CharId:              68,          // 照片形象
		Name:                "骑兵",        // 能力全部合格即可
		Power:               AbilityStdS, // SBABB -2
		Wisdom:              AbilityStdB,
		Leadership:          AbilityStdA,
		Agility:             AbilityStdB,
		Luck:                AbilityStdB,
		CharExtraHP:         0,
		CharExtraMP:         0,
		BattleFieldFigureId: 0, // 小兵都没有特殊的战场形象，依靠Unit的战场形象
	},
	{
		CharId:              68,
		Name:                "弓兵",        // 能力全部合格即可
		Power:               AbilityStdA, // ABBBS -2
		Wisdom:              AbilityStdB,
		Leadership:          AbilityStdB,
		Agility:             AbilityStdB,
		Luck:                AbilityStdS,
		CharExtraHP:         0,
		CharExtraMP:         0,
		BattleFieldFigureId: 0,
	},
	{
		CharId:              68,
		Name:                "弓骑兵",       // 能力全部合格即可
		Power:               AbilityStdS, // SBBBA -2
		Wisdom:              AbilityStdB,
		Leadership:          AbilityStdB,
		Agility:             AbilityStdB,
		Luck:                AbilityStdA,
		CharExtraHP:         0,
		CharExtraMP:         0,
		BattleFieldFigureId: 0,
	},
	{
		CharId:              68,
		Name:                "武道家",       // 能力全部合格即可
		Power:               AbilityStdA, // ACASB -2
		Wisdom:              AbilityStdC,
		Leadership:          AbilityStdA,
		Agility:             AbilityStdS,
		Luck:                AbilityStdB,
		CharExtraHP:         0,
		CharExtraMP:         0,
		BattleFieldFigureId: 0,
	},
	{
		CharId:              68,
		Name:                "贼兵",        // 能力全部合格即可
		Power:               AbilityStdS, // SCBBS -2
		Wisdom:              AbilityStdC,
		Leadership:          AbilityStdB,
		Agility:             AbilityStdB,
		Luck:                AbilityStdS,
		CharExtraHP:         0,
		CharExtraMP:         0,
		BattleFieldFigureId: 0,
	},
	{
		CharId:              24,
		Name:                "舞娘",        // 能力全部合格即可
		Power:               AbilityStdA, // ABBSB -2
		Wisdom:              AbilityStdB,
		Leadership:          AbilityStdB,
		Agility:             AbilityStdS,
		Luck:                AbilityStdB,
		CharExtraHP:         0,
		CharExtraMP:         0,
		BattleFieldFigureId: 0,
	},
	{
		CharId:              68,
		Name:                "炮车",        // 能力全部合格即可
		Power:               AbilityStdS, // SBACA -2
		Wisdom:              AbilityStdB,
		Leadership:          AbilityStdA,
		Agility:             AbilityStdC,
		Luck:                AbilityStdA,
		CharExtraHP:         0,
		CharExtraMP:         0,
		BattleFieldFigureId: 0,
	},
	{
		CharId:              68,
		Name:                "道士",        // 能力全部合格即可
		Power:               AbilityStdC, // CSBAB -3
		Wisdom:              AbilityStdS,
		Leadership:          AbilityStdB,
		Agility:             AbilityStdA,
		Luck:                AbilityStdB,
		CharExtraHP:         0,
		CharExtraMP:         0,
		BattleFieldFigureId: 0,
	},
	{
		CharId:              68,
		Name:                "策士",        // 能力全部合格即可
		Power:               AbilityStdB, // BSBBB -3
		Wisdom:              AbilityStdS,
		Leadership:          AbilityStdB,
		Agility:             AbilityStdB,
		Luck:                AbilityStdB,
		CharExtraHP:         0,
		CharExtraMP:         0,
		BattleFieldFigureId: 0,
	},
	{
		CharId:              68,
		Name:                "风水士",       // 能力全部合格即可
		Power:               AbilityStdC, // CSCAA -3
		Wisdom:              AbilityStdS,
		Leadership:          AbilityStdC,
		Agility:             AbilityStdA,
		Luck:                AbilityStdA,
		CharExtraHP:         0,
		CharExtraMP:         0,
		BattleFieldFigureId: 0,
	},
	{
		CharId:              68,
		Name:                "骑马策士",      // 能力全部合格即可
		Power:               AbilityStdA, // ASBBC -3
		Wisdom:              AbilityStdS,
		Leadership:          AbilityStdB,
		Agility:             AbilityStdB,
		Luck:                AbilityStdC,
		CharExtraHP:         0,
		CharExtraMP:         0,
		BattleFieldFigureId: 0,
	},
	{
		CharId:              68,
		Name:                "西凉骑兵",      // 能力全部合格即可
		Power:               AbilityStdS, // SCSBB -2
		Wisdom:              AbilityStdC,
		Leadership:          AbilityStdS,
		Agility:             AbilityStdB,
		Luck:                AbilityStdB,
		CharExtraHP:         0,
		CharExtraMP:         0,
		BattleFieldFigureId: 0,
	},
	{
		CharId:              68,
		Name:                "黄巾贼",       // 能力全部合格即可
		Power:               AbilityStdB, // BCBBC -7
		Wisdom:              AbilityStdC,
		Leadership:          AbilityStdB,
		Agility:             AbilityStdB,
		Luck:                AbilityStdC,
		CharExtraHP:         0,
		CharExtraMP:         0,
		BattleFieldFigureId: 0,
	},
	{
		CharId:              68,
		Name:                "海盗",        // 能力全部合格即可
		Power:               AbilityStdS, // SBBAB -2
		Wisdom:              AbilityStdB,
		Leadership:          AbilityStdB,
		Agility:             AbilityStdA,
		Luck:                AbilityStdB,
		CharExtraHP:         0,
		CharExtraMP:         0,
		BattleFieldFigureId: 0,
	},
	{
		CharId:              68,
		Name:                "都督",        // 能力全部合格即可
		Power:               AbilityStdA, // ASBBB -2
		Wisdom:              AbilityStdS,
		Leadership:          AbilityStdB,
		Agility:             AbilityStdB,
		Luck:                AbilityStdB,
		CharExtraHP:         0,
		CharExtraMP:         0,
		BattleFieldFigureId: 0,
	},
	{
		CharId:              68,
		Name:                "咒术士",       // 能力全部合格即可
		Power:               AbilityStdC, // CSBBA -3
		Wisdom:              AbilityStdS,
		Leadership:          AbilityStdS,
		Agility:             AbilityStdB,
		Luck:                AbilityStdB,
		CharExtraHP:         0,
		CharExtraMP:         0,
		BattleFieldFigureId: 0,
	},
	{
		CharId:              68,
		Name:                "仙人",        // 能力全部合格即可
		Power:               AbilityStdC, // CSCAS -2
		Wisdom:              AbilityStdS,
		Leadership:          AbilityStdC,
		Agility:             AbilityStdA,
		Luck:                AbilityStdS,
		CharExtraHP:         0,
		CharExtraMP:         0,
		BattleFieldFigureId: 0,
	},
	{
		CharId:              68,
		Name:                "驯熊师",       // 能力全部合格即可
		Power:               AbilityStdS, // SCABB -3
		Wisdom:              AbilityStdC,
		Leadership:          AbilityStdA,
		Agility:             AbilityStdB,
		Luck:                AbilityStdB,
		CharExtraHP:         0,
		CharExtraMP:         0,
		BattleFieldFigureId: 0,
	},
	{
		CharId:              68,
		Name:                "驯虎师",       // 能力全部合格即可
		Power:               AbilityStdA, // ACASC -3
		Wisdom:              AbilityStdC,
		Leadership:          AbilityStdA,
		Agility:             AbilityStdS,
		Luck:                AbilityStdC,
		CharExtraHP:         0,
		CharExtraMP:         0,
		BattleFieldFigureId: 0,
	},
	{
		CharId:              68,
		Name:                "木人",        // 能力全部合格即可
		Power:               AbilityStdA, // ACASC -3
		Wisdom:              AbilityStdS,
		Leadership:          AbilityStdA,
		Agility:             AbilityStdS,
		Luck:                AbilityStdC,
		CharExtraHP:         0,
		CharExtraMP:         0,
		BattleFieldFigureId: 0,
	},
	{
		CharId:              68,
		Name:                "土偶",        // 能力全部合格即可
		Power:               AbilityStdS, // SCSBB -3
		Wisdom:              AbilityStdC,
		Leadership:          AbilityStdS,
		Agility:             AbilityStdB,
		Luck:                AbilityStdB,
		CharExtraHP:         0,
		CharExtraMP:         0,
		BattleFieldFigureId: 0,
	},
	{
		CharId:              68,
		Name:                "驯熊师",       // 能力全部合格即可
		Power:               AbilityStdS, // SCABB -2
		Wisdom:              AbilityStdC,
		Leadership:          AbilityStdA,
		Agility:             AbilityStdB,
		Luck:                AbilityStdB,
		CharExtraHP:         0,
		CharExtraMP:         0,
		BattleFieldFigureId: 0,
	},
	{
		CharId:              68,
		Name:                "驯虎师",       // 能力全部合格即可
		Power:               AbilityStdA, // ACASC -2
		Wisdom:              AbilityStdC,
		Leadership:          AbilityStdA,
		Agility:             AbilityStdS,
		Luck:                AbilityStdC,
		CharExtraHP:         0,
		CharExtraMP:         0,
		BattleFieldFigureId: 0,
	},
	{
		CharId:              1,
		Name:                "曹操",
		Power:               82,
		Wisdom:              92,
		Leadership:          98,
		Agility:             80,
		Luck:                84,
		CharExtraHP:         2,
		CharExtraMP:         0,
		BattleFieldFigureId: 100,
	},
	{
		CharId:              2,
		Name:                "夏侯惇",
		Power:               98,
		Wisdom:              64,
		Leadership:          82,
		Agility:             90,
		Luck:                66,
		CharExtraHP:         6,
		CharExtraMP:         0,
		BattleFieldFigureId: 101,
	},
	{
		CharId:              3,
		Name:                "张辽",
		Power:               92,
		Wisdom:              86,
		Leadership:          84,
		Agility:             78,
		Luck:                94,
		CharExtraHP:         2,
		CharExtraMP:         0,
		BattleFieldFigureId: 100,
	},
	{
		CharId:              1,
		Name:                "关羽",
		Power:               96,
		Wisdom:              90,
		Leadership:          98,
		Agility:             68,
		Luck:                62,
		CharExtraHP:         2,
		CharExtraMP:         0,
		BattleFieldFigureId: 100,
	},
	{
		CharId:              1,
		Name:                "曹操",
		Power:               82,
		Wisdom:              90,
		Leadership:          94,
		Agility:             AbilityStdA,
		Luck:                78,
		CharExtraHP:         2,
		CharExtraMP:         0,
		BattleFieldFigureId: 100,
	},
	{
		CharId:              1,
		Name:                "刘备",
		Power:               96,
		Wisdom:              90,
		Leadership:          98,
		Agility:             68,
		Luck:                62,
		CharExtraHP:         2,
		CharExtraMP:         0,
		BattleFieldFigureId: 100,
	},
	{
		CharId:              1,
		Name:                "关羽",
		Power:               96,
		Wisdom:              90,
		Leadership:          98,
		Agility:             68,
		Luck:                62,
		CharExtraHP:         2,
		CharExtraMP:         0,
		BattleFieldFigureId: 100,
	},
	{
		CharId:              1,
		Name:                "张飞",
		Power:               96,
		Wisdom:              90,
		Leadership:          98,
		Agility:             68,
		Luck:                62,
		CharExtraHP:         2,
		CharExtraMP:         0,
		BattleFieldFigureId: 100,
	},
}
