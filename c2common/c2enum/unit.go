package c2enum

type UnitId int

const (
	We int = iota // Wo
	Friend
	Enemy
)

const (
	Move int = iota
	attack
	defend
	weak
	gain
)

const (
	Junzhu1       HeroId = iota // QunXiong
	Junzhu2                     // YingXiong
	Junzhu3                     // BaWang
	BuBing1                     // QingBuBing
	BuBing2                     // ZhongBuBing
	BuBing3                     // JinWeiBing
	QiBing1                     // QingQiBing
	QiBing2                     // ZhongQiBing
	QiBing3                     // QinWeiDui
	GongBing1                   // GongBing
	GongBing2                   // NuBing
	GongBing3                   // LianNuBing
	GongQiBing1                 // GongQiBing
	GongQiBing2                 // NuQiBing
	GongQiBing3                 // LianNuQiBing
	DaoZei1                     // DaoZei
	DaoZei2                     // ShanZei
	DaoZei3                     // HaoJie
	WuShuJia1                   // WuShuJia
	WuShuJia2                   // QuanShi
	WuShuJia3                   // QuanSheng
	WuNiang1                    // WuNiang
	WuNiang2                    // WuJi
	WuNiang3                    // WuNv
	PaoChe1                     // QingPaoChe
	PaoChe2                     // ZhongPaoChe
	PaoChe3                     // PiLiChe
	CeShi1                      // CeShi
	CeShi2                      // CanMou
	CeShi3                      // JunShi
	QiMaCeShi1                  // QiMaCeShi
	QiMaCeShi2                  // QiMaCanMou
	QiMaCeShi3                  // QiMaJunShi
	FengShuiShi1                // FengShuiShi
	FengShuiShi2                // FangShuShi
	FengShuiShi3                // XianShuShi
	DaoShi1                     // DaoShi
	DaoShi2                     // HuanShuShi
	DaoShi3                     // YaoShuShi
	XiLiangQiBing               // XiLiangQiBibng
	ZhouShuShi                  // ZhouShuShi
	XianRen                     // XianRen
	DuDu                        // DuDu
	HaiDao                      // HaiDao
	XunHuShi                    // XunHuShi
	XunXiongShi                 // XunXiongShi
	YunShuDui                   // YunShuDui
	NanManBing                  // NanManBing
	BaiXing                     // BaiXing
	HuangDi                     // HuangDi
	UNIT_END                    // END
)
