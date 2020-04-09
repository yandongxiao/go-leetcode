package contains_duplicate_217

func containsDuplicate(nums []int) bool {
	mapping := make(map[int]struct{})
	for _, x := range nums {
		if _, ok := mapping[x]; ok {
			return true
		}
		mapping[x] = struct{}{}
	}
	return false
}
