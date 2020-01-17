package main

import "testing"

func Test_bonus(t *testing.T) {
	tests := []struct {
		name string
		sales []int
		want int
	}{
		//{"bonus received", []int{12_000, 8_000, 15_000, 8_000}, 350},
		{"for all", []int{12_000, 8_000, 15_000, 8_000}, 350},
		{"Nothing", []int{10_000, 10_000,}, 0},
	}
	for _, test := range tests {
		got := bonus(test.sales)
		if got != test.want {
			t.Error("For amount of bonus test:", test.name, "got:", got, "want:", test.want)
		}
	}
}