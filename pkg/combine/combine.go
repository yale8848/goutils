package combine

import "strings"

var defSep = "_"

func CartCombine(data [][]string, sep string) []string {
	var _sep = defSep
	if sep != "" {
		_sep = sep
	}
	var _r []string
	lens := func(i int) int { return len(data[i]) }
	for i := make([]int, len(data)); i[0] < lens(0); next(i, lens) {
		var r []string
		for j, k := range i {
			r = append(r, data[j][k])
		}
		_r = append(_r, strings.Join(r, _sep))
	}
	return _r
}

func next(i []int, lens func(i int) int) {
	for j := len(i) - 1; j >= 0; j-- {
		i[j]++
		if j == 0 || i[j] < lens(j) {
			return
		}
		i[j] = 0
	}
}
