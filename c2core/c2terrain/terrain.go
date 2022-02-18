package c2terrain

type Terrain uint8

const ( // 这里改成普通枚举类型
	TerrainPlain       Terrain = iota // 平原
	TerrainPrairie                    // 草原
	TerrainWoods                      // 树林
	TerrainWasteland                  // 荒地
	TerrainMoutains                   // 山地
	TerrainIwayama                    // 岩山 // --
	TerrainCliff                      // 山崖 // --
	TerrainSnowland                   // 雪原
	TerrainBridge                     // 桥梁
	TerrainShoal                      // 浅滩
	TerrainSwamp                      // 沼泽
	TerrainPood                       // 池塘 // --
	TerrainRivulet                    // 小河 // --
	TerrainRiver                      // 大河
	TerrainFence                      // 栅栏 // --
	TerrainCityWall                   // 城墙 // --
	TerrainCity                       // 城内
	TerrainCityGate                   // 城门 // --
	TerrainCityAndMoat                // 城池
	TerrainPass                       // 关隘
	TerrainAbatis                     // 鹿砦
	TerrainVillge                     // 村庄
	TerrainBarracks                   // 兵营
	TerrainHouse                      // 民居
	TerrainTreasury                   // 宝物库
	TerrainPool                       // 水池 // --
	TerrainFire                       // 火 // --
	TerrainBoat                       // 船 // --
	TerrainEnd         = 31
)
