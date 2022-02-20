package compare

import (
	"math"
)

type Merge struct {
	Arr []int
	Aux []int
}

func MergeConstruct(a []int) {
	merge := &Merge{
		Arr: a,
		Aux: make([]int, len(a)),
	}
	//merge.sortFr(0, len(a)-1)
	merge.sortBu(len(a))
}

// 自底向上
func (m *Merge) sortBu(N int) {
	for i := 1; i < N; i = i + i {
		for j := 0; j < N-i; j += i + i {
			m.merge(j, j+i-1, int(math.Min(float64(j+i+i-1), float64(N-1))))
		}
	}
}

// 自顶向下
// 拆分切片直到不能切为止
func (m *Merge) sortFr(lo int, hi int) {
	if hi <= lo {
		return
	}
	mid := (lo + hi) >> 1
	m.sortFr(lo, mid)
	m.sortFr(mid+1, hi)
	m.merge(lo, mid, hi)
}

func (m *Merge) merge(lo int, mid int, hi int) {
	left := lo
	right := mid + 1
	// 将原始数组塞到辅助数组中
	for i := 0; i <= hi; i++ {
		m.Aux[i] = m.Arr[i]
	}
	for i := lo; i <= hi; i++ {
		// 左边用完，取右边
		if left > mid {
			m.Arr[i] = m.Aux[right]
			right++
			// 右边用完取左边
		} else if right > hi {
			m.Arr[i] = m.Aux[left]
			left++
			// 右边小于左边，放右边
		} else if m.less(right, left) {
			m.Arr[i] = m.Aux[right]
			right++
			// 左边小于右边放左边
		} else {
			m.Arr[i] = m.Aux[left]
			left++
		}
	}
}

func (this *Merge) less(i, j int) bool {
	return this.Aux[i] < this.Aux[j]
}
