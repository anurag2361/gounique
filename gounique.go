package unique

import "sort"

//Int returns unique int values in a slice
func Int(slice []int) []int {
	uMap := make(map[int]struct{})
	result := []int{}
	for _, val := range slice {
		uMap[val] = struct{}{}
	}
	for key := range uMap {
		result = append(result, key)
	}
	sort.Ints(result)
	return result
}

//Strings returns unique string values in a slice
func Strings(slice []string) []string {
	uMap := make(map[string]struct{})
	result := []string{}
	for _, val := range slice {
		uMap[val] = struct{}{}
	}
	for key := range uMap {
		result = append(result, key)
	}
	sort.Strings(result)
	return result
}
