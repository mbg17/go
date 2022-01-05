package compare

type Compare interface {
	CompareTo()
}

//KMP判断是否为子串

func Kmp(a, b string) bool {
	bIndex := 0
	for i := 0; i < len(a) && bIndex < len(b); i++ {
		if a[i] == b[bIndex] {
			bIndex++
		} else {
			bIndex = 0
		}
	}
	return bIndex == len(b)
}
