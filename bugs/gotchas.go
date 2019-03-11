package main

import (
	"strconv"
	"sort"
)

func ReturnInt() (int) {
	return 1
}

func ReturnFloat() (float32) {
	return float32(1.1)
}

func ReturnIntArray() ([3]int) {
	return [3]int{1, 3, 4}
}

func ReturnIntSlice() ([]int) {
	return []int{1, 2, 3}
}

func IntSliceToString(input []int) (string) {
	var result string
	for _, val := range input {
		temp := strconv.Itoa(val)
		result += temp
	}
	return result
}

func MergeSlices(slice1 []float32, slice2 []int32) ([]int) {
	result := make([]int, 0, len(slice1) + len(slice2))
	for _, val := range slice1 {
		result = append(result, int(val))
	}
	for _, val := range slice2 {
		result = append(result, int(val))
	}
	return result
}

func GetMapValuesSortedByKey(input map[int]string) ([]string) {
	mapLen := len(input)
	result := make([]string, 0, mapLen)
	
	keys := make([]int, 0, mapLen)
	for key, _ := range input {
        keys = append(keys, key)
    }

    sort.Ints(keys)

    for _, key := range keys {
    	result = append(result, input[key])
    }
	return result
}