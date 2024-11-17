package gulc


// 计算数字number在base进制下的位数
func DigitNumber(number int64, base int) int {
	digit := 0
	for number > 0 {
		digit++
		number /= int64(base)
	}
	return digit
}


// AddNumber 字符串加法，base指明数字的位数
func AddNumber(num1, num2 string, base int) string {
	m := len(num1)
	n := len(num2)
	k := max(m, n)
	ans := make([]byte, k + 1)
	i := m - 1
	j := n - 1
	c := 0

	// char2num 在进制数学上将字符转换为对应的数字，如'1' -> 1, 'a' -> 10
	var char2num func(ch byte) int = func(ch byte) int {
		if ch >= '0' && ch <= '9' {
			return int(ch - '0')
		}
		return int(ch - 'a')
	}

	var num2Char func(num int) byte = func(num int) byte {
		if num >= 0 && num <= 9 {
			return byte(num) + '0'
		}
		return byte(num) + 'a'
	}

	for i >= 0 || j >= 0 || c > 0 {
		sum := c
		if i >= 0 {
			sum += char2num(num1[i])
		}
		if j >= 0 {
			sum += char2num(num2[j])
		}
		ans[k] = num2Char(sum % base)
		i--
		j--
		k--
		c = sum / base
	}

	return string(ans[k + 1:])
}