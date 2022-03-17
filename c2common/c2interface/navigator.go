package c2interface

import "github.com/BismarckDD/Caocao/c2common/c2gs"

// Navigator is used for transitioning between game screens
type Navigator interface {
	ToMainMenu()                                                // 主界面
	ToBattle(battleId uint16, gameState ...*c2gs.GameState)     // 进入战斗场景
	ToCampsite(campSiteId uint16, gameState ...*c2gs.GameState) // 进入战前准备场景
	ToScenario(scenarioId uint16, gameState ...*c2gs.GameState) // 进入剧情场景
	ToBattleTest()                                              // 战斗测试界面 debug
	EndGame()                                                   // 游戏结束
}
