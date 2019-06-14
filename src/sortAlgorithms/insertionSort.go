package sortAlgorithms

func InsertionSort(sli []int) {
	var j int
	for i, v := range sli {
		j = i
		for j > 0 {
			if v < sli[j-1] {
				sli[j] = sli[j-1]
			} else {
				break
			}
			j--
		}
		sli[j] = v
	}
}
