package c2status

type UnitStatus uint16

const (
	UnitStatusToxic         UnitStatus = 0x1
	UnitStatusChaos         UnitStatus = 0x2
	UnitStatusHold          UnitStatus = 0x4
	UnitStatusForbidden     UnitStatus = 0x8
	UnitStatusAttackUp      UnitStatus = 0x10
	UnitStatusSpiritUp      UnitStatus = 0x20
	UnitStatusDefenceUp     UnitStatus = 0x40
	UnitStatusExplosiveUp   UnitStatus = 0x80
	UnitStatusMoraleUp      UnitStatus = 0x100
	UnitStatusAttackDown    UnitStatus = 0x200
	UnitStatusSpiritDown    UnitStatus = 0x400
	UnitStatusDefenceDown   UnitStatus = 0x800
	UnitStatusExplosiveDown UnitStatus = 0x100
	UnitStatusMoraleDown    UnitStatus = 0x200
	UnitStatusMoveUp        UnitStatus = 0x400
	UnitStatusMoveDown      UnitStatus = 0x800

	UnitStatusAllException UnitStatus = 0xf
	UnitStatusAllUp        UnitStatus = 0x1f0
	UnitStatusAllDown      UnitStatus = 0x360
)

type StrikeAssist uint8

const (
	StrikeAssistForbiddenCS     StrikeAssist = 0x1  // forbidden counter-strike
	StrikeAssistEnableCCS       StrikeAssist = 0x2  // enable counter-counter-strike
	StrikeAssistEnableLS        StrikeAssist = 0x4  // enable leading-strike
	StrikeAssistEnablePS        StrikeAssist = 0x8  // enable pierce-strike
	StrikeAssistWithToxic       StrikeAssist = 0x10 // attack with toxic
	StrikeAssistWithChaos       StrikeAssist = 0x20 // attack with chaos
	StrikeAssistWithHold        StrikeAssist = 0x40 // attack with hold
	StrikeAssistWithProhibition StrikeAssist = 0x80 // attack with prohibition
)

type SpellAssist uint8

const (
	SpellAssistLandEnhance  SpellAssist = 0x1  // enable wind magic damage
	SpellAssistWaterEnhance SpellAssist = 0x2  // enable wind magic damage
	SpellAssistFireEnhance  SpellAssist = 0x4  // enable wind magic damage
	SpellAssistWindEnhance  SpellAssist = 0x8  // enable wind magic damage
	SpellAssistMPReduce     SpellAssist = 0x10 // save mp when spell
)
