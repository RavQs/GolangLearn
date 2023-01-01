package main

import "testing"

func TestConvert(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		str, expres := "abbbcc", "a1bbb3cc2"

		res := Convert(str)

		if expres != res {
			t.Errorf("%s expected to be %s", expres, res)
		}
	})
	t.Run("medium", func(t *testing.T) {
		str, expres := "abbbbbbbbbbbbccccccccccccc",
			"a1bbbbbbbbbbbb12ccccccccccccc13"

		res := Convert(str)

		if expres != res {
			t.Errorf("%s expected to be %s", expres, res)
		}
	})
	t.Run("simple", func(t *testing.T) {
		str, expres := "abbbbbbbbbbbbcccccccccccccddddddrrrrrrrrrrrrrrrrrrrffffffffeRRRRRRRRR",
			"a1bbbbbbbbbbbb12ccccccccccccc13dddddd6rrrrrrrrrrrrrrrrrrr19ffffffff8e1RRRRRRRRR9"

		res := Convert(str)

		if expres != res {
			t.Errorf("%s expected to be %s", expres, res)
		}
	})
}
