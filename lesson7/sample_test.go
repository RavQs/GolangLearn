package main

import "testing"

func TestMultiple(t *testing.T) { //в терминал можно передавать параметры определнных тестов с именем
	//go test -run TestMultiple/simple или 'simple'
	t.Run("simple", func(t *testing.T) {
		t.Parallel()
		t.Log("simple")
		var x, y, expResult = 2, 2, 4

		testResult := Multiple(x, y)

		if expResult != testResult {
			t.Errorf("%d expected to be %d", testResult, expResult)
		}

		t.Run("more simple", func(t *testing.T) { //Иерархия тестов
			var x, y, expResult = 1, 1, 1

			testResult := Multiple(x, y)

			if expResult != testResult {
				t.Errorf("%d expected to be %d", testResult, expResult)
			}
		})
	})

	t.Run("medium", func(t *testing.T) {
		t.Parallel()
		t.Log("medium")
		var x, y, expResult = 222, 222, 49284

		testResult := Multiple(x, y)

		if expResult != testResult {
			t.Errorf("%d expected to be %d", testResult, expResult)
		}
	})

	t.Run("big", func(t *testing.T) {
		t.Parallel()
		t.Log("big")
		var x, y, expResult = 140000, 545, 76300000

		testResult := Multiple(x, y)

		if expResult != testResult {
			t.Errorf("%d expected to be %d", testResult, expResult)
		}
	})

	t.Run("negative", func(t *testing.T) {
		t.Parallel()
		t.Log("negative")
		var x, y, expResult = -2, -4, 8

		testResult := Multiple(x, y)

		if expResult != testResult {
			t.Errorf("%d expected to be %d", testResult, expResult)
		}
	})
}
