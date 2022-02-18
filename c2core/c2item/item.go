package c2item

import (
	"github.com/BismarckDD/Caocao/c2core/c2status"
	"github.com/BismarckDD/Caocao/c2core/c2unit_category"
)

type AdaptPosEnum uint8

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
	HP          int16 // 恢复用豆、恢复用米、恢复用桃
	MP          int16 // 神秘水、神秘酒
	Power       uint8 // 武力果
	Leadership  uint8 // 统帅果
	Wisdom      uint8 // 智力果
	Agility     uint8 // 敏捷果
	Luck        uint8 // 好运果
	Exp         uint8 // 经验果
	Status      uint8 // 各种药: 解除中毒、定身、禁咒、混乱、全部异常状态
	Price       uint8 // 实际价格 Price * 100
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
	Id                             uint8  // Id & 0x80 == 0x80 = > treasure.
	Level                          uint8  // 1 - 9
	MaxLevel                       uint8  // 3 or 9
	Exp                            uint8  // 0 - 100
	Name                           string //
	AdaptPosEnum                          // 可以装备的位置
	c2unit_category.UnitCategoryId        // 可以装备哪些部队
	// 数值提升
	// 基础属性数值
	BaseAttack                 int16 // attack abs-value when lv1
	BaseSpirit                 int16 // spirit abs-value when lv1
	BaseDefence                int16 // defence abs-value when lv1
	BaseExplosive              int16 // agility abs-value when lv1
	BaseMorale                 int16 // luck abs-value when lv1
	BaseHPInc                  int16 // base hp inc when lv1
	BaseMPInc                  int16 // base mp inc when lv1
	BaseHPRecover              int16 // base hp inc when lv1
	BaseMPRecover              int16 // base mp inc when lv1
	BaseExpInc                 int16 // base mp inc when lv1
	BaseWeaponExpInc           int16 // base mp inc when lv1
	BaseArmorExpInc            int16 // base mp inc when lv1
	BaseAttackRatio            int8  // percentage increase -- power
	BaseDefenceRatio           int8  // percentage increase -- leadership
	BaseSpiritRatio            int8  // percentage increase -- wisdom
	BaseExplosiveRatio         int8  // percentage increase -- agility
	BaseMoraleRatio            int8  // percentage increase -- luck
	BaseHPIncRatio             int8  // base hp inc when lv1
	BaseMPIncRatio             int8  // base mp inc when lv1
	BaseHPRecoverRatio         int8  // base hp inc when lv1
	BaseMPRecoverRatio         int8  // base mp inc when lv1
	BaseExpIncRatio            int8  // base mp inc when lv1
	BaseWeaponExpIncRatio      int8  // base mp inc when lv1
	BaseArmorExpIncRatio       int8  // base mp inc when lv1
	BaseHitChance              uint8 // normal strike succeed prob. PiShouTao + 15
	BaseCriticalStrikeChance   uint8 // critical strike trigger prob. YuXi + 100
	BaseDoubleStrikeChance     uint8 // double strike trigger prob. no equip.
	BaseDodgeChance            uint8 // prob dodge to normal hit. YiTianJian +15
	BaseRemoteDodgeChance      uint8 // prob dodge to remote strike. JingKai +100
	BaseCriticalDodgeChance    uint8 // prob dodge to critical strike. HuangJinKai +100
	BaseMagicHitChance         uint8 // prob enhance to magic spell. QiXingJian +15
	BaseMagicDodgeChance       uint8 // prob dodge to magic spell. no equip.
	BaseMagicDamageReduceRatio uint8 // MDR (ratio). BaiYinKai +50

	// 升级时增加的属性数值
	UpAttack                 int16 // attack per up level
	UpDefence                int16 // defence per up level
	UpSpirit                 int16 // absolute value increase -- wisdom
	UpExplosive              int16 // absolute value increase -- agility
	UpMorale                 int16 // absolute value increase -- luck
	UpHPInc                  int16 // base hp inc when lv1
	UpMPInc                  int16 // base mp inc when lv1
	UpExpInc                 int16 // base mp inc when lv1
	UpWeaponExpInc           int16 // base mp inc when lv1
	UpArmorExpInc            int16 // base mp inc when lv1
	UpHPRecover              int16
	UpMPRecover              int16
	UpAttackRatio            int8 // attack ratio increase per
	UpDefenceRatio           int8 // percentage increase -- leadership
	UpSpiritRatio            int8 // percentage increase -- wisdom
	UpExplosiveRatio         int8 // percentage increase -- agility
	UpMoraleRatio            int8 // percentage increase -- luck
	UpHPIncRatio             int8 // base hp inc when lv1
	UpMPIncRatio             int8 // base mp inc when lv1
	UpHPRecoverRatio         int8
	UpMPRecoverRatio         int8
	UpExpIncRatio            int8  // base mp inc when lv1
	UpWeaponExpIncRatio      int8  // base mp inc when lv1
	UpArmorExpIncRatio       int8  // base mp inc when lv1
	UpHitChance              uint8 // normal strike succeed prob
	UpCriticalStrikeChance   uint8 // critical strike trigger prob
	UpDoubleStrikeChance     uint8 // double strike trigger prob
	UpDodgeChance            uint8 // dodge to normal hit, YiTianJian + 15
	UpRemoteDodgeChance      uint8 // prob dodge to remote strike. JingKai +100
	UpCriticalDodgeChance    uint8 // dodge to critical hit. HuangJinKai +100
	UpMagicHitChance         uint8 // prob enhance to magic spell. QiXingJian +15
	UpMagicDodgeChance       uint8 // dodge to magic spell
	UpMagicDamageReduceRatio uint8 // MDR (ratio)

	// 附加攻击状态
	c2status.UnitStatus   // 物品上有这个属性代表可以恢复上述状态
	c2status.StrikeAssist // 可以附加中毒、混乱、定身、禁咒, 封杀反击、反击后反击、穿刺攻击、引导攻击
	c2status.SpellAssist  //
	// 附加移动力
	Move uint8
	//
	c2unit_category.AttackRange
	//
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

