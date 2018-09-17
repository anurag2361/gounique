package unique

//Int returns unique int values
func Int(slice []int) map[int]struct{} {
	uMap := make(map[int]struct{})
	for key := range slice {
		uMap[key] = struct{}{}
	}
	return uMap
}
