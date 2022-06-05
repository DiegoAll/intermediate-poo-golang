package main

import "testing"

func TestSum(t *testing.T) {
	/* 	total := Sum(5, 6)

	   	if total != 20 {
	   		t.Errorf("Sum was incorrect, got %d expected %d", total, 20)
	   	} */

	tables := []struct {
		a int
		b int
		n int
	}{
		{1, 2, 3},
		{2, 2, 4},
		{25, 26, 51},
	}

	for _, item := range tables {
		total := Sum(item.a, item.b)

		if total != item.n {
			t.Errorf("Sum was incorrect, got %d expected %d", total, item.n)
		}
	}
}

func TestMax(t *testing.T) {
	tables := []struct {
		a int
		b int
		n int
	}{
		{4, 2, 4},
		{3, 2, 3},
	}

}
