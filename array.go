package gulc


// 返回有序数组中最小的不小于target的数的下标
func FindSmallestEqualGreat(nums []int, target int) int {
	n := len(nums)
	if target > nums[n - 1] {
		return -1
	}
	low := 0
	high := n - 1
	for low < high {
		mid := (low + high) / 2
		if nums[mid] >= target {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}


// FindLargestEqualLess 返回数组中最大的小于等于target的数的下标
func FindLargestEqualLess(nums []int, target int) int {
	n := len(nums)
	if nums[0] > target {
		return -1
	}
	low := 0
	high := n - 1
	for low < high {
		mid := (low + high + 1) / 2
		if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid
		}
	}
	return low
}



// matrixPow以O(logn)的时间复杂度计算matrix的order次方
func MatrixPow(matrix [][]int, order int) [][]int {
	n := len(matrix)
	// 定义一个n阶的单位矩阵
	E := make([][]int, n)
	for i := 0; i < n; i++ {
		E[i] = make([]int, n)
		E[i][i] = 1
	}
	// 使用快速幂算法
	for order > 0 {
		if (order & 1) > 0 {
			E = MatrixMultiply(E, matrix)
		}
		order >>= 1
		matrix = MatrixMultiply(matrix, matrix)
	}
	return E
}


// matrixMultiply返回两个矩阵的乘积矩阵
func MatrixMultiply(A, B [][]int) [][]int {
	m := len(A)
	n := len(A[0])
	p := len(B)
	q := len(B[0])
	if n != p {
		return nil
	}
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, q)
	}
	// 矩阵相乘
	for i := 0; i < m; i++ {
		for j := 0; j < q; j++ {
			sum := 0
			for k := 0; k < n; k++ {
				sum += A[i][k] * B[k][j]
			}
			res[i][j] = sum
		}
	}
	return res
}


// FindKthNumber 返回数组nums中第k小的数O(n)
func FindKthNumber(nums []int, k int) int {
	n := len(nums)
	var dfs func(int, int, int) int 
	// 快速选择算法
	dfs = func(left, right, k int) int {
		if left == right {
			return nums[left]
		}
		num := nums[left]
		ptr := left
		for i := left + 1; i <= right; i++ {
			if nums[i] < num {
				ptr++
				if ptr != i {
					nums[ptr], nums[i] = nums[i], nums[ptr]
				}
			}
		}
		nums[left] = nums[ptr]
		nums[ptr] = num
		if ptr - left + 1 == k {
			return num
		} else if ptr - left + 1 > k {
			return dfs(left, ptr - 1, k)
		} else {
			return dfs(ptr + 1, right, k - ptr + left - 1)
		}
	}
	return dfs(0, n - 1, k)
}



// CountInversePairs 返回数组nums中的逆序对数目
func CountInversePairs(nums []int) int {
	res := 0
	n := len(nums)
	var mergeSort func(int, int)
	mergeSort = func(low, high int) {
		if low >= high {
			return
		}
		mid := (low + high) / 2
		mergeSort(low, mid)
		mergeSort(mid + 1, high)
		tempArr := make([]int, high - low + 1)
		k := 0
		i := low
		j := mid + 1
		for i <= mid && j <= high {
			if nums[j] < nums[i] {
				res += (mid - i + 1)
				tempArr[k] = nums[j]
				k++
				j++
			} else {
				tempArr[k] = nums[i]
				k++
				i++
			}
		}
		for i <= mid {
			tempArr[k] = nums[i]
			k++
			i++
		}
		for j <= high {
			tempArr[k] = nums[j]
			k++
			j++
		}
		for k := 0; k < len(tempArr); k++ {
			nums[low + k] = tempArr[k]
		}
	}
	mergeSort(0, n - 1)
	return res
}


