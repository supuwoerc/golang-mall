package helper

// RemoveDuplicateStrings 字符串切片去重
func RemoveDuplicateStrings(slides []string) []string {
	result := make([]string, 0, len(slides))
	temp := make(map[string]struct{})
	for _, item := range slides {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}
