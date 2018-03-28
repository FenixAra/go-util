package arrays

func Contains(sa []string, s string) bool {
	for _, v := range sa {
		if v == s {
			return true
		}
	}
	return false
}

func ContainsAny(sa, s []string) bool {
	m := make(map[string]struct{})
	for _, v := range s {
		m[v] = struct{}{}
	}

	for _, v := range sa {
		if _, ok := m[v]; ok {
			return true
		}
	}
	return false
}

func RemoveDuplicates(sa []string) []string {
	m := make(map[string]struct{})
	var res []string
	for _, v := range sa {
		if _, ok := m[v]; !ok {
			res = append(res, v)
			m[v] = struct{}{}
		}
	}

	return res
}

func AppendWithoutDuplicates(src []string, s []string) []string {
	m := make(map[string]struct{})
	for _, v := range src {
		m[v] = struct{}{}
	}

	for _, v := range s {
		if _, ok := m[v]; !ok {
			src = append(src, v)
		}
	}
	return src
}

func RemoveFromArray(src []string, s []string) []string {
	m := make(map[string]struct{})
	for _, v := range s {
		m[v] = struct{}{}
	}

	var newArray []string

	for _, v := range src {
		if _, ok := m[v]; !ok {
			newArray = append(newArray, v)
		}
	}
	return newArray
}

func RemoveFirstElement(src []string) []string {
	return RemoveNthElement(src, 1)
}

func RemoveNthElement(src []string, n int) []string {
	var a []string
	for i, _ := range src {
		if i != (n - 1) {
			a = append(a, src[i])
		}
	}
	return a
}
