package main

func searchLinear(arr []int, key int) int {
	for i, a := range arr {
		if a == key {
			return i
		}
	}
	return -1
}
