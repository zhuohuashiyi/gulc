package gulc


// substrImplKMP kmp算法实现子字符串查找
func SubstrImplKMP(str, pattern string) []int {
	n := len(str)
	m := len(pattern)
	result := []int{}

	// 计算模式字符串的最长前缀后缀（next）数组
	next := computeNext(pattern)
	i, j := 0, 0
	for i < n {
		if pattern[j] == str[i] {
			i++
			j++
		}
		if j == m {
			// 找到匹配项
			result = append(result, i-j)
			j = next[j-1]
		} else if i < n && pattern[j] != str[i] {
			if j != 0 {
				j = next[j-1]
			} else {
				i++
			}
		}
	}
	return result
}


func computeNext(str string) []int {
	m := len(str)
	next := make([]int, m)
	length := 0
	i := 1
	for i < m {
		if str[i] == str[length] {
			length++
			next[i] = length
			i++
		} else {
			if length != 0 {
				length = next[length-1]
			} else {
				next[i] = 0
				i++
			}
		}
	}

	return next
}


// 滚动哈希实现子字符串查找O(n)
func SubstrImplHash(str, pattern string) []int {
	m := len(str)
	n := len(pattern)
	res := make([]int, 0)
	if m < n {
		return res
	}
	var base, mod int64 = 31, 1000000007
	var h1, h2, h3 int64 = 0, 0, 1
	h1 = int64(computeStrHash(str[m - n: ], base, mod))
	h2 = int64(computeStrHash(pattern, base, mod))
	for i := 0; i < n-1; i++ {
		h3 *= base
		h3 %= mod
	}
	if h1 == h2 {
		res = append(res, m - n)
	}
	for i := m - n - 1; i >= 0; i-- {
		h1 = (h1 - int64(str[i + n] - 'a' + 1) * h3 % mod + mod) % mod
		h1 = h1 * base + int64(str[i] - 'a' + 1) 
		h1 %= mod
		if h1 == h2 {
			res = append(res, i)
		}
	}
	return res
}

// 计算字符串的哈希值
func computeStrHash(str string, base, mod int64) int {
	var hash int64 = 0
	n := len(str)
	for i := n - 1; i >= 0; i-- {
		hash = hash*base + int64(str[i]-'a'+1)
		hash %= mod
	}
	return int(hash)
}


// isSubseq 利用二分查找优化判断子序列
func IsSubseq(seq, subSeq string) bool {
	pos := make([][]int, 26)
	for i := 0; i < 26; i++ {
		pos[i] = make([]int, 0)
	}
	n := len(seq)
	m := len(subSeq)
	for i := 0; i < n; i++ {
		pos[int(seq[i] - 'a')] = append(pos[int(seq[i] - 'a')], i)
	}
	i := 0
	j := 0
	for j < m {
		k := FindSmallestEqualGreat(pos[int(subSeq[j] - 'a')], i)
		if k == -1 {
			return false
		}
		i = pos[int(subSeq[j] - 'a')][k] + 1
		j++
	}
	return true
}