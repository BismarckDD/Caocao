package c2item

import "github.com/BismarckDD/Caocao/c2core/c2status"

type AdaptPosEnum uint8
type AdaptUnitEnum uint32 // at most 32 unit

const (
	AdaptPosEnumWeapon AdaptPosEnum = iota
	AdaptPosEnumArmor
	AdaptPosEnumAssit
)

// Save file will save this
var ConsumeItemWareHouse = make([]uint8, 16)

// Save file not save this.
// 似乎ConsumeItem在战场上的使用范围都是9宫格
var ConsumeItemList = []ConsumeItem{
	ConsumeItem{
		Name:        "恢复用豆",
		Description: "恢复30点HP",
		HP:          30,
		Price:       1,
	},
	ConsumeItem{
		Name:        "恢复用米",
		Description: "恢复80点HP",
		HP:          80,
		Price:       3,
	},
	ConsumeItem{
		Name:        "恢复用桃",
		Description: "恢复200点HP",
		HP:          200,
		Price:       10,
	},
	ConsumeItem{
		Name:        "神秘水",
		Description: "恢复30点MP",
		MP:          30,
		Price:       5,
	},
	ConsumeItem{
		Name:        "神秘酒",
		Description: "恢复80点HP",
		HP:          80,
		Price:       10,
	},
	ConsumeItem{
		Name:        "解毒药",
		Description: "解除中毒状态",
		Status:      0x01,
		Price:       3,
	},
	ConsumeItem{
		Name:        "兴奋剂",
		Description: "解除混乱状态",
		Status:      0x01,
	},

	ConsumeItem{
		Name:        "膏药",
		Description: "解除定身状态",
		Status:      0x04,
	},
	ConsumeItem{
		Name:        "止咳药",
		Description: "解除禁咒状态",
		Status:      0x08,
		Price:       3,
	},
	ConsumeItem{
		Name:        "万能药",
		Description: "解除所有异常状态",
		Status:      0xff,
		Price:       10,
	},
	ConsumeItem{
		Name:        "印绶",
		Description: "兵种上升",
		HP:          5,
		MP:          2,
		Price:       10,
	},
	ConsumeItem{
		Name:        "武力果",
		Description: "武力上升",
		Power:       2,
		Price:       0,
	},
	ConsumeItem{
		Name:        "统帅果",
		Description: "统帅上升",
		Leadership:  2,
		Price:       0,
	},
	ConsumeItem{
		Name:        "智力果",
		Description: "智力上升",
		Wisdom:      2,
		Price:       0,
	},
	ConsumeItem{
		Name:        "敏捷果",
		Description: "敏捷上升",
		Agility:     2,
		Price:       0,
	},
	ConsumeItem{
		Name:        "运气果",
		Description: "运气上升",
		Luck:        2,
		Price:       0,
	},
}

type ConsumeItem struct {
	Name        string
	Description string
	MaxHP       uint16 // 印绶 +5
	HP          uint16 // 恢复用豆、恢复用米、恢复用桃
	MaxMP       uint16 // 印绶 +2
	MP          uint16 // 神秘水、神秘酒
	Power       uint8  // 武力果
	Leadership  uint8  // 统帅果
	Wisdom      uint8  // 智力果
	Agility     uint8  // 敏捷果
	Luck        uint8  // 好运果
	Exp         uint8  // 经验果
	Status      uint8  // 各种药: 解除中毒、定身、禁咒、混乱、全部异常状态
	Price       uint8  // 实际价格 Price * 100
}

func (item *ConsumeItem) GetPower() uint16 {
	return uint16(item.Power)
}

func (item *ConsumeItem) GetLeadership() uint16 {
	return uint16(item.Leadership)
}

func (item *ConsumeItem) GetWisdom() uint16 {
	return uint16(item.Wisdom)
}

func (item *ConsumeItem) GetAgility() uint16 {
	return uint16(item.Agility)
}

func (item *ConsumeItem) GetLuck() uint16 {
	return uint16(item.Luck)
}

func (item *ConsumeItem) GetExp() uint16 {
	return uint16(item.Exp)
}

func (item *ConsumeItem) GetPrice() uint16 {
	return uint16(item.Price) * 100
}

// 使用时直接和unit.Status做与操作即可
func (item *ConsumeItem) GetStatus() uint8 {
	return ^item.Status
}

// Save file will save this
var ItemList = [255]Item{}

