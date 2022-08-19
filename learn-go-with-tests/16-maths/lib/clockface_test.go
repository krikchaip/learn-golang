package clockface_test

import (
	clockface "16-maths/lib"
	"encoding/xml"
	"math"
	"testing"
	"time"
)

func TestSecondHand(t *testing.T) {
	// time -> unit vector
	cases := []struct {
		time  time.Time
		point clockface.Point
	}{
		{clock(0, 0, 0), degreeToPoint(0)},
		{clock(0, 0, 30), degreeToPoint(180)},
		{clock(0, 0, 45), degreeToPoint(270)},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			_, _, got := clockface.MakeHands(c.time)
			want := c.point.ShiftLength(clockface.SecondHandLength)

			if !got.RoughlyEqual(want) {
				t.Errorf("Got %v, wanted %v", got, want)
			}
		})
	}
}

func TestMinuteHand(t *testing.T) {
	// time -> unit vector
	cases := []struct {
		time  time.Time
		point clockface.Point
	}{
		{clock(0, 0, 0), degreeToPoint(0)},
		{clock(0, 30, 0), degreeToPoint(180)},
		{clock(0, 45, 0), degreeToPoint(270)},
		{clock(0, 0, 30), degreeToPoint(3)},
		{clock(0, 45, 15), degreeToPoint(271.5)},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			_, got, _ := clockface.MakeHands(c.time)
			want := c.point.ShiftLength(clockface.MinuteHandLength)

			if !got.RoughlyEqual(want) {
				t.Errorf("Got %v, wanted %v", got, want)
			}
		})
	}
}

func TestHourHand(t *testing.T) {
	// time -> unit vector
	cases := []struct {
		time  time.Time
		point clockface.Point
	}{
		{clock(0, 0, 0), degreeToPoint(0)},
		{clock(0, 30, 0), degreeToPoint(15)},
		{clock(12, 30, 0), degreeToPoint(15)},
		{clock(0, 45, 0), degreeToPoint(22.5)},
		{clock(0, 0, 30), degreeToPoint(.25)},
		{clock(0, 45, 15), degreeToPoint(22.625)},
		{clock(21, 0, 0), degreeToPoint(270)},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got, _, _ := clockface.MakeHands(c.time)
			want := c.point.ShiftLength(clockface.HourHandLength)

			if !got.RoughlyEqual(want) {
				t.Errorf("Got %v, wanted %v", got, want)
			}
		})
	}
}

func TestClockSVG(t *testing.T) {
	cases := []struct {
		time       time.Time
		secondHand Line
		minuteHand Line
		hourHand   Line
	}{
		{
			clock(0, 0, 0),
			degreeToLine(0, clockface.SecondHandLength),
			degreeToLine(0, clockface.MinuteHandLength),
			degreeToLine(0, clockface.HourHandLength),
		},
		{
			clock(0, 0, 30),
			degreeToLine(180, clockface.SecondHandLength),
			degreeToLine(3, clockface.MinuteHandLength),
			degreeToLine(.25, clockface.HourHandLength),
		},
		{
			clock(0, 45, 15),
			degreeToLine(90, clockface.SecondHandLength),
			degreeToLine(271.5, clockface.MinuteHandLength),
			degreeToLine(22.625, clockface.HourHandLength),
		},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			bytes := []byte(clockface.ClockSVG(c.time))
			svg := &SVG{}

			// ?? svg = xml.Parse(bytes)
			xml.Unmarshal(bytes, svg)

			if !containsLine(c.secondHand, svg.Line) {
				t.Errorf(
					"Expected to find the second hand line %+v, in the SVG lines %+v",
					c.secondHand,
					svg.Line,
				)
			}

			if !containsLine(c.minuteHand, svg.Line) {
				t.Errorf(
					"Expected to find the minute hand line %+v, in the SVG lines %+v",
					c.minuteHand,
					svg.Line,
				)
			}

			if !containsLine(c.hourHand, svg.Line) {
				t.Errorf(
					"Expected to find the hour hand line %+v, in the SVG lines %+v",
					c.hourHand,
					svg.Line,
				)
			}
		})
	}
}

// ?? generated from https://github.com/miku/zek
type SVG struct {
	// XMLName xml.Name `xml:"svg"`
	// Text    string   `xml:",chardata"`
	// Xmlns   string   `xml:"xmlns,attr"`
	// Width   string   `xml:"width,attr"`
	// Height  string   `xml:"height,attr"`
	ViewBox string `xml:"viewBox,attr"`
	// Version string   `xml:"version,attr"`
	Circle Circle `xml:"circle"`
	Line   []Line `xml:"line"`
}

type Circle struct {
	// Text  string `xml:",chardata"`
	Cx float64 `xml:"cx,attr"`
	Cy float64 `xml:"cy,attr"`
	R  float64 `xml:"r,attr"`
	// Style string `xml:"style,attr"`
}

type Line struct {
	// Text  string `xml:",chardata"`
	X1 float64 `xml:"x1,attr"`
	Y1 float64 `xml:"y1,attr"`
	X2 float64 `xml:"x2,attr"`
	Y2 float64 `xml:"y2,attr"`
	// Style string `xml:"style,attr"`
}

func clock(hour, minute, second int) time.Time {
	return time.Date(1337, time.January, 1, hour, minute, second, 0, time.UTC)
}

func testName(t time.Time) string {
	return t.Format("15:04:05")
}

func roundFloat(val float64) float64 {
	ratio := math.Pow(10, float64(clockface.Precision))
	return math.Round(val*ratio) / ratio
}

func containsLine(l Line, ls []Line) bool {
	for _, line := range ls {
		l = Line{roundFloat(l.X1), roundFloat(l.Y1), roundFloat(l.X2), roundFloat(l.Y2)}
		if l == line {
			return true
		}
	}
	return false
}

func degreeToPoint(degree float64) clockface.Point {
	radian := (math.Pi/180)*math.Mod(degree, 360) - math.Pi/2
	Y, X := math.Sincos(radian)
	return clockface.Point{X, Y}
}

func degreeToLine(degree, length float64) Line {
	p := degreeToPoint(degree).ShiftLength(length)
	return Line{
		clockface.OriginX,
		clockface.OriginY,
		p.X,
		p.Y,
	}
}
