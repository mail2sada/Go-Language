package bubble

func BubbleSort(dataList []int) (list []int) {

	var sorted, swapped bool = false, false

	for !sorted {

		swapped = false

		for count := 0; count < len(dataList)-1; count++ {
			if dataList[count+1] < dataList[count] {
				dataList[count+1], dataList[count] = swap(dataList[count+1], dataList[count])
				swapped = true
			}

		}
		if !swapped {
			sorted = true
		}

	}
	list = dataList

	return
}

func swap(a, b int) (int, int) {
	return b, a
}
