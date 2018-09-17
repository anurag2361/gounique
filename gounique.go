package unique

//Int returns unique int values
func Int(slice []int) map[int]struct{} {
	uMap := make(map[int]struct{})
	for _, val := range slice {
		uMap[val] = struct{}{}
	}
	return uMap
}
