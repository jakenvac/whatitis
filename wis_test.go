package whatitis

import "testing"

func Add(a, b int) int {
	return a + b
}

func Multiply(a, b int) int {
	return a * b
}

func Subtract(a, b int) int {
	return a - b
}

func Divide(a, b int) int {
	return a / b
}

func TestHello(t *testing.T) {
	What(t, "Testing positive operations", func(w *WhatContext) {

		w.It("Should successfully add numbers", func(i *ItContext) {
			expected := 2
			actual := Add(1, 1)
			i.Is(actual, EqualTo, expected)
			expected = 4
			actual = Add(2, 2)
			i.Is(actual, EqualTo, expected)
			expected = 5
			actual = Add(4, 1)
			i.Is(actual, EqualTo, expected)
			expected = 3
			actual = Add(1, 1)
			i.Is(actual, EqualTo, expected)
			expected = 5
			actual = Add(4, 1)
			i.Is(actual, EqualTo, expected)
		})

		w.It("Should successfully multiply numbers", func(i *ItContext) {
			expected := 2
			actual := Multiply(1, 2)
			i.Is(actual, EqualTo, expected)
			expected = 4
			actual = Multiply(2, 2)
			i.Is(actual, EqualTo, expected)
			expected = 2
			actual = Multiply(1, 1)
			i.Is(actual, EqualTo, expected)
		})

	})

	What(t, "Testing negative operations", func(w *WhatContext) {
		w.It("Should successfully subtract numbers", func(i *ItContext) {
			expected := 2
			actual := Subtract(3, 1)
			i.Is(actual, EqualTo, expected)
			expected = 4
			actual = Subtract(200, 196)
			i.Is(actual, EqualTo, expected)
		})

		w.It("Should successfully divide numbers", func(i *ItContext) {
			expected := 2
			actual := Divide(4, 2)
			i.Is(actual, EqualTo, expected)
			expected = 4
			actual = Divide(8, 2)
			i.Is(actual, EqualTo, expected)
			expected = 2
			actual = Divide(8, 4)
			i.Is(actual, EqualTo, expected)
		})
	})
}