type Item struct {
	// meta
	Id            uint8  // Id & 0x80 == 0x80 = > treasure.
	Level         uint8  // 1 - 9
	MaxLevel      uint8  // 3 or 9
	Exp           uint8  // 0 - 100
	Name          string //
	AdaptPosEnum         // 可以装备的位置
	AdaptUnitEnum        // 可以装备哪些部队
	// 数值提升
	// 基础属性数值
	BaseAttack                 int16 // attack increase when lv1
	BaseDefence                int16 // defence increase when lv1
	BaseSpirit                 int16 // absolute value increase -- wisdom
	BaseExplosive              int16 // absolute value increase -- agility
	BaseMorale                 int16 // absolute value increase -- luck
	BaseAttackRatio            int8  // percentage increase -- power
	BaseDefenceRatio           int8  // percentage increase -- leadership
	BaseSpiritRatio            int8  // percentage increase -- wisdom
	BaseExplosiveRatio         int8  // percentage increase -- agility
	BaseMoraleRatio            int8  // percentage increase -- luck
	BaseHitChance              uint8 // normal strike succeed prob
	BaseCrticalHitChance       uint8 // critical strike trigger prob
	BaseDoubleHitChance        uint8 // double strike trigger prob
	BaseDodgeChance            uint8 // prob dodge to normal hit, YiTianJian + 15
	BaseCriticalHitDodgeChance uint8 // prob dodge to critical hit
	BaseMagicDodgeChange       uint8 // prob dodge to magic spell
	BaseMagicDamageReduce      uint8 // MDR (ratio)
	// 升级时增加的属性数值
	UpAttack                 int16 // attack per up level
	UpDefence                int16 // defence per up level
	UpSpirit                 int16 // absolute value increase -- wisdom
	UpExplosive              int16 // absolute value increase -- agility
	UpMorale                 int16 // absolute value increase -- luck
	UpAttackRatio            int8  // attack ratio increase per
	UpDefenceRatio           int8  // percentage increase -- leadership
	UpSpiritRatio            int8  // percentage increase -- wisdom
	UpExplosiveRatio         int8  // percentage increase -- agility
	UpMoraleRatio            int8  // percentage increase -- luck
	UpHitChance              uint8 // normal strike succeed prob
	UpCrticalHitChance       uint8 // critical strike trigger prob
	UpDoubleHitChance        uint8 // double strike trigger prob
	UpDodgeChance            uint8 // dodge to normal hit, YiTianJian + 15
	UpCriticalHitDodgeChance uint8 // dodge to critical hit
	UpMagicDodgeChange       uint8 // dodge to magic spell
	UpMagicDamageReduce      uint8 // MDR (ratio)

	// 附加攻击状态
	c2status.UnitStatus   // 可以附加中毒、混乱、定身、禁咒状态
	c2status.StrikeStatus //

	// 附加移动力
	Move uint8
}

func (item *Item) GetAttack() int16 {
	return item.BaseAttack + item.UpAttack*int16(item.Level-1)
}

func (item *Item) GetDefence() int16 {
	return item.BaseDefence + item.UpDefence*int16(item.Level-1)
}

func (item *Item) GetSpirit() int16 {
	return item.BaseSpirit + item.UpAttack*int16(item.Level-1)
}

func (item *Item) GetExplosive() int16 {
	return item.BaseExplosive + item.UpExplosive*int16(item.Level-1)
}

func (item *Item) GetMorale() int16 {
	return item.BaseMorale + item.UpMorale*int16(item.Level-1)
}

func (item *Item) GetAttackRatio() int8 {
	return item.BaseAttackRatio + item.UpAttackRatio*int8(item.Level-1)
}

func (item *Item) GetDefenceRatio() int8 {
	return item.BaseDefenceRatio + item.UpDefenceRatio*int8(item.Level-1)
}

func (item *Item) GetSpiritRatio() int8 {
	return item.BaseSpiritRatio + item.UpSpiritRatio*int8(item.Level-1)
}

func (item *Item) GetExplosiveRatio() int8 {
	return item.BaseExplosiveRatio + item.UpExplosiveRatio*int8(item.Level-1)
}

func (item *Item) GetMoraleRatio() int8 {
	return item.BaseMoraleRatio + item.UpMoraleRatio*int8(item.Level-1)
}

func (item *Item) GetHitChance() uint8 {
	return item.BaseHitChance + item.UpHitChance*uint8(item.Level-1)
}

func (item *Item) GetCriticalChance() uint8 {
	return item.BaseHitChance + item.UpHitChance*uint8(item.Level-1)
}

