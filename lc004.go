package algorithm

/*
给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
算法的时间复杂度应该为 O(log (m+n)) 。

示例 1：
输入：nums1 = [1,3], nums2 = [2]
输出：2.00000
解释：合并数组 = [1,2,3] ，中位数 2

示例 2：
输入：nums1 = [1,2], nums2 = [3,4]
输出：2.50000
解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
*/

/*
寻找中位数：
假设是一个有序数组 S，共 N 个数，若 N 为奇数，中位数 = S[n/2+1] , 若 N 为偶数，中位数 = (S[n/2] + S[n/2+1])/2
现在是两个有序数组，我们需要找其中第 k 位小的数

从两个数组分别取 S1[k/2], S2[k/2]
if S1[k/2] > S2[k/2]; k大的数肯定不在 S2[:k/2] 中， 同理
if S1[k/2] < S2[k/2]; k大的数肯定不在 S1[:k/2] 中， 否则
k大的数 = S1[k/2] = S2[k/2]
*/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n := len(nums1) + len(nums2)
	if n%2 == 1 {
		return float64(findKth(nums1, nums2, n/2+1))
	} else {
		return float64(findKth(nums1, nums2, n/2)+findKth(nums1, nums2, n/2+1)) / 2.0
	}
}

// 递归查询
func findKth(nums1 []int, nums2 []int, k int) int {
	len1, len2 := len(nums1), len(nums2)
	// 确保 nums1 是较短的数组
	if len1 > len2 {
		return findKth(nums2, nums1, k)
	}

	// 如果 nums1 为空，直接返回 nums2 中的第 k 个元素
	if len1 == 0 {
		return nums2[k-1]
	}

	// 走到这里，nums1 一定是非空的，因为确保了 nums1 是较短的数组，nums 一定非空
	if k == 1 {
		return min(nums1[0], nums2[0])
	}

	// 在 nums1 和 nums2 中找 k 个元素出来比较，nums1 取前 i 个，nums2 取前 j 个
	i := min(len1, k/2)
	j := k - i

	// 如果 nums2[j-1] 大，说明 nums1 中的前 i 个元素都小于 nums2[j-1]，因此可以排除 nums1 的前 i 个元素
	// 如果 nums1[i-1] 大，说明 nums2 中的前 j 个元素都小于 nums1[i-1]，因此可以排除 nums2 的前 j 个元素
	// 如果相等，说明找到了第 k 个元素
	if nums1[i-1] < nums2[j-1] {
		return findKth(nums1[i:], nums2, k-i)
	} else if nums1[i-1] > nums2[j-1] {
		return findKth(nums1, nums2[j:], k-j)
	} else {
		return nums1[i-1]
	}
}

// 迭代查询
func getKth(nums1, nums2 []int, k int) int {
	// 初始化查询左边界
	index1, index2 := 0, 0
	for {
		// 已经全部排除
		if index1 == len(nums1) {
			return nums2[index2+k-1]
		}
		if index2 == len(nums2) {
			return nums1[index1+k-1]
		}
		// 直接比较
		if k == 1 {
			return min(nums1[index1], nums2[index2])
		}

		// 二分查找
		newIndex1 := min(index1+k/2, len(nums1)) - 1
		newIndex2 := min(index2+k/2, len(nums2)) - 1
		pivot1, pivot2 := nums1[newIndex1], nums2[newIndex2]
		if pivot1 <= pivot2 {
			k -= (newIndex1 - index1 + 1)
			index1 = newIndex1 + 1
		} else {
			k -= (newIndex2 - index2 + 1)
			index2 = newIndex2 + 1
		}
	}
	return 0
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
