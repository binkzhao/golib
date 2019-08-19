package ds

import (
	"sort"
)

/**
* 使用golang实现常见排序算法，以及使用sort包来实现这些算法：统一都是升序排序
 */

// 冒泡排序
func BulleSort(items []int) {
	for i := 0; i < len(items); i++ {
		for j := 0; j < len(items)-i-1; j++ {
			if items[j+1] < items[j] {
				swap(&items[j], &items[j+1])
			}
		}
	}
}

// 冒泡排序 使用sort包
func BubbleSortUsingSortPackage(data sort.Interface) {
	len := data.Len()
	for i := 0; i < len-1; i++ {
		for j := 0; j < len-1-i; j++ {
			if data.Less(j+1, j) {
				data.Swap(j+1, j)
			}
		}
	}
}

// 插入排序
func InsertSort(items []int) {
	for i := 1; i < len(items); i++ {
		for j := i; j > 0; j-- {
			if items[j-1] <= items[j] {
				// 已经是有序，不需要再继续比较
				break
			}

			swap(&items[j-1], &items[j])
		}
	}
}

// 插入排序 使用sort包
func InsertSortUsingSortPackage(data sort.Interface) {
	r := data.Len() - 1
	for i := 1; i <= r; i++ {
		for j := i; j > 0 && data.Less(j, j-1); j-- {
			data.Swap(j, j-1)
		}
	}
}

// 简单选择排序
func SelectSort(items []int) {
	for i := 0; i < len(items)-1; i++ {
		var minIdx = i
		for j := i + 1; j < len(items); j++ {
			if items[j] < items[minIdx] {
				minIdx = j
			}
		}
		swap(&items[i], &items[minIdx])
	}
}

// 简单选择排序 使用sort包
func SelectSortUsingSortPackage(data sort.Interface) {
	r := data.Len() - 1
	for i := 0; i < r; i++ {
		min := i
		for j := i + 1; j <= r; j++ {
			if data.Less(j, min) {
				min = j
			}
		}
		data.Swap(i, min)
	}
}

// 快速排序
func QuickSort(src []int, first, last int) {
	flag := first
	left := first
	right := last

	if first >= last {
		return
	}

	for first < last {
		for first < last {
			if src[last] >= src[flag] {
				last--
				continue
			}
			swap(&src[last], &src[flag])
			flag = last
			break
		}

		for first < last {
			if src[first] <= src[flag] {
				first++
				continue
			}
			swap(&src[first], &src[flag])
			flag = first
			break
		}
	}

	QuickSort(src, left, flag-1)
	QuickSort(src, flag+1, right)
}

// 归并排序
func MergeSort(items []int) []int {
	len := len(items)
	if len == 1 {
		return items
	}

	middle := len / 2
	left, right := items[:middle], items[middle:]

	return merge(MergeSort(left), MergeSort(right))
}

func merge(left, right []int) (result []int) {
	result = make([]int, len(right)+len(left))
	var i, j, p int
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result[p] = left[i]
			i++
		} else {
			result[p] = right[j]
			j++
		}
		p++
	}

	for i < len(left) {
		result[p] = left[i]
		p++
		i++
	}

	for j < len(right) {
		result[p] = right[j]
		p++
		j++
	}
	return
}

// 交换两个元素的值
func swap(a, b *int) {
	*a, *b = *b, *a
}
