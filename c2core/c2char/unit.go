package c2char

import (
	"github.com/BismarckDD/Caocao/c2core/c2item"
	"github.com/BismarckDD/Caocao/c2core/c2status"
	"github.com/BismarckDD/Caocao/c2core/c2terrain"
	"github.com/BismarckDD/Caocao/c2core/c2unit_category"
)

const (
	HUNDRED          int16 = 100
	AbilityNormalEff int16 = 100
	AbilityUpEff     int16 = 120
	AbilityDownEff   int16 = 70
)

type Unit struct {
	Char  // each unit is a char
	Level uint8
	HP    int16
	MaxHP int16
	MP    int16
	MaxMP int16
	// unit
	c2unit_category.UnitCategory // Unit基础数据
	c2status.UnitStatus          // toxic, chaos, hold, prohibition,
	Direction                    uint8
	BaseAttack                   int16
	BaseSpirit                   int16
	BaseDefence                  int16
	BaseExplosive                int16
	BaseMorale                   int16
	Weapon                       *c2item.Item
	Armor                        *c2item.Item
	Assist                       *c2item.Item
	// 地图相关属性
	c2terrain.Terrain // 该 unit 所在的地形
}

// param
// CharId: e.g. Caocao
// CategoryId: e.g. 3
// Level: e.g. 3
func CreateUnit() *Unit {
	return nil
}

// HP(MP) = 部队基本HP(MP) + 武将加成 + HP(MP)加成[等级 + 2(使用用印授的次数)]
func (unit *Unit) LevelUp() bool {
	unit.Level = unit.Level + 1
	unit.MaxHP = int16(unit.UnitBaseHP) + int16(unit.CharExtraHP) + int16(unit.UnitExtraHP)*int16(unit.Level+2*unit.CategoryLevel)
	unit.MaxMP = int16(unit.UnitBaseMP) + int16(unit.CharExtraMP) + int16(unit.UnitExtraMP)*int16(unit.Level+2*unit.CategoryLevel)
	return true
}

func (unit *Unit) Transfer() bool {
	unit.UnitCategory = c2unit_category.UnitCategoryList[unit.UnitCategory.TransferId]
	unit.MaxHP = int16(unit.UnitBaseHP) + int16(unit.CharExtraHP) + int16(unit.UnitExtraHP)*int16(unit.Level+2*unit.CategoryLevel)
	unit.MaxMP = int16(unit.UnitBaseMP) + int16(unit.CharExtraMP) + int16(unit.UnitExtraMP)*int16(unit.Level+2*unit.CategoryLevel)
	unit.HP = unit.MaxHP      // 恢复HP
	unit.MP = unit.MaxMP      // 恢复MP
	unit.UnitStatus &= 0xfff0 // 恢复异常状态
	return true
}

func (unit *Unit) GetAttackStatusEff() int16 {
	if unit.UnitStatus&c2status.UnitStatusAttackUp != 0 {
		return AbilityUpEff
	} else if unit.UnitStatus&c2status.UnitStatusAttackDown != 0 {
		return AbilityDownEff
	} else {
		return AbilityNormalEff
	}
}

func (unit *Unit) GetSpiritStatusEff() int16 {
	if unit.UnitStatus&c2status.UnitStatusSpiritUp != 0 {
		return AbilityUpEff
	} else if unit.UnitStatus&c2status.UnitStatusSpiritDown != 0 {
		return AbilityDownEff
	} else {
		return AbilityNormalEff
	}
}

func (unit *Unit) GetDefenceStatusEff() int16 {
	if unit.UnitStatus&c2status.UnitStatusDefenceUp != 0 {
		return AbilityUpEff
	} else if unit.UnitStatus&c2status.UnitStatusDefenceDown != 0 {
		return AbilityDownEff
	} else {
		return AbilityNormalEff
	}
}

func (unit *Unit) GetExplosiveStatusEff() int16 {
	if unit.UnitStatus&c2status.UnitStatusExplosiveUp != 0 {
		return AbilityUpEff
	} else if unit.UnitStatus&c2status.UnitStatusExplosiveDown != 0 {
		return AbilityDownEff
	} else {
		return AbilityNormalEff
	}
}

func (unit *Unit) GetMoraletatusEff() int16 {
	if unit.UnitStatus&c2status.UnitStatusMoraleUp != 0 {
		return AbilityUpEff
	} else if unit.UnitStatus&c2status.UnitStatusMoraleDown != 0 {
		return AbilityDownEff
	} else {
		return AbilityNormalEff
	}
}

