package c2char

import (
	"github.com/BismarckDD/Caocao/c2core/c2item"
	"github.com/BismarckDD/Caocao/c2core/c2status"
)

const (
	HUNDRED          int16 = 100
	AbilityNormalEff int16 = 100
	AbilityUpEff     int16 = 120
	AbilityDownEff   int16 = 70
)

type UnitCategoryId uint32

const (
	// JunZhu
	UnitCategoryQunXiong UnitCategoryId = 1

	// WuJiang
	UnitCategoryBuBing     UnitCategoryId = 1 << 1
	UnitCategoryQiBing     UnitCategoryId = 1 << 2
	UnitCategoryGongBing   UnitCategoryId = 1 << 3
	UnitCategoryGongQiBing UnitCategoryId = 1 << 4

	// 穿文官衣服的武将, 奋起、鼓舞、气合提升攻击力
	UnitCategoryWuDaoJia UnitCategoryId = 1 << 5
	UnitCategoryZeiBing  UnitCategoryId = 1 << 6
	UnitCategoryWuNiang  UnitCategoryId = 1 << 7
	UnitCategoryPaoChe   UnitCategoryId = 1 << 8

	// WenGuan
	UnitCategoryDaoShi      UnitCategoryId = 1 << 9  // XuanWuBaoYu
	UnitCategoryCeShi       UnitCategoryId = 1 << 10 // ZhuQueBaoYu
	UnitCategoryFengShuiShi UnitCategoryId = 1 << 11 // BaiHuBaoYu
	UnitCategoryQiMaCeShi   UnitCategoryId = 1 << 12 // QingLongBaoYu

	// Special
	UnitCategoryXiLiangQiBing UnitCategoryId = 1 << 13
	UnitCategoryHuangJinZei   UnitCategoryId = 1 << 14
	UnitCategoryHaiDao        UnitCategoryId = 1 << 15
	UnitCategoryDuDu          UnitCategoryId = 1 << 16 //
	UnitCategoryZhouShuShi    UnitCategoryId = 1 << 17 // All skills but Four-Shen, Four-High, Weather.
	UnitCategoryXianRen       UnitCategoryId = 1 << 18 // All skills
	UnitCategoryXunXiongShi   UnitCategoryId = 1 << 19 // SCABB = QiBing
	UnitCategoryXunHuShi      UnitCategoryId = 1 << 20 // ACASC = WuShuJia
	UnitCategoryMuRen         UnitCategoryId = 1 << 21 // ACASC = WuShuJia
	UnitCategoryTuOu          UnitCategoryId = 1 << 22 // SCSBB = XiLiangQiBing
	UnitCategoryBaiXing       UnitCategoryId = 1 << 23 //
	UnitCategoryHuangDi       UnitCategoryId = 1 << 24 //
	UnitCategoryYunShuDui     UnitCategoryId = 1 << 25 // CBCCB
	// ZiZhongDui    UnitCategoryId = 1 << 23
	// LiangMoDui    UnitCategoryId = 1 << 24
	// 还有6个兵种空余
)

type Unit struct {
	Char  // each unit is a char
	Level uint8
	HP    uint16
	MaxHP uint16
	MP    uint16
	MaxMP uint16
	// unit
	UnitCategory
	UnitCategoryId uint32 // 兵种类型
	// unit ability
	c2status.UnitStatus // toxic, chaos, hold, prohibition,
	Direction           uint8
	Move                uint8
	BaseAttack          int16
	BaseSpirit          int16
	BaseDefence         int16
	BaseExplosive       int16
	BaseMorale          int16
	Weapon              *c2item.Item
	Armor               *c2item.Item
	Assist              *c2item.Item
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

func (unit *Unit) GetBaseAttack() int16 {
	return 0
}

func (unit *Unit) GetBaseSpirit() int16 {
	return 0
}

func (unit *Unit) GetBase() int16 {
	return 0
}

func (unit *Unit) GetAttack() int16 {
	attack := unit.BaseAttack + unit.Weapon.GetAttack() + unit.Armor.GetAttack() + unit.Assist.GetAttack()
	ratio := int16(unit.Weapon.GetAttackRatio()) + int16(unit.Armor.GetAttackRatio()) + int16(unit.Assist.GetAttackRatio())
	attack += unit.Assist.GetAttack() + int16(ratio/HUNDRED)
	attack = int16(attack * unit.GetAttackStatusEff() / HUNDRED)
	return attack
}

func (unit *Unit) GetSpirit() int16 {
	attack := unit.BaseAttack + unit.Weapon.GetAttack() + unit.Armor.GetAttack() + unit.Assist.GetAttack()
	ratio := int16(unit.Weapon.GetAttackRatio()) + int16(unit.Armor.GetAttackRatio()) + int16(unit.Assist.GetAttackRatio())
	attack += unit.Assist.GetAttack() + int16(ratio/HUNDRED)
	attack = int16(attack * unit.GetAttackStatusEff() / HUNDRED)
	return attack
}

func (unit *Unit) GetDefence() int16 {
	attack := unit.BaseAttack + unit.Weapon.GetAttack() + unit.Armor.GetAttack() + unit.Assist.GetAttack()
	ratio := int16(unit.Weapon.GetAttackRatio()) + int16(unit.Armor.GetAttackRatio()) + int16(unit.Assist.GetAttackRatio())
	attack += unit.Assist.GetAttack() + int16(ratio/HUNDRED)
	attack = int16(attack * unit.GetAttackStatusEff() / HUNDRED)
	return attack
}

func (unit *Unit) GetAgility() int16 {
	attack := unit.BaseAttack + unit.Weapon.GetAttack() + unit.Armor.GetAttack() + unit.Assist.GetAttack()
	ratio := int16(unit.Weapon.GetAttackRatio()) + int16(unit.Armor.GetAttackRatio()) + int16(unit.Assist.GetAttackRatio())
	attack += unit.Assist.GetAttack() + int16(ratio/HUNDRED)
	attack = int16(attack * unit.GetAttackStatusEff() / HUNDRED)
	return attack
}

func (unit *Unit) GetMorale() int16 {
	attack := unit.BaseAttack + unit.Weapon.GetAttack() + unit.Armor.GetAttack() + unit.Assist.GetAttack()
	ratio := int16(unit.Weapon.GetAttackRatio()) + int16(unit.Armor.GetAttackRatio()) + int16(unit.Assist.GetAttackRatio())
	attack += unit.Assist.GetAttack() + int16(ratio/HUNDRED)
	attack = int16(attack * unit.GetAttackStatusEff() / HUNDRED)
	return attack
}

func (unit *Unit) GetMoveFromStatus() int16 {
	if unit.UnitStatus&c2status.UnitStatusMoveUp != 0 {
		return 2
	} else {
		return 0
	}
}
func (unit *Unit) GetMove() uint8 {

	return unit.Move + unit.Weapon.Move + unit.Armor.Move + unit.Assist.Move
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
