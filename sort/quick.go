package sort

// 额外开辟空间
func QuickSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	} else {
		splitdata := arr[0]       //以第一个元素为基准元素
		low := make([]int, 0, 0)  //存储比基准元素小的元素
		high := make([]int, 0, 0) //存储比基准元素大的元素
		mid := make([]int, 0, 0)  //存储与基准元素相等的元素
		mid = append(mid, arr[0]) //先将基准元素压入
		for i := 1; i < length; i++ {
			if arr[i] < splitdata {
				low = append(low, arr[i])
			} else if arr[i] > splitdata {
				high = append(high, arr[i])
			} else {
				mid = append(mid, arr[i])
			}
		}
		low, high = QuickSort(low), QuickSort(high)   //递归处理剩余的无序数据
		myarr := append(append(low, mid...), high...) //合并所有的有续集
		return myarr
	}

}

//
func QuickSort2(array []int, begin, end int) {
	if begin < end {
		// 进行切分
		loc := partition(array, begin, end)
		// 对左边进行快排
		QuickSort2(array, begin, loc-1)
		// 对右边进行快排
		QuickSort2(array, loc+1, end)
	}
}

// 切分函数，并返回切分元素的下标
func partition(array []int, beign, end int) int {
	i := beign + 1 // 将array[begin]作为基准数，因此从array[begin+1]开始与基准数比较！
	j := end       //数组最后一位

	// 没有重合之前
	for i < j {
		if array[i] > array[beign] {
			array[i], array[j] = array[j], array[i] //交换
			j--
		} else {
			i++
		}
	}
	/* 跳出while循环后，i = j。
	   * 此时数组被分割成两个部分 --> array[begin+1] ~ array[i-1] < array[begin]
	   * --> array[i+1] ~ array[end] > array[begin]
	   * 这个时候将数组array分成两个部分，再将array[i]与array[begin]进行比较，决定
	   array[i]的位置。
	   * 最后将array[i]与array[begin]交换，进行两个分割部分的排序！以此类推，直到
	   最后i = j不满足条件就退出！
	*/
	if array[i] >= array[beign] { // 这里必须要取等“>=”，否则数组元素由相同的值组成时，会出现错误！
		i--
	}
	array[beign], array[i] = array[i], array[beign]
	return i
}

// 算法改进
func QuickSort3(array []int, begin, end int) {
	if begin < end {
		// 当数组小于 4 时使用直接插入排序
		if end-begin <= 4 {
			InsertSort(array[begin : end+1])
			return
		}
		// 进行切分
		loc := partition(array, begin, end)
		// 对左部分进行快排
		QuickSort3(array, begin, loc-1)
		// 对右部分进行快排
		QuickSort3(array, loc+1, end)
	}
}

// 三切分的快速排序
func QuickSort4(array []int, begin, end int) {
	if begin < end {
		// 三向切分函数，返回左边和右边下标
		lt, gt := partition3(array, begin, end)
		// 从lt到gt的部分是三切分的中间数列
		// 左边三向快排
		QuickSort4(array, begin, lt-1)
		// 右边三向快排
		QuickSort4(array, gt+1, end)
	}
}

// 切分函数，并返回切分元素的下标
func partition3(array []int, begin, end int) (int, int) {
	lt := begin       // 左下标从第一位开始
	gt := end         // 右下标是数组的最后一位
	i := begin + 1    // 中间下标，从第二位开始
	v := array[begin] // 基准数
	// 以中间坐标为准
	for i <= gt {
		if array[i] > v { // 大于基准数，那么交换，右指针左移
			array[i], array[gt] = array[gt], array[i]
			gt--
		} else if array[i] < v { // 小于基准数，那么交换，左指针右移
			array[i], array[lt] = array[lt], array[i]
			lt++
			i++
		} else {
			i++
		}
	}
	return lt, gt
}

func QuickSort5(array []int, begin, end int) {
	for begin < end {
		// 进行切分
		loc := partition(array, begin, end)
		// 那边元素少先排哪边
		if loc-begin < end-loc {
			// 先排左边
			QuickSort5(array, begin, loc-1)
			begin = loc + 1
		} else {
			// 先排右边
			QuickSort5(array, loc+1, end)
			end = loc - 1
		}
	}
}
