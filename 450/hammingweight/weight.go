package hammingweight

func hammingWeight(n int) int {
	count := 0
	for n != 0 {
		count += int(n & 1) // Add the least significant bit
		n >>= 1             // Right shift to process the next bit
	}
	return count
}
func countBits(n int) []int {
	result := make([]int, n+1)
	for i := 1; i <= n; i++ {
		result[i] = result[i/2] + i%2
	}
	return result
}

func reverseBits(n uint32) uint32 {
	var result uint32 = 0
	for n > 0 {
		result <<= 1
		result |= n & 1
		n >>= 1
	}
	return result
}

func missingNumber(nums []int) int {
	n := len(nums)
	ans := 0
	for i := 1; i < n+1; i++ {
		ans = ans ^ i
	}
	for _, num := range nums {
		ans = ans ^ num
	}
	return ans
}
