package clockface

import (
	"math"
	"time"
)

// A Point represents a two dimensional Cartesian coordinate
type Point struct {
	X, Y float64
}

// Shift a unit vector p by Origin constants and a specified Length l
func (p Point) ShiftLength(l float64) Point {
	p.X = OriginX + l*p.X
	p.Y = OriginY + l*p.Y
	return p
}

func (a Point) RoughlyEqual(b Point) bool {
	const threshold = 1e-7
	eq := func(a, b float64) bool {
		return math.Abs(a-b) < threshold
	}
	return eq(a.X, b.X) && eq(a.Y, b.Y)
}

const (
	OriginX          float64 = 150
	OriginY          float64 = 150
	SecondHandLength float64 = 90
	MinuteHandLength float64 = 80
	HourHandLength   float64 = 40
)

// SecondHand is the unit vector of the second hand of an analogue clock at time `t`
// represented as a Point.
func SecondHand(t time.Time) Point {
	radian := secondsInRadian(t.Second())
	unitY, unitX := math.Sincos(radian)
	return Point{unitX, unitY}.ShiftLength(SecondHandLength)
}

func secondsInRadian(second int) float64 {
	s := float64(second)
	return 2*math.Pi*(s/60) - math.Pi/2
}
