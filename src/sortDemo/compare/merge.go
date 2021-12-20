package compare

type Merge struct {
	Arr []int
	Aux []int
}

func MergeConstruct(a []int) {
	merge := &Merge{
		Arr: a,
		Aux: make([]int, len(a)),
	}
	merge.sort(a, 0, len(a)-1)
}

func (m *Merge) sort(a []int, lo int, hi int) {
	if hi <= lo {
		return
	}
	mid := (lo + hi) >> 1
	m.sort(a, lo, mid)
	m.sort(a, mid+1, hi)
	m.merge(a, lo, mid, hi)
}

func (m *Merge) merge(a []int, lo int, mid int, hi int) {
	left := lo
	right := mid + 1
	for i := 0; i <= hi; i++ {
		m.Aux[i] = a[i]
	}
	for i := lo; i <= hi; i++ {
		if left > mid {
			m.Arr[i] = m.Aux[right]
			right++
		} else if right > hi {
			m.Arr[i] = m.Aux[left]
			left++
		} else if m.less(right, left) {
			m.Arr[i] = m.Aux[right]
			right++
		} else {
			m.Arr[i] = m.Aux[left]
			left++
		}
	}
}

func (this *Merge) less(i, j int) bool {
	return this.Aux[i] < this.Aux[j]
}
