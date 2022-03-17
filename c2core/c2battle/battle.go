package c2battle

/*

 */
type Battle struct {
	charWe       []*BattleUnit
	charFriend   []*BattleUnit
	charEnemy    []*BattleUnit
	battleMap    *BattleMap
	maxRound     uint8 // 最大回合数
	winCondition uint8 //
}

// 配合脚本添加Unit
func (battle *Battle) AddUnit() {

	// 补充音效
}

func (battle *Battle) UpdateMap(x, y int) {

	// 补充音效
}
