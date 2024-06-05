package dynamic

type memoized struct {
	f func(int) int
	memo map[int]int
}

func Memoize(f func(int) int) *memoized {
	return &memoized{
		f: f,
		memo: make(map[int]int),
	}
}

func (m *memoized) Call(n int) int {
	if val, ok := m.memo[n]; ok {
		return val
	}
	result := m.f(n)
	m.memo[n] = result
	return result
}


func Tribonacci(n int) int {
	if n == 0 || n == 1{
		return 0
	}
	if n == 2 {
		return 1
	}
	return Memoize(Tribonacci).Call(n-1) + Memoize(Tribonacci).Call(n-2) + Memoize(Tribonacci).Call(n-3)

}

