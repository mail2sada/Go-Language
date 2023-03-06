package interpolationsearch

func InterpolationSearch(list []int, key int) (pos int, found bool) {
	min, max := list[0], list[len(list)-1]

	pos = -1
	found = false

	low, high := 0, len(list)-1

	for {
		if key < min {
			return
		}

		if key > max {
			return
		}

		// make a guess of the location
		var guess int
		if high == low {
			guess = high
		} else {
			size := high - low
			offset := int(float64(size-1) * (float64(key-min) / float64(max-min)))
			guess = low + offset
		}

		// maybe we found it?
		if list[guess] == key {
			// scan backwards for start of value range
			for guess > 0 && list[guess-1] == key {
				guess--
			}
			pos = guess
			found = true
			return
		}

		// if we guessed to high, guess lower or vice versa
		if list[guess] > key {
			high = guess - 1
			max = list[high]
		} else {
			low = guess + 1
			min = list[low]
		}
	}
}