// Save file will save this
var ItemWareHouse = []Item{
	{
		Id:         0,
		Level:      1,
		MaxLevel:   3,
		Exp:        0,
		Name:       "短剑",
		BaseAttack: 10,
		UpAttack:   10,
	},
	{
		Id:         1,
		Level:      1,
		MaxLevel:   3,
		Exp:        0,
		Name:       "大剑",
		BaseAttack: 40,
		UpAttack:   10,
	},
	{
		Id:         2,
		Level:      1,
		MaxLevel:   3,
		Exp:        0,
		Name:       "钢剑",
		BaseAttack: 80,
		UpAttack:   10,
	},
	{
		Id:         3,
		Level:      1,
		MaxLevel:   3,
		Exp:        0,
		Name:       "短枪",
		BaseAttack: 10,
		UpAttack:   10,
	},
	{
		Id:         4,
		Level:      1,
		MaxLevel:   3,
		Exp:        0,
		Name:       "长枪",
		BaseAttack: 40,
		UpAttack:   10,
	},
	{
		Id:         5,
		Level:      1,
		MaxLevel:   3,
		Exp:        0,
		Name:       "钢剑",
		BaseAttack: 80,
		UpAttack:   10,
	},
	{
		Id:         6,
		Level:      1,
		MaxLevel:   3,
		Exp:        0,
		Name:       "短弓",
		BaseAttack: 10,
		UpAttack:   10,
	},
	{
		Id:         7,
		Level:      1,
		MaxLevel:   3,
		Exp:        0,
		Name:       "大弓",
		BaseAttack: 40,
		UpAttack:   10,
	},
	{
		Id:         8,
		Level:      1,
		MaxLevel:   3,
		Exp:        0,
		Name:       "铁弓",
		BaseAttack: 80,
		UpAttack:   10,
	},
	{
		Id:         9,
		Level:      1,
		MaxLevel:   3,
		Exp:        0,
		Name:       "短棍",
		BaseAttack: 10,
		UpAttack:   10,
	},
	{
		Id:         10,
		Level:      1,
		MaxLevel:   3,
		Exp:        0,
		Name:       "长棍",
		BaseAttack: 40,
		UpAttack:   10,
	},
	{
		Id:         11,
		Level:      1,
		MaxLevel:   3,
		Exp:        0,
		Name:       "铁棍",
		BaseAttack: 80,
		UpAttack:   10,
	},

	{
		Id:         128,
		Level:      1,
		MaxLevel:   9,
		Exp:        0,
		Name:       "雌雄双剑",
		BaseAttack: 10,
		UpAttack:   10,
	},
	{
		Id:         129,
		Level:      1,
		MaxLevel:   9,
		Exp:        0,
		Name:       "青钢剑",
		BaseAttack: 10,
		UpAttack:   10,
	},
	{
		Id:         130,
		Level:      1,
		MaxLevel:   9,
		Exp:        0,
		Name:       "倚天剑",
		BaseAttack: 10,
		UpAttack:   10,
	},
	{
		Id:            131,
		Level:         1,
		MaxLevel:      9,
		Exp:           0,
		Name:          "古锭刀",
		BaseAttack:    10,
		UpAttack:      10,
		BaseExplosive: 10,
	},
	{
		Id:         132,
		Level:      1,
		MaxLevel:   9,
		Exp:        0,
		Name:       "青龙偃月刀",
		BaseAttack: 10,
		UpAttack:   10,
	},
	{
		Id:         133,
		Level:      1,
		MaxLevel:   3,
		Exp:        0,
		Name:       "钢剑",
		BaseAttack: 80,
		UpAttack:   10,
	},
	{
		Id:         6,
		Level:      1,
		MaxLevel:   3,
		Exp:        0,
		Name:       "短弓",
		BaseAttack: 10,
		UpAttack:   10,
	},
	{
		Id:         7,
		Level:      1,
		MaxLevel:   3,
		Exp:        0,
		Name:       "大弓",
		BaseAttack: 40,
		UpAttack:   10,
	},
	{
		Id:         8,
		Level:      1,
		MaxLevel:   3,
		Exp:        0,
		Name:       "铁弓",
		BaseAttack: 80,
		UpAttack:   10,
	},
	{
		Id:         9,
		Level:      1,
		MaxLevel:   3,
		Exp:        0,
		Name:       "短棍",
		BaseAttack: 10,
		UpAttack:   10,
	},
	{
		Id:         10,
		Level:      1,
		MaxLevel:   3,
		Exp:        0,
		Name:       "长棍",
		BaseAttack: 40,
		UpAttack:   10,
	},
	{
		Id:         11,
		Level:      1,
		MaxLevel:   3,
		Exp:        0,
		Name:       "铁棍",
		BaseAttack: 80,
		UpAttack:   10,
	},
}
