package main

import (
	"errors"
	"fmt"
)

type Value struct {
	value            int
	powerOfTen       int
	physicalQuantity rune
}

// Add adds two Values together and returns the result
// It returns an error if the physical quantity of the two summands are not the same. The calculation will still be correct.
func Add(a, b Value) (Value, error) {
	v := Value{}

	if a.powerOfTen > b.powerOfTen {
		v.powerOfTen = b.powerOfTen
		v.value = a.value*(a.value-b.value) + b.value
	} else {
		v.powerOfTen = a.powerOfTen
		v.value = a.value + b.value*(b.value-a.value)
	}

	if a.physicalQuantity != b.physicalQuantity {
		v.physicalQuantity = '#'
		return v, errors.New("two added values have to have the same physical quantity")
	}
	v.physicalQuantity = a.physicalQuantity
	return v, nil
}

// Subtract subtracts two Values and returns the result.
// It returns an error if the physical quantity of the minuend and the subtrahend isn't the same. The calculation will
// still be correct.
func Subtract(a, b Value) (Value, error) {
	v := Value{}

	if a.powerOfTen > b.powerOfTen {
		v.powerOfTen = b.powerOfTen
		v.value = a.value*(a.value-b.value) - b.value
	} else {
		v.powerOfTen = a.powerOfTen
		v.value = a.value - b.value*(b.value-a.value)
	}

	if a.physicalQuantity != b.physicalQuantity {
		v.physicalQuantity = '#'
		return v, errors.New("two subtracted values have to have the same physical quantity")
	}
	v.physicalQuantity = a.physicalQuantity
	return v, nil
}

type Point struct {
	x, y, z Value
}

func (p Point) Equals(p2 Point) bool {
	return p.x == p2.x && p.y == p2.y && p.z == p2.z
}

type Beam struct {
	start, end Point
	material   Material
	section    Section
}

type Material struct {
	name           string
	group          string
	elasticModulus Value
	strength       Value
}

type Section struct {
	area   Value
	Iy, Iz Value
	h, b   Value
	ys, zs Value
	geometry
}

// triangulation: ear clipping
type Sketch2D struct {
	sketchElements []Segment2D
}

func (s Sketch2D) IsClosed() bool {

	var start = s.sketchElements[0].getStart()
	var currentEnd = s.sketchElements[0].getEnd()
	s.sketchElements = s.sketchElements[1:]

	for len(s.sketchElements) > 0 {

	}

	return true
}

// SegmentContainsPoint checks if any of the provided segments starts or ends at the specified point and returns a slice
// without the segment, which contains the point, and the other point of this segment.
// It returns an error if none of the segments contains the point.
func SegmentContainsPoint(elements []Segment2D, point Point) ([]Segment2D, Point, error) {
	for i := 0; i < len(elements); i++ {
		if point.Equals(elements[i].getEnd()) {
			return append(elements[:i], elements[i+1:]...), elements[i].getStart(), nil
		} else if point.Equals(elements[i].getStart()) {
			return append(elements[:i], elements[i+1:]...), elements[i].getEnd(), nil
		}
	}
	return elements, Point{}, errors.New("none of the segments contained in this point")
}

type Segment2D interface {
	getStart() Point
	getEnd() Point
}

type Line struct {
	start, end Point
}

func (l Line) getStart() Point {
	return l.start
}

func (l Line) getEnd() Point {
	return l.end
}
