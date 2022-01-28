package c2magic

import "github.com/BismarckDD/Caocao/c2core/c2char"

type MagicTargetEnum uint8

const (
	MagicTargetWeAndFriends MagicTargetEnum = 0x03
	MagicTargetEnemies      MagicTargetEnum = 0x04
	MagicTargetAll          MagicTargetEnum = 0x07
)

/* total size is 64 */
// 施法距离: 一般为3个距离
// 策略影响范围: 1,5,9,line2,line4,all
// 有效目标：敌、我友、全，暂时是这三类
// 消耗 MP, 消耗 HP(施法者, 固定值)
// 恢复 HP, 恢复 MP(目标，HP为系数，MP为固定值)
// 减少 HP, 减少 MP(目标，HP为系数，MP为固定值)
// 命中系数H(uint8)，H越大，策略命中衰减越小。H可以大于1

// 策略对HP影响
// 恢复 K = 0.1/0.5, B=30/20 => K,B 压缩为一个uint8
// 40 + A.Spirit *0.1 = 0.1(A.Spirit + 100) + 30
// 70 + A.Spirit *0.5 = 0.5(A.Spirit + 100) + 20
// HP = K / 10 * (A.Spirit + 100) + B
// 减少HP的影响 K,B都可以使用uint8表示 K,B=uint8/10 max=25.5
// (A.Lv + spirit.Diff / 3) * K + B
// HaiXiao, XuanYun: K=25.5, B=25.5
// JuYan K=1.2, B=30
// ZhuQue K=1, B=25
// LieHuo,BaoYan,JiLiu,ShanLan,DiLong K=0.9, B=22.5
// ZhuoRe,HuoLong,LongJuan,ZhuoLiu,ShuiLong,LuoShi,DiZhen,DuYan,YouHuo,QingLong: K=0.7, B=17.5
// HuoZhen,XuanFeng,FengLong,ShuiZhen,DingShen: K=0.5, B=12.5
// FengZhen,DuWu: K=0,4, B=9.6
// DingJun,ShaBao: K=0.2, B=5
// DieBao && YouHuo 这两个计策等于说有两个目标，recover为自身，reduce为目标
// 判断条件: 1. MPReduce!=0 && MPRecover!=0; 2. HPReduce!=0 && HPRecover!=0.
type Magic struct {
	MagicId          uint8                    // MagicId is corresponding to its AnimationId
	MagicTarget      MagicTargetEnum          // we && friends, enemies, all
	CallAnimationId  uint8                    // == 0, not exist
	EffAnimationId   uint8                    // must exist
	MagicSpellRange  c2char.AttackRange       // uint8
	MagicEffectRange c2char.AttackEffectRange // uint8
	HPCost           uint8
	MPCost           uint8
	HPRecoverK       uint8 // A.Spirit.Diff * K / 100 + B
	HPRecoverB       uint8 // use as uint8
	MPRecoverK       uint8 // (A.Lv + (A.Spirit.Diff - B.Spirit) / 3) * K + B
	MPRecoverB       uint8 // use as float8 (0~25.5)
	HPReduceK        uint8 // (A.Lv + (A.Spirit.Diff - B.Spirit) / 3) * K + B
	HPReduceB        uint8 // use as float8 (0~25.5)
	MPReduceK        uint8 // (A.Lv + (A.Spirit.Diff - B.Spirit) / 3) * K + B
	MPReduceB        uint8 // use as float8 (0~25.5)
	MagicCate        uint8 // 0x1:Land 0x2:Water 0x4:Fire 0x8:Wind 0x10:Assist 0x20:Saint
	// Cate is used for magic reinforcement
	// ...17bytes above
	MagicName [15]byte // 15 bytes
} // 2022.02.17 => 32个uint8

func ConvertTo16Bytes(str string) [16]byte {
	res := [16]byte{}
	temp := []byte(str)
	for i := 0; i < len(temp); i++ {
		res[i] = temp[i]
	}
	return res
}

func LoadMagicInfo() {

}

func SaveMagicInfo() {

}

func (magic *Magic) Spell(attacker, defender c2char.Unit) bool {

	if attacker.HP < magic.HPCost+attacker.MaxHP*0.2 || attacker.MP < magic.MPCost {

		return false
	}
	attacker.HP = attacker.HP - magic.HPCost
	attacker.MP = attacker.MP - magic.MPCost
	if (magic.HPReduceB != 0 || magic.HPReduceK != 0) && (magic.HPRecoverB != 0 || magic.HPRecoverK != 0) {
		// DieBao or YouHuo
		attacker.HP - magic.HPCost
		attacker.HP
		return true
	}
	return true
}
