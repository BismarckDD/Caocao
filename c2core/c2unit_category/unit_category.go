package c2unit_category

// Level相关常量
const (
	MaxLevel         = 50
	TransferLevelOne = 15
	TransferLevelTwo = 30
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

	UnitCategoryWenGuan = UnitCategoryDaoShi | UnitCategoryCeShi | UnitCategoryFengShuiShi | UnitCategoryQiMaCeShi |
		UnitCategoryDuDu | UnitCategoryZhouShuShi | UnitCategoryXianRen

	UnitCategoryWuJiang = UnitCategoryBuBing | UnitCategoryQiBing | UnitCategoryGongBing | UnitCategoryGongQiBing |
		UnitCategoryXiLiangQiBing | UnitCategoryQunXiong

	UnitCategoryZongHe = UnitCategoryWuDaoJia | UnitCategoryZeiBing | UnitCategoryWuNiang | UnitCategoryPaoChe |
		UnitCategoryHuangJinZei | UnitCategoryHaiDao | UnitCategoryXunXiongShi | UnitCategoryXunHuShi |
		UnitCategoryHuangDi | UnitCategoryBaiXing | UnitCategoryMuRen | UnitCategoryTuOu | UnitCategoryYunShuDui

	UnitCategoryForSword = UnitCategoryQunXiong | UnitCategoryBuBing | UnitCategoryZeiBing | UnitCategoryHaiDao | UnitCategoryHuangJinZei
	UnitCategoryForLance = UnitCategoryQiBing | UnitCategoryXiLiangQiBing
	UnitCategoryForBow   = UnitCategoryGongBing | UnitCategoryGongQiBing
	UnitCategoryForClub  = UnitCategoryWuDaoJia | UnitCategoryWuNiang | UnitCategoryXunHuShi | UnitCategoryXunXiongShi |
		UnitCategoryMuRen | UnitCategoryTuOu | UnitCategoryBaiXing | UnitCategoryHuangDi
	UnitCategoryForCannon     = UnitCategoryPaoChe
	UnitCategoryForMagicSword = UnitCategoryDaoShi | UnitCategoryFengShuiShi | UnitCategoryXianRen | UnitCategoryDuDu
	UnitCategoryForMagicFan   = UnitCategoryCeShi | UnitCategoryQiMaCeShi | UnitCategoryZhouShuShi

	UnitCategoryForArmor = UnitCategoryWuJiang
	UnitCategoryForCloth = UnitCategoryWenGuan | UnitCategoryZongHe
	UnitCategoryForAll   = UnitCategoryForArmor | UnitCategoryForCloth
)

// size: 64
type UnitCategoryBase struct {
	UnitCategoryId uint32 // 支持32个兵种分类, 每个兵种可以转职
	AttackInc      uint8  // X:5 S:4 A:3 B:2 C:1
	DefenseInc     uint8  // X:5 S:4 A:3 B:2 C:1
	SpiritInc      uint8  // X:5 S:4 A:3 B:2 C:1
	ExplosiveInc   uint8  // X:5 S:4 A:3 B:2 C:1
	MoraleInc      uint8  // X:5 S:4 A:3 B:2 C:1
	UnitBaseHP     uint8  // HP when lv1
	UnitBaseMP     uint8  // MP when lv1
	UnitExtraHP    uint8  // inc when levelup
	UnitExtraMP    uint8  // inc
}

// UnitCategory 相对于 UnitCategoryBase
// 弓兵、弩兵、连弩兵 -> 弓兵
type UnitCategory struct {
	UnitCategoryBase
	CategoryLevel uint8
	AttackRange
	AttackEffectRange
	Move          uint8 // 移动力上限255
	TransferId    uint8 // 转职目标在CategoryList的下标
	TransferLevel uint8 // 转职需要的等级需求
	BattleFieldId uint8 // Battle field vision.
	Name          string
}

