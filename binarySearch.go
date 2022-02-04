package main

func searchBinary(arr []int, key int) int {
	halfLen := len(arr) / 2
	val := arr[halfLen]
	if val < key {
		return searchBinary(arr[halfLen:], key)
	} else if val > key {
		return searchBinary(arr[:halfLen], key)
	} else if val == key {
		return val
	}
	return -1
}
