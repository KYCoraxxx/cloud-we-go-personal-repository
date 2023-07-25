package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var total = len(nums1) + len(nums2)
	if total%2 == 1 {
		return float64(findKth(nums1, nums2, total/2+1))
	}
	return float64(findKth(nums1, nums2, total/2)+findKth(nums1, nums2, total/2+1)) / 2
}

func findKth(nums1 []int, nums2 []int, i int) float64 {
	if len(nums1) > len(nums2) {
		return findKth(nums2, nums1, i)
	}
	if len(nums1) == 0 {
		return float64(nums2[i-1])
	}
	if i == 1 {
		return float64(min(nums1[0], nums2[0]))
	}
	var k = min(i/2, len(nums1))
	var j = i - k
	if nums1[k-1] > nums2[j-1] {
		return findKth(nums1, nums2[j:], i-j)
	}
	return findKth(nums1[k:], nums2, i-k)
}

func min(i int, i2 int) int {
	if i < i2 {
		return i
	}
	return i2
}
