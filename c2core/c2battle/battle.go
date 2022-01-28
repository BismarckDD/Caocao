package c2battle

/*

 */
type Battle struct {
	charWe     []*BattleUnit
	charFriend []*BattleUnit
	charEnemy  []*BattleUnit
	battleMap  *BattleMap
}

func (battle *Battle) AddChar()
