package clockface

import (
	"math"
	"time"
)

// A Point represents a two dimensional Cartesian coordinate
type Point struct {
	X, Y float64
}

const (
	OriginX          float64 = 150
	OriginY          float64 = 150
	SecondHandLength float64 = 90
)

// SecondHand is the unit vector of the second hand of an analogue clock at time `t`
// represented as a Point.
func SecondHand(t time.Time) Point {
	radian := secondsInRadian(t.Second())
	unitY, unitX := math.Sincos(radian)

	return Point{
		OriginX + SecondHandLength*unitX,
		OriginY + SecondHandLength*unitY,
	}
}

func secondsInRadian(second int) float64 {
	s := float64(second)
	return 2*math.Pi*(s/60) - math.Pi/2
}
