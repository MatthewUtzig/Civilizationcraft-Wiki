package main

/*
*	Type structure contains all information
*   to be fed into the Structure Information template
*
*
*
 */

type Structure struct {
	ChunkWidth int
	ChunkLength int
	Cost int
	Upkeep int
	Hammer int
	Hours int
	Minutes int
	Points int
	HitPoints int
	TownLimit int
	CivLimit int
	Special []string
	IsCamp bool
}

var Structures = map[string]Structure {
	"camp": {
		ChunkWidth:   2,
		ChunkLength:  2,
		Cost:         -1,
		Upkeep:       -1,
		Hammer:       -1,
		Hours:        0,
		Minutes:      5,
		Points:       1000,
		HitPoints:    500,
		TownLimit:        -1,
		CivLimit: -1,
		Special:      []string {
			"First Building of a Civilization",
		},
		IsCamp:       true,
	},

	"capitol": {
		ChunkWidth:   4,
		ChunkLength:  4,
		Cost:         -1,
		Upkeep:       0,
		Hammer:       1000,
		Hours:        -1,
		Minutes:      -1,
		Points:       6000,
		HitPoints:    -1,
		TownLimit:    -1,
		CivLimit:     1,
		Special:      []string {
			"Primary building of a civilization",
		},
		IsCamp:       false,
	},

	"townhall": {
		ChunkWidth:   2,
		ChunkLength:  2,
		Cost:         -1,
		Upkeep:       0,
		Hammer:       1000,
		Hours:        -1,
		Minutes:      -1,
		Points:       1000,
		HitPoints:    -1,
		TownLimit:    1,
		CivLimit:     -1,
		Special:      []string {
			"Primary building of a town",
		},
		IsCamp:       false,
	},

	"arrowboat": {
		ChunkWidth:   1,
		ChunkLength:  1,
		Cost:         -1,
		Upkeep:       0,
		Hammer:       1000,
		Hours:        -1,
		Minutes:      -1,
		Points:       1000,
		HitPoints:    -1,
		TownLimit:    1,
		CivLimit:     -1,
		Special:      []string {
			"Primary building of a town",
		},
		IsCamp:       false,
	},


}