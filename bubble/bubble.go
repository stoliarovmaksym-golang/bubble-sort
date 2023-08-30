package bubble

func Sort(s []int) []int {
	sorted := false
	for !sorted {
		sorted = true
		for i := 0; i < len(s)-1; i++ {
			if s[i] > s[i+1] {
				sorted = false
				s[i], s[i+1] = s[i+1], s[i]
			}
		}
	}
	return s
}
