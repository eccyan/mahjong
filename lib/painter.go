package lib

import (
	"fmt"
)

// A Painter can draw anythings to display
type Painter interface {
	// Draw is drawing for character or graphics
	Draw()
}

// CharacterPainter can draw a character
type CharacterPainter struct {
	character rune
}

// NewCharacterPainter creates *Painter from character
func NewCharacterPainter(character rune) Painter {
	return &CharacterPainter{character}
}

func (characterPainter CharacterPainter) Character() rune {
	return characterPainter.character
}

// Draw draw
func (painter CharacterPainter) Draw() {
	fmt.Print(string(painter.Character()))
}
