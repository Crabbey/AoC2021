package common

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
)

var _ = spew.Dump


type Coords struct {
	Row int
	Col int
}

func (c *Coords) Print() (string) {
	return fmt.Sprintf("%v, %v", c.Col, c.Row)
}

func (c *Coords) GetCoordsInDir(dir string, distance int) *Coords {
	newCoords := Coords{
		Row: c.Row,
		Col: c.Col,
	}
	switch dir {
		case "left":
			newCoords.Translate(&Coords{Row:0, Col: -1}, distance)
		case "right":
			newCoords.Translate(&Coords{Row:0, Col: 1}, distance)
		case "up":
			newCoords.Translate(&Coords{Row:-1, Col: 0}, distance)
		case "down":
			newCoords.Translate(&Coords{Row:1, Col: 0}, distance)
		case "upleft":
			newCoords.Translate(&Coords{Row:-1, Col: -1}, distance)
		case "upright":
			newCoords.Translate(&Coords{Row:-1, Col: 1}, distance)
		case "downleft":
			newCoords.Translate(&Coords{Row:1, Col: -1}, distance)
		case "downright":
			newCoords.Translate(&Coords{Row:1, Col: 1}, distance)
	}
	return &newCoords
}

func (c *Coords) Translate(vector *Coords, distance int) {
	c.Row += vector.Row * distance
	c.Col += vector.Col * distance
}

func (c *Coords) Diff(second *Coords) *Coords {
	ret := &Coords{
		Row: c.Row - second.Row,
		Col: c.Col - second.Col,
	}
	return ret
}

func (c *Coords) UniqueReference() string {
	return fmt.Sprintf("%v|%v", c.Row, c.Col)
}