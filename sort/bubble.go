package sort

func BubbleSort(list []int) {
	length := len(list)
	didSwap := false //标志在前面几轮是否有被交换过
	// 控制交换的轮数，进行n-1轮
	for i := length - 1; i > 0; i-- {
		// 每次从第一位开始比较，比较到第i位就不需要了，前一轮已经比较好了
		for j := 0; j < i; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
				didSwap = true
			}
		}
		// 如果一轮都没有比较，呢么就是有序的
		if !didSwap {
			return
		}
	}

}
