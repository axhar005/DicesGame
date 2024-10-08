package main

import "strings"

func contains(slice []int, item int) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func remove(slice []int, item int) []int {
	for i, v := range slice {
		if v == item {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

func centerText(text string, width int) string{
	padding := (width - len(text)) / 2
	if padding < 0 {
		padding = 0
	}
	return	strings.Repeat(" ", padding) + text + strings.Repeat(" ", padding)
}
