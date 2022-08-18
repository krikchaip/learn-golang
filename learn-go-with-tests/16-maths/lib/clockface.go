package clockface

import (
	"fmt"
	"math"
	"strings"
	"time"
)

const (
	OriginX          float64 = 150
	OriginY          float64 = 150
	SecondHandLength float64 = 90
	MinuteHandLength float64 = 80
	HourHandLength   float64 = 40
)

const tag = `
<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
		width="100%%"
		height="100%%"
		viewBox="0 0 300 300"
		version="2.0">

	<!-- bezel -->
	<circle cx="%[1]f" cy="%[2]f" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>

	<!-- hour hand -->
	<!-- <line x1="%[1]f" y1="%[2]f" x2="" y2=""
				style="fill:none;stroke:#000;stroke-width:7px;"/> -->

	<!-- minute hand -->
	<!-- <line x1="%[1]f" y1="%[2]f" x2="" y2=""
				style="fill:none;stroke:#000;stroke-width:7px;"/> -->

	<!-- second hand -->
	<line x1="%[1]f" y1="%[2]f" x2="%[3]f" y2="%[4]f"
				style="fill:none;stroke:#f00;stroke-width:3px;"/>
</svg>
`

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

func ClockSVG(t time.Time) string {
	secondHand := SecondHand(t)
	return fmt.Sprintf(
		strings.TrimLeft(tag, "\n "),
		OriginX, OriginY,
		secondHand.X, secondHand.Y,
	)
}
