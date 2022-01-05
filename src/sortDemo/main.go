package main

import (
	"fmt"
	"math/rand"
	"sort"
	"sortDemo/compare"
	"strings"
)

type Transaction struct {
	amount float64
}

func (t Transaction) String() string {
	return fmt.Sprintf("%v", t.amount)
}

type Transactions []Transaction

func (a Transactions) Len() int           { return len(a) }
func (a Transactions) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Transactions) Less(i, j int) bool { return a[i].amount < a[j].amount }

func main() {
	s := compare.Construct([]int{-9, -8, 7, -4, 1, 5, 3, 6, 7, 0, 3, 5})
	s.ChooseSort()
	fmt.Println(s.Arr)
	s = compare.Construct([]int{-9, -8, 7, -4, 1, 5, 3, 6, 7, 0, 3, 5})
	s.InsertSort()
	fmt.Println(s.Arr)
	s = compare.Construct([]int{-9, -8, 7, -4, 1, 5, 3, 6, 7, 0, 3, 5})
	s.ShellSort()
	fmt.Println(s.Arr)
	transactions := make([]Transaction, 0)
	for i := 0; i < 10; i++ {
		f := rand.Float64()
		transactions = append(transactions, Transaction{amount: f})
	}
	fmt.Println(transactions)
	sort.Sort(Transactions(transactions))
	fmt.Println(transactions)
	ints := []int{-9, -8, 7, -4, 1, 5, 3, 6, 7, 0, 3, 5}
	compare.MergeConstruct(ints)
	fmt.Println(ints)
	kmp := compare.Kmp("abc", "bc")
	fmt.Println(kmp)
}

func repeatedStringMatch(a string, b string) int {
	s := ""
	ans := 0
	max := len(a) + len(b)
	for len(s) < max {
		s += a
		ans++
		if strings.Contains(s, b) {
			return ans
		}
	}
	return -1
}
