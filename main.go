package main

import (
	"fmt"
	"github.com/eccyan/mahjong/lib"
	"github.com/fwhappy/mahjong/wall"
)

func main() {
	tiles := lib.NewTiles()
	w := wall.NewWall()
	w.SetTiles(tiles.Ids())
	w.Shuffle()
	w.ForwardDraw()

	for _, id := range w.GetTiles() {
		tile, contains := tiles.ById(id)
		if contains {
			tile.Painter.Draw()
		} else {
			tiles.Back().Painter.Draw()
		}
	}
	fmt.Println("")
}
