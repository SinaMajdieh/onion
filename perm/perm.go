package perm

func comb(raw [][]int, depth int, res [][]int, combo []int) [][]int {
	if depth >= len(raw) {
		cpy := []int{0, 0, 0}
		copy(cpy, combo)
		res = append(res, cpy)
		return res
	}
	for i := 0; i < len(raw[depth]); i++ {
		combo[depth] = raw[depth][i]
		res = comb(raw, depth+1, res, combo)
	}
	return res
}

func Permutation(data [][]int) [][]int {
	var res [][]int
	combo := []int{0, 0, 0}
	return comb(data, 0, res, combo)
}
