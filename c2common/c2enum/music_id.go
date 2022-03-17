package c2enum

// WavIdIdType represents a region ID
type WavIdType uint16 // many changed needed when changing to ID

// WavIds
const (
	WavIdNone WavIdType = iota
	WavSee00            // Bird song
	WavSee01            // Chan song
	WavSee02            // Bird song2
	WavSee03            // Wind song
	WavSee04            // Rain and thunder
	WavSee05            // Bird song3
	WavSee06            // children and house-cart
	WavSee07            // house song
	WavSee08            // huashui
	WavSem00            // magic-fire
	WavSem01            // magic-fire-long
	WavSem02            // magic-water
	WavSem03            // magic-water-long
	WavSem04            // magic-wind
	WavSem05            // magic-wind-long
	WavSem06            // magic-land
	WavSem07            // magic-land-long
	WavSem09            // magic-fire
	WavSem10            // magic-fire-long
	WavSem11            // magic-recover
	WavSem12            // magic-update
	WavSem13            // magic-houjiao
	WavSem14            // magic-
	WavSem15            // magic-land
	WavSem16            // magic-land-long

	WavSe0  // item-left-click
	WavSe1  // item-left-click-block
	WavSe2  // item-left-click-block
	WavSe3  // ???
	WavSe4  // navigate-to-battle
	WavSe5  // Troop
	WavSe6  // walk in battle.
	WavSe7  // shoot in battle.
	WavSe8  // magic-fire
	WavSe9  // buy item
	WavSe10 // sell item
	WavSe11 //
	WavSe12 // up level, weapon-level, armor-level
	WavSe13 // magic-houjiao
	WavSe14 // get item in battle.
	WavSe15 // magic-land
	WavSe16 // magic-land-long
	WavSe34 // Sword, Lance use.
	WavSe35 // Sword, Lance hit.
	WavSe36 // Hit
	WavSe37 // Bow use
	WavSe38 // Bow hit
	WavSe39 // Recover music effect.
	WavSe40 // Lose in Battle
	WavSe41 // Buf music effect.
	WavSe42 // Debuf music effect.

	Track2  // battle 昂扬的气氛
	Track3  // battle 诡异的气氛
	Track4  // battle 紧张的气氛
	Track5  // battle 快活的气氛
	Track6  // battle 紧张的气氛
	Track7  // battle 昂扬的气氛2
	Track8  // battle 紧张的气氛2
	Track9  // battle Win
	Track10 // battle 紧张的气氛3
	Track11 // battle ????
	Track12 // battle 斜谷之战
	Track13 // scenario 阴谋
	Track14 // scenario 阴谋
	Track15 // scenario 阴谋
	Track16 // battle 不知道
	Track17 // battle 定军山
	Track18 // scenario normal
	Track19 // scenario 悲伤的气氛
	Track20 // scenario
	Track21 // end 游戏结束
	Track22 // failure 悲伤的气氛

)
