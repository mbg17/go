package compare

type Sort struct {
	Arr []int
}

func Construct(arr []int) *Sort {
	return &Sort{
		Arr: arr,
	}
}
func (this *Sort) less(i, j int) bool {
	return this.Arr[i] < this.Arr[j]
}
func (this *Sort) each(i, j int) {
	temp := this.Arr[i]
	this.Arr[i] = this.Arr[j]
	this.Arr[j] = temp
}
func (this *Sort) ChooseSort() {
	for i := range this.Arr {
		min := i
		for j := i + 1; j < len(this.Arr); j++ {
			if this.less(j, min) {
				min = j
			}
		}
		this.each(min, i)
	}
}

func (this *Sort) InsertSort() {
	for i := 1; i < len(this.Arr); i++ {
		for j := i; j > 0; j-- {
			if this.less(j, j-1) {
				this.each(j, j-1)
			}
		}
	}
}

func (this *Sort) ShellSort() {
	n := len(this.Arr)
	h := 1
	for h < n/3 {
		h = 3*h + 1
	}
	for h >= 1 {
		for i := h; i < n; i += h {
			for j := i; j >= h && this.less(j, j-h); j -= h {
				this.each(j, j-h)
			}
		}
		h /= 3
	}
}
