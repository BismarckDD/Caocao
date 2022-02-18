package c2battle

type Direction int

type BattleUnit struct {
	UnitId      int     // 对应一个兵种的属性
	PosX        int     // 战场位置
	PosY        int     // 战场位置
	pBattle     *Battle // 对应
	NDirection  Direction
	BaseAttack  int // 攻击力
	BaseDefence int // 防御力
	BaseSpirit  int // 精神力
	BaseDamage  int // 攻击力
	BaseLuck    int // 运气
}

// 获得攻击力
func (battleUnit *BattleUnit) GetDamage() int {
	return unit.GetAttack()
}

// 获得精神力
func (battleUnit *BattleUnit) GetSpirit() int {
	return battleUnit.BaseDamage
}

// 获得敏捷

func (battleUnit *BattleUnit) GetDamage() int {
	return battleUnit.BaseDamage
}

func (battleUnit *BattleUnit) GetDamage() int {
	return battleUnit.BaseDamage
}

// move
func (battleUnit *BattleUnit) Move() error {
	return nil
}

// 攻击
func (battleUnit *BattleUnit) Attack() error {
	return nil
}

// 防御
func (battleUnit *BattleUnit) Defend() error {
	return nil
}

// 施展魔法
func (battleUnit *BattleUnit) Spell() error {
	return nil
}

//
func (battleUnit *BattleUnit) Excite() error {

}
