package main

import (
	"github.com/RavQs/GolangLearn/practices/myPractice"
	"testing"
)

func TestConvert(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		str, expres := "abbbcc", "a1bbb3cc2"

		res := myPractice.Convert(str)

		if expres != res {
			t.Errorf("%s expected to be %s", expres, res)
		}
	})
	t.Run("medium", func(t *testing.T) {
		str, expres := "abbbbbbbbbbbbccccccccccccc",
			"a1bbbbbbbbbbbb12ccccccccccccc13"

		res := myPractice.Convert(str)

		if expres != res {
			t.Errorf("%s expected to be %s", expres, res)
		}
	})
	t.Run("simple", func(t *testing.T) {
		str, expres := "abbbbbbbbbbbbcccccccccccccddddddrrrrrrrrrrrrrrrrrrrffffffffeRRRRRRRRR",
			"a1bbbbbbbbbbbb12ccccccccccccc13dddddd6rrrrrrrrrrrrrrrrrrr19ffffffff8e1RRRRRRRRR9"

		res := myPractice.Convert(str)

		if expres != res {
			t.Errorf("%s expected to be %s", expres, res)
		}
	})
}

func TestFindOutlier(t *testing.T) {
	t.Run("simpleFirst", func(t *testing.T) {
		arrToTest := []int{2, 4, 0, 100, 4, 11, 2602, 36}
		expRes := 11

		res := myPractice.FindOutlier(arrToTest)

		if expRes != res {
			t.Errorf("%d expected to be %d", expRes, res)
		}
	})
	t.Run("simpleSecond", func(t *testing.T) {
		arrToTest := []int{160, 3, 1719, 19, 11, 13, -21}
		expRes := 160

		res := myPractice.FindOutlier(arrToTest)

		if expRes != res {
			t.Errorf("%d expected to be %d", expRes, res)
		}
	})

	t.Run("ownTest", func(t *testing.T) {
		arrToTest := []int{15, 24, 77, 15, -3, 7, 17}
		expRes := 24

		res := myPractice.FindOutlier(arrToTest)

		if expRes != res {
			t.Errorf("%d expected to be %d", expRes, res)
		}
	})
}

func TestDigitalRoot(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		numToTest := 16
		expRes := 7

		res := myPractice.DigitalRoot(numToTest)

		if expRes != res {
			t.Errorf("%d expected to be %d", expRes, res)
		}
	})
	t.Run("medium", func(t *testing.T) {
		numToTest := 942
		expRes := 6

		res := myPractice.DigitalRoot(numToTest)

		if expRes != res {
			t.Errorf("%d expected to be %d", expRes, res)
		}
	})
	t.Run("difficult", func(t *testing.T) {
		numToTest := 493193
		expRes := 2

		res := myPractice.DigitalRoot(numToTest)

		if expRes != res {
			t.Errorf("%d expected to be %d", expRes, res)
		}
	})
}