var UnitCategoryBaseList = []UnitCategoryBase{
	{ // JunZhu
		UnitCategoryId: 0x1,
		AttackInc:      3,
		SpiritInc:      3,
		DefenseInc:     3,
		ExplosiveInc:   3,
		MoraleInc:      3,
		UnitBaseHP:     100,
		UnitBaseMP:     30,
		UnitExtraHP:    5,
		UnitExtraMP:    1,
	},
	{ // BuBing
		UnitCategoryId: 0x2,
		AttackInc:      2,
		SpiritInc:      3,
		DefenseInc:     4,
		ExplosiveInc:   2,
		MoraleInc:      2,
		UnitBaseHP:     110,
		UnitBaseMP:     10,
		UnitExtraHP:    6,
		UnitExtraMP:    1,
	},
	{ // QiBing
		UnitCategoryId: 0x4,
		AttackInc:      4,
		SpiritInc:      2,
		DefenseInc:     3,
		ExplosiveInc:   2,
		MoraleInc:      2,
		UnitBaseHP:     100,
		UnitBaseMP:     10,
		UnitExtraHP:    5,
		UnitExtraMP:    1,
	},
	{ // GongBing
		UnitCategoryId: 0x8,
		AttackInc:      3,
		SpiritInc:      2,
		DefenseInc:     2,
		ExplosiveInc:   2,
		MoraleInc:      4,
		UnitBaseHP:     90,
		UnitBaseMP:     10,
		UnitExtraHP:    4,
		UnitExtraMP:    1,
	},
	{ // GongQiBing
		UnitCategoryId: 0x10,
		AttackInc:      4,
		SpiritInc:      2,
		DefenseInc:     2,
		ExplosiveInc:   2,
		MoraleInc:      3,
		UnitBaseHP:     100,
		UnitBaseMP:     10,
		UnitExtraHP:    5,
		UnitExtraMP:    1,
	},
	{ // WuDaoJia
		UnitCategoryId: 0x20,
		AttackInc:      3,
		SpiritInc:      1,
		DefenseInc:     3,
		ExplosiveInc:   4,
		MoraleInc:      2,
		UnitBaseHP:     90,
		UnitBaseMP:     20,
		UnitExtraHP:    4,
		UnitExtraMP:    1,
	},
	{ // ZeiBing
		UnitCategoryId: 0x40,
		AttackInc:      4,
		SpiritInc:      1,
		DefenseInc:     2,
		ExplosiveInc:   2,
		MoraleInc:      4,
		UnitBaseHP:     100,
		UnitBaseMP:     20,
		UnitExtraHP:    5,
		UnitExtraMP:    1,
	},
	{ // WuNiang
		UnitCategoryId: 0x80,
		AttackInc:      3,
		SpiritInc:      2,
		DefenseInc:     2,
		ExplosiveInc:   4,
		MoraleInc:      2,
		UnitBaseHP:     90,
		UnitBaseMP:     20,
		UnitExtraHP:    3,
		UnitExtraMP:    1,
	},
	{ // PaoChe
		UnitCategoryId: 0x100,
		AttackInc:      4,
		SpiritInc:      2,
		DefenseInc:     3,
		ExplosiveInc:   1,
		MoraleInc:      3,
		UnitBaseHP:     90,
		UnitBaseMP:     10,
		UnitExtraHP:    4,
		UnitExtraMP:    1,
	},
	{ // DaoShi
		UnitCategoryId: 0x200,
		AttackInc:      1,
		SpiritInc:      4,
		DefenseInc:     2,
		ExplosiveInc:   3,
		MoraleInc:      2,
		UnitBaseHP:     80,
		UnitBaseMP:     40,
		UnitExtraHP:    3,
		UnitExtraMP:    2,
	},
	{ // CeShi
		UnitCategoryId: 0x400,
		AttackInc:      2,
		SpiritInc:      4,
		DefenseInc:     2,
		ExplosiveInc:   2,
		MoraleInc:      2,
		UnitBaseHP:     90,
		UnitBaseMP:     40,
		UnitExtraHP:    4,
		UnitExtraMP:    2,
	},
	{ // FengShuiShi
		UnitCategoryId: 0x800,
		AttackInc:      1,
		SpiritInc:      4,
		DefenseInc:     1,
		ExplosiveInc:   3,
		MoraleInc:      3,
		UnitBaseHP:     80,
		UnitBaseMP:     50,
		UnitExtraHP:    3,
		UnitExtraMP:    2,
	},
	{ // QiMaCeShi
		UnitCategoryId: 0x1000,
		AttackInc:      3,
		SpiritInc:      4,
		DefenseInc:     2,
		ExplosiveInc:   2,
		MoraleInc:      1,
		UnitBaseHP:     100,
		UnitBaseMP:     40,
		UnitExtraHP:    5,
		UnitExtraMP:    2,
	},
	{ // XiLiangQiBing
		UnitCategoryId: 0x2000,
		AttackInc:      4,
		SpiritInc:      1,
		DefenseInc:     4,
		ExplosiveInc:   2,
		MoraleInc:      2,
		UnitBaseHP:     110,
		UnitBaseMP:     50,
		UnitExtraHP:    6,
		UnitExtraMP:    1,
	},
	{ // HuangJinZei
		UnitCategoryId: 0x4000,
		AttackInc:      2,
		SpiritInc:      1,
		DefenseInc:     2,
		ExplosiveInc:   2,
		MoraleInc:      1,
		UnitBaseHP:     90,
		UnitBaseMP:     40,
		UnitExtraHP:    4,
		UnitExtraMP:    1,
	},
	{ // HaiDao
		UnitCategoryId: 0x8000,
		AttackInc:      4,
		SpiritInc:      2,
		DefenseInc:     2,
		ExplosiveInc:   3,
		MoraleInc:      2,
		UnitBaseHP:     90,
		UnitBaseMP:     20,
		UnitExtraHP:    4,
		UnitExtraMP:    1,
	},
	{ // DuDu
		UnitCategoryId: 0x10000,
		AttackInc:      3,
		SpiritInc:      4,
		DefenseInc:     2,
		ExplosiveInc:   2,
		MoraleInc:      2,
		UnitBaseHP:     90,
		UnitBaseMP:     30,
		UnitExtraHP:    4,
		UnitExtraMP:    2,
	},
	{ // ZhouShuShi
		UnitCategoryId: 0x20000,
		AttackInc:      1,
		SpiritInc:      4,
		DefenseInc:     2,
		ExplosiveInc:   2,
		MoraleInc:      3,
		UnitBaseHP:     80,
		UnitBaseMP:     60,
		UnitExtraHP:    3,
		UnitExtraMP:    3,
	},
	{ // XianRen
		UnitCategoryId: 0x40000,
		AttackInc:      1,
		SpiritInc:      4,
		DefenseInc:     1,
		ExplosiveInc:   3,
		MoraleInc:      4,
		UnitBaseHP:     80,
		UnitBaseMP:     60,
		UnitExtraHP:    3,
		UnitExtraMP:    3,
	},
	{ // XunXiongShi
		UnitCategoryId: 0x80000,
		AttackInc:      4,
		SpiritInc:      2,
		DefenseInc:     3,
		ExplosiveInc:   2,
		MoraleInc:      2,
		UnitBaseHP:     110,
		UnitBaseMP:     5,
		UnitExtraHP:    6,
		UnitExtraMP:    1,
	},
	{ // XunHuShi
		UnitCategoryId: 0x10000,
		AttackInc:      3,
		SpiritInc:      1,
		DefenseInc:     3,
		ExplosiveInc:   4,
		MoraleInc:      1,
		UnitBaseHP:     90,
		UnitBaseMP:     5,
		UnitExtraHP:    4,
		UnitExtraMP:    1,
	},
}
var UnitCategoryList = []UnitCategory{
	{ // QunXiong
		UnitCategoryBase:  UnitCategoryBaseList[UnitCategoryQunXiong],
		CategoryLevel:     1,
		AttackRange:       AttackRange1,
		AttackEffectRange: AttackEffectRange1Grid,
		Move:              6,
		TransferId:        1,
		TransferLevel:     TransferLevelOne,
		BattleFieldId:     0,
		Name:              "群雄",
	},
	{
		UnitCategoryBase:  UnitCategoryBaseList[UnitCategoryQunXiong],
		CategoryLevel:     2,
		AttackRange:       AttackRange4,
		AttackEffectRange: AttackEffectRange1Grid,
		Move:              6,
		TransferId:        2,
		TransferLevel:     TransferLevelTwo,
		BattleFieldId:     1,
		Name:              "英雄",
	},
	{
		UnitCategoryBase:  UnitCategoryBaseList[UnitCategoryQunXiong],
		CategoryLevel:     3,
		AttackRange:       AttackRange4,
		AttackEffectRange: AttackEffectRange1Grid,
		Move:              7,
		TransferId:        0,
		TransferLevel:     0,
		BattleFieldId:     2,
		Name:              "霸王",
	},
	{
		UnitCategoryBase:  UnitCategoryBaseList[UnitCategoryBuBing],
		CategoryLevel:     1,
		AttackRange:       AttackRange4,
		AttackEffectRange: AttackEffectRange1Grid,
		Move:              4,
		TransferId:        4,
		TransferLevel:     TransferLevelOne,
		BattleFieldId:     3,
		Name:              "轻步兵",
	},
	{
		UnitCategoryBase:  UnitCategoryBaseList[UnitCategoryBuBing],
		CategoryLevel:     2,
		AttackRange:       AttackRange4,
		AttackEffectRange: AttackEffectRange1Grid,
		Move:              4,
		TransferId:        5,
		TransferLevel:     TransferLevelTwo,
		BattleFieldId:     4,
		Name:              "重步兵",
	},
	{
		UnitCategoryBase:  UnitCategoryBaseList[UnitCategoryBuBing],
		CategoryLevel:     3,
		AttackRange:       AttackRange4,
		AttackEffectRange: AttackEffectRange1Grid,
		Move:              5,
		TransferId:        0,
		TransferLevel:     0,
		BattleFieldId:     5,
		Name:              "近卫兵",
	},
	{
		UnitCategoryBase:  UnitCategoryBaseList[UnitCategoryQiBing],
		CategoryLevel:     1,
		AttackRange:       AttackRange1,
		AttackEffectRange: AttackEffectRange1Grid,
		Move:              6,
		TransferId:        7,
		TransferLevel:     TransferLevelOne,
		BattleFieldId:     6,
		Name:              "轻骑兵",
	},
	{
		UnitCategoryBase:  UnitCategoryBaseList[UnitCategoryQiBing],
		CategoryLevel:     2,
		AttackRange:       AttackRange1,
		AttackEffectRange: AttackEffectRange1Grid,
		Move:              6,
		TransferId:        8,
		TransferLevel:     TransferLevelTwo,
		BattleFieldId:     7,
		Name:              "重骑兵",
	},
	{
		UnitCategoryBase:  UnitCategoryBaseList[UnitCategoryQiBing],
		CategoryLevel:     3,
		AttackRange:       AttackRange4,
		AttackEffectRange: AttackEffectRange1Grid,
		Move:              7,
		TransferId:        0,
		TransferLevel:     0,
		BattleFieldId:     8,
		Name:              "亲卫队",
	},
	{
		UnitCategoryBase:  UnitCategoryBaseList[UnitCategoryGongBing],
		CategoryLevel:     1,
		AttackRange:       AttackRange5,
		AttackEffectRange: AttackEffectRange1Grid,
		Move:              4,
		TransferId:        10,
		TransferLevel:     TransferLevelOne,
		BattleFieldId:     9,
		Name:              "弓兵",
	},
	{
		UnitCategoryBase:  UnitCategoryBaseList[UnitCategoryGongBing],
		CategoryLevel:     2,
		AttackRange:       AttackRange6,
		AttackEffectRange: AttackEffectRange1Grid,
		Move:              4,
		TransferId:        11,
		TransferLevel:     TransferLevelTwo,
		BattleFieldId:     10,
		Name:              "弩兵",
	},
	{
		UnitCategoryBase:  UnitCategoryBaseList[UnitCategoryQiBing],
		CategoryLevel:     3,
		AttackRange:       AttackRange8,
		AttackEffectRange: AttackEffectRange1Grid,
		Move:              5,
		TransferId:        0,
		TransferLevel:     0,
		BattleFieldId:     11,
		Name:              "连弩兵",
	},
	{
		UnitCategoryBase:  UnitCategoryBaseList[UnitCategoryGongQiBing],
		CategoryLevel:     1,
		AttackRange:       AttackRange2,
		AttackEffectRange: AttackEffectRange1Grid,
		Move:              6,
		TransferId:        13,
		TransferLevel:     TransferLevelOne,
		BattleFieldId:     12,
		Name:              "弓骑兵",
	},
	{
		UnitCategoryBase:  UnitCategoryBaseList[UnitCategoryGongQiBing],
		CategoryLevel:     2,
		AttackRange:       AttackRange5,
		AttackEffectRange: AttackEffectRange1Grid,
		Move:              6,
		TransferId:        14,
		TransferLevel:     TransferLevelTwo,
		BattleFieldId:     13,
		Name:              "弩骑兵",
	},
	{
		UnitCategoryBase:  UnitCategoryBaseList[UnitCategoryGongQiBing],
		CategoryLevel:     3,
		AttackRange:       AttackRange6,
		AttackEffectRange: AttackEffectRange1Grid,
		Move:              7,
		TransferId:        0,
		TransferLevel:     0,
		BattleFieldId:     14,
		Name:              "连弩骑兵",
	},
	{
		UnitCategoryBase:  UnitCategoryBaseList[UnitCategoryWuDaoJia],
		CategoryLevel:     1,
		AttackRange:       AttackRange1,
		AttackEffectRange: AttackEffectRange1Grid,
		Move:              5,
		TransferId:        16,
		TransferLevel:     TransferLevelOne,
		BattleFieldId:     15,
		Name:              "武道家",
	},
	{
		UnitCategoryBase:  UnitCategoryBaseList[UnitCategoryWuDaoJia],
		CategoryLevel:     2,
		AttackRange:       AttackRange4,
		AttackEffectRange: AttackEffectRange1Grid,
		Move:              5,
		TransferId:        17,
		TransferLevel:     TransferLevelTwo,
		BattleFieldId:     16,
		Name:              "拳士",
	},
	{
		UnitCategoryBase:  UnitCategoryBaseList[UnitCategoryWuDaoJia],
		CategoryLevel:     3,
		AttackRange:       AttackRange4,
		AttackEffectRange: AttackEffectRange1Grid,
		Move:              6,
		TransferId:        0,
		TransferLevel:     0,
		BattleFieldId:     17,
		Name:              "拳圣",
	},
	{
		UnitCategoryBase:  UnitCategoryBaseList[UnitCategoryZeiBing],
		CategoryLevel:     1,
		AttackRange:       AttackRange1,
		AttackEffectRange: AttackEffectRange1Grid,
		Move:              5,
		TransferId:        19,
		TransferLevel:     TransferLevelOne,
		BattleFieldId:     18,
		Name:              "贼兵",
	},
	{
		UnitCategoryBase:  UnitCategoryBaseList[UnitCategoryZeiBing],
		CategoryLevel:     2,
		AttackRange:       AttackRange4,
		AttackEffectRange: AttackEffectRange1Grid,
		Move:              5,
		TransferId:        20,
		TransferLevel:     TransferLevelTwo,
		BattleFieldId:     19,
		Name:              "义贼",
	},
	{
		UnitCategoryBase:  UnitCategoryBaseList[UnitCategoryZeiBing],
		CategoryLevel:     3,
		AttackRange:       AttackRange4,
		AttackEffectRange: AttackEffectRange1Grid,
		Move:              6,
		TransferId:        0,
		TransferLevel:     0,
		BattleFieldId:     20,
		Name:              "豪杰",
	},
	{
		UnitCategoryBase:  UnitCategoryBaseList[UnitCategoryWuNiang],
		CategoryLevel:     1,
		AttackRange:       AttackRange1,
		AttackEffectRange: AttackEffectRange1Grid,
		Move:              5,
		TransferId:        22,
		TransferLevel:     TransferLevelOne,
		BattleFieldId:     21,
		Name:              "舞娘",
	},
	{
		UnitCategoryBase:  UnitCategoryBaseList[UnitCategoryWuNiang],
		CategoryLevel:     2,
		AttackRange:       AttackRange4,
		AttackEffectRange: AttackEffectRange1Grid,
		Move:              5,
		TransferId:        23,
		TransferLevel:     TransferLevelTwo,
		BattleFieldId:     22,
		Name:              "舞伎",
	},
	{
		UnitCategoryBase:  UnitCategoryBaseList[UnitCategoryWuNiang],
		CategoryLevel:     3,
		AttackRange:       AttackRange4,
		AttackEffectRange: AttackEffectRange1Grid,
		Move:              6,
		TransferId:        0,
		TransferLevel:     0,
		BattleFieldId:     23,
		Name:              "巫女",
	},
	{
		UnitCategoryBase:  UnitCategoryBaseList[UnitCategoryPaoChe],
		CategoryLevel:     1,
		AttackRange:       AttackRange7,
		AttackEffectRange: AttackEffectRange5Grid,
		Move:              3,
		TransferId:        25,
		TransferLevel:     TransferLevelOne,
		BattleFieldId:     24,
		Name:              "轻炮车",
	},
	{
		UnitCategoryBase:  UnitCategoryBaseList[UnitCategoryPaoChe],
		CategoryLevel:     2,
		AttackRange:       AttackRange7,
		AttackEffectRange: AttackEffectRange1Grid,
		Move:              3,
		TransferId:        26,
		TransferLevel:     TransferLevelTwo,
		BattleFieldId:     25,
		Name:              "重炮车",
	},
	{
		UnitCategoryBase:  UnitCategoryBaseList[UnitCategoryPaoChe],
		CategoryLevel:     3,
		AttackRange:       AttackRange9,
		AttackEffectRange: AttackEffectRange5Grid,
		Move:              4,
		TransferId:        0,
		TransferLevel:     0,
		BattleFieldId:     26,
		Name:              "霹雳车",
	},
	{
		UnitCategoryBase:  UnitCategoryBaseList[UnitCategoryDaoShi],
		CategoryLevel:     1,
		AttackRange:       AttackRange1,
		AttackEffectRange: AttackEffectRange1Grid,
		Move:              4,
		TransferId:        28,
		TransferLevel:     TransferLevelOne,
		BattleFieldId:     27,
		Name:              "道士",
	},
	{
		UnitCategoryBase:  UnitCategoryBaseList[UnitCategoryDaoShi],
		CategoryLevel:     2,
		AttackRange:       AttackRange1,
		AttackEffectRange: AttackEffectRange1Grid,
		Move:              4,
		TransferId:        29,
		TransferLevel:     TransferLevelTwo,
		BattleFieldId:     28,
		Name:              "幻术士",
	},
	{
		UnitCategoryBase:  UnitCategoryBaseList[UnitCategoryDaoShi],
		CategoryLevel:     3,
		AttackRange:       AttackRange1,
		AttackEffectRange: AttackEffectRange1Grid,
		Move:              5,
		TransferId:        0,
		TransferLevel:     0,
		BattleFieldId:     29,
		Name:              "妖术士",
	},
	{
		UnitCategoryBase:  UnitCategoryBaseList[UnitCategoryCeShi],
		CategoryLevel:     1,
		AttackRange:       AttackRange1,
		AttackEffectRange: AttackEffectRange1Grid,
		Move:              4,
		TransferId:        31,
		TransferLevel:     TransferLevelOne,
		BattleFieldId:     30, // also fo QiMaCeShi
		Name:              "策士",
	},
	{
		UnitCategoryBase:  UnitCategoryBaseList[UnitCategoryCeShi],
		CategoryLevel:     2,
		AttackRange:       AttackRange1,
		AttackEffectRange: AttackEffectRange1Grid,
		Move:              4,
		TransferId:        32,
		TransferLevel:     TransferLevelTwo,
		BattleFieldId:     31, // also for QiMaCanMou
		Name:              "参谋",
	},
	{
		UnitCategoryBase:  UnitCategoryBaseList[UnitCategoryCeShi],
		CategoryLevel:     3,
		AttackRange:       AttackRange1,
		AttackEffectRange: AttackEffectRange1Grid,
		Move:              5,
		TransferId:        0,
		TransferLevel:     0,
		BattleFieldId:     32, // also for DuDu, ZhouShuShi, QiMaJunShi
		Name:              "军师",
	},
	{
		UnitCategoryBase:  UnitCategoryBaseList[UnitCategoryFengShuiShi],
		CategoryLevel:     1,
		AttackRange:       AttackRange1,
		AttackEffectRange: AttackEffectRange1Grid,
		Move:              4,
		TransferId:        34,
		TransferLevel:     TransferLevelOne,
		BattleFieldId:     33,
		Name:              "风水士",
	},
	{
		UnitCategoryBase:  UnitCategoryBaseList[UnitCategoryFengShuiShi],
		CategoryLevel:     2,
		AttackRange:       AttackRange1,
		AttackEffectRange: AttackEffectRange1Grid,
		Move:              5,
		TransferId:        35,
		TransferLevel:     TransferLevelTwo,
		BattleFieldId:     34,
		Name:              "方术士",
	},
	{
		UnitCategoryBase:  UnitCategoryBaseList[UnitCategoryFengShuiShi],
		CategoryLevel:     3,
		AttackRange:       AttackRange1,
		AttackEffectRange: AttackEffectRange1Grid,
		Move:              6,
		TransferId:        0,
		TransferLevel:     0,
		BattleFieldId:     35,
		Name:              "仙术士",
	},
	{
		UnitCategoryBase:  UnitCategoryBaseList[UnitCategoryQiMaCeShi],
		CategoryLevel:     1,
		AttackRange:       AttackRange1,
		AttackEffectRange: AttackEffectRange1Grid,
		Move:              6,
		TransferId:        37,
		TransferLevel:     TransferLevelOne,
		BattleFieldId:     36,
		Name:              "骑马策士",
	},
	{
		UnitCategoryBase:  UnitCategoryBaseList[UnitCategoryQiMaCeShi],
		CategoryLevel:     2,
		AttackRange:       AttackRange1,
		AttackEffectRange: AttackEffectRange1Grid,
		Move:              6,
		TransferId:        38,
		TransferLevel:     TransferLevelTwo,
		BattleFieldId:     37,
		Name:              "骑马参谋",
	},
	{
		UnitCategoryBase:  UnitCategoryBaseList[UnitCategoryQiMaCeShi],
		CategoryLevel:     3,
		AttackRange:       AttackRange1,
		AttackEffectRange: AttackEffectRange1Grid,
		Move:              7,
		TransferId:        0,
		TransferLevel:     0,
		BattleFieldId:     38,
		Name:              "骑马军师",
	},
	{ // XiLiang && HuangJin
		UnitCategoryBase:  UnitCategoryBaseList[UnitCategoryXiLiangQiBing],
		CategoryLevel:     0,
		AttackRange:       AttackRange4,
		AttackEffectRange: AttackEffectRange1Grid,
		Move:              6,
		TransferId:        0,
		TransferLevel:     0,
		BattleFieldId:     39,
		Name:              "西凉骑兵",
	},
	{
		UnitCategoryBase:  UnitCategoryBaseList[UnitCategoryHuangJinZei],
		CategoryLevel:     0,
		AttackRange:       AttackRange1,
		AttackEffectRange: AttackEffectRange1Grid,
		Move:              4,
		TransferId:        0,
		TransferLevel:     0,
		BattleFieldId:     40,
		Name:              "黄巾贼",
	},
	{ // DongWuShuiJun
		UnitCategoryBase:  UnitCategoryBaseList[UnitCategoryHaiDao],
		CategoryLevel:     0,
		AttackRange:       AttackRange1,
		AttackEffectRange: AttackEffectRange1Grid,
		Move:              5,
		TransferId:        0,
		TransferLevel:     0,
		BattleFieldId:     41,
		Name:              "海盗",
	},
	{
		UnitCategoryBase:  UnitCategoryBaseList[UnitCategoryDuDu],
		CategoryLevel:     0,
		AttackRange:       AttackRange1,
		AttackEffectRange: AttackEffectRange1Grid,
		Move:              5,
		TransferId:        0,
		TransferLevel:     0,
		BattleFieldId:     32, // use JunShi vision
		Name:              "都督",
	},
	{ // XunShouShi
		UnitCategoryBase:  UnitCategoryBaseList[UnitCategoryXunXiongShi],
		CategoryLevel:     0,
		AttackRange:       AttackRange1,
		AttackEffectRange: AttackEffectRange1Grid,
		Move:              4,
		TransferId:        0,
		TransferLevel:     0,
		BattleFieldId:     43,
		Name:              "驯熊师",
	},
	{
		UnitCategoryBase:  UnitCategoryBaseList[UnitCategoryXunHuShi],
		CategoryLevel:     0,
		AttackRange:       AttackRange1,
		AttackEffectRange: AttackEffectRange1Grid,
		Move:              5,
		TransferId:        0,
		TransferLevel:     0,
		BattleFieldId:     44,
		Name:              "驯虎师",
	},
	{ // ZhouShuShi && XianRen
		UnitCategoryBase:  UnitCategoryBaseList[UnitCategoryZhouShuShi],
		CategoryLevel:     0,
		AttackRange:       AttackRange1,
		AttackEffectRange: AttackEffectRange1Grid,
		Move:              5,
		TransferId:        0,
		TransferLevel:     0,
		BattleFieldId:     32, // use JunShi vision
		Name:              "咒术士",
	},
	{
		UnitCategoryBase:  UnitCategoryBaseList[UnitCategoryXianRen],
		CategoryLevel:     0,
		AttackRange:       AttackRange1,
		AttackEffectRange: AttackEffectRange1Grid,
		Move:              7,
		TransferId:        0,
		TransferLevel:     0,
		BattleFieldId:     32, // use JunShi vision
		Name:              "仙人",
	},
	{ // MuRen && TuOU
		UnitCategoryBase:  UnitCategoryBaseList[UnitCategoryMuRen],
		CategoryLevel:     0,
		AttackRange:       AttackRange1,
		AttackEffectRange: AttackEffectRange1Grid,
		Move:              4,
		TransferId:        0,
		TransferLevel:     0,
		BattleFieldId:     47,
		Name:              "木人",
	},
	{
		UnitCategoryBase:  UnitCategoryBaseList[UnitCategoryTuOu],
		CategoryLevel:     0,
		AttackRange:       AttackRange1,
		AttackEffectRange: AttackEffectRange1Grid,
		Move:              4,
		TransferId:        0,
		TransferLevel:     0,
		BattleFieldId:     48,
		Name:              "土偶",
	},
	{ // BaiXing && HuangDi
		UnitCategoryBase:  UnitCategoryBaseList[UnitCategoryBaiXing],
		CategoryLevel:     0,
		AttackRange:       AttackRange0, // cannot attack
		AttackEffectRange: AttackEffectRange1Grid,
		Move:              4,
		TransferId:        0,
		TransferLevel:     0,
		BattleFieldId:     49,
		Name:              "百姓",
	},
	{
		UnitCategoryBase:  UnitCategoryBaseList[UnitCategoryHuangDi],
		CategoryLevel:     0,
		AttackRange:       AttackRange0, // cannot attack
		AttackEffectRange: AttackEffectRange1Grid,
		Move:              4,
		TransferId:        0,
		TransferLevel:     0,
		BattleFieldId:     50,
		Name:              "皇帝",
	},
	{ // YunShuDui
		UnitCategoryBase:  UnitCategoryBaseList[UnitCategoryYunShuDui],
		CategoryLevel:     0,
		AttackRange:       AttackRange0, // cannot attack
		AttackEffectRange: AttackEffectRange1Grid,
		Move:              4,
		TransferId:        0,
		TransferLevel:     0,
		BattleFieldId:     51,
		Name:              "辎重队",
	},
	{
		UnitCategoryBase:  UnitCategoryBaseList[UnitCategoryYunShuDui],
		CategoryLevel:     0,
		AttackRange:       AttackRange0, // cannot attack
		AttackEffectRange: AttackEffectRange1Grid,
		Move:              4,
		TransferId:        0,
		TransferLevel:     0,
		BattleFieldId:     52,
		Name:              "粮秣队",
	},
}

func LoadUnitCategory(loadPath string) []*UnitCategory {

	unitCategory := []*UnitCategory{}
	return unitCategory
}

func SaveUnitCategory(savePath string, unitCategory []*UnitCategory) bool {

	return true
}
