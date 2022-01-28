package c2initialize

import (
	"github.com/BismarckDD/Caocao/c2core/c2char"
	"github.com/BismarckDD/Caocao/c2core/c2magic"
)

/*
Damage.
(1) HaiXiao, XuanYun: all hp
(2) JuYan: (A.Level + (A.Spirit - D.Spirit) / 3) * 1.2 + 30
(3) ZhuQue: (A.Level + (A.Spirit - D.Spirit) / 3) + 25
(4) LieHuo,BaoYan,JiLiu,ShanLan,DiLong: (A.Level + (A.Spirit - D.Spirit) / 3 + 5) * 0.9 + 18
(5) ZhuoRe,HuoLong,LongJuan,ZhuoLiu,ShuiLong,LuoShi,DiZhen,DuYan,YouHuo,QingLong
(6) HuoZhen, XuanFeng, FengLong, ShuiZhen, DingShen: (A.Level + (A.Spirit - D.Spirit) / 3 + 1) * 0.5 + 12
(7) FengZhen, DuWu (A.Level + (A.Spirit - D.Spirit) / 3 + 4) * 0.4 + 8
(8) DingJun, ShaBao: (A.Level + (A.Spirit - D.Spirit) / 3 + 4) * 0.4 + 8


Debuf.
(1) DieBao: (A.Level + (A.Spirit - D.Spirit) / 3) * 0.2 + 5
(2) DunBing, DunDui: Explosive -> 70%
(3) XuTuo, YanZhan: Morale -> 70%
(4) YaPo, WeiHe: Defense -> 70%
(5) DunBing, DunDui: Explosive -> 70%
(6) ZhouFeng, ZhouYin  -> status: JinZhou
(7) HuangYan, HuangBao -> status: Chaos
(8) BaZhenTu
(9) XuanWu

Poison -> 10% hp reduce, when round start.

Buf.
(1) QiangXing: move + 2
(2) LianBing, LianDui: Explosive -> 120%
(3) AngYang, HuanSheng: Morale -> 120%
(4) FenQi, GuWu: Attack/Spirit -> 120%
(5) JianGu, QiangZhen: Defense -> 120%
(6) QiHe: Attack/Spirit -> 120%
(7) BaQi: All -> 120%
(8) JianGu, QiangZhen: Defense -> 120%
(9) HuiGui:
(10) MingXiang: -50HP, +25MP

Recover.
(1) JianYan, XianCe: +24MP, +48MP
(2) JueXing, DaJueXing:
(3) XiaoBuJi, YuanDui:
HP = 40 + A.Spirit / 10
(4) DaBuJi, YuanJun, ShuSong
HP = 70 + A.Spirit / 2
(5) BaiHu, DaBuJi + JueXing


Attack Calculation:






















*/

var MagicNames = []string{
	"落石", "地阵", "山岚", "地龙", "巨岩",
	"浊流", "水阵", "激流", "水龙", "海啸",
	"灼热", "火阵", "烈火", "火龙", "爆焰",
	"旋风", "风阵", "龙卷", "风龙", "沙暴",
	"小补给", "大补给", "输送", "援队", "援军",
	"冥想", "建言", "献策", "强行",
	"觉醒", "大觉醒", "气合", "霸气",
	"奋起", "鼓舞", "坚固", "强阵", "练兵", "练队", "昂扬", "欢声",
	"压迫", "威吓", "咒骂", "挑拨", "钝兵", "钝队", "虚脱", "厌战",
	"定身", "定军", "咒封", "咒印", "毒烟", "毒雾", "谎报", "谎言",
	"八阵图", "诱惑", "眩晕", "谍报",
	"晴明", "阴天", "豪雨", "大雪",
	"青龙", "白虎", "朱雀", "玄武",
}

var DefaultMagicList = []c2magic.Magic{
	c2magic.Magic{ // 小补给
		MagicId:          21,
		MagicTarget:      c2magic.MagicTargetWeAndFriends,
		CallAnimationId:  0, // no call animation.
		EffAnimationId:   1, // each effect
		MagicSpellRange:  c2char.AttackRange0,
		MagicEffectRange: c2char.AttackEffectRange1Grid,
		MagicName:        c2magic.ConvertTo16Bytes(MagicNames[21]),
	},
	c2magic.Magic{ // 大补给
		MagicId:          22,
		MagicTarget:      c2magic.MagicTargetWeAndFriends,
		CallAnimationId:  0, // no call animation.
		EffAnimationId:   1, // each effect
		MagicSpellRange:  c2char.AttackRange0,
		MagicEffectRange: c2char.AttackEffectRange1Grid,
		MagicName:        c2magic.ConvertTo16Bytes(MagicNames[22]),
	},
	c2magic.Magic{ //
		MagicId:          23,
		MagicTarget:      c2magic.MagicTargetWeAndFriends,
		CallAnimationId:  0, // no call animation.
		EffAnimationId:   1, // each effect
		MagicSpellRange:  c2char.AttackRange0,
		MagicEffectRange: c2char.AttackEffectRange1Grid,
		MagicName:        c2magic.ConvertTo16Bytes(MagicNames[23]),
	},
	c2magic.Magic{ //
		MagicId:          24,
		MagicTarget:      MagicTargetWeAndFriends,
		CallAnimationId:  0, // no call animation.
		EffAnimationId:   1, // each effect
		MagicSpellRange:  c2char.AttackRange0,
		MagicEffectRange: c2char.AttackEffectRange1Grid,
		MagicName:        convertTo16Bytes(MagicNames[1]),
	},
	c2magic.Magic{
		MagicId:          1,
		MagicTarget:      MagicTargetWeAndFriends,
		CallAnimationId:  0, // no call animation.
		EffAnimationId:   1, // each effect
		MagicSpellRange:  c2char.AttackRange0,
		MagicEffectRange: c2char.AttackEffectRange1Grid,
		MagicName:        convertTo16Bytes(MagicNames[1]),
	},
	Magic{
		MagicId:          1,
		MagicTarget:      MagicTargetWeAndFriends,
		CallAnimationId:  0, // no call animation.
		EffAnimationId:   1, // each effect
		MagicSpellRange:  c2char.AttackRange0,
		MagicEffectRange: c2char.AttackEffectRange1Grid,
		MagicName:        convertTo16Bytes(MagicNames[1]),
	},
	Magic{
		MagicId:          1,
		MagicTarget:      MagicTargetWeAndFriends,
		CallAnimationId:  0, // no call animation.
		EffAnimationId:   1, // each effect
		MagicSpellRange:  c2char.AttackRange0,
		MagicEffectRange: c2char.AttackEffectRange1Grid,
		MagicName:        convertTo16Bytes(MagicNames[1]),
	},
	Magic{
		MagicId:          1,
		MagicTarget:      MagicTargetWeAndFriends,
		CallAnimationId:  0, // no call animation.
		EffAnimationId:   1, // each effect
		MagicSpellRange:  c2char.AttackRange0,
		MagicEffectRange: c2char.AttackEffectRange1Grid,
		MagicName:        convertTo16Bytes(MagicNames[1]),
	},
}
