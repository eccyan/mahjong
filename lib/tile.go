package lib

// Tile contains information for displaying
type Tile struct {
	Painter Painter
}

type Tiles struct {
	withId map[int]Tile
	ids    []int
	back   Tile
}

func NewTiles() Tiles {
	tiles := Tiles{withId: newTilesWithId(), back: newTile('🀫')}
	tiles.ids = getAllIds(tiles)
	return tiles
}

func (tiles Tiles) ById(id int) (Tile, bool) {
	tile, contains := tiles.withId[id]
	return tile, contains
}

func (tiles Tiles) Ids() []int {
	return tiles.ids
}

func (tiles Tiles) Back() Tile {
	return tiles.back
}

func newTile(character rune) Tile {
	return Tile{&CharacterPainter{character}}
}

// NewTiles creates tiles for mahjong
//
// 0       : Back
// 1  - 9  : 1 - 9 Characters
// 11 - 19 : 1 - 9 Bamboos
// 21 - 29 : 1 - 9 Dots
// 31 - 34 : South, West, East and North Winds
// 41 - 43 : Green, Red and White dragons
// 51 ~ 54 : 1 Spring, 2 Summer, 3 Autumn and 4 Winter Seasons
// 61 ~ 64 : 1 Plum, 2 Orchid, 3 Mum, 4 Bamboo Flowers
// 71      : Joker
func newTilesWithId() map[int]Tile {
	return map[int]Tile{
		1: newTile('🀇'), 2: newTile('🀈'), 3: newTile('🀉'), 4: newTile('🀊'), 5: newTile('🀋'),
		6: newTile('🀌'), 7: newTile('🀍'), 8: newTile('🀎'), 9: newTile('🀏'),
		11: newTile('🀐'), 12: newTile('🀑'), 13: newTile('🀒'), 14: newTile('🀓'), 15: newTile('🀔'),
		16: newTile('🀕'), 17: newTile('🀖'), 18: newTile('🀗'), 19: newTile('🀘'),
		21: newTile('🀙'), 22: newTile('🀚'), 23: newTile('🀛'), 24: newTile('🀜'), 25: newTile('🀝'),
		26: newTile('🀞'), 27: newTile('🀟'), 28: newTile('🀠'), 29: newTile('🀡'),
		31: newTile('🀀'), 32: newTile('🀁'), 33: newTile('🀂'), 34: newTile('🀃'),
		41: newTile('🀅'), 42: newTile('🀄'), 43: newTile('🀆'),
		51: newTile('🀦'), 52: newTile('🀧'), 53: newTile('🀨'), 54: newTile('🀩'),
		61: newTile('🀢'), 62: newTile('🀣'), 63: newTile('🀤'), 64: newTile('🀥'),
		71: newTile('🀪'),
	}
}

func getAllIds(tiles Tiles) []int {
	ids := make([]int, len(tiles.withId))

	i := 0
	for id := range ids {
		ids[i] = id
		i++
	}
	return append(append(append(ids, ids...), ids...), ids...)
}
