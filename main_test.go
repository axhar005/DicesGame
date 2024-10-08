package main

import "testing"

func TestScoreCount(t *testing.T){
	{
		slice := []int{1, 5, 5, 4}
		result := scoreCount(slice)
		expected := 0
		if result != expected {
			t.Errorf("error: scoreCount(%v) = %5d; want: %5d", slice, result, expected)
		}
	}

	{
		slice := []int{1}
		result := scoreCount(slice)
		expected := 100
		if result != expected {
			t.Errorf("error: scoreCount(%v) = %5d; want: %5d", slice, result, expected)
		}
	}

	{
		slice := []int{3,3,3}
		result := scoreCount(slice)
		expected := 300
		if result != expected {
			t.Errorf("error: scoreCount(%v) = %5d; want: %5d", slice, result, expected)
		}
	}

	{
		slice := []int{3,3,3,3,3}
		result := scoreCount(slice)
		expected := 1200
		if result != expected {
			t.Errorf("error: scoreCount(%v) = %5d; want: %5d", slice, result, expected)
		}
	}

	{
		slice := []int{3,3,3,3,3,3}
		result := scoreCount(slice)
		expected := 2400
		if result != expected {
			t.Errorf("error: scoreCount(%v) = %5d; want: %5d", slice, result, expected)
		}
	}

	{
		slice := []int{1,1,1,1}
		result := scoreCount(slice)
		expected := 1100
		if result != expected {
			t.Errorf("error: scoreCount(%v) = %5d; want: %5d", slice, result, expected)
		}
	}

	{
		slice := []int{1,1,1,1,1,1}
		result := scoreCount(slice)
		expected := 2000
		if result != expected {
			t.Errorf("error: scoreCount(%v) = %5d; want: %5d", slice, result, expected)
		}
	}
}