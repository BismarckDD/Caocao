package c2unit_category

type EculidRange uint8

// for attack range && spell range
type AttackRange uint8

// combination of its EculidRange
type AttackEffectRange uint8

const (
	EculidRange0   EculidRange = iota // e.g. Magic: BaQi
	EculidRange1                      // left, top, right, bottom
	EculidRange21                     // left-2, top-2, right-2, bottom-2
	EculidRange22                     // lt, lb, rt, rb
	EculidRange3                      // eculid-3 range, NuBing
	EculidRange4                      // eculid-4 range, LianNuBing
	EculidRange5                      // eculid-5 range, PiLiChe
	EculidRangeInf                    // e.g. Magic: ShuSong
)

const (
	AttackRangeNone AttackRange = 0                                   // e.g. Magic: Weather
	AttackRangeInf  AttackRange = 1 << EculidRangeInf                 // e.g. Magic: ShuSong
	AttackRange0    AttackRange = 1 << EculidRange0                   // e.g. Magic: BaQi
	AttackRange1    AttackRange = 1 << EculidRange1                   // (Qing or Zhong)Qibing 4-grid
	AttackRange2    AttackRange = 1 << EculidRange21                  // e.g. GongQiBing
	AttackRange3    AttackRange = AttackRange1 + AttackRange2         // e.g. Magic: BaoYan
	AttackRange4    AttackRange = AttackRange1 + (1 << EculidRange22) // e.g. QinWeiDui 8-grid
	AttackRange5    AttackRange = AttackRange2 + (1 << EculidRange22) // e.g. GongBing eculid-2
	AttackRange6    AttackRange = AttackRange5 + (1 << EculidRange3)  // e.g. NuBing eculid-2 + eculid-3
	AttackRange7    AttackRange = (1 << EculidRange4)                 // e.g. (Qing or Zhong)PaoChe eculid-4
	AttackRange8    AttackRange = AttackRange6 + AttackRange7         // e.g. LianNuBing eculid-2 + eculid-3 + eculid-4
	AttackRange9    AttackRange = AttackRange7 + (1 << EculidRange5)  // e.g. PiLiChe
	AttackRange10   AttackRange = AttackRange3 + (1 << EculidRange22) // e.g. Magic: XuanYun
	AttackRange11   AttackRange = AttackRange10 + (1 << EculidRange3) // e.g. Magic: DaBuji
	AttackRange12   AttackRange = AttackRange11 + (1 << EculidRange4) // e.g. Magic: XiaoBuji
)

const (
	AttackEffectRangeNone  AttackEffectRange = 0                                             // e.g. Magic: Weather
	AttackEffectRangeInf   AttackEffectRange = 1 << EculidRangeInf                           // e.g. Magic: ShaBao
	AttackEffectRange1Grid AttackEffectRange = 1 << EculidRange0                             // Normal, for most unit, 1-grid
	AttackEffectRange5Grid AttackEffectRange = AttackEffectRange1Grid + (1 << EculidRange1)  // e.g. Char:PiLiChe; Magic: HuoLong
	AttackEffectRange9Grid AttackEffectRange = AttackEffectRange5Grid + (1 << EculidRange22) // e.g. Magic: YuanJun, ZhuQue
	AttackEffectRange2Line AttackEffectRange = 1 << EculidRange0
	AttackEffectRange4Line AttackEffectRange = 1 << EculidRange0 // attack for a line contains 4 grid
	AttackEffectRange5Line AttackEffectRange = 1 << EculidRange0 // attack for a line contains 4 grid
)
