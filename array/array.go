package array

// Erfen
//
//	@Description: 二分法查找目标数字下标
//	@param nums 顺序数组
//	@param target 目标数字
//	@return int 数组下标
func Erfen(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	for l <= r {
		m := (l + r) / 2
		if target > nums[m] {
			l = m + 1
		} else if target < nums[m] {
			r = m - 1
		} else {
			return m
		}
	}

	return -1
}

// RemoveElement
//
//	@Description: 删除nums中值为target的元素，返回删除后的数组
//	@param nums 数组
//	@param target 目标值
//	@return int 删除后的长度
func RemoveElement(nums []int, target int) int {
	l := len(nums)
	r := 0
	for i := 0; i < l; i++ {
		if nums[i] != target {
			nums[r] = nums[i]
			r++
		}
	}
	return r
}

// SortedSquares
//
//	@Description:计算顺序数组中每个元素的平方，并且排序
//	@param nums 顺序数组
//	@return []int 数组
func SortedSquares(nums []int) []int {
	l := 0
	r := len(nums) - 1
	nums1 := make([]int, len(nums))
	i := r
	for l <= r {
		i1, i2 := nums[l]*nums[l], nums[r]*nums[r]
		if i1 > i2 {
			nums1[i] = i1
			l++
		} else {
			nums1[i] = i2
			r--
		}
		i--
	}
	return nums1
}

func MinSubArrayLen(nums []int, target int) int {
	l, r := 0, 0
	length := 0
	result := 0
	sum := 0
	for r < len(nums) && l <= r {
		sum1 := nums[r] + sum
		if sum1 <= target {
			sum += nums[r]
			r++
		}
		if sum1 > target {
			length = r - l
			if length < result || result == 0 {
				result = length
			}
			sum -= nums[l]
			l++
		}
	}
	return result

}
