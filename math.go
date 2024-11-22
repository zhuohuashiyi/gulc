package gulc



// Pow 基于快速幂实现a^b， a, b都必须为正整数且程序不会考虑结果的溢出
func Pow(a, b int) int {
	res := 1
	for b > 0 {
		if (b & 1) > 0 {
			res *= a
		}
		a *= a
		b >>= 1
	}
	return res
}