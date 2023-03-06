package binarysearch

// Binary search requires dataList to be sorted
// in this example we will assume dataList is sorted

// dataList := []int{3, 4, 5, 6, 7, 8, 21, 33, 44, 79, 86}
func BinarySearch(dataList []int, key int) (pos int, found bool) {

	pos = -1
	found = false

	low := 0
	high := len(dataList) - 1
	for low < high {

		median := (low + high) / 2
		//fmt.Println("High: ", high, "\t low:", low, "\t median:", median, "\t dataList[median]:", dataList[median], "\t key", key)
		if dataList[median] < key {
			low = median + 1
		} else if dataList[median] > key {
			high = median - 1
		} else {
			pos = median
			found = true
			break
		}
	}
	return
}
