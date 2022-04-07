package sort

func ShellSOrt(list []int) {
	n := len(list)

	//每次减半，直到步长为1
	for step := n / 2; step >= 1; step /= 2 {
		//开始插入排序，每轮步长为step
		for i := step; i < n; i += step {
			for j := i - step; j >= 0; j -= step {
				// 满足插入就交换元素
				if list[j+step] < list[j] {
					list[j], list[j+step] = list[j+step], list[j]
					continue
				}
				break
			}
		}
	}
}
