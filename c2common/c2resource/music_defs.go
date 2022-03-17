package c2resource

import (
	"github.com/BismarckDD/Caocao/c2common/c2enum"
)

// MusicDef stores the music definitions of a region
type MusicDef struct {
	Region    c2enum.RegionIdType
	InTown    bool
	MusicFile string
}

func getMusicDefs() []MusicDef {
	return []MusicDef{
		{c2enum.Battle1, false, BGMScenario1},
		{c2enum.RegionWilderness, false, BGMScenario2},
		{c2enum.RegionCave, false, BGMCaves},
		{c2enum.RegionCrypt, false, BGMCrypt},
		{c2enum.RegionMonestary, false, BGMMonastery},
		{c2enum.RegionCourtyard, false, BGMMonastery},
		{c2enum.RegionBarracks, false, BGMMonastery},
		{c2enum.RegionJail, false, BGMMonastery},
		{c2enum.RegionCathedral, false, BGMMonastery},
		{c2enum.RegionCatacombs, false, BGMMonastery},
		{c2enum.RegionTristram, false, BGMTristram},
		{c2enum.RegionTown, false, BGMTown2},
		{c2enum.RegionSewer, false, BGMSewer},
		{c2enum.RegionHarem, false, BGMHarem},
		{c2enum.RegionBasement, false, BGMHarem},
		{c2enum.RegionDesert, false, BGMDesert},
		{c2enum.RegionTomb, false, BGMTombs},
		{c2enum.RegionLair, false, BGMLair},
		{c2enum.RegionArcane, false, BGMSanctuary},
		{c2enum.RegionTown, false, BGMTown3},
		{c2enum.RegionJungle, false, BGMJungle},
		{c2enum.RegionKurast, false, BGMKurast},
		{c2enum.RegionSpider, false, BGMSpider},
		{c2enum.RegionDungeon, false, BGMKurastSewer},
		{c2enum.RegionSewer, false, BGMKurastSewer},
		{c2enum.RegionTown, false, BGMTown4},
		{c2enum.RegionMesa, false, BGMMesa},
		{c2enum.RegionLava, false, BGMMesa},
		{c2enum.RegonTown, false, BGMXTown},
	}
}

// GetMusicDef returns the MusicDef of the given region
func GetMusicDef(regionType c2enum.RegionIdType) *MusicDef {
	musicDefs := getMusicDefs()
	for idx := range musicDefs {
		if musicDefs[idx].Region != regionType {
			continue
		}

		return &musicDefs[idx]
	}

	return &musicDefs[0]
}
