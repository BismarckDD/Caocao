package c2attack

import "github.com/BismarckDD/Caocao/c2core/c2unit_category"

var AttackBonusLookUp map[c2unit_category.UnitCategoryId]map[c2unit_category.UnitCategoryId]int8 = map[c2unit_category.UnitCategoryId]map[c2unit_category.UnitCategoryId]int8{
	c2unit_category.UnitCategoryBuBing: map[c2unit_category.UnitCategoryId]int8{
		c2unit_category.UnitCategoryQiBing:   -50,
		c2unit_category.UnitCategoryGongBing: 50,
		c2unit_category.UnitCategoryPaoChe:   50,
	},
	// c2unit_category.UnitCategoryQiBing: map[c2unit_category.UnitCategoryId]int8{
	// 	c2unit_category.UnitCategoryBuBing: 50,
	// },
	c2unit_category.UnitCategoryGongBing: map[c2unit_category.UnitCategoryId]int8{
		c2unit_category.UnitCategoryQiBing: 50,
		// c2unit_category.UnitCategoryXiLiangQiBing: 50,
		// c2unit_category.UnitCategoryGongQiBing: 50,
		c2unit_category.UnitCategoryXunXiongShi: 50,
		c2unit_category.UnitCategoryXunHuShi:    50,
	},
	c2unit_category.UnitCategoryGongQiBing: map[c2unit_category.UnitCategoryId]int8{
		c2unit_category.UnitCategoryQiBing: 50,
		// c2unit_category.UnitCategoryXiLiangQiBing: 50,
		c2unit_category.UnitCategoryXunXiongShi: 50,
		c2unit_category.UnitCategoryXunHuShi:    50,
	},
	c2unit_category.UnitCategoryPaoChe: map[c2unit_category.UnitCategoryId]int8{
		c2unit_category.UnitCategoryQiBing: 50,
		// c2unit_category.UnitCategoryXiLiangQiBing: 50,
		// c2unit_category.UnitCategoryGongQiBing: 50,
		c2unit_category.UnitCategoryXunXiongShi: 50,
		c2unit_category.UnitCategoryXunHuShi:    50,
	},
}
