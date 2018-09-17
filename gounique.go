package unique

//Int returns unique int values in a map paired with an empty object
func Int(slice []int) map[int]struct{} {
	uMap := make(map[int]struct{})
	for _, val := range slice {
		uMap[val] = struct{}{}
	}
	return uMap
}

//Strings returns unique string values in a map paired with an empty object
func Strings(slice []string) map[string]struct{} {
	uMap := make(map[string]struct{})
	for _, val := range slice {
		uMap[val] = struct{}{}
	}
	return uMap
}
