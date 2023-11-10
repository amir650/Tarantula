package engine

import (
	"Tarantula/engine"
	"testing"
)

func TestKnightCreation(t *testing.T) {

	p1 := *engine.NewKnight(engine.WHITE, 30, false)
	p2 := *engine.NewKnight(engine.WHITE, 30, false)

	if !p1.Equals(p2) {
		t.Errorf("Should refer to the same thing!")
	}

}

func TestBishopCreation(t *testing.T) {

	p1 := engine.NewBishop(engine.WHITE, 30, false)
	p2 := engine.NewBishop(engine.WHITE, 30, false)

	if p1 != p2 {
		t.Errorf("Should refer to the same thing!")
	}

}

func TestRookCreation(t *testing.T) {

	p1 := engine.NewRook(engine.WHITE, 30, false)
	p2 := engine.NewRook(engine.WHITE, 30, false)

	if p1 != p2 {
		t.Errorf("Should refer to the same thing!")
	}

}

func TestQueenCreation(t *testing.T) {

	p1 := engine.NewQueen(engine.WHITE, 30, false)
	p2 := engine.NewQueen(engine.WHITE, 30, false)

	if p1 != p2 {
		t.Errorf("Should refer to the same thing!")
	}

}