func TerrainBonus(val int16, unit *Unit) int16 {

	terrainBounsData, ok := c2unit_category.TerrainAbilityBonusLookUp[c2unit_category.UnitCategoryId(unit.UnitCategoryId)]
	if !ok {
		return val
	}
	return val * int16(terrainBounsData[unit.Terrain]) / HUNDRED
}

func (unit *Unit) GetAttack() int16 {
	// 物品绝对值加成
	attack := unit.BaseAttack + unit.Weapon.GetAttack() + unit.Armor.GetAttack() + unit.Assist.GetAttack()
	ratio := int16(unit.Weapon.GetAttackRatio()) + int16(unit.Armor.GetAttackRatio()) + int16(unit.Assist.GetAttackRatio())
	// 物品百分比加成
	attack += unit.Assist.GetAttack() + int16(ratio/HUNDRED)
	// 状态加成
	attack = int16(attack * unit.GetAttackStatusEff() / HUNDRED)
	// 地形加成
	return TerrainBonus(attack, unit)
}

func (unit *Unit) GetSpirit() int16 {
	spirit := unit.BaseSpirit + unit.Weapon.GetSpirit() + unit.Armor.GetSpirit() + unit.Assist.GetSpirit()
	ratio := int16(unit.Weapon.GetSpiritRatio()) + int16(unit.Armor.GetSpiritRatio()) + int16(unit.Assist.GetSpiritRatio())
	spirit += int16(unit.BaseSpirit * ratio / HUNDRED)
	spirit = int16(spirit * unit.GetSpiritStatusEff() / HUNDRED)
	return spirit
}

func (unit *Unit) GetDefence() int16 {
	defence := unit.BaseDefence + unit.Weapon.GetDefence() + unit.Armor.GetDefence() + unit.Assist.GetDefence()
	ratio := int16(unit.Weapon.GetDefenceRatio()) + int16(unit.Armor.GetDefenceRatio()) + int16(unit.Assist.GetDefenceRatio())
	defence += int16(unit.BaseDefence * ratio / HUNDRED)
	defence = int16(defence * unit.GetDefenceStatusEff() / HUNDRED)
	return defence
}

func (unit *Unit) GetExplosive() int16 {
	explosive := unit.BaseExplosive + unit.Weapon.GetExplosive() + unit.Armor.GetExplosive() + unit.Assist.GetExplosive()
	ratio := int16(unit.Weapon.GetExplosiveRatio()) + int16(unit.Armor.GetExplosiveRatio()) + int16(unit.Assist.GetExplosiveRatio())
	explosive += int16(unit.BaseExplosive * ratio / HUNDRED)
	explosive = int16(explosive * unit.GetExplosiveStatusEff() / HUNDRED)
	return explosive
}

func (unit *Unit) GetMorale() int16 {
	morale := unit.BaseMorale + unit.Weapon.GetMorale() + unit.Armor.GetMorale() + unit.Assist.GetMorale()
	ratio := int16(unit.Weapon.GetMoraleRatio()) + int16(unit.Armor.GetMoraleRatio()) + int16(unit.Assist.GetMoraleRatio())
	morale += int16(unit.BaseMorale * ratio / HUNDRED)
	morale = int16(morale * unit.GetAttackStatusEff() / HUNDRED)
	return morale
}

func (unit *Unit) GetMoveFromStatus() uint8 {
	if unit.UnitStatus&c2status.UnitStatusMoveUp != 0 {
		return 2
	} else {
		return 0
	}
}
func (unit *Unit) GetMove() uint8 {

	return unit.Move + unit.Weapon.Move + unit.Armor.Move + unit.Assist.Move +
		unit.GetMoveFromStatus()
}

// 同类的ConsumeItem是完全一致
// 比如99个豆，每个都一样
func (unit *Unit) Use(item *c2item.ConsumeItem) bool {
	unit.HP += item.HP
	unit.MP += item.MP
	return true
}

// 每一个Item都是不同的
// 比如两把钢枪，一把lv1，一把lv3
func (unit *Unit) Equip(item *c2item.Item) bool {
	if item.AdaptPosEnum == c2item.AdaptPosEnumWeapon {

	} else if item.AdaptPosEnum == c2item.AdaptPosEnumArmor {

	} else if item.AdaptPosEnum == c2item.AdaptPosEnumAssit {

	}
	return true
}
