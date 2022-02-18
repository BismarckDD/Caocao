package c2attack

import (
	"math/rand"
	"time"

	"github.com/BismarckDD/Caocao/c2core/c2char"
)
const (
	RandSource := rand.NewSource(time.Now())
	RandIntGen := rand.New(RandSource)
)


type AttackCategory uint8

const (
	AttackCategoryNormalStrike   AttackCategory = 0x1
	AttackCategoryCriticalStrike AttackCategory = 0x2
	AttackCategoryDoubleStrike   AttackCategory = 0x4
	AttackCategoryCounterStrike  AttackCategory = 0x4
)

/*
普通攻击命中率:(HitRatio)
IF X.Explosive > Y.Explosive * 2, THEN HR = 100;
ELSE IF X.Explosive > YF, THEN HR = (X.Explosive - Y.Explosive) * 10 / Y.Explosive + 90;
ELSE IF 2 * X.Explosive > Y.Explosive, THEN HR = (2 * X.Explosive - Y.Explosive) * 30 / Y.Explosive + 60;
ELSE HR = 3 * (X.Explosive * 3 - Y.Explosive) * 10 / Y.Explosive + 30;
MinHitRatio is 30.
MaxHitRatio is 100.

普通攻击伤害值(StrikeDamage):
IF X.Attack > Y.Defence then SD = X.Lv + 25 + (X.Attack - Y.Defence) / 2
ELSE SD = X.Lv + 25 - (Y.Defence - X.Attack) / 2.

暴击概率:
如果X.Morale < Y.Morale，那么 CSC = 1;
如果Y.Morale <= X.Morale < 2 * Y.Morale，那么 CSC = 2 + 18 *(X.Morale / X.Morale - 1);
如果Y.Morale * 2 <= X.Morale < 3 * Y.Morale，那么 CSC = 20 + 80 * (X.Morale / Y.Morale - 2);
// 此段H上升很快。所以尽量让我方武将攻击那些爆发力差不多或接近其三倍的敌部队。
如果X.Morale >= 3 * Y.Morale，那么 CSC = 100;

暴击伤害值(CriticalStrikeDamage):
CSD = SD * 1.7

双击概率(同暴击概率计算，Morale -> Explosive)
*/

func GetBonus(x, y *c2char.Unit) int16 {
	lookUp, ok := AttackBonusLookUp[x.UnitCategoryId] 
	if !ok { // 没有查到
		return 100
	}
    bonus, ok := lookUp[y.UnitCategoryId]
	if !ok { // 没有查到
		return 100
	}
	return bonus
}

func CalculateDamage(x, y *c2char.Unit) int16 {
	var damage int16 
	if x.GetAttack() > y.GetDefence() {
		damage = x.Level + 25 + (x.GetAttack() - y.GetDefence()) / 2
	} else {
		damage = x.Level + 25 - (y.GetDefence() - x.GetAttack()) / 2
	}
	// 考虑兵种的克制关系
	damage = damage * (100 + GetBonus(x, y)) / 100
	if damage < 0 {
		damage := 1
	}
	return damage
}

func CalculateHitRatio(attacker, defender *c2char.Unit) int16 {
	return CalculateDouble100(attacker.GetExplosive(), defender.GetExplosive())
}

func CalculateCriticalStrikeChance(attacker, defender *c2char.Unit) int16 {
	return CalculateTriple100(attacker.GetMorale(), defender.GetMorale())
}

func CalculateDoubleStrikeChance(attacker, defender *c2char.Unit) int16 {
	return CalculateTriple100(attacker.GetExplosive(), defender.GetExplosive())
}

func CalculateDouble100(x, y int16) int16 {
	if x > y*2 {
		return 100
	} else if x > y {
		return (x-y)*10/y + 90
	} else if x*2 > y {
		return (x*2-y)*30/y + 60
	} else if x*3 > y {
		return (x*3-y)*30/y + 30
	} else {
		return 30
	}
}

func CalculateTriple100(x, y int16) int16 {
	if x < y {
		return 1
	} else if x < 2*y {
		return 2 + 18*(x*100/y-100)/100
	} else if x < 3*y {
		return 20 + 80*(x*100/y-100)/100
	} else {
		return 100
	}
}

// 实际发生攻击
func TriggerAttack() bool {

}

func ProcessAttack(attacker, defender *c2char.Unit) bool {

	hitRatio := CalculateHitRatio(attacker, defender)
	damage := CalculateDamage(attacker, defender)
	// (1) 判断是否暴击 并 计算命中率（普通攻击命中率和暴击命中率不同）
	csc := CalculateCriticalStrikeChance(attacker, defender)
	prob := int16(RandIntGen.Int31n(c2char.HUNDRED))
	var csf bool 
	if prob <= csc {
csf = true
	} else {
		csf = false
	}
	// (2) 判断是否命中，结算伤害

	// (3) 判断是否能反击（反击只能暴击、不能连击）

	// (4) 判断反击后是否反击（这次攻击属于反击）

	// (5) 判断是否双击 并 计算命中率(双击属于普通攻击)
	return true

}

func PreviewAttack(attacker, defender *c2char.Unit) (hitRatio, damage int16) {
	return CalculateHitRatio(attacker, defender), CalculateDamage(attacker, defender)
}