func (item *Item) GetCriticalStrikeChance() uint8 {
	return item.BaseCriticalStrikeChance + item.UpCriticalStrikeChance*uint8(item.Level-1)
}

func (item *Item) GetDoubleStrikeChance() uint8 {
	return item.BaseDoubleStrikeChance + item.UpDoubleStrikeChance*uint8(item.Level-1)
}

func (item *Item) GetDodgeChance() uint8 {
	return item.BaseDodgeChance + item.UpDodgeChance*uint8(item.Level-1)
}

func (item *Item) GetRemoteDodgeChance() uint8 {
	return item.BaseRemoteDodgeChance + item.UpRemoteDodgeChance*uint8(item.Level-1)
}

func (item *Item) GetCriticalDodgeChance() uint8 {
	return item.BaseCriticalDodgeChance + item.UpCriticalDodgeChance*uint8(item.Level-1)
}

func (item *Item) GetMagicHitChance() uint8 {
	return item.BaseMagicHitChance + item.UpMagicHitChance*uint8(item.Level-1)
}

func (item *Item) GetMagicDodgeChance() uint8 {
	return item.BaseMagicDodgeChance + item.UpMagicDodgeChance*uint8(item.Level-1)
}

func (item *Item) GetMagicDamageReudceRatio() uint8 {
	return item.BaseMagicDamageReduceRatio + item.UpMagicDamageReduceRatio*uint8(item.Level-1)
}

func (item *Item) GetHPInc() int16 {
	return item.BaseHPInc + item.UpHPInc*int16(item.Level-1)
}

func (item *Item) GetMPInc() int16 {
	return item.BaseMPInc + item.UpMPInc*int16(item.Level-1)
}

func (item *Item) GetHPIncRatio() int8 {
	return item.BaseHPIncRatio + item.UpHPIncRatio*int8(item.Level-1)
}

func (item *Item) GetMPIncRatio() int8 {
	return item.BaseMPIncRatio + item.UpMPIncRatio*int8(item.Level-1)
}

func (item *Item) GetHPRecover() int16 {
	return item.BaseHPRecover + item.UpHPRecover*int16(item.Level-1)
}

func (item *Item) GetMPRecover() int16 {
	return item.BaseMPRecover + item.UpMPRecover*int16(item.Level-1)
}

func (item *Item) GetHPRecoverRatio() int8 {
	return item.BaseHPRecoverRatio + item.UpHPRecoverRatio*int8(item.Level-1)
}

func (item *Item) GetMPRecoverRatio() int8 {
	return item.BaseMPRecoverRatio + item.UpMPRecoverRatio*int8(item.Level-1)
}

// 发生在每回合开始时的经验变化
func (item *Item) GetExpInc() int16 {
	return item.BaseExpInc + item.UpExpInc*int16(item.Level-1)
}

func (item *Item) GetWeaponExpInc() int16 {
	return item.BaseWeaponExpInc + item.UpWeaponExpInc*int16(item.Level-1)
}

func (item *Item) GetArmorExpInc() int16 {
	return item.BaseArmorExpInc + item.UpArmorExpInc*int16(item.Level-1)
}

// 发生在攻防发生时的经验变化
func (item *Item) GetExpIncRatio() int16 {
	return item.BaseExpInc + item.UpExpInc*int16(item.Level-1)
}

func (item *Item) GetWeaponExpIncRatio() int16 {
	return item.BaseWeaponExpInc + item.UpWeaponExpInc*int16(item.Level-1)
}

func (item *Item) GetArmorExpIncRatio() int16 {
	return item.BaseArmorExpInc + item.UpArmorExpInc*int16(item.Level-1)
}
