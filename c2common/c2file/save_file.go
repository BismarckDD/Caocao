package c2file

type SaveFileHeader struct {
	Magic           [4]byte
	MainString      string
	ViceString      string
	HeaderSize      uint32
	ScenarioId      uint32
	CharTableOffset uint32
	CharTableNum    uint32
	BoolFlagOffset  uint32
	BoolFlagLength  uint32
	IntFlagOffset   uint32
	IntFlagLength   uint32
}

type Char
