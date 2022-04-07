package sort

func selectSort(list []int) {
	n := len(list)
	// 进行N-1轮迭代
	for i := 0; i < n-1; i++ {
		//每次从第i位开始,找最小的元素
		min := list[i]
		minIndex := i
		for j := i + 1; j < n; j++ {
			if list[j] < min {
				min = list[j]
				minIndex = j
			}
		}
		//如果在i轮找到更小的数，就进行交换
		if i != minIndex {
			list[i], list[minIndex] = list[minIndex], list[i]
		}

	}
}

// 我们每一轮，除了找最小数之外，还找最大数，然后分别和前面和后面的元素交换，这样循环次数减少一半
func selectSortBetter(list []int) {
	n := len(list)
	// 进行N-1轮迭代
	for i := 0; i < n-1; i++ {
		//每次从第i位开始,找最小的元素
		minIndex := i
		maxIndex := i
		// 找最大值和最小值的下标
		for j := i + 1; j < n; j++ {
			if list[j] > list[maxIndex] {
				maxIndex = j
				continue
			}
			if list[j] < list[minIndex] {
				minIndex = j
			}

		}
		//交换
		if maxIndex == i && minIndex != n-i-1 {
			// 如果最大值是开头的元素，而最小值不是最尾的元素
			// 先将最大值和最尾的元素交换
			list[n-i-1], list[maxIndex] = list[maxIndex], list[n-i-1]
			// 然后最小的元素放在最开头
			list[i], list[minIndex] = list[minIndex], list[i]
		} else if maxIndex == i && minIndex == n-i-1 {
			// 如果最大值在开头，最小值在结尾，直接交换
			list[minIndex], list[maxIndex] = list[maxIndex], list[minIndex]
		} else {
			// 否则先将最小值放在开头，再将最大值放在结尾
			list[i], list[minIndex] = list[minIndex], list[i]
			list[n-i-1], list[maxIndex] = list[maxIndex], list[n-i-1]
		}

	}
}
