package sortAlgorithms

func BaboSort(sli []int) {
	numOfSlice := len(sli)
	for i := 0; i < numOfSlice; i++ {
		for j := 1; j < numOfSlice-i; j++ {
			if sli[j-1] > sli[j] {
				sli[j-1], sli[j] = sli[j], sli[j-1]
			}
		}
	}
}
