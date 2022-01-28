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
)

type StrikeStatus uint8

const (
	StrikeStatusForbiddenCS StrikeStatus = 0x1 // forbidden counter-strike
	StrikeStatusEnableCCS   StrikeStatus = 0x2 // enable counter-counter-strike
	StrikeStatusEnableLS    StrikeStatus = 0x4 // enable leading-strike
	StrikeStatusEnablePS    StrikeStatus = 0x4 // enable pierce-strike
)
