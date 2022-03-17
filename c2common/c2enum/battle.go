package c2enum

// WavIdIdType represents a region ID
type CampaignId uint16 // many changed needed when changing to ID

// WavIds
const (
	CampaignNone CampaignId = iota
	// First Chapter. 霸王诞生
	CampaignYingchuan        // 颍川之战
	CampaignSishuiGate       // 泗水关之战
	CampaignHulaoGate        // 虎牢关之战
	CampaignPursueDongZhuo   // 董卓追击战
	CampaignCrusadeHuangJin  // 青州黄巾军讨伐战
	CampaignCrusadeXuZhou    // 徐州讨伐战
	CampaignPuyang1          // 濮阳之战一
	CampaignPuyang2          // 濮阳之战二
	CampaignPuyang3          // 濮阳之战三
	CampaignSaveEmperor      // 献帝救出战
	CampaignCrusadeZhangXiu1 // 张绣讨伐战
	CampaignCrusadeYuanShu   // 袁术讨伐战
	CampaignCrusadeZhangXiu2 // 张绣讨伐战二
	CampaignXiapi            // 下坯之战
	CampaignSurroundLvBu     // 吕布包围战

	// Second Chapter. 河北之争
	CampaignInvadeXuZhou // 徐州入侵战
	CampaignBaima        // 白马之战
	CampaignYanjin       // 延津之战
	CampaignGuandu       // 官渡之战
	CampaignCangting     // 仓亭之战
	CampaignRangshan     // 穰山之战
	CampaignLiyang       // 黎阳之战
	CampaignYeCity       // 邺城之战
	CampaignNanpi        // 南皮攻略战
	CampaignLiuCity      // 柳城平定战

	// Third Chapter. 三足鼎立
	CampaignBowangSlope   // 博望坡之战
	CampaignChangbanSlope // 长坂坡之战
	CampaignChibi         // 赤壁之战
	CampaignHuarongWay    // 华容道之战
	CampaignHefei         // 合肥之战
	CampaignCounterMachao // 马超迎击战

	// 奸臣线
	CampaignCounterXiaoyaojin // 逍遥津之战
	CampaignCounterRuxukouRed // 濡须口之战(红)

	// Fourth Chapter. 天下统一
	CampaignCounterDingjunMoutainRed // 定军山之战(红)
	CampaignCounterHanshui           // 汉水之战
	CampaignCounterXiegu             // 斜谷之战
	CampaignCounterFanCity           // 樊城救援战
	CampaignCounterYangpingGateRed   // 阳平关之战(红)
	CampaignCounterJiangeRed         // 剑阁之战(红)
	CampaignCounterInvadeChendu      // 成都入侵战
	CampaignCounterBaidiCityRed      // 白帝城之战(红)
	CampaignChibi2                   // 赤壁之战二
	CampaignCounterRuxukou2          // 濡须口之战二
	CampaignInvadeJianYeRed          // 建业入侵战(红)

	// 忠臣线
	CampaignCounterDingjunMoutainBlue // 定军山之战(蓝)
	CampaignCounterSaveMaiCity        // 麦城救援战
	CampaignJiangling                 // 江陵之战
	CampaignRuxukouBlue               // 濡须口之战(蓝)
	CampaignInvadeJianYeBlue          // 建业入侵战(蓝)

	// Fourth Chapter. 魔王复活
	CampaignCounterYufupu           // 鱼腹浦之战
	CampaignCounterBaidiCityBlue    // 白帝城之战(蓝)
	CampaignCounterJiangeBlue       // 剑阁之战(蓝)
	CampaignCounterYangpingGateBlue // 阳平关之战(蓝)
	CampaignCounterWuzhangyuan1     // 五丈原之战一
	CampaignCounterWuzhangyuan2     // 五丈原之战二

)
