package circular

const one int64 = 1

func Get(circle []int64, index int64) int64 {
	return circle[index%len64(circle)]
}

func Sum(circle []int64, start, end int64) int64 {
	l := len64(circle)
	return sum(circle[start%l:]) + sum(circle[:end%l]) + ((end/l - start/l - 1) * sum(circle))
}

func SumSlice(circle []int64, startFrom, start, end int64) []int64 {
	size := int(end - start)
	r := make([]int64, size)
	v := Sum(circle, startFrom, start+one)
	for i := 0; i < size; i++ {
		r[i] = v
		v += Get(circle, start+int64(i)+one)
	}
	return r
}

func len64(input []int64) int64 {
	return int64(len(input))
}

func sum(input []int64) (Σ int64) {
	for i := range input {
		Σ += input[i]
	}
	return
}
