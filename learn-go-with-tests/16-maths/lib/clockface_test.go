package clockface_test

import (
	clockface "16-maths/lib"
	"encoding/xml"
	"testing"
	"time"
)

func TestSecondHand(t *testing.T) {
	// time -> unit vector
	cases := []struct {
		time  time.Time
		point clockface.Point
	}{
		{clock(0, 0, 0), clockface.Point{0, -1}},
		{clock(0, 0, 30), clockface.Point{0, 1}},
		{clock(0, 0, 45), clockface.Point{-1, 0}},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := clockface.SecondHand(c.time)
			want := c.point.ShiftLength(clockface.SecondHandLength)

			if !got.RoughlyEqual(want) {
				t.Errorf("Got %v, wanted %v", got, want)
			}
		})
	}
}

func TestClockSVG(t *testing.T) {
	tm := clock(0, 0, 0)

	bytes := []byte(clockface.ClockSVG(tm))
	svg := &SVG{}

	// ?? svg = xml.Parse(bytes)
	xml.Unmarshal(bytes, svg)

	want := Line{150, 150, 150, 60}

	for _, line := range svg.Line {
		if line == want {
			return
		}
	}

	t.Errorf(
		"Expected to find the second hand line %+v, in the SVG lines %+v",
		want,
		svg.Line,
	)
}

// ?? generated by https://github.com/miku/zek
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
