package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestExampleCleanup(t *testing.T) {
	x := 0
	Convey("A", t, func() {
		x++
		So(1, ShouldEqual, 1)
		So(4, ShouldEqual, 4)
		Convey("A-B", func() {
			x++
			So("bb", ShouldEqual, "bb")
			Convey("A-B-C1", func() {
				x++
				So(x, ShouldEqual, 3) //После этого convey начинается с А секции, а var x остается неизменной
			})
			Convey("A-B-C2", func() {
				x++
				So(x, ShouldEqual, 6)
			})
		})
		Reset(func() {
			t.Log("finish")
		})
	})

}
